package domain

import "github.com/google/uuid"

type Candidate struct {
	Id string
	Name string
	Surname string
	Contacts []Contact
}

func CreateCandidate(name, surname string, contacts []Contact) *Candidate {
	return &Candidate{
		Id:       uuid.NewString(),
		Name:     name,
		Surname:  surname,
		Contacts: contacts,
	}
}