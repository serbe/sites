package sites

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

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
	reS := regexp.MustCompile(`((?:\w=\d;)+)`)
	if !reS.Match(body) {
		return ips
	}
	sub := strings.Split(string(reS.FindSubmatch(body)[1]), ";")
	hm := make(map[string]string)
	for _, s := range sub {
		if s != "" {
			v := strings.Split(s, "=")
			hm[v[0]] = v[1]
		}
	}
	dom.Find("table#proxylist.data tbody tr").Each(func(_ int, s *goquery.Selection) {
		ip := s.Find("td").Eq(1).Text()
		ip = strings.Replace(ip, `document.write(":"`, ":", -1)
		ip = strings.Replace(ip, ")", "", -1)
		ip = strings.Replace(ip, "+", "", -1)
		for k, v := range hm {
			ip = strings.Replace(ip, k, v, -1)
		}
		ip = strings.Replace(ip, "+", "", -1)
		ip = "http://" + ip
		if len(ip) > 8 {
			ips = append(ips, ip)
		}
	})
	return ips
}
