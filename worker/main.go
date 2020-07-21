package worker

import "fmt"

// Worker - struct represented worker info
type Worker struct {
	ID          uint64
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Possition   string
	Company     string
	Salary      float32
}

// NewWorker - safe method to create worker struct
func NewWorker(id uint64, firstName, lastName, phoneNumber, email, possition, company string, salary float32) *Worker {
	return &Worker{ID: id, FirstName: firstName, LastName: lastName, PhoneNumber: phoneNumber, Email: email, Possition: possition, Company: company, Salary: salary}
}

func main() {
	safeWorker := NewWorker(0, "Kate", "Ostin", "380976482598", "kateost@gmail.com", "Manager", "Teamdream", 1000.)
	simpleWorker := Worker{0, "Tom", "Ostin", "380974320191", "roewa@gmail.com", "Manager", "Teamdream", 2000.}
	defaultWorker := Worker{}

	fmt.Println(safeWorker)
	fmt.Println(simpleWorker)
	fmt.Println(defaultWorker)
}
