package main

/**
	njson example - Copyright 2015 Joseph Lewis <joseph@josephlewis.net>
	Licensed under the BSD 3-clause license.
**/

import (
	"fmt"

	"github.com/connectordb/njson"
)

var (
	one = `
{
    "RealName": "John",
    "Username": "john123",
    "Age": 33
}
`
	two = `
{
	"name"	 : "Bob",
	"Username" : "Frank",
	"Age"	 : 33
}
`
)

type User struct {
	RealName string `publicdata:"-"    privatedata:"name"`
	Username string `publicdata:"name" privatedata:"-"`
	Age      int    `json:"-"`
}


func main() {
	// Identity
	doDecode("", one, User{"John", "john123", 33})
	// no matching fields except age
	doDecode("publicdata", one, User{"", "", 33})
	// no matching fields except age
	doDecode("privatedata", one, User{"", "", 33})
	// We can't know the user's age either
	doDecode("json", one, User{"John", "john123", 0})

	// User not specified
	doDecode("", two, User{"", "Frank", 33})
	// we don't know the user's real name
	doDecode("publicdata", two, User{"", "Bob", 33})
	// we don't know the user's username
	doDecode("privatedata", two, User{"Bob", "", 33})
	// We can't know the user's age
	doDecode("json", two, User{"", "Frank", 0})

}


func doDecode(key, value string, expected User) {
	bytes := []byte(value)

	output := User{}

	njson.UnmarshalTag(bytes, &output, key)

	fmt.Printf("=======================\n")
	fmt.Printf("Key: %v\n", key)
	fmt.Println()
	fmt.Println("Original")
	fmt.Println(value)
	fmt.Println("Decoded")
	fmt.Println(output)

	fmt.Println("Matches expected?")
	fmt.Printf("%v \n", output == expected)

}
/**
func doQuery(key string, title string) {
	u := User{"Joseph", "josephlewis42", 99}

	njson.

	bytes, _ := njson.MarshalIndentWithTag(u, "", "\t", key)
	fmt.Printf("%s (%s):\n", title, key)
	fmt.Printf("%s\n", string(bytes))
	fmt.Println()
}**/
