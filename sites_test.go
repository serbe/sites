package sites

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func printIPS(ips []string) {
	for _, item := range ips {
		fmt.Println(item)
	}
	fmt.Println(len(ips))
}

func checkURI(ips []string, t *testing.T) {
	for _, uri := range ips {
		_, err := url.ParseRequestURI(uri)
		assert.NoError(t, err, uri)
	}
}
