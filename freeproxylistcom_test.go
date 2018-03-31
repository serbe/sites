package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_freeProxyListCom(t *testing.T) {
	for _, link := range freeProxyListComLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "freeProxyListCom crawl", link)
		assert.NotEmpty(t, freeProxyListComIPS(body), "freeProxyListComIPS empty", link)
	}
}
