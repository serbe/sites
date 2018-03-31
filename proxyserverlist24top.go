package sites

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func proxyserverlist24top() []string {
	var ips []string
	for _, l := range proxyserverlist24topList() {
		body, err := crawl(l)
		if err != nil {
			errmsg("proxyserverlist24top crawl", err)
			return ips
		}
		for _, link := range proxyserverlist24topLinks(body) {
			body, err := crawl(link)
			if err != nil {
				errmsg("proxyserverlist24top crawl", err)
				continue
			}
			scheme := HTTP
			if strings.Contains(link, "socks") {
				scheme = SOCKS5
			}
			ips = append(ips, ipsFromBytes(body, scheme)...)
		}
	}
	return ips
}

func proxyserverlist24topList() []string {
	list := []string{
		"http://www.proxyserverlist24.top",
		"http://www.live-socks.net/",
	}
	return list
}

func proxyserverlist24topLinks(body []byte) []string {
	var links []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("proxyserverlist24topLinks NewDocumentFromReader", err)
		return links
	}
	dom.Find("div.jump-link").Each(func(_ int, s *goquery.Selection) {
		href, exist := s.Find("a").Attr("href")
		if exist {
			links = append(links, href)
		}
	})
	return links
}
