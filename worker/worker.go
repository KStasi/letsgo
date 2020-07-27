package worker

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
func NewWorker(id uint64,
	firstName string,
	lastName string,
	phoneNumber string,
	email string,
	possition string,
	company string,
	salary float32) *Worker {
	return &Worker{ID: id, FirstName: firstName, LastName: lastName, PhoneNumber: phoneNumber, Email: email, Possition: possition, Company: company, Salary: salary}
}
