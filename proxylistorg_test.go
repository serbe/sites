package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_proxylistorg(t *testing.T) {
	for _, link := range proxylistorgLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "proxylistorg crawl", link)
		assert.NotEmpty(t, proxylistorgIPS(body), "proxylistorgIPS empty", link)
	}
}
