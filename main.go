package main

import "fmt"

func main() {
	token, _ := GetToken()
	fmt.Println(token)

	c, _ := ParseConfig()
	fmt.Printf("%+v\n", c)

	c.SharedTestConfig.Token = token
	fmt.Println(c.SharedTestConfig.Token)
}
