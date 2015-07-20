package main

/**
	njson example - Copyright 2015 Joseph Lewis <joseph@josephlewis.net>
	Licensed under the BSD 3-clause license.

	Output:

		View to oauth clients (oauthView):
		{
			"name": "josephlewis42"
		}

		found in cache userView
		View to user (userView):
		{
			"my-name": "Joseph",
			"Username": "josephlewis42",
			"my-age": 99
		}

		found in cache
		all view ():
		{
			"RealName": "Joseph",
			"Username": "josephlewis42",
			"Age": 99
		}


**/

import (
	"fmt"
	"github.com/connectordb/njson"
)

type User struct {
	RealName string `oauthView:"-"    userView:"my-name"`
	Username string `oauthView:"name" `
	Age      int    `oauthView:"-" userView:"my-age"`
}

func decode(key string, title string) {
	u := User{"Joseph", "josephlewis42", 99}

	bytes, _ := njson.MarshalIndentWithTag(u, "", "\t", key)
	fmt.Printf("%s (%s):\n", title, key)
	fmt.Printf("%s\n", string(bytes))
	fmt.Println()
}

func main() {
	decode("oauthView", "View to oauth clients")
	decode("userView", "View to user")
	decode("", "all view")
}
