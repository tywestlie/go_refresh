package main

import "fmt"

func main() {
	dogArray := []string{"George", "Todd", "Lou", "Ada"}

	// dogArray[0] = "George"
	// dogArray[1] = "Todd"

	fmt.Println(len(dogArray))
	fmt.Println(dogArray[1:3])
}
