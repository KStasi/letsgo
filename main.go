package main

import (
	"fmt"

	. "github.com/KStasi/letsgo/workplace"
)

func main() {

	simpleWorker := Worker{0,
		"Tom",
		"Ostin",
		"380974320191",
		"roewa@gmail.com",
		"Manager",
		"Teamdream",
		2000.,
	}
	defaultWorker := Worker{}

	fmt.Println(simpleWorker)
	fmt.Println(defaultWorker)
}
