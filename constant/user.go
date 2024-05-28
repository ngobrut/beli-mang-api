package constant

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

var ValidGender = map[string]bool{
	string(Male):   true,
	string(Female): true,
}

var Genders = []string{
	string(Male),
	string(Female),
}
