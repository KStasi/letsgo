package main

import (
	"fmt"

	. "github.com/KStasi/letsgo/workplace"
)

func main() {

	worker := Worker{
		HumanInfo{
			1023232,
			"Tom"
			"Ostin",
			Male,
			180,
			70
		},
		"Notary",
		"Teamdream",
		2000.,
		"380974320191",
		"roewa@gmail.com",
	}
	workerMockup := Worker{}
	lawyer := Lawyer{
		Workers : worker,
		Type : Social,
		LegislationArrea : "Agro",
	}

	fmt.Println(workerMockup)
	fmt.Println(lawyer)
}
