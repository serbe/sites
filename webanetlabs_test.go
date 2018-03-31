package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_webanetlabs(t *testing.T) {
	for _, l := range webanetlabsList() {
		body, err := crawl(l)
		assert.NoError(t, err, "webanetlabs crawl", l)
		assert.NotEmpty(t, webanetlabsIPS(body), "webanetlabsIPS empty", l)
		links := webanetlabsLinks(body)
		assert.NotEmpty(t, links, "webanetlabsLinks empty", links)
		for _, link := range links {
			body, err := crawl(link)
			assert.NoError(t, err, "webanetlabs crawl", link)
			assert.NotEmpty(t, webanetlabsIPS(body), "webanetlabsIPS empty", link)
		}
	}
}
