package main

import (
	"fmt"
	"golearning/algorithm"
)

func main() {
	fmt.Println("主程序启动成功......")

	islandGrid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println("岛屿数量(DFS)：", algorithm.NumIslands(copyGrid(islandGrid)))
	fmt.Println("岛屿数量(BFS)：", algorithm.NumIslandsBFS(copyGrid(islandGrid)))

	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("字母异位词分组：", algorithm.GroupAnagrams(words))




	fmt.Println("启动程序结束......")
}

func copyGrid(src [][]byte) [][]byte {
	dst := make([][]byte, len(src))
	for i := range src {
		dst[i] = append([]byte(nil), src[i]...)
	}
	return dst
}
