package sites

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func xicidailiCom() []string {
	var ips []string
	for _, link := range xicidailiComLinks() {
		body, err := crawl(link)
		if err != nil {
			errmsg("xicidailiCom crawl", err)
			continue
		}
		ips = append(ips, xicidailiComIPS(body)...)
	}
	return ips
}

func xicidailiComLinks() []string {
	var links []string
	for page := 1; page < 7; page++ {
		links = append(links, fmt.Sprintf("http://www.xicidaili.com/nn/%d", page))
	}
	return links
}

func xicidailiComIPS(body []byte) []string {
	var ips []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("xicidailiComIPS NewDocumentFromReader", err)
		return ips
	}
	dom.Find("table#ip_list tr").Each(func(_ int, s *goquery.Selection) {
		td := s.Find("td")
		ip := strings.ToLower(td.Eq(5).Text()) + "://" + td.Eq(1).Text() + ":" + td.Eq(2).Text()
		if len(ip) > 8 {
			ips = append(ips, ip)
		}
	})
	return ips
}
