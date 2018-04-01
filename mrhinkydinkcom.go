package sites

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

func mrHinkyDinkCom() []string {
	var ips []string
	for _, link := range mrHinkyDinkComLinks() {
		body, err := crawl(link)
		if err != nil {
			errmsg("mrHinkyDinkCom crawl", err)
			continue
		}
		ips = append(ips, mrHinkyDinkComIPS(body)...)
	}
	return ips
}

func mrHinkyDinkComLinks() []string {
	links := []string{
		"http://www.mrhinkydink.com/proxies.htm",
	}
	return links
}

func mrHinkyDinkComIPS(body []byte) []string {
	var ips []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("mrHinkyDinkComIPS NewDocumentFromReader", err)
		return ips
	}
	dom.Find("tr.text").Each(func(_ int, s *goquery.Selection) {
		td := s.Find("td")
		if td.Length() == 8 {
			ips = append(ips, "http://"+td.Eq(0).Text()+":"+td.Eq(1).Text())
		}
	})
	return ips
}
