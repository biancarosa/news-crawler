package main

import (
	"os"

	logging "github.com/op/go-logging"
	"github.com/wwgobr/news-crawler/crawler"
)

func init() {
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

func main() {
	var url = os.Getenv("URL")
	nc := new(crawler.NewsCrawler)
	nc.SetUrl(url)
	nc.Start()
}
