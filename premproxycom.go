package sites

// import (
// 	"bytes"
// 	"regexp"
// 	"strings"

// 	"github.com/PuerkitoBio/goquery"
// )

// func premProxyCom() []string {
// 	var ips []string
// 	body, err := crawl("https://premproxy.com/list")
// 	if err != nil {
// 		errmsg("premProxyCom crawl", err)
// 		return ips
// 	}
// 	ips = append(ips, premProxyComIPS(body)...)
// 	return ips
// }

// func premProxyComIPS(body []byte) []string {
// 	var ips []string
// 	r := bytes.NewReader(body)
// 	dom, err := goquery.NewDocumentFromReader(r)
// 	if err != nil {
// 		errmsg("premProxyComIPS NewDocumentFromReader", err)
// 		return ips
// 	}
// 	re, err := regexp.Compile(`<script src="(/js/.+?\.js)"></script>`)
// 	if err != nil || !re.Match(body) {
// 		errmsg("premProxyComIPS Match js", err)
// 		return ips
// 	}
// 	uri := re.FindSubmatch(body)
// 	jsBody, err := crawl("https://premproxy.com" + string(uri[1]))
// 	if err != nil {
// 		return ips
// 	}
// 	re, err := regexp.Compile(`<script src="(/js/.+?\.js)"></script>`)
// 	if err != nil || !re.Match(body) {
// 		errmsg("premProxyComIPS Compile", err)
// 		return ips
// 	}
// 	dom.Find("ul").Each(func(_ int, s *goquery.Selection) {
// 		scheme := strings.ToLower(s.Find("li.https").Text())
// 		ip := decodeBase64(s.Find("li.proxy").Text())
// 		if ip != "" && (scheme == HTTP || scheme == HTTPS || scheme == SOCKS5) {
// 			ips = append(ips, scheme+"://"+ip)
// 		}
// 	})
// 	return ips
// }
