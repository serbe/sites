package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_coolProxyNetCom(t *testing.T) {
	for _, link := range coolProxyNetComLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "coolProxyNetCom crawl", link)
		assert.NotEmpty(t, coolProxyNetComIPS(body), "coolProxyNetComIPS empty", link)
	}
}
