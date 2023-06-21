package main
import (
    "fmt"
    "time")


// The board dimensions.
const num_rows = 8
const num_cols = num_rows

// Whether we want an open or closed tour.
const require_closed_tour = false

// Value to represent a square that we have not visited.
const unvisited = -1

// Define offsets for the knight's movement.
type Offset struct {
    dr, dc int
}

var move_offsets []Offset

var num_calls int64


func initialize_offsets() {
    move_offsets = []Offset {
        Offset {-2, -1},
        Offset {-1, -2},
        Offset {+2, -1},
        Offset {+1, -2},
        Offset {-2, +1},
        Offset {-1, +2},
        Offset {+2, +1},
        Offset {+1, +2},
    }
}


func make_board(num_rows, num_cols int) [][]int {
    board := make([][]int, num_rows)
    for i := 0; i < len(board); i++ {
        board[i] = make([]int, num_cols)
        for j := 0; j < num_cols; j++ {
            board[i][j] = unvisited
        }
    }
    return board
}


func dump_board(board [][]int) {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            fmt.Printf("%02d ", board[i][j])
        }
        fmt.Println()
    }
}


// Try to extend a knight's tour starting at (start_row, start_col).
// Return true or false to indicate whether we have found a solution.
func find_tour(board [][]int, num_rows, num_cols, cur_row, cur_col, num_visited int) bool {
    num_calls += 1
    if num_visited == num_rows * num_cols {
        if require_closed_tour {
            for _, offset := range move_offsets {
                row := cur_row + offset.dr
                col := cur_col + offset.dc
                if row >= 0 && row < num_rows && col >= 0 && col < num_cols && board[row][col] == 0 {
                    return true
                }
            }
            return false
        } else {
            return true
        }
    } else {
        for _, offset := range move_offsets {
            row := cur_row + offset.dr
            col := cur_col + offset.dc

            //skip where target is off board or already visited
            if row < 0 || row >= num_rows || col < 0 || col >= num_cols || board[row][col] != unvisited {
                continue
            }

            //valid row and col giving unvisited target
            board[row][col] = num_visited
            if find_tour(board, num_rows, num_cols, row, col, num_visited + 1) {
                return true
            }

            //backtrack - undo the move as this state can't find solution
            board[row][col] = unvisited
        }
        return false
    }
}


func main() {
    num_calls = 0

    // Initialize the move offsets.
    initialize_offsets()

    // Create the blank board.
    board := make_board(num_rows, num_cols)

    // Try to find a tour.
    start := time.Now()
    board[0][0] = 0
    if find_tour(board, num_rows, num_cols, 0, 0, 1) {
        fmt.Println("Success!")
    } else {
        fmt.Println("Could not find a tour.")
    }
    elapsed := time.Since(start)
    dump_board(board)
    fmt.Printf("%f seconds\n", elapsed.Seconds())
    fmt.Printf("%d calls\n", num_calls)
}
