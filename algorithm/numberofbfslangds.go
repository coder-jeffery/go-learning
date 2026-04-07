package algorithm

//岛屿数量问题

/***

解法 2：广度优先搜索 BFS（队列实现）
核心思路
遇到陆地，岛屿数量 +1
用队列存储陆地坐标
循环出队，淹掉陆地，并把上下左右的陆地入队
直到队列为空，完成一个岛屿的遍历


*/
// 定义坐标结构体
type Point struct {
	x, y int
}

func NumIslandsBFS(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0
	// 上下左右四个方向
	dirs := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				// 初始化队列
				queue := []Point{{i, j}}
				grid[i][j] = '0' // 淹掉

				// BFS 循环
				for len(queue) > 0 {
					// 出队
					cur := queue[0]
					queue = queue[1:]

					// 遍历四个方向
					for _, d := range dirs {
						nx, ny := cur.x+d.x, cur.y+d.y
						if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == '1' {
							grid[nx][ny] = '0'
							queue = append(queue, Point{nx, ny})
						}
					}
				}
			}
		}
	}
	return count
}
