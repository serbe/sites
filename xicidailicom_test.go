package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_xicidailiCom(t *testing.T) {
	for _, link := range xicidailiComLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "xicidailiCom crawl", link)
		assert.NotEmpty(t, xicidailiComIPS(body), "xicidailiComIPS empty", link)
		checkURI(xicidailiComIPS(body), t)
		// printIPS(xicidailiComIPS(body))
	}
}
