package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_proxyListOrg(t *testing.T) {
	for _, link := range proxyListOrgLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "proxyListOrg crawl", link)
		assert.NotEmpty(t, proxyListOrgIPS(body), "proxyListOrgIPS empty", link)
		checkURI(proxyListOrgIPS(body), t)
	}
}
