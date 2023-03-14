package pkg

import "strings"

func indexOf(data string, find string, index int) int {
	if index == 0 {
		curr_index := strings.Index(data, find)
		return curr_index
	}
	next_split := data[index:len(data)]
	curr_index := strings.Index(next_split, find)
	if curr_index == -1 {
		return -1
	}
	return curr_index + index
}
