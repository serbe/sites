package sites

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

func cnProxyCom() []string {
	var ips []string
	for _, link := range cnProxyComLinks() {
		body, err := crawl(link)
		if err != nil {
			errmsg("cnProxyCom crawl", err)
			continue
		}
		ips = append(ips, cnProxyComIPS(body)...)
	}
	return ips
}

func cnProxyComLinks() []string {
	var links = []string{
		"http://cn-proxy.com/",
	}
	return links
}

func cnProxyComIPS(body []byte) []string {
	var ips []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("cnProxyCom NewDocumentFromReader", err)
		return ips
	}
	dom.Find("tbody tr").Each(func(_ int, s *goquery.Selection) {
		td := s.Find("td")
		ips = append(ips, "http://"+td.Eq(0).Text()+":"+td.Eq(1).Text())
	})
	return ips
}
