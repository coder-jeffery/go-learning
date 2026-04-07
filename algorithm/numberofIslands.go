package algorithm

//岛屿数量问题


/**

解法 1：深度优先搜索 DFS（推荐，代码极简）
核心思路
遍历网格每一个格子
遇到陆地 '1'，岛屿数量 +1
立即把当前陆地 + 所有相邻陆地全部淹掉（改为 '0'），避免重复统计
递归 / 循环处理上下左右四个方向


*/


func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)    // 网格行数
	cols := len(grid[0]) // 网格列数
	count := 0           // 岛屿数量

	// 遍历所有格子
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 发现陆地
			if grid[i][j] == '1' {
				count++
				// DFS 淹掉整个岛屿
				Dfs(grid, i, j, rows, cols)
			}
		}
	}
	return count
}

// DFS 递归函数：把当前及相邻的陆地全部置为0
func Dfs(grid [][]byte, i, j, rows, cols int) {
	// 越界 或 不是陆地，直接返回
	if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] != '1' {
		return
	}

	// 淹掉当前陆地（标记为已访问）
	grid[i][j] = '0'

	// 上下左右四个方向递归
	Dfs(grid, i+1, j, rows, cols) // 下
	Dfs(grid, i-1, j, rows, cols) // 上
	Dfs(grid, i, j+1, rows, cols) // 右
	Dfs(grid, i, j-1, rows, cols) // 左
}

//func main() {
//	// 测试用例
//	grid := [][]byte{
//		{'1', '1', '0', '0', '0'},
//		{'1', '1', '0', '0', '0'},
//		{'0', '0', '1', '0', '0'},
//		{'0', '0', '0', '1', '1'},
//	}
//	fmt.Println("岛屿数量：", numIslands(grid)) // 输出：3
//}
