package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_freeproxylist(t *testing.T) {
	links := freeproxylistLinks()
	for _, link := range links {
		body, err := crawl(link)
		assert.NoError(t, err, "freeproxylist crawl", link)
		assert.NotEmpty(t, freeproxylistIPS(body), "freeproxylistIPS empty", link)
	}
}
