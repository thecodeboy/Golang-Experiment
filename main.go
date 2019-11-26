package main

import (
	"fmt"

	"github.com/EasyDB-io/Golang/client"
)

func main() {
	c := client.Connect("9b1f4ce4-7a13-42f2-9806-3ff30e608c0e", "3da03b86-7762-4083-8789-749485a6d134")
	c.Put("517", "https://www.google.com/search?q=bytes+to+string+golang&oq=bytes+to+string+golang&aqs=chrome..69i57j0l5.3769j0j4&sourceid=chrome&ie=UTF-8")
	s, err := c.Get("517")
	if err != nil {
		fmt.Println("Error fetching value.")
	}
	fmt.Println(string(s))
	c.Delete("517")
	fmt.Println(c.List())
}
