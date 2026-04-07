//go:build ignore

package main

import (
	"fmt"
	"sort"
)

/*3.合并区间 输入： [[1,3],[2,6],[8,10],[15,18]] 输出：[[1,6],[8,10],[15,18]] 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].*/

// Interval 定义区间类型
type Interval struct {
	Start int
	End   int
}

// merge 合并重叠区间
func Merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	// 按区间起始值排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	res := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := &res[len(res)-1]
		curr := intervals[i]
		if curr.Start <= last.End {
			// 重叠，合并区间
			if curr.End > last.End {
				last.End = curr.End
			}
		} else {
			// 不重叠，加入结果
			res = append(res, curr)
		}
	}
	return res
}

func main() {
	intervals := []Interval{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	merged := Merge(intervals)
	fmt.Printf("%+v\n", merged) // 输出: [{1 6} {8 10} {15 18}]
}
