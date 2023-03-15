package main

import (
	"strings"

	"github.com/AJRDRGZ/module-examples/slices"
	"rsc.io/quote"
)

func main() {
	list := []string{"EDteam", "gophers", "golang", quote.Hello()}

	slices.Filter(list, func(item string) bool {
		return strings.HasPrefix(strings.ToLower(item), "h")
	})
}
