package domain

type CandidatePage struct {
	Items      []Candidate
	TotalCount int
	Page       int
	Size       int
}

type CandidateRepository interface {
	Save(candidate *Candidate) error
	FindById(id string) (*Candidate, error)
	FindAll(pageable Pageable) (CandidatePage, error)
}
