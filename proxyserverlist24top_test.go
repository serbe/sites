package sites

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_proxyserverlist24top(t *testing.T) {
	for _, l := range proxyserverlist24topList() {
		body, err := crawl(l)
		assert.NoError(t, err, "proxyserverlist24top crawl", l)
		links := proxyserverlist24topLinks(body)
		assert.NotEmpty(t, links, "proxyserverlist24topLinks empty", l)
		for _, link := range links {
			body, err := crawl(link)
			assert.NoError(t, err, "proxyserverlist24top crawl", link)
			scheme := HTTP
			if strings.Contains(link, "socks") {
				scheme = SOCKS5
			}
			assert.NotEmpty(t, ipsFromBytes(body, scheme), "ipsFromBytes empty", link)
		}
	}
}
