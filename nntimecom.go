package sites

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func nnTimeCom() []string {
	var ips []string
	for _, link := range nnTimeComLinks() {
		body, err := crawl(link)
		if err != nil {
			errmsg("nnTimeCom crawl", err)
			continue
		}
		ips = append(ips, nnTimeComIPS(body)...)
	}
	return ips
}

func nnTimeComLinks() []string {
	var links []string
	for page := 1; page < 6; page++ {
		links = append(links, fmt.Sprintf("http://nntime.com/proxy-updated-0%d.htm", page))
	}
	return links
}

func nnTimeComIPS(body []byte) []string {
	var ips []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("nnTimeComIPS NewDocumentFromReader", err)
		return ips
	}
	dom.Find("table#proxylist.data tbody tr").Each(func(_ int, s *goquery.Selection) {
		ip := "http://" + s.Find("td").Eq(1).Text()
		if len(ip) > 8 {
			ips = append(ips, ip)
		}
	})
	return ips
}
