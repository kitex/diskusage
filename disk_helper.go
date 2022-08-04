package main

import (
	"os"
	"path/filepath"
	"sort"
)

type Pair struct {
	Key   string
	Value int64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func sortByValue(kvarr map[string]int64) []Pair {

	p := make(PairList, len(kvarr))

	i := 0
	for k, v := range kvarr {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(p))

	return p

}

// du -sh * | sort -rh sorted output
func sortByKey(varr map[string]int64) []string {

	keys := make([]string, 0, len(varr))
	for k := range varr {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func dirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
