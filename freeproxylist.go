package sites

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func freeProxyList() []string {
	var ips []string
	for _, link := range freeProxyListLinks() {
		body, err := crawl(link)
		if err != nil {
			errmsg("freeProxyList crawl", err)
			continue
		}
		ips = append(ips, freeProxyListIPS(body)...)
	}
	return ips
}

func freeProxyListLinks() []string {
	var links = []string{
		"https://www.us-proxy.org/",
		"http://free-proxy-list.net/",
		"https://www.sslproxies.org/",
		"https://www.socks-proxy.net/",
	}
	return links
}

func freeProxyListIPS(body []byte) []string {
	var ips []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("freeProxyListIPS NewDocumentFromReader", err)
		return ips
	}
	dom.Find("tr").Each(func(_ int, s *goquery.Selection) {
		td := s.Find("td")
		if td.Length() == 8 {
			if strings.ToLower(td.Eq(4).Text()) == SOCKS5 {
				ips = append(ips, "socks5://"+td.Eq(0).Text()+":"+td.Eq(1).Text())
			} else if strings.ToLower(td.Eq(4).Text()) == SOCKS4 {
				ips = append(ips, "socks4://"+td.Eq(0).Text()+":"+td.Eq(1).Text())
			} else {
				if td.Eq(6).Text() == "yes" {
					ips = append(ips, "https://"+td.Eq(0).Text()+":"+td.Eq(1).Text())
				} else {
					ips = append(ips, "http://"+td.Eq(0).Text()+":"+td.Eq(1).Text())
				}
			}
		}
	})
	return ips
}
