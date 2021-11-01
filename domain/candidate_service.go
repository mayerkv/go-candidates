package domain

import "errors"

var (
	ErrCandidateNotExists = errors.New("candidate not exists")
)

type CandidateService struct {
	candidateRepository CandidateRepository
}

func NewCandidateService(candidateRepository CandidateRepository) *CandidateService {
	return &CandidateService{candidateRepository: candidateRepository}
}

func (s *CandidateService) CreateCandidate(name, surname string, contacts []Contact) (*Candidate, error) {
	candidate := CreateCandidate(name, surname, contacts)

	if err := s.candidateRepository.Save(candidate); err != nil {
		return nil, err
	}

	return candidate, nil
}

func (s *CandidateService) GetCandidate(id string) (*Candidate, error) {
	candidate, err := s.candidateRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if candidate == nil {
		return nil, ErrCandidateNotExists
	}

	return candidate, nil
}

func (s *CandidateService) SearchCandidate(pageable Pageable) (CandidatePage, error) {
	return s.candidateRepository.FindAll(pageable)
}
