[![Build Status](https://travis-ci.org/connectordb/njson.svg)](https://travis-ci.org/connectordb/njson)[![Coverage Status](https://coveralls.io/repos/connectordb/njson/badge.svg?branch=master&service=github)](https://coveralls.io/github/connectordb/njson?branch=master)[![GoDoc](https://godoc.org/github.com/connectordb/njson?status.svg)](http://godoc.org/github.com/connectordb/njson)

# njson
A JSON parser derived from go's that supports user-defined tags.

For example, let's say you wanted two different views on an object for your
users and for oauth clients.


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

Output
------

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
