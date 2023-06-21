package main
import (
    "fmt"
    "time")


// Make a board filled with periods.
func make_board(num_rows int) [][]string {
    num_cols := num_rows
    board := make([][]string, num_rows)
    for r := range board {
        board[r] = make([]string, num_cols)
        for c := 0; c < num_cols; c++ {
            board[r][c] = "."
        }
    }
    return board
}


// Display the board.
func dump_board(board [][]string, num_rows int) {
    for r := 0; r < len(board); r++ {
        for c := 0; c < len(board[r]); c++ {
            fmt.Printf("%s ", board[r][c])
        }
        fmt.Println()
    }
}


// Return true if this series of squares contains at most one queen.
func series_is_legal(board [][]string, num_rows, r0, c0, dr, dc int) bool {
    num_cols := num_rows
    has_queen := false

    r := r0
    c := c0
    for {
        if board[r][c] == "Q" {
            // If we already have a queen on this row,
            // then this board is not legal.
            if has_queen { return false }

            // Remember that we have a queen on this row.
            has_queen = true
        }

        // Move to the next square in the series.
        r += dr
        c += dc

        // If we fall off the board, then the series is legal.
        if  r >= num_rows ||
            c >= num_cols ||
            r < 0 ||
            c < 0 {
                return true
        }
    }
}


// Return true if the board is legal.
func board_is_legal(board [][]string, num_rows int) bool {
    // See if each row is legal.
    for r := 0; r < num_rows; r++ {
        if !series_is_legal(board, num_rows, r, 0, 0, 1) { return false }
    }

    // See if each column is legal.
    for c := 0; c < num_rows; c++ {
        if !series_is_legal(board, num_rows, 0, c, 1, 0) { return false }
    }

    // See if diagonals down to the right are legal.
    for r := 0; r < num_rows; r++ {
        if !series_is_legal(board, num_rows, r, 0, 1, 1) { return false }
    }
    for c := 0; c < num_rows; c++ {
        if !series_is_legal(board, num_rows, 0, c, 1, 1) { return false }
    }

    // See if diagonals down to the left are legal.
    for r := 0; r < num_rows; r++ {
        if !series_is_legal(board, num_rows, r, num_rows - 1, 1, -1) { return false }
    }
    for c := 0; c < num_rows; c++ {
        if !series_is_legal(board, num_rows, 0, c, 1, -1) { return false }
    }

    // If we survived this long, then the board is legal.
    return true
}


// Return true if the board is legal and a solution.
func board_is_a_solution(board [][]string, num_rows int) bool {
    // See if it is legal.
    if !board_is_legal(board, num_rows) { return false }

    // See if the board contains exactly num_rows queens.
    num_queens := 0
    for r := 0; r < num_rows; r++ {
        for c := 0; c < num_rows; c++ {
            if board[r][c] == "Q" { num_queens += 1 }
        }
    }
    return num_queens == num_rows
}


// Try placing a queen at position [r][c].
// Return true if we find a legal board.
func place_queens_1(board [][]string, num_rows, r, c int) bool {
    if r >= num_rows {
        return board_is_a_solution(board, num_rows)
    }

    // Find the next square.
    next_r := r
    next_c := c + 1
    if next_c >= num_rows {
        next_r += 1
        next_c = 0
    }

    if place_queens_1(board, num_rows, next_r, next_c) {
        return true
    }

    board[r][c] = "Q"
    if place_queens_1(board, num_rows, next_r, next_c) {
        return true
    }

    board[r][c] = "."
    return false
}


func main() {
    const num_rows = 5
    board := make_board(num_rows)

    start := time.Now()
    success := place_queens_1(board, num_rows, 0, 0)
    //success := place_queens_2(board, num_rows, 0, 0, 0)
    //success := place_queens_3(board, num_rows, 0, 0, 0)

    elapsed := time.Since(start)
    if success {
        fmt.Println("Success!")
        dump_board(board, num_rows)
    } else {
        fmt.Println("No solution")
    }
    fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}
