package main

import "fmt"

func main() {
	result := longestOnes([]int {1,1,1,0,0,0,1,1,1,1}, 0)
	fmt.Printf("Result: %v", result)
}

func longestOnes(nums []int, k int) int {
	// Fill initial k zeroes
	length := len(nums)
	start, end := 0, 0
	if (k == 0) {
		end++
	} else {			
		start, end = fillZeroes(nums, length, 0, k)
	}
	max := checkAndUpdateMax(0, start, end)

	// Loop until end reaches the end of the slice
	for end + 1 < length {
        if nums[end + 1] == 1 {
            if k == 0 && nums[start] == 1 {
				max = checkAndUpdateMax(max, start, end - 1)
				end++
			} else {
				end++ 
				max = checkAndUpdateMax(max, start, end)
			}
        } else if nums[start] == 0 || k == 0 {
            start++
			end++
        } else {			
			start, end = fillZeroes(nums, length, 0, k)
        }
    }	

	max = checkAndUpdateMax(max, start, end)
    return max
}

func fillZeroes(nums []int, length, start, k int) (int, int) {    
	// Fill k zeroes
	kCount := 0
	end := start
	for end < length && kCount < k{
        if nums[end] == 0 {
            kCount++
        }
		end++
    }

	// return end decremented by 1 as it was already increased
    return start, end - 1
}

func checkAndUpdateMax(max, start, end int) int {
	if end + 1 - start > max {
		max = end + 1 - start
	}
	return max;
}
