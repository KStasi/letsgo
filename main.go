package main

import (
	"fmt"

	. "github.com/KStasi/letsgo/worker"
)

func main() {
	safeWorker := NewWorker(0,
		"Kate",
		"Ostin",
		"380976482598",
		"kateost@gmail.com",
		"Manager",
		"Teamdream",
		1000.,
	)
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

	fmt.Println(safeWorker)
	fmt.Println(simpleWorker)
	fmt.Println(defaultWorker)
}
