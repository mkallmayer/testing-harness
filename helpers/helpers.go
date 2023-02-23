package helpers

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type ErrorBlacklist struct {
	Blacklist []string
}

func loadBaseCorpus(files []string) ([][]byte, error) {
	corpus := make([][]byte, len(files))
	for i, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			corpus[i] = b
		} else {
			fmt.Printf("Error while reading corpus file %s: %v", f, err)
			return corpus, err
		}
	}
	return corpus, nil
}

func addToCorpus(f *testing.F, corpus [][]byte) {
	for _, b := range corpus {
		f.Add(b)
	}
}

func BaseCorpus(f *testing.F, files []string) {
	corpus, err := loadBaseCorpus(files)
	if err != nil {
		fmt.Printf("Proceeding with empty initial corpus.")
		return
	}
	addToCorpus(f, corpus)
}

func (b ErrorBlacklist) IsSerious(err error) bool {
	var serious = true
	for _, s := range b.Blacklist {
		serious = serious && !strings.Contains(err.Error(), s)
	}
	return serious
}
