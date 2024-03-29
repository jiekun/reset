// There is a street with n * 2 plots, where there are n plots on each side of the street. The plots on each side are numbered from 1 to n. On each plot, a house can be placed.

// Return the number of ways houses can be placed such that no two houses are adjacent to each other on the same side of the street. Since the answer may be very large, return it modulo 109 + 7.

// Note that if a house is placed on the ith plot on one side of the street, a house can also be placed on the ith plot on the other side of the street.

// Example 1:

// Input: n = 1
// Output: 4
// Explanation:
// Possible arrangements:
// 1. All plots are empty.
// 2. A house is placed on one side of the street.
// 3. A house is placed on the other side of the street.
// 4. Two houses are placed, one on each side of the street.
// Example 2:

// Input: n = 2
// Output: 9
// Explanation: The 9 possible arrangements are shown in the diagram above.

// Constraints:

// 1 <= n <= 104

var p int64 = 1000000007

func countHousePlacements(n int) int {
    // house placed dp1[i]
    dp1 := make([]int64, n, n)
    // house not placed dp1[i]
    dp2 := make([]int64, n, n)
    dp1[0] = 1
    dp2[0] = 1
    for i := 1; i < n; i++ {
        dp2[i] = (dp1[i-1] + dp2[i-1]) % p
        dp1[i] = (dp2[i-1]) % p
    }

    return int((dp1[n-1] + dp2[n-1]) * (dp1[n-1] + dp2[n-1]) % p)
}
