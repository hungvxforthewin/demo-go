package main

import "fmt"

func main() {
	// Creating string
	// String is array of char
	var greeting =  "Hello world!"

	fmt.Printf("normal string: ")
	fmt.Printf("%s", greeting)
	fmt.Printf("\n")
	fmt.Printf("hex bytes: ")

	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x ", greeting[i])
	}
	fmt.Printf("\n")
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%d ", greeting[i])
	}

	fmt.Printf("\n")
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%c ", greeting[i])
	}

	fmt.Printf("\n")
	const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	/*q flag escapes unprintable characters, with + flag it escapses non-ascii
	  characters as well to make output unambigous */
	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", sampleText)
	fmt.Printf("\n")


	// method of string
	/*
		len
		concat
		split
		template literal
		join
	*/
}
