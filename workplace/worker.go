package workplace

// Gender - type of person's gender
type Gender int

// ProfessionType - type of profession
type ProfessionType int

// WorkType - type of profession
type WorkType int

const (
	// Male - man
	Male Gender = iota
	// Female - woman
	Female
	// Hidden - Unspecified gender
	Hidden
)
const (
	// Remote - man
	Remote WorkType = iota
	// Office - woman
	Office
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

// IWorker - interface to indroduce own profession, position and name
type IWorker interface {
	Introduce() (string, string, string)
}

// Human - struct represented person info
type Human struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Sex       Gender `json:"gender"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
	Age       int    `json:"age"`
}

// Worker - struct represented worker info
type Worker struct {
	Possition   string  `json:"possition"`
	Company     string  `json:"company"`
	Salary      float32 `json:"salary"`
	PhoneNumber string  `json:"phone"`
	Email       string  `json:"email"`
	Human
}

// Doctor - struct represented occupation info
type Doctor struct {
	Type         ProfessionType `json:"type"`
	Department   string         `json:"department"`
	Education    string         `json:"education"`
	WorkingHours string         `json:"working_hours"`
	Worker
}

// Developer - struct represented occupation info
type Developer struct {
	Type     ProfessionType `json:"type"`
	WorkType WorkType       `json:"work_type"`
	Worker
}

// Lawyer - struct represented occupation info
type Lawyer struct {
	Type             ProfessionType `json:"type"`
	LegislationArrea string         `json:"legislation_arrea"`
	Worker
}

// Sportsman - struct represented occupation info
type Sportsman struct {
	Type    ProfessionType `json:"type"`
	Rewards []string       `json:"rewards"`
	Rank    string         `json:"rank"`
	Worker
}

// Nurse - struct represented occupation info
type Nurse struct {
	Type           ProfessionType `json:"type"`
	Experience     int32          `json:"experience"`
	HasOwnChildren bool           `json:"has_children"`
	Worker
}

// CelebrateBirthday - method to make person older
func (receiver *Human) CelebrateBirthday() {
	receiver.Age++
}

// UpdateShape - method to modify person's shape
func (receiver *Human) UpdateShape(newHeight int, newWeight int) {
	receiver.Height = newHeight
	receiver.Weight = newWeight
}

// Introduce - interface IWorker implemented for Doctor
func (d Doctor) Introduce() (string, string, string) {
	return "Doctor", d.Possition, d.FirstName
}

// Introduce - interface IWorker implemented for Developer
func (d Developer) Introduce() (string, string, string) {
	return "Developer", d.Possition, d.FirstName
}

// Introduce - interface IWorker implemented for Lawyer
func (l Lawyer) Introduce() (string, string, string) {
	return "Lawyer", l.Possition, l.FirstName
}

// Introduce - interface IWorker implemented for Sportsman
func (s Sportsman) Introduce() (string, string, string) {
	return "Sportsman", s.Possition, s.FirstName
}

// Introduce - interface IWorker implemented for Nurse
func (n Nurse) Introduce() (string, string, string) {
	return "Nurse", n.Possition, n.FirstName
}
