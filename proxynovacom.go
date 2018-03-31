package sites

// func proxyNovaCom() []string {
// 	var ips []string
// 	body, err := crawl("https://www.proxynova.com/proxy-server-list/")
// 	if err != nil {
// 		errmsg("proxyNovaCom crawl", err)
// 		return ips
// 	}
// 	ips = append(ips, proxyNovaComIPS(body)...)
// 	return ips
// }

// func proxyNovaComIPS(body []byte) []string {
// 	var ips []string
// 	r := bytes.NewReader(body)
// 	dom, err := goquery.NewDocumentFromReader(r)
// 	if err != nil {
// 		errmsg("proxyNovaComIPS NewDocumentFromReader", err)
// 		return ips
// 	}
// 	dom.Find("tr").Each(func(_ int, s *goquery.Selection) {
// 		// ip, exist := s.Find("td abbr").Attr("title")
// 		ip := s.Find("td abbr").Text()
// 		port := s.Find("td a").Text()
// 		// fmt.Println(ip, port)
// 		fmt.Println(ip)
// 		// if exist && port != "" {
// 		ips = append(ips, "http://"+ip+":"+port)
// 		// }
// 	})
// 	return ips
// }
