package workplace

// Gender - type of person's gender
type Gender int

// ProfessionType - type of profession
type ProfessionType int

const (
	// Male - man
	Male Gender = iota
	// Female - woman
	Female
	// Hidden - Unspecified gender
	Hidden
)

const (
	// Social - man
	Social ProfessionType = iota
	// Technical - woman
	Technical
	// Medical - Unspecified gender
	Medical
	// Econimical - Unspecified gender
	Econimical
	// Agronomical - Unspecified gender
	Agronomical
	// Creative - Unspecified gender
	Creative
)

// Human - struct represented person info
type Human struct {
	ID          uint64 `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Sex         Gender `json:"gender"`
	PhoneNumber string `json:"phone"`
	Email       string `json:"email"`
}

// Worker - struct represented worker info
type Worker struct {
	HumanInfo Human   `json:"human"`
	Possition string  `json:"possition"`
	Company   string  `json:"company"`
	Salary    float32 `json:"salary"`
}

// Profession - struct represented occupation info
type Profession struct {
	WorkerInfo Worker         `json:"worker"`
	Duties     string         `json:"duties"`
	Type       ProfessionType `json:"type"`
}

// NewWorker - safe method to create worker struct
// func NewWorker(id uint64,
// 	firstName string,
// 	lastName string,
// 	phoneNumber string,
// 	email string,
// 	possition string,
// 	company string,
// 	salary float32) Worker {
// 	return Worker{ID: id, FirstName: firstName, LastName: lastName, PhoneNumber: phoneNumber, Email: email, Possition: possition, Company: company, Salary: salary}
// }
