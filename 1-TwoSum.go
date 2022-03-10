package main

//double loop
// func twoSum(nums []int, target int) []int {
// 	res := []int{}
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				res = append(res, i)
// 				res = append(res, j)
// 				return res
// 			}
// 		}
// 	}
// 	return res
// }

//map
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		residual := target - nums[i]
		if _, ok := m[residual]; ok {
			return []int{m[residual], i}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}

// func main() {
// 	nums := []int{3, 2, 4}
// 	target := 6
// 	res := twoSum(nums, target)
// 	fmt.Println(res)
// }
