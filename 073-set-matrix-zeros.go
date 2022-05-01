// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

// You must do it in place.

// Example 1:

// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
// Output: [[1,0,1],[0,0,0],[1,0,1]]
// Example 2:

// Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

// Constraints:

// m == matrix.length
// n == matrix[0].length
// 1 <= m, n <= 200
// -231 <= matrix[i][j] <= 231 - 1

// Follow up:

// A straightforward solution using O(mn) space is probably a bad idea.
// A simple improvement uses O(m + n) space, but still not the best solution.
// Could you devise a constant space solution?

func setZeroes(matrix [][]int) {
    n := len(matrix)
    m := len(matrix[0])

    markRow, markCol := false, false
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if matrix[i][j] == 0 {
                if i == 0 {
                    markRow = true
                }
                if j == 0 {
                    markCol = true
                }

                matrix[0][j] = 0
                matrix[i][0] = 0
            }
        }
    }

    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }

    if markRow {
        for j := 0; j < m; j++ {
            matrix[0][j] = 0
        }
    }

    if markCol {
        for i := 0; i < n; i++ {
            matrix[i][0] = 0
        }
    }
    return
}