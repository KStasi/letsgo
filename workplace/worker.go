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

// Human - struct represented person info
type Human struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Sex       Gender `json:"gender"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
}

// Worker - struct represented worker info
type Worker struct {
	HumanInfo   Human   `json:"human"`
	Possition   string  `json:"possition"`
	Company     string  `json:"company"`
	Salary      float32 `json:"salary"`
	PhoneNumber string  `json:"phone"`
	Email       string  `json:"email"`
}

// Doctor - struct represented occupation info
type Doctor struct {
	Workers      Worker         `json:"workers"`
	Type         ProfessionType `json:"type"`
	Department   string         `json:"department"`
	Education    string         `json:"education"`
	WorkingHours string         `json:"working_hours"`
}

// Developer - struct represented occupation info
type Developer struct {
	Workers  Worker         `json:"workers"`
	Type     ProfessionType `json:"type"`
	WorkType WorkType       `json:"work_type"`
}

// Lawyer - struct represented occupation info
type Lawyer struct {
	Workers          Worker         `json:"workers"`
	Type             ProfessionType `json:"type"`
	LegislationArrea string         `json:"legislation_arrea"`
}

// Sportsman - struct represented occupation info
type Sportsman struct {
	Workers Worker         `json:"workers"`
	Type    ProfessionType `json:"type"`
	Rewards []string       `json:"rewards"`
	Rank    string         `json:"rank"`
}

// Nurse - struct represented occupation info
type Nurse struct {
	Workers        Worker         `json:"workers"`
	Type           ProfessionType `json:"type"`
	Experience     int32          `json:"experience"`
	HasOwnChildren bool           `json:"has_children"`
}
