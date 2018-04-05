package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fakeMyIPInfo(t *testing.T) {
	for _, link := range fakeMyIPInfoLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "fakeMyIPInfo crawl", link)
		assert.NotEmpty(t, fakeMyIPInfoIPS(body), "fakeMyIPInfoIPS empty", link)
		checkURI(fakeMyIPInfoIPS(body), t)
	}
}
