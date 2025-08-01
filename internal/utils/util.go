package utils

import (
	"math/rand"
	"strings"
)

const randomTitle = "test example sample first"

func init() {
	rand.Intn(len(randomTitle))
}

func RandomTitle(n int) string {
	words := strings.Fields(randomTitle)
	k := len(words)

	var result string
	for i := 0; i < n; i++ {
		c := rand.Intn(k)
		result = words[c]
	}
	return result
}
