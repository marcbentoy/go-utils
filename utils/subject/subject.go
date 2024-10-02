package sdksubject

type Subject struct {
	Name string
}

func NewSubject(name string) *Subject {
	return &Subject{
		Name: name,
	}
}
