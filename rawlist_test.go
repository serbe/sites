package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rawList(t *testing.T) {
	for _, link := range rawListLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "rawList crawl", link)
		assert.NotEmpty(t, ipsFromBytes(body, HTTP), "rawListIPS empty", link)
		checkURI(ipsFromBytes(body, HTTP), t)
	}
}
