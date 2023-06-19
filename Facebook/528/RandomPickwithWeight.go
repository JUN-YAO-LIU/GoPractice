type Solution struct {
    prefixSum []int
    totalSum  int
}


func Constructor(w []int) Solution {
    totalSum := 0
    prefixSum := make([]int, len(w))
    // we make sure w lenght is greater than 0
    if len(w) <= 0 {
        return Solution { 
            prefixSum: prefixSum,
            totalSum: totalSum,
        }
    }
    prefixSum[0] = w[0]
    totalSum = w[0]
    for i := 1; i < len(w); i++ {
        prefixSum[i] = prefixSum[i-1] + w[i]
        totalSum += w[i]
    }
    return Solution {
        prefixSum: prefixSum,
        totalSum: totalSum,
    }
}


func (this *Solution) PickIndex() int {
    // binary search the left boundary
    target := rand.Intn(this.totalSum)
    idx := searchIndexGreaterThan(this.prefixSum, target)
    return idx
}

func searchIndexGreaterThan(nums []int, target int) int{
    left := 0
    right := len(nums)-1

    for left <= right {
        mid := left + (right - left) / 2
        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    if left >= len(nums) {
        return -1
    }

    return left
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */