package main

import (
	"fmt"

	"github.com/KStasi/letsgo/workplace"
)

func main() {

	worker := workplace.Worker{
		workplace.Human{
			1023232,
			"Tom",
			"Ostin",
			workplace.Male,
			180,
			70,
		},
		"Notary",
		"Teamdream",
		2000.,
		"380974320191",
		"roewa@gmail.com",
	}
	lawyer := workplace.Lawyer{
		Workers:          worker,
		Type:             workplace.Social,
		LegislationArrea: "Agro",
	}

	fmt.Println(lawyer)
	fmt.Printf("%+v\n", lawyer)
}
