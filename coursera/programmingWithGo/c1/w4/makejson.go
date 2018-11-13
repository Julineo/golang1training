package main

import (
	"bufio"
	"os"
	"fmt"
	"encoding/json"
	"log"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := make(map[string]string)

	//read name
	scanner.Scan()
	m["name"] = scanner.Text()
	//read address
	scanner.Scan()
	m["address"] = scanner.Text()

	json, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}
