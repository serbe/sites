package sites

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func webanetlabs() []string {
	var ips []string
	for _, l := range webanetlabsList() {
		body, err := crawl(l)
		if err != nil {
			errmsg("webanetlabs crawl", err)
			return ips
		}
		ips = append(ips, webanetlabsIPS(body)...)
		for _, link := range webanetlabsLinks(body) {
			body, err := crawl(link)
			if err != nil {
				errmsg("webanetlabs crawl", err)
				continue
			}
			ips = append(ips, ipsFromBytes(body, HTTP)...)
		}
	}
	return ips
}

func webanetlabsList() []string {
	var list = []string{
		"http://webanetlabs.net/publ/24",
	}
	return list
}

func webanetlabsLinks(body []byte) []string {
	var links []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("webanetlabsLinks NewDocumentFromReader", err)
		return links
	}
	dom.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, ".txt") {
			links = append(links, "http://webanetlabs.net"+href)
		}
	})
	return links
}

func webanetlabsIPS(body []byte) []string {
	var ips []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("webanetlabsIPS NewDocumentFromReader", err)
		return ips
	}
	dom.Find("p").Each(func(_ int, s *goquery.Selection) {
		sp := strings.Split(s.Text(), "\n")
		if len(sp) > 9 {
			for _, ip := range sp {
				if len(ip) > 10 {
					ips = append(ips, "http://"+ip)
				}
			}
		}
	})
	return ips
}
