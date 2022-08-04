package main

import (
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

	//p is sorted

	result := make(map[string]int64)

	for _, k := range p {
		result[k.Key] = k.Value
		//fmt.Printf("%v\t%v\n", k.Key, k.Value)
	}

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
