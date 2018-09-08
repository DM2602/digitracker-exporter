#!/bin/bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o digitracker-exporter .
sudo docker build . -t ckevi/digitracker-exporter:latest
sudo docker tag ckevi/digitracker-exporter:latest ckevi/digitracker-exporter:0.9-beta
sudo docker push ckevi/digitracker-exporter:latest
sudo docker push ckevi/digitracker-exporter:0.9-beta