package helpers

import (
	"os"
	"strings"
)

type ErrorBlacklist struct {
	blacklist []string
}

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

func (b ErrorBlacklist) IsSerious(err error) bool {
	var serious = true
	for _, s := range b.blacklist {
		serious = serious && !strings.Contains(err.Error(), s)
	}
	return serious
}
