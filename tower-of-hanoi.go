package main
import "fmt"
//import "time"


const num_disks = 9
var num_moves = 0


// Add a disk to beginning of the post.
func push(post []int, disk int) []int {
    return append([]int{disk}, post...)
}


// Remove the first disk from the post.
// Return that disk and the revised post.
func pop(post []int) (int, []int) {
    return post[0], post[1:]
}


// Move one disk from from_post to to_post.
func move_disk(posts [][]int, from_post, to_post int) {
    var disk int
    disk, posts[from_post] = pop(posts[from_post])
    posts[to_post] = push(posts[to_post], disk)
}


// Draw the posts by showing the size of the disk at each level.
func draw_posts(posts [][]int) {
    // Add 0s to the end of each post so they all have num_disks entries.
    for p := 0; p < 3; p++ {
        for len(posts[p]) < num_disks {
            posts[p] = push(posts[p], 0)
        }
    }

    // Draw the posts.
    for row := 0; row < num_disks; row++ {
        // Draw this row.
        for p := 0; p < 3; p++ {
            // Draw the disk on post p's row.
            fmt.Printf("%d ", posts[p][row])
        }
        fmt.Println()
    }
    // Draw a line between moves.
    fmt.Println("-----")

    // Remove the 0s.
    for p := 0; p < 3; p++ {
        for len(posts[p]) > 0 && posts[p][0] == 0 {
            _, posts[p] = pop(posts[p])
        }
    }
}


// Move the disks from from_post to to_post
// using temp_post as temporary storage.
func move_disks(posts [][]int, num_to_move, from_post, to_post, temp_post int) {
    //fmt.Println(num_disks)
    //time.Sleep(500 * time.Millisecond)
    if num_to_move > 0 {
        move_disks(posts, num_to_move - 1, from_post, temp_post, to_post)
        move_disk(posts, from_post, to_post)
        num_moves += 1
        draw_posts(posts)
        if num_to_move > 1 {
            move_disks(posts, num_to_move - 1, temp_post, to_post, from_post)
        }
    }
}


func main() {
    // Make three posts.
    posts := [][]int {}

    // Push the disks onto post 0 biggest first.
    posts = append(posts, []int{})
    for disk := num_disks; disk > 0; disk-- {
        posts[0] = push(posts[0], disk)
    }

    // Make the other posts empty.
    for p := 1; p < 3; p++ {
        posts = append(posts, []int{})
    }

    // Draw the initial setup.
    draw_posts(posts)

    // Move the disks.
    move_disks(posts, num_disks, 0, 1, 2)
    
    fmt.Printf("It took %d moves for %d disk(s).\n", num_moves, num_disks)
}
