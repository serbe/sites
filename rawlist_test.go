package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rawlist(t *testing.T) {
	for _, link := range rawlistLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "rawlist crawl", link)
		assert.NotEmpty(t, ipsFromBytes(body, HTTP), "rawlistIPS empty", link)
	}
}
