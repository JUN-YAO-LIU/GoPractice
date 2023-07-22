// 不用刪除 只是需要給指標讓他可以指得到就好
// 操作array的兩個指標
// 不刪除就是換位置
// 不相同往前排，就換到前面的位置
func removeElement(nums []int, val int) int {

	j := 0

	for i := 0; i < len(nums); i++{
		if nums[i] != val {

            fmt.Println("--1--",nums[i],nums[j])

            nums[j]= nums[i]
            fmt.Println("--2--",nums[i],nums[j])

			nums[i] = nums[j]
            fmt.Println("--3--",nums[i],nums[j])
            
			j++
		}
        
	}

	return j
}