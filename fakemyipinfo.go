package sites

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fakeMyIPInfo() []string {
	var ips []string
	for _, link := range fakeMyIPInfoLinks() {
		body, err := crawl(link)
		if err != nil {
			errmsg("fakeMyIPInfo crawl", err)
			continue
		}
		ips = append(ips, fakeMyIPInfoIPS(body)...)
	}
	return ips
}

func fakeMyIPInfoLinks() []string {
	links := []string{
		"http://www.fakemyip.info/elite-proxies",
		"http://www.fakemyip.info/anonymous-proxies",
		"http://www.fakemyip.info/transparent-proxies",
	}

	return links
}

func fakeMyIPInfoIPS(body []byte) []string {
	var ips []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("fakeMyIPInfoIPS NewDocumentFromReader", err)
		return ips
	}
	dom.Find("tr").Each(func(_ int, s *goquery.Selection) {
		td := s.Find("td")
		if td.Length() == 6 {
			ip := td.Eq(0).Text()
			port := strings.TrimLeft(td.Eq(1).Text(), "0")
			ips = append(ips, "http://"+ip+":"+port)
		}
	})
	return ips
}
