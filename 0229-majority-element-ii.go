// Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

// Example 1:

// Input: nums = [3,2,3]
// Output: [3]
// Example 2:

// Input: nums = [1]
// Output: [1]
// Example 3:

// Input: nums = [1,2]
// Output: [1,2]

// Constraints:

// 1 <= nums.length <= 5 * 104
// -109 <= nums[i] <= 109

// Follow up: Could you solve the problem in linear time and in O(1) space?

func majorityElement(nums []int) []int {
    n := len(nums)
    standard := n / 3

    memory := map[int]int{}

    for i := range nums {
        memory[nums[i]] += 1
    }

    var result []int
    for k, v := range memory {
        if v > standard {
            result = append(result, k)
        }
    }
    return result
}