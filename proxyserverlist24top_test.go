package sites

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_proxyServerList24Top(t *testing.T) {
	for _, l := range proxyServerList24TopList() {
		body, err := crawl(l)
		assert.NoError(t, err, "proxyServerList24Top crawl", l)
		links := proxyServerList24TopLinks(body)
		assert.NotEmpty(t, links, "proxyServerList24TopLinks empty", l)
		for _, link := range links {
			body, err := crawl(link)
			assert.NoError(t, err, "proxyServerList24Top crawl", link)
			scheme := HTTP
			if strings.Contains(link, "socks") {
				scheme = SOCKS5
			}
			assert.NotEmpty(t, ipsFromBytes(body, scheme), "ipsFromBytes empty", link)
			checkURI(ipsFromBytes(body, scheme), t)
		}
	}
}
