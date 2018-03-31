package sites

func rawlist() []string {
	var ips []string
	for _, link := range rawlistLinks() {
		body, err := crawl(link)
		if err != nil {
			errmsg("rawlist crawl", err)
			continue
		}
		ips = append(ips, ipsFromBytes(body, HTTP)...)
	}
	return ips
}

func rawlistLinks() []string {
	links := []string{
		"https://www.rmccurdy.com/scripts/proxy/good.txt",
		"http://www.proxylists.net/http_highanon.txt",
		"http://ab57.ru/downloads/proxylist.txt",
		"http://multiproxy.org/txt_all/proxy.txt",
	}
	return links
}
