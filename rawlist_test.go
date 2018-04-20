package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rawList(t *testing.T) {
	for _, link := range rawListLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "rawList crawl", link)
		assert.NotEmpty(t, ipsFromBytes(body, "all"), "rawListIPS empty", link)
		checkURI(ipsFromBytes(body, "all"), t)
		// printIPS(ipsFromBytes(body, "all"))
	}
}
