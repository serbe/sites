package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_freeproxylist(t *testing.T) {
	for _, link := range freeproxylistLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "freeproxylist crawl", link)
		assert.NotEmpty(t, freeproxylistIPS(body), "freeproxylistIPS empty", link)
	}
}
