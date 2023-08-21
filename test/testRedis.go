package main

import "fmt"

func main() {
	//utils.InitConfig()
	//utils.InitRedisClient()
	//username := utils.HGet("stu-1", "username")
	//fmt.Println(username)

	//nums := []int{1, 2, 3, 4}
	//
	//fmt.Println(nums[:0])
	//print()
	//numberOfSteps(14)
	//t1 := "2021-01-01 00:00:00"            //外部传入的时间字符串
	//timeTemplate1 := "2006-01-02 15:04:05" //常规类型
	//localtime, _ := time.ParseInLocation(timeTemplate1, t1, time.Local)
	//mTime := localtime.UnixNano()
	//tie := fmt.Sprintf("%b", mTime)
	//fmt.Println(tie)

}
func testStruct() {
	type Set map[string]struct{}
	set := make(Set)

	for _, item := range []string{"A", "A", "B", "C"} {
		set[item] = struct{}{}
	}
	fmt.Println(len(set)) // 3
	if _, ok := set["A"]; ok {
		fmt.Println("A exists") // A exists
	}
}
func equalPairs(grid [][]int) int {
	temSlice := make([]int, 0)
	num := 0
	for key, _ := range grid {
		for _, v := range grid {
			temSlice = append(temSlice, v[key])
		}
		for _, v1 := range grid {
			if LoopCompare(v1, temSlice) {
				num++
			}
		}

	}
	return num

}
func LoopCompare(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func numberOfSteps(num int) int {
	n := 0
	for num != 0 {
		n++
		if (num % 2) == 0 {
			num = num / 2
		} else {
			num = (num - 1) / 2
		}
	}
	return n
}
func maximumWealth(accounts [][]int) int {
	temp := 0
	for _, account := range accounts {
		sum := 0

		for _, val := range account {
			sum += val
		}
		if sum > temp {
			temp = sum
		}
	}
	return temp
}
