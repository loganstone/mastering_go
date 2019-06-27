package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	fmt.Println("User id:", os.Getuid())

	var u *user.User
	u, _ = user.Current()

	group, _ := u.GroupIds()

	for _, i := range group {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
