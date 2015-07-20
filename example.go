package main

/**
	njson example - Copyright 2015 Joseph Lewis <joseph@josephlewis.net>
	Licensed under the BSD 3-clause license.
**/

import (
	"fmt"

	"github.com/connectordb/njson"
)

type User struct {
	RealName string `publicdata:"-"    privatedata:"name"`
	Username string `publicdata:"name" privatedata:"-"`
	Age      int    `json:"-"`
}

func doQuery(key string, title string) {
	u := User{"Joseph", "josephlewis42", 99}

	bytes, _ := njson.MarshalIndentWithTag(u, "", "\t", key)
	fmt.Printf("%s (%s):\n", title, key)
	fmt.Printf("%s\n", string(bytes))
	fmt.Println()

	{
		bytes, _ := njson.MarshalIndent(u, "", "\t")
		fmt.Printf("%s (%s):\n", "full", "json")
		fmt.Printf("%s\n", string(bytes))
		fmt.Println()
	}
}

func main() {
	//doQuery("publicdata", "Public View")
	//doQuery("privatedata", "Private View")
	//doQuery("json", "json tag")
	doQuery("publicdata", "non-existing tag")
}
