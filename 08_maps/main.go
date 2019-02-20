package main

import "fmt"

func main() {
	// emails := make(map[string]string)
	//
	//
	// emails["Bug"] = "bug@gmail.com"
	// emails["Bub"] = "bub@gmail.com"
	// emails["Bog"] = "bog@gmail.com"

	emails := map[string]string{"Bug": "bug@gmail.com", "Bub": "bub@gmail.com", "Boug": "boug@gmail.com", "Bing": "bing@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bug"])

	delete(emails, "Bug")
	fmt.Println(emails)
}
