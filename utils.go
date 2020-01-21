package sites

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// constant schemes
const (
	HTTP   = "http"
	HTTPS  = "https"
	SOCKS4 = "socks4"
	SOCKS5 = "socks5"
)

var (
	useDebug bool
	useError bool
)

func randomUA() string {
	userAgents := []string{
		"Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 7.0; SM-G892A Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/60.0.3112.107 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 7.0; SM-G930VC Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/58.0.3029.83 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 6P Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.83 Mobile Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1",
		"Mozilla/5.0 (Apple-iPhone7C2/1202.466; U; CPU like Mac OS X; en) AppleWebKit/420+ (KHTML, like Gecko) Version/3.0 Mobile/1A543 Safari/419.3",
		"Mozilla/5.0 (Windows Phone 10.0; Android 6.0.1; Microsoft; RM-1152) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Mobile Safari/537.36 Edge/15.15254",
		"Mozilla/5.0 (Linux; Android 7.0; SM-T827R4 Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.116 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
		"Mozilla/5.0 (X11; CrOS x86_64 8172.45.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.64 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/601.3.9 (KHTML, like Gecko) Version/9.0.2 Safari/601.3.9",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.111 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:64.0) Gecko/20100101 Firefox/64.0",
		"Mozilla/5.0 (X11; Linux i686; rv:64.0) Gecko/20100101 Firefox/64.0",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1",
		"Links (2.7; Linux 3.7.9-2-ARCH x86_64; GNU C 4.7.1; text)",
		"Lynx/2.8.8dev.3 libwww-FM/2.14 SSL-MM/1.4.1",
		"Opera/9.80 (X11; Linux i686; Ubuntu/14.10) Presto/2.12.388 Version/12.16",
		"Opera/9.80 (Windows NT 6.0) Presto/2.12.388 Version/12.14",
		"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		"Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
		"Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)",
		"DuckDuckBot/1.0; (+http://duckduckgo.com/duckduckbot.html)",
		"Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
		"Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)",
		"Sogou Pic Spider/3.0( http://www.sogou.com/docs/help/webmasters.htm#07)",
		"Sogou head spider/3.0( http://www.sogou.com/docs/help/webmasters.htm#07)",
		"Sogou web spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)",
		"Sogou Orion spider/3.0( http://www.sogou.com/docs/help/webmasters.htm#07)",
		"Sogou-Test-Spider/4.0 (compatible; MSIE 5.5; Windows 98)",
		"Mozilla/5.0 (compatible; Konqueror/3.5; Linux) KHTML/3.5.5 (like Gecko) (Exabot-Thumbnails)",
		"Mozilla/5.0 (compatible; Exabot/3.0; +http://www.exabot.com/go/robot)",
		"facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)",
		"ia_archiver (+http://www.alexa.com/site/help/webmasters; crawler@alexa.com)",
	}
	return userAgents[rand.Intn(len(userAgents))]
}

func crawl(target string) ([]byte, error) {
	timeout := time.Duration(15000) * time.Millisecond
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", randomUA())
	req.Header.Set("Connection", "close")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Referer", "https://www.google.com/")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	// fname := strings.ReplaceAll(target, "/", "")
	// fname = strings.ReplaceAll(fname, ":", "")
	// fname = strings.ReplaceAll(fname, ".", "")
	// fname = fname + ".html"
	// _ = ioutil.WriteFile(fname, body, 0644)
	return body, err
}

func ipsFromBytes(body []byte, scheme string) []string {
	var ips []string
	reIP := `((?:(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]):\d{2,5})`
	reIPWScheme := `([http|https|socks4|socks5]://(?:(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]):\d{2,5})`
	re := regexp.MustCompile(reIP)
	if scheme == "" {
		re = regexp.MustCompile(reIPWScheme)
	}
	if !re.Match(body) {
		return ips
	}
	results := re.FindAllSubmatch(body, -1)
	for _, res := range results {
		proxy := string(res[1])
		switch scheme {
		case "all":
			ips = append(ips, HTTP+"://"+proxy, HTTPS+"://"+proxy, SOCKS5+"://"+proxy)
		case "":
			ips = append(ips, proxy)
		default:
			ips = append(ips, scheme+"://"+proxy)
		}
	}
	return ips
}

func decodeBase64(src string) string {
	out, _ := base64.StdEncoding.DecodeString(src)
	return string(out)
}

func errmsg(str string, err error) {
	if useError {
		log.Println("Error in", str, err)
	}
}

func debugmsg(str ...interface{}) {
	if useDebug {
		log.Println(str...)
	}
}

// func rot13(b rune) rune {
// 	var a, z rune
// 	switch {
// 	case 'a' <= b && b <= 'z':
// 		a, z = 'a', 'z'
// 	case 'A' <= b && b <= 'Z':
// 		a, z = 'A', 'Z'
// 	default:
// 		return b
// 	}
// 	return (b-a+13)%(z-a+1) + a
// }

func fixURI(ips []string) []string {
	for i, res := range ips {
		split := strings.Split(res, "://")
		scheme := split[0]
		if scheme != HTTP && scheme != HTTPS && scheme != SOCKS4 && scheme != SOCKS5 {
			ips[i] = "http://" + split[1]
		}
	}
	return ips
}
