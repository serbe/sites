package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cyberSyndromeNet(t *testing.T) {
	for _, link := range cyberSyndromeNetLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "cyberSyndromeNet crawl", link)
		assert.NotEmpty(t, cyberSyndromeNetIPS(body), "cyberSyndromeNetIPS empty", link)
		checkURI(cyberSyndromeNetIPS(body), t)
	}
}
