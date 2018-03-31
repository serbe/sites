package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_freeProxyList(t *testing.T) {
	for _, link := range freeProxyListLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "freeProxyList crawl", link)
		assert.NotEmpty(t, freeProxyListIPS(body), "freeProxyListIPS empty", link)
	}
}
