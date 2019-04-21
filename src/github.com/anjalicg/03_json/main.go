// https://www.json-generator.com/
// https://tutorialedge.net/golang/parsing-json-with-golang/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	jsonFilePath := "./generated.json"

	fmt.Println(jsonFilePath)
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Json file opened successfully")
	fmt.Println(jsonFile)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// fmt.Println(byteValue)
	// var result map[string]interface{}
	var result1 []map[string]interface{}
	var result2 interface{}
	json.Unmarshal([]byte(byteValue), &result1)
	json.Unmarshal([]byte(byteValue), &result2)
	fmt.Println(result1[0]["isActive"])
	fmt.Println(result2)

	// Parsing JSON with structs:-
	type Social struct {
		Facebook string `json:"facebook"`
		Twitter  string `json:"twitter"`
	}

	type User struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Age    int    `json:"Age"`
		Social Social `json:"social"`
	}

	type Users struct {
		Users []User `json:"users"`
	}
	jsonFilePath2 := "./users.json"
	jsonFile2, err2 := os.Open(jsonFilePath2)
	if err2 != nil {
		fmt.Println(err)
	}
	defer jsonFile2.Close()
	byteValue2, _ := ioutil.ReadAll(jsonFile2)
	var users Users
	json.Unmarshal(byteValue2, &users)
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

	return
	//VCP-N certification

}
