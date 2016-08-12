package main

import (
	"math/rand"
	"strings"
)

func GenOborona(s [][]string) string {
	var res []string
	for _, lst := range s {
		res = append(res, lst[rand.Intn(len(lst))])
	}
	return strings.Join(res, " ")
}
