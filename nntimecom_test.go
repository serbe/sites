package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nnTimeCom(t *testing.T) {
	for _, link := range nnTimeComLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "nnTimeCom crawl", link)
		assert.NotEmpty(t, nnTimeComIPS(body), "nnTimeComIPS empty", link)
		checkURI(nnTimeComIPS(body), t)
	}
}
