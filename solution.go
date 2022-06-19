package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	HelloWorld := emoji.Sprint(":world_map:")
	return "Hello " + HelloWorld + "!"
}
