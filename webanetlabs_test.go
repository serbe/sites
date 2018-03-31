package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_webanetLabs(t *testing.T) {
	for _, l := range webanetLabsList() {
		body, err := crawl(l)
		assert.NoError(t, err, "webanetLabs crawl", l)
		assert.NotEmpty(t, webanetLabsIPS(body), "webanetLabsIPS empty", l)
		links := webanetLabsLinks(body)
		assert.NotEmpty(t, links, "webanetLabsLinks empty", links)
		for _, link := range links {
			body, err := crawl(link)
			assert.NoError(t, err, "webanetLabs crawl", link)
			assert.NotEmpty(t, webanetLabsIPS(body), "webanetLabsIPS empty", link)
		}
	}
}
