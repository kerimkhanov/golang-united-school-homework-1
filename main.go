package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	HelloWorld := emoji.Sprint(":world_map:")
	fmt.Println("Hello " + HelloWorld + "!")
}
