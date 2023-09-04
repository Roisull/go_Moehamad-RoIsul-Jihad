package main

import "fmt"

func generate(numRows int) [][]int {
    if numRows <= 0 {
        return [][]int{}
    }
    
    result := make([][]int, numRows)
    
    for i := 0; i < numRows; i++ {
        result[i] = make([]int, i+1)
        result[i][0], result[i][i] = 1, 1
        
        for j := 1; j < i; j++ {
            result[i][j] = result[i-1][j-1] + result[i-1][j]
        }
    }
    
    return result
}

func main() {
    numRows := 5
    triangle := generate(numRows)
    
    for _, row := range triangle {
        fmt.Println(row)
    }
}