package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kuaidaili(t *testing.T) {
	for _, link := range kuaidailiLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "kuaidaili crawl", link)
		assert.NotEmpty(t, kuaidailiIPS(body), "kuaidailiIPS empty", link)
		checkURI(kuaidailiIPS(body), t)
		// printIPS(kuaidailiIPS(body))
	}
}
