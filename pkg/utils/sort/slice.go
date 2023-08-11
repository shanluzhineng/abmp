// Package sort provides utility that sort slice by length
package sort

import (
	"sort"
)

type byLen []string

// get slice length
func (a byLen) Len() int {
	return len(a)
}

// Less check which element is less
func (a byLen) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

// Swap swap elements
func (a byLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ByLen sort by length
func ByLen(s []string) {

	sort.Sort(byLen(s))

}
