# prom-hasp-exporter
 Simple Prometeus Hasp Lic Exporter
 Экспортер Hasp лицензий для Prometeus
# Сборка

`make`
#  Установка
`make install`

### Запуск
`./hasp-exporter` или 
`/usr/bin/local/hasp-exporter`

### Установка как сервис 
sudo cp hasp-exporter.service /etc/systemd/system/
systemctl enable hasp-exporter
systemctl start hasp-exporter

## Результат
http://prometeus-host:8181/metrics
### HELP free_executions Number of license executions remaining
### TYPE free_executions gauge
free_executions{Feature="Simple CODE",Host="10.10.10.10:1947",Product="Simple License 1",Vendor="100000"} 71562
free_executions{Feature="Simple CODE",Host="hasp-server.company:1947",Product="Simple License 1",Vendor="100000"} 71562
### HELP free_sessions Number of free sessions of license to hasp connect
### TYPE free_sessions gauge
free_sessions{Feature="",Host="10.10.10.10:1947",Product="",Vendor=""} -1
free_sessions{Feature="",Host="10.10.10.10:1947",Product="",Vendor="100000"} -1
free_sessions{Feature="",Host="hasp-server.company:1947",Product="",Vendor=""} -1
free_sessions{Feature="",Host="hasp-server.company:1947",Product="",Vendor="100000"} -1
free_sessions{Feature="",Host="hasp-server.company:1947",Product="",Vendor="000000001"} -1
free_sessions{Feature="",Host="hasp-server.company:1947",Product="000000001 Product 20",Vendor="000000001"} -1
free_sessions{Feature="Revert Code",Host="10.10.10.10:1947",Product="Simple License 1",Vendor="100000"} 3
free_sessions{Feature="Test Segment",Host="hasp-server.company:1947",Product="Simple License 1",Vendor="100000"} 3
free_sessions{Feature="LOL_CODE",Host="hasp-server.company:1947",Product="DataLib Lincense 3",Vendor="000000001"} 1
free_sessions{Feature="LOL_CASCADE",Host="hasp-server.company:1947",Product="DataLib Lincense 3",Vendor="000000001"} 1
free_sessions{Feature="LOL_NETWORK",Host="hasp-server.company:1947",Product="DataLib Lincense 3",Vendor="000000001"} 1
free_sessions{Feature="Simple License 1 Core",Host="10.10.10.10:1947",Product="Simple License 1",Vendor="100000"} 5
free_sessions{Feature="Simple License 1 Core",Host="hasp-server.company:1947",Product="Simple License 1",Vendor="100000"} 5
free_sessions{Feature="Simple License 1 Full",Host="10.10.10.10:1947",Product="Simple License 1",Vendor="100000"} 5
free_sessions{Feature="Simple License 1 Full",Host="hasp-server.company:1947",Product="Simple License 1",Vendor="100000"} 5
free_sessions{Feature="ExtensionTime",Host="hasp-server.company:1947",Product="2048",Vendor="000000001"} 15
free_sessions{Feature="ServiceLimit",Host="hasp-server.company:1947",Product="2048",Vendor="000000001"} 15
free_sessions{Feature="Services Long",Host="hasp-server.company:1947",Product="2048",Vendor="000000001"} 15
free_sessions{Feature="Version",Host="hasp-server.company:1947",Product="2048",Vendor="000000001"} 20
free_sessions{Feature="DATA_SDK_TYPE",Host="hasp-server.company:1947",Product=Data SDK",Vendor="000000001"} 3

free_sessions{Feature="data_code",Host="hasp-server.company:1947",Product=Data SDK",Vendor="000000001"} 3
