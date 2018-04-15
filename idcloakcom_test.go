package sites

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_idcloakCom(t *testing.T) {
	for _, link := range idcloakComLinks() {
		body, err := crawl(link)
		assert.NoError(t, err, "idcloakCom crawl", link)
		assert.NotEmpty(t, idcloakComIPS(body), "idcloakComIPS empty", link)
		checkURI(idcloakComIPS(body), t)
	}
}
