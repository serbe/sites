package sites

import (
	"regexp"
	"strconv"
	"strings"
)

func cyberSyndromeNet() []string {
	var ips []string
	for _, link := range cyberSyndromeNetLinks() {
		body, err := crawl(link)
		if err != nil {
			errmsg("cyberSyndromeNet crawl", err)
			continue
		}
		ips = append(ips, cyberSyndromeNetIPS(body)...)
	}
	return ips
}

func cyberSyndromeNetLinks() []string {
	var links = []string{
		"http://www.cybersyndrome.net/pla6.html",
		"http://www.cybersyndrome.net/pld6.html",
	}
	return links
}

func cyberSyndromeNetIPS(body []byte) []string {
	var ips []string
	reAS := regexp.MustCompile(`var as=\[([\d,]+?)\]`)
	rePS := regexp.MustCompile(`var ps=\[([\d,]+?)\]`)
	reN := regexp.MustCompile(`var n=(\(.+?);`)
	if !reAS.Match(body) || !rePS.Match(body) || !reN.Match(body) {
		return ips
	}
	as := strings.Split(string(reAS.FindSubmatch(body)[1]), ",")
	ps := strings.Split(string(rePS.FindSubmatch(body)[1]), ",")
	n := string(reN.FindSubmatch(body)[1])
	rePSNum := regexp.MustCompile(`(ps\[\d+\])`)
	fPS := rePSNum.FindAllString(n, -1)
	for _, item := range fPS {
		n = strings.Replace(n, item, ps2s(ps, item), 1)
	}
	res, _ := strconv.Atoi(calc(n))
	as = rotate(as, res)
	for i, port := range ps {
		ips = append(ips, "http://"+as[i*4]+"."+as[i*4+1]+"."+as[i*4+2]+"."+as[i*4+3]+":"+port)
	}
	return ips
}
