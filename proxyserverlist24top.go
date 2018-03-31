package sites

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func proxyServerList24Top() []string {
	var ips []string
	for _, l := range proxyServerList24TopList() {
		body, err := crawl(l)
		if err != nil {
			errmsg("proxyServerList24Top crawl", err)
			return ips
		}
		for _, link := range proxyServerList24TopLinks(body) {
			body, err := crawl(link)
			if err != nil {
				errmsg("proxyServerList24Top crawl", err)
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

func proxyServerList24TopList() []string {
	list := []string{
		"http://www.proxyserverlist24.top",
		"http://www.live-socks.net/",
	}
	return list
}

func proxyServerList24TopLinks(body []byte) []string {
	var links []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("proxyServerList24TopLinks NewDocumentFromReader", err)
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
