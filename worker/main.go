package worker

import "fmt"

// Worker - struct represented worker info
type worker struct {
	ID          uint64
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Possition   string
	Company     string
	Salary      float32
}

func NewWorker(id uint64, FirstName, LastName, PhoneNumber, Email, Possition, Company string, Salary float32) *Worker {
	return &worker{id, FirstName, LastName, PhoneNumber, Email, Possition, Company, Salary}
}

func main() {
	v := NewWorker(0, "Kate", "Ostin", "380976482598", "kateost@gmail.com", "Manager", "Teamdream", 1000.)
	fmt.Println(v)
}
