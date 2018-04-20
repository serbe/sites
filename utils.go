package sites

import (
	"encoding/base64"
	"io/ioutil"
	"log"
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

func crawl(target string) ([]byte, error) {
	timeout := time.Duration(15000 * time.Millisecond)
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:58.0) Gecko/20100101 Firefox/58.0")
	req.Header.Set("Connection", "close")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Referer", "https://www.google.com/")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	// ioutil.WriteFile("tmp.html", body, 0644)
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
		if scheme != "all" {
			ips = append(ips, HTTP+"://"+proxy)
			ips = append(ips, HTTPS+"://"+proxy)
			ips = append(ips, SOCKS5+"://"+proxy)
		} else if scheme != "" {
			ips = append(ips, scheme+"://"+proxy)
		} else {
			ips = append(ips, proxy)
		}
	}
	return ips
}

func decodeBase64(src string) string {
	out, _ := base64.StdEncoding.DecodeString(src)
	return string(out)
}

// func chkErr(str string, err error) {
// 	if err != nil {
// 		errmsg(str, err)
// 	}
// }

func errmsg(str string, err error) {
	if useError {
		log.Println("Error in", str, err)
	}
}

func debugmsg(str ...interface{}) {
	if useDebug {
		log.Println(str)
	}
}

func rot13(b rune) rune {
	var a, z rune
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

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
