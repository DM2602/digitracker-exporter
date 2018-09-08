# digitracker-exporter

## Introduction

This project is a Prometheus exporter to provide TSDB data of digitec.ch products. It took some inspiration of the already existing blackbox exporter made by the Prometheus team.

## Usage

To get the data in Prometheus you must configure the job in the prometheus.yml file as follows:

```yml
   - job_name: 'digitracker'
     metrics_path: /probe
     static_configs:
     - targets:
       - https://www.digitec.ch/de/s1/product/hyperx-fury-2x-8gb-ddr4-2133-dimm-288-arbeitsspeicher-5724752
       - https://www.digitec.ch/de/s1/product/noctua-nh-u12s-1580cm-cpu-kuehler-2479413?tagIds=76-526
     relabel_configs:
       - source_labels: [__address__]
         target_label: __param_target
       - source_labels: [__param_target]
         target_label: instance
       - target_label: __address__
         replacement: 127.0.0.1:7979  # The digitracker-exporter's real hostname:port.
```

It is recommended to set the scrape interval as high as possbile since the prices don't change every second. This is only to prevent the data piling up on your disk.