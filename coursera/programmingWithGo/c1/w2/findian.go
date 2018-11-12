/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
*/
package main

import (
	"bufio"
	"os"
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^[iI].*[aA].*[nN]$")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter string: ")
	str, _ := reader.ReadString('\n')
	b := []byte(str[:len(str)-1])

	if len(re.Find(b)) > 0 {
		fmt.Println("Found!")
		return
	}
	fmt.Println("Not Found!")
}

/*
//versions:

func main() {
	fmt.Println("Please enter a string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Errorf("Problems reading input: %e.\n", err)
	}
	s = strings.ToLower(s)
	if strings.Contains(s, "a") && strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func main()
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func main() {
	const foundString = "Found!"
	const notFoundString = "Not Found!"

	fmt.Println("Enter string:")
	r := bufio.NewReader(os.Stdin)
	input, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(notFoundString)
	} else {
		cleanedInput := strings.ToLower(input)
		var startsWithI = strings.HasPrefix(cleanedInput, "i")
		var containsAInBetween = strings.Index(cleanedInput, "a") > 0
		var endsWithN = strings.HasSuffix(cleanedInput, "n\n")

		if startsWithI && containsAInBetween && endsWithN {
			fmt.Println(foundString)
		} else {
			fmt.Println(notFoundString)
		}
	}
}
*/
