package sites

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func proxyListOrg() []string {
	var ips []string
	for _, link := range proxyListOrgLinks() {
		body, err := crawl(link)
		if err != nil {
			errmsg("proxyListOrg crawl", err)
			continue
		}
		ips = append(ips, proxyListOrgIPS(body)...)
	}
	return ips
}

func proxyListOrgIPS(body []byte) []string {
	var ips []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("proxyListOrg NewDocumentFromReader", err)
		return ips
	}
	dom.Find("ul").Each(func(_ int, s *goquery.Selection) {
		scheme := strings.ToLower(s.Find("li.https").Text())
		ip := s.Find("li.proxy").Text()
		ip = strings.Replace(ip, "Proxy('", "", -1)
		ip = strings.Replace(ip, "')", "", -1)
		ip = decodeBase64(ip)
		if ip != "" && (scheme == HTTP || scheme == HTTPS || scheme == SOCKS5) {
			ips = append(ips, scheme+"://"+ip)
		}
	})
	return ips
}

func proxyListOrgLinks() []string {
	var links []string
	for page := 1; page < 5; page++ {
		links = append(links, fmt.Sprintf("https://proxy-list.org/english/index.php?p=%d", page))
	}
	links = append(links, `https://proxy-list.org/english/search.php?search=anonymous-and-elite&country=any&type=anonymous-and-elite&port=any&ssl=any`)
	return links
}
