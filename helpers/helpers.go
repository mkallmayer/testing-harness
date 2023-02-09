package helpers

import (
	"fmt"
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
			return nil, fmt.Errorf("Could not read %s: %s", f, err.Error())
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
