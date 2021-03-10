package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	personInformation := make(map[string]string)

	userInput := bufio.NewScanner(os.Stdin)

	// get name
	fmt.Println("Please input your name")

	userInput.Scan()
	personName := userInput.Text()

	personInformation["name"] = personName

	// get address
	fmt.Println("Please input your address")

	userInput.Scan()
	personAddr := userInput.Text()

	personInformation["address"] = personAddr

	// convert Golang map to JSON object
	json_object, err := json.Marshal(personInformation)

	if err == nil {
		fmt.Println("JSON object of personInformation: ", json_object)
		fmt.Println("Golang sttring of JSON object: ", string(json_object))
	}

	return
}
