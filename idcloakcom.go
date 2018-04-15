package sites

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

func idcloakCom() []string {
	var ips []string
	for _, link := range idcloakComLinks() {
		body, err := crawl(link)
		if err != nil {
			errmsg("idcloakCom crawl", err)
			continue
		}
		ips = append(ips, idcloakComIPS(body)...)
	}
	return ips
}

func idcloakComLinks() []string {
	links := []string{
		"http://www.idcloak.com/proxylist/free-proxy-ip-list.html",
	}
	return links
}

func idcloakComIPS(body []byte) []string {
	var ips []string
	r := bytes.NewReader(body)
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		errmsg("idcloakComIPS NewDocumentFromReader", err)
		return ips
	}
	dom.Find("tr").Each(func(_ int, s *goquery.Selection) {
		td := s.Find("td")
		if td.Length() == 8 {
			ips = append(ips, td.Eq(5).Text()+"://"+td.Eq(7).Text()+":"+td.Eq(6).Text())
		}
	})
	return ips
}
