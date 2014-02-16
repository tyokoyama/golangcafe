package main

import (
	"fmt"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Current User: %v\n", u)

	// ユーザIDを自分で調べて書き換えて下さい。
	// 調べ方例 for Mac：finger yokoyamaなど。
	u, err = user.LookupId("501")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("LookupId User: %v\n", u)

	u, err = user.Lookup("root")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Printf("Lookup User: %v\n", u)

}