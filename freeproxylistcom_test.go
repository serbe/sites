package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_freeproxylistcom(t *testing.T) {
	links := freeproxylistcomLinks()
	for _, link := range links {
		body, err := crawl(link)
		assert.NoError(t, err, "freeproxylistcom crawl", link)
		assert.NotEmpty(t, freeproxylistcomIPS(body), "freeproxylistcomIPS empty", link)
	}
}
