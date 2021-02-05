package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type requestParams map[string]interface{}

var params = requestParams{
	"Haspid":     "0",
	"Featureid":  "-1",
	"Vendorid":   "0",
	"Productid":  "0",
	"Filterfrom": "1",
	"Filterto":   "200",
}

type config struct {
	HaspUris []string `env:"HASP_URIS"`
}

var Cfg config
var proport int

func init() {
	flag.IntVar(&proport, "proport", 8181, "exporter port")
}

type MetricJson struct {
	Id       string `json:"ndx"`
	Vendor   string `json:"ven"`
	Name     string `json:"haspname"`
	Haspid   string `json:"haspid"`
	Ip       string `json:"ip"`
	Feature  string `json:"fn"`
	Limit    string `json:"logl"`
	License  string `json:"lic"`
	Sessions string `json:"sesc"`
	Product  string `json:"prname"`
}

var replacer = strings.NewReplacer("<br>", " ", "</br>", " ", "&nbsp;", " ", "<nobr>", " ", "</nobr>", " ")
var (
	// CommitHash используется в Makefile для проставления хеша коммита
	CommitHash = "111"
	// Version используется в Makefile аналогично с CommitHash
	Version = "1.0"
	// get free sessions
	freeSessions = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "free_sessions",
		Help: "Number of free sessions of license to hasp connect",
	},
		[]string{"Product", "Feature", "Host", "Vendor"},
	)
	// get unused license executions
	freeExecutions = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "free_executions",
		Help: "Number of license executions remaining",
	},
		[]string{"Product", "Feature", "Host", "Vendor"},
	)
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	output := fmt.Sprintf("version - %s\nhashcommit - %s", Version, CommitHash)
	log.Printf("Versionhandler return version %s and hash %s", Version, CommitHash)
	fmt.Fprint(w, output)
}

func collectData() {
	digitPattern := regexp.MustCompile(`\d+`)
	go func() {
		for {
			for _, uri := range Cfg.HaspUris {
				u, err := url.Parse(uri)
				if err != nil {
					panic(err)
				}
				var datas []MetricJson
				q := u.Query()
				for k, v := range params {
					q.Set(k, v.(string))
				}
				u.RawQuery = q.Encode()
				err = json.Unmarshal([]byte(parseHasp(u.String())), &datas)
				if err != nil {
					panic(err)
				}
				for _, feat := range datas {
					licenses := -1
					//println(feat.Feature, feat.Id, feat.Haspid, feat.Limit, feat.Sessions, feat.License)
					limit, _ := strconv.Atoi(feat.Limit)
					sessions, _ := strconv.Atoi(feat.Sessions)
					licenses, _ = strconv.Atoi(digitPattern.FindString(feat.License))
					if limit > 0 {
						freeSessions.With(prometheus.Labels{"Host": u.Host, "Vendor": feat.Vendor, "Product": feat.Product, "Feature": feat.Feature}).Set(float64(limit - sessions))
					} else {
						freeSessions.With(prometheus.Labels{"Host": u.Host, "Vendor": feat.Vendor, "Product": feat.Product, "Feature": feat.Feature}).Set(-1)
					}
					if licenses >= 0 && strings.Contains(feat.License, "Executions") {
						freeExecutions.With(prometheus.Labels{"Host": u.Host, "Vendor": feat.Vendor, "Product": feat.Product, "Feature": feat.Feature}).Set(float64(licenses))
					}
				}
			}
		}
	}()
}

func main() {
	Cfg = config{}
	env.Parse(&Cfg)
	flag.Parse()
	prometheus.MustRegister(freeSessions)
	prometheus.MustRegister(freeExecutions)
	collectData()
	router := mux.NewRouter()
	router.HandleFunc("/version", versionHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", router)

	level := os.Getenv("TRACE")
	if level == "true" {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
	// Only log the warning severity or above.

	log.Warnln("Server is listening...")
	log.Fatalln(http.ListenAndServe(":"+strconv.Itoa(proport), nil))
}

func parseHasp(url string) string {
	leftStrip := "/*JSON:features*/"
	rightStrip := "\n/*\n <admin_status>\n  <code>0</code>\n  <text>SNTL_ADMIN_STATUS_OK</text>\n </admin_status>\n*/\n\n"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	dataString := replacer.Replace(string(body))
	out := strings.TrimLeft(strings.TrimRight(dataString, rightStrip), leftStrip)
	out = "[" + out + "]"
	time.Sleep(time.Second)
	return out
}
