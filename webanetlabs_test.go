package sites

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_webanetLabs(t *testing.T) {
	for _, l := range webanetLabsList() {
		body, err := crawl(l)
		assert.NoError(t, err, "webanetLabs crawl", l)
		assert.NotEmpty(t, webanetLabsIPS(body), "webanetLabsIPS empty", l)
		assert.NotEmpty(t, webanetLabsLinks(body), "webanetLabsLinks empty")
		for _, link := range webanetLabsLinks(body) {
			body, err := crawl(link)
			assert.NoError(t, err, "webanetLabs crawl", link)
			scheme := HTTP
			if strings.Contains(link, "socks") {
				scheme = SOCKS5
			}
			assert.NotEmpty(t, ipsFromBytes(body, scheme), "ipsFromBytes empty", link)
			checkURI(ipsFromBytes(body, scheme), t)
		}
	}
}
