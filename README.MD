# Prometheus metrics

Simple example how to use prometheus/client_golang


To launch

`go build -o app && ./app`

use for create requests (time sleep in milliseconds)

`curl --request GET --url 'http://localhost:8080/?sleep=5' `

use for retrieve metrics
```
curl --request GET --url 'http://localhost:9090/metrics'

curl --request GET --url 'http://localhost:9090/metrics' | grep http_request
```