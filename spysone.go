package sites

// import (
// 	"bytes"
// 	"fmt"

// 	"github.com/PuerkitoBio/goquery"
// )

// func spySone() []string {
// 	var ips []string
// 	for _, link := range spySoneLinks() {
// 		body, err := crawl(link)
// 		if err != nil {
// 			errmsg("spySone crawl", err)
// 			continue
// 		}
// 		ips = append(ips, spySoneIPS(body)...)
// 	}
// 	return ips
// }

// func spySoneLinks() []string {
// 	var links = []string{
// 		"http://spys.one/free-proxy-list/CN/",
// 	}
// 	return links
// }

// func spySoneIPS(body []byte) []string {
// 	var ips []string
// 	r := bytes.NewReader(body)
// 	dom, err := goquery.NewDocumentFromReader(r)
// 	if err != nil {
// 		errmsg("usProxy NewDocumentFromReader", err)
// 		return ips
// 	}
// 	dom.Find("tr").Each(func(_ int, s *goquery.Selection) {
// 		td := s.Find("td")
// 		if td.Length() == 9 {
// 			// if strings.ToLower(td.Eq(4).Text()) == SOCKS5 {
// 			// 	ips = append(ips, "socks5://"+td.Eq(0).Text()+":"+td.Eq(1).Text())
// 			// } else if strings.ToLower(td.Eq(4).Text()) != "socks4" {
// 			// 	if td.Eq(6).Text() == "yes" {
// 			// 		ips = append(ips, "https://"+td.Eq(0).Text()+":"+td.Eq(1).Text())
// 			// 	} else {
// 			// 		ips = append(ips, "http://"+td.Eq(0).Text()+":"+td.Eq(1).Text())
// 			// 	}
// 			// }
// 			fmt.Println(td.Eq(0).Text())
// 			fmt.Println(td.Eq(1).Text())
// 			fmt.Println(td.Eq(2).Text())
// 		}
// 	})
// 	return ips
// }
