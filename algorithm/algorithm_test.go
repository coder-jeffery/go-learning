package algorithm

import (
	"reflect"
	"sort"
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	if got := NumIslands(copyGrid(grid)); got != 3 {
		t.Fatalf("NumIslands() = %d, want 3", got)
	}

	if got := NumIslandsBFS(copyGrid(grid)); got != 3 {
		t.Fatalf("NumIslandsBFS() = %d, want 3", got)
	}
}

func TestGroupAnagrams(t *testing.T) {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	got := normalizeGroups(GroupAnagrams(words))
	want := normalizeGroups([][]string{
		{"eat", "tea", "ate"},
		{"tan", "nat"},
		{"bat"},
	})

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("GroupAnagrams() = %v, want %v", got, want)
	}
}

func copyGrid(src [][]byte) [][]byte {
	dst := make([][]byte, len(src))
	for i := range src {
		dst[i] = append([]byte(nil), src[i]...)
	}
	return dst
}

func normalizeGroups(groups [][]string) [][]string {
	normalized := make([][]string, len(groups))
	for i := range groups {
		normalized[i] = append([]string(nil), groups[i]...)
		sort.Strings(normalized[i])
	}

	sort.Slice(normalized, func(i, j int) bool {
		if len(normalized[i]) == 0 || len(normalized[j]) == 0 {
			return len(normalized[i]) < len(normalized[j])
		}
		return normalized[i][0] < normalized[j][0]
	})

	return normalized
}
