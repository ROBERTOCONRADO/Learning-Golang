package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10, 11, 12, 13}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	//len=5 cap=5 [2 4 6 8 10]

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	//len=0 cap=5 []

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
	//len=4 cap=5 [2 4 6 8]

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])
	//len=3 cap=3 [6 8 10]

	s = append(s, 100)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
	//len=2 cap=16 [2 4]
	
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[2:])
	//len=2 cap=16 [6 8 10 11 12 13 100]
}

