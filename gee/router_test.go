package gee

import "testing"

//单元测试 xxx_test.go
func TwoSum(nums []int, target int) []int {
 m := make(map[int]int)
 for i, num := range nums {
   key := target - num
   if j, ok := m[key]; ok {
     return []int{j, i}
   }
   m[nums[i]] = i
 }
 return []int{}
}

func TestTwoSum(t *testing.T) {
  t.Log(TwoSum([]int{2, 7, 11, 15}, 9))
}
//go test two_sum_test.go two_sum.go
