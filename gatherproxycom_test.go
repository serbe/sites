package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gatherproxycom(t *testing.T) {
	for _, link := range gatherproxycomLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "gatherproxycom crawl", link)
		assert.NotEmpty(t, gatherproxycomIPS(body), "gatherproxycomIPS empty", link)
	}
}
