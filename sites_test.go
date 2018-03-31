package sites

import (
	"fmt"
)

func printIPS(ips []string) {
	for _, item := range ips {
		fmt.Println(item)
	}
	fmt.Println(len(ips))
}
