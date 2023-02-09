package helpers

import (
	"os"
	"strings"
)

func BaseCorpus(files []string) ([][]byte, error) {
	corpus := make([][]byte, len(files))
	for i, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			corpus[i] = b
		} else {
			corpus[i] = make([]byte, 1)
		}
	}
	return corpus, nil
}

func IsSerious(err error, blacklist []string) bool {
	var serious bool = true
	for _, b := range blacklist {
		serious = serious && !strings.Contains(err.Error(), b)
	}
	return serious
}
