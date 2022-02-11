# Prometheus metrics

Simple example hot wo use prometheus/client_golang


To launch

`go build -o app prom_metrics_example && ./app`

use for create requests (time sleep in milliseconds)

`curl -i http://localhost:8080/?sleep=5`

use for retrieve metrics
```
curl -i http://localhost:9090/metrics

curl -i http://localhost:9090/metrics | grep http_request
```