package main

import "fmt"

func main() {
	// ids := []int{33, 76, 85, 53, 21, 32, 6}

	emails := map[string]string{"Bug": "bug@gmail.com", "Bub": "bub@gmail.com", "Boug": "boug@gmail.com", "Bing": "bing@gmail.com"}

	// for i, id := range ids {
	// 	fmt.Printf("%d - ID: %d\n", i, id)
	// }
	//
	// for _, id := range ids {
	// 	fmt.Printf("ID: %d\n", id)
	// }
	//
	// sum := 0
	//
	// for _, id := range ids {
	// 	sum += id
	// }
	// fmt.Println("Sum:", sum)

	for key, value := range emails {
		fmt.Printf("name: %v, email: %v\n", key, value)
	}

}
