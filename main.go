package main

import (
	"fmt"

	"github.com/zxylon/olaf/cmd/olaf"
)

func main() {
	err := olaf.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
