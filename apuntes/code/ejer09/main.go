package main

import (
	"fmt"
	"time"

	us "./user"
)

type pepe struct {
	us.User
}

func main() {
	user := new(pepe)
	user.SetUser(1, "pepe", time.Now(), true)
	fmt.Println(user.User)
}
