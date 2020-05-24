package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := "xxx"
	s2 := "eidbadffsdgfgsdfgsfdgfgsdgfgsdfgfgsdfgooo"
	// m1 := map[int]int{2: 2, 3: 3}
	// m2 := map[int]int{3: 3, 2: 2}
	// fmt.Println(m1, "--", m2, reflect.DeepEqual(m1, m2))
	// s := "achin"
	// fmt.Printf("%T", s[0:1])
	fmt.Println(checkInclusion(s1, s2))
	// a, b := m1[1

}

func checkInclusion(s1 string, s2 string) bool {

	base_map := create_freq_map(s1)
	for i := 0; i < len(s2)-len(s1)+1; i++ {
		sub_str := s2[i : i+len(s1)]
		sub_map := create_freq_map(sub_str)
		if reflect.DeepEqual(base_map, sub_map) {
			return true
		}
	}
	return false
}

func create_freq_map(str string) map[string]int {
	freq := map[string]int{}
	for _, c := range str {
		if _, ok := freq[string(c)]; !ok {
			freq[string(c)] = 0
		}
		freq[string(c)] += 1
	}
	return freq
}
