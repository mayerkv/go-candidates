package domain

type ContactType int

const (
	ContactTypePhone ContactType = iota
	ContactTypeEmail
)

type Contact struct {
	Type ContactType
	Value string
}
