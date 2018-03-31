package sites

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func webanetLabs() []string {
	var ips []string
	for _, l := range webanetLabsList() {
		body, err := crawl(l)
		if err != nil {
			errmsg("webanetLabs crawl", err)
			return ips
		}
		ips = append(ips, webanetLabsIPS(body)...)
		for _, link := range webanetLabsLinks(body) {
			body, err := crawl(link)
			if err != nil {
				errmsg("webanetLabs crawl", err)
				continue
			}
			ips = append(ips, ipsFromBytes(body, HTTP)...)
		}
	}
	return ips
}

func webanetLabsList() []string {
	var list = []string{
		"http://webanetLabs.net/publ/24",
	}
	return list
}

func webanetLabsLinks(body []byte) []string {
	var links []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("webanetLabsLinks NewDocumentFromReader", err)
		return links
	}
	dom.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.Contains(href, ".txt") {
			links = append(links, "http://webanetLabs.net"+href)
		}
	})
	return links
}

func webanetLabsIPS(body []byte) []string {
	var ips []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("webanetLabsIPS NewDocumentFromReader", err)
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
