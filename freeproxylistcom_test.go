package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_freeproxylistcom(t *testing.T) {
	for _, link := range freeproxylistcomLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "freeproxylistcom crawl", link)
		assert.NotEmpty(t, freeproxylistcomIPS(body), "freeproxylistcomIPS empty", link)
	}
}
