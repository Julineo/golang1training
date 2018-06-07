package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
	
	data2 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data2)) // `["one" "three"]`
	fmt.Printf("%q\n", data2)           // `["one" "three" "three"]`
	
	data3 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty3(data3)) // `["one" "three"]`
	fmt.Printf("%q\n", data3)           // `["one" "three" "three"]`
}


func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		fmt.Println(s)
		if s != "" {
			fmt.Println(s)
			out = append(out, s)
		}
	}
	return out
}

//in-place function to eliminate adjacent duplicates in a []string slice
func nonempty3(strings []string) []string {
	out := make([]string, 0) // zero-length slice or this will work as well: []string{}
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}