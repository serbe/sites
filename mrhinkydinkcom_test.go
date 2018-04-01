package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mrHinkyDinkCom(t *testing.T) {
	for _, link := range mrHinkyDinkComLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "mrHinkyDinkCom crawl", link)
		assert.NotEmpty(t, mrHinkyDinkComIPS(body), "mrHinkyDinkComLinks empty", link)
	}
}
