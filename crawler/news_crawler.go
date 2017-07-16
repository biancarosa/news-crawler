package crawler

type Crawler interface {
	Start()
	SetUrl(url string)
}

type NewsCrawler struct {
	url string
}

func (nc *NewsCrawler) Start() {
	log.Infof("Starting to crawl [%s]", nc.url)
}

func (nc *NewsCrawler) SetUrl(url string) {
	log.Infof("Setting URL [%s]", url)
	nc.url = url
}
