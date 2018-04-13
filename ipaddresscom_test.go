package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ipaddressCom(t *testing.T) {
	for _, link := range ipaddressComLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "ipaddressCom crawl", link)
		assert.NotEmpty(t, ipaddressComIPS(body), "ipaddressComIPS empty", link)
		checkURI(ipaddressComIPS(body), t)
	}
}
