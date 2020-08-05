package main

import (
	"fmt"

	"github.com/KStasi/letsgo/workplace"
)

func main() {
	worker := workplace.Worker{
		Human: workplace.Human{
			ID:        1023232,
			FirstName: "Tom",
			LastName:  "Ostin",
			Sex:       workplace.Male,
			Height:    180,
			Weight:    70,
			Age:       20,
		},
		Possition:   "Notary",
		Company:     "Teamdream",
		Salary:      2000.,
		PhoneNumber: "380974320191",
		Email:       "roewa@gmail.com",
	}
	lawyer := workplace.Lawyer{
		Worker:           worker,
		Type:             workplace.Social,
		LegislationArrea: "Agro",
	}
	fmt.Printf("%+v\n", lawyer)
	lawyer.CelebrateBirthday()
	fmt.Printf("%+v\n", lawyer)
}
