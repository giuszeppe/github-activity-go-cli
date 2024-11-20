package main

import (
	"fmt"
	"strings"

	"github.com/giuszeppe/github-activity-go-cli/service"
)

func main() {
	out, err := service.GetActivityForUsername("giuszeppe")
	if err != nil {
		return
	}

	fmt.Println("Output:\n", strings.Join(out, "\n"))
}
