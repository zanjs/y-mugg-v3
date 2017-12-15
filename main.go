package main

import (
	"fmt"

	"github.com/zanjs/y-mugg-v3/app/router"
)

func main() {

	qm := 4001 % 40

	fmt.Println(qm)

	router.InitRoute()

}
