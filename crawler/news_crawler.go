package crawler

import (
	"strings"
	"sync"

	"github.com/wwgobr/news-crawler/requests"
)

type Crawler interface {
	Start()
	SetUrl(url string)
}

type NewsCrawler struct {
	url string
}

func isImage(url string) bool {
	if strings.HasSuffix(url, ".gif") || strings.HasSuffix(url, ".jpg") ||
		strings.HasSuffix(url, ".jpeg") || strings.HasSuffix(url, ".png") {
		return true
	}
	return false
}

func (nc *NewsCrawler) Start() {
	var wg sync.WaitGroup
	log.Infof("Starting to crawl [%s]", nc.url)
	body, err := requests.GetBody(nc.url)
	if err != nil {
		return
	}
	urls := findUrls(*body)
	for i := 0; i < len(urls); i++ {
		url := urls[i]
		if isImage(url) {
			log.Info(url)
			wg.Add(1)
			go func() {
				defer wg.Done()
				body, err := requests.GetBody(url)
				name := strings.Split(url, "/")
				filename := "imgs/" + name[len(name)-1]
				if err != nil {
					f, err := Open(filename)
					defer f.Close()
					if err != nil {
						log.Errorf("Error opening file %s - %s", filename, err.Error())
					} else if body != nil {
						_, err := f.WriteString(*body)
						if err != nil {
							log.Errorf("Error writing file %s - %s", filename, err.Error())
						}
					}
				}
			}()
		}
	}
	wg.Wait()
}

func (nc *NewsCrawler) SetUrl(url string) {
	log.Infof("Setting URL [%s]", url)
	nc.url = url
}
