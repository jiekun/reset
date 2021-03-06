// Given an integer n, return a list of all simplified fractions between 0 and 1 (exclusive) such that the denominator is less-than-or-equal-to n. You can return the answer in any order.

// Example 1:

// Input: n = 2
// Output: ["1/2"]
// Explanation: "1/2" is the only unique fraction with a denominator less-than-or-equal-to 2.
// Example 2:

// Input: n = 3
// Output: ["1/2","1/3","2/3"]
// Example 3:

// Input: n = 4
// Output: ["1/2","1/3","1/4","2/3","3/4"]
// Explanation: "2/4" is not a simplified fraction because it can be simplified to "1/2".

// Constraints:

// 1 <= n <= 100

func simplifiedFractions(n int) []string {
    m := map[float64]bool{}

    var result []string

    for i := 2; i <= n; i++ {
        for j := 1; j < i; j++ {
            t := float64(j) / float64(i)
            if m[t] {
                continue
            }
            m[t] = true
            result = append(result, fmt.Sprintf("%d/%d", j, i))
        }
    }

    return result
}
