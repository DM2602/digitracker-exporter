package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func priceHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	target := params.Get("target")
	if target == "" {
		http.Error(w, "Target parameter is missing", http.StatusBadRequest)
		return
	}

	productPrice := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "product_price",
		Help: "Displays the current price of the product",
	}, []string{"product"})

	registry := prometheus.NewRegistry()
	registry.MustRegister(productPrice)

	productPriceAmount, productTitle := getProduct(target)
	productPrice.WithLabelValues(productTitle).Set(productPriceAmount)

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

func getProduct(url string) (float64, string) {
	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	var replacer = strings.NewReplacer(".–", "")
	var removeNewline = strings.NewReplacer(" - digitec", "", "\n", "")
	doc := soup.HTMLParse(resp)
	price := strings.Trim(doc.Find("div", "class", "product-price").Text(), " ")
	title := removeNewline.Replace(doc.Find("title").Text())
	fixedPrice := replacer.Replace(price)
	result, err := strconv.ParseFloat(fixedPrice, 64)
	if err != nil {
		return 0, "N/A"
	}
	return result, title
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html>
    <head><title>Digitracker Exporter</title></head>
    <body>
	<h1>Digitracker Exporter</h1>
	<h5>Version: 0.9-beta-2</h5>
	<p>An exporter for prometheus to track the prices of products on digitec.ch</p>
    <p><a href="/probe?target=https://www.digitec.ch/de/s1/product/intel-core-i5-8600k-lga-1151-360ghz-unlocked-prozessor-6448577?tagIds=76">Probe the price of a digitec product</a></p>
    <p><a href="/metrics">Metrics</a></p></body></html>`))
	})

	http.HandleFunc("/probe", func(w http.ResponseWriter, r *http.Request) {
		priceHandler(w, r)
	})
	log.Fatal(http.ListenAndServe(":7979", nil))
}
