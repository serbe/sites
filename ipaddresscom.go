package sites

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

func ipaddressCom() []string {
	var ips []string
	for _, link := range ipaddressComLinks() {
		body, err := crawl(link)
		if err != nil {
			errmsg("ipaddressCom crawl", err)
			continue
		}
		ips = append(ips, ipaddressComIPS(body)...)
	}
	return ips
}

func ipaddressComLinks() []string {
	links := []string{
		"https://www.ipaddress.com/proxy-list/",
	}
	return links
}

func ipaddressComIPS(body []byte) []string {
	var ips []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("ipaddressComIPS NewDocumentFromReader", err)
		return ips
	}
	dom.Find("tr").Each(func(_ int, s *goquery.Selection) {
		td := s.Find("td")
		if td.Length() == 4 {
			ips = append(ips, "http://"+td.Eq(0).Text())
		}
	})
	return ips
}
