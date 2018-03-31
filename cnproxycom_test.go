package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cnProxyCom(t *testing.T) {
	for _, link := range cnProxyComLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "cnProxyCom crawl", link)
		assert.NotEmpty(t, cnProxyComIPS(body), "cnProxyComIPS empty", link)
	}
}
