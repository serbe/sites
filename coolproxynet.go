package sites

// import (
// 	"bytes"
// 	"fmt"
// 	"strings"

// 	"github.com/PuerkitoBio/goquery"
// )

// func coolProxyNet() []string {
// 	var ips []string
// 	for _, link := range coolProxyNetLinks() {
// 		body, err := crawl(link)
// 		if err != nil {
// 			errmsg("coolProxyNet crawl", err)
// 			continue
// 		}
// 		ips = append(ips, coolProxyNetIPS(body)...)
// 	}
// 	return ips
// }

// func coolProxyNetLinks() []string {
// 	var links []string
// 	for page := 1; page < 4; page++ {
// 		links = append(links, fmt.Sprintf("https://www.cool-proxy.net/proxies/http_proxy_list/sort:score/direction:desc/page:%d", page))
// 	}
// 	return links
// }

// func coolProxyNetIPS(body []byte) []string {
// 	var ips []string
// 	r := bytes.NewReader(body)
// 	dom, err := goquery.NewDocumentFromReader(r)
// 	if err != nil {
// 		errmsg("coolProxyNetIPS NewDocumentFromReader", err)
// 		return ips
// 	}
// 	dom.Find("tbody tr").Each(func(_ int, s *goquery.Selection) {
// 		td := s.Find("td")
// 		ip := td.Eq(0).Text()
// 		ip = strings.Replace(ip, `document.write(Base64.decode(str_rot13("`, "", -1)
// 		ip = strings.Replace(ip, `")))`, "", -1)
// 		ip = strings.Map(rot13, ip)
// 		ip = decodeBase64(ip)
// 		port := td.Eq(1).Text()
// 		if len(port) > 1 {
// 			ips = append(ips, "http://"+ip+":"+port)
// 		}
// 	})
// 	return ips
// }
