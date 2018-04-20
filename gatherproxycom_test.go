package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gatherProxyCom(t *testing.T) {
	for _, link := range gatherProxyComLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "gatherProxyCom crawl", link)
		assert.NotEmpty(t, gatherProxyComIPS(body), "gatherProxyComIPS empty", link)
		checkURI(gatherProxyComIPS(body), t)
		// printIPS(gatherProxyComIPS(body))
	}
}
