package grpc_service

import (
	"context"
	"github.com/mayerkv/go-candidates/domain"
)

type CandidateServiceImpl struct {
	candidateService *domain.CandidateService
}

func NewCandidateServiceImpl(candidateService *domain.CandidateService) CandidatesServiceServer {
	return &CandidateServiceImpl{candidateService: candidateService}
}

func (s *CandidateServiceImpl) CreateCandidate(ctx context.Context, request *CreateCandidateRequest) (*CreateCandidateResponse, error) {
	candidate, err := s.candidateService.CreateCandidate(request.Name, request.Surname, mapFromContacts(request.Contacts))
	if err != nil {
		return nil, err
	}

	return &CreateCandidateResponse{
		Id: candidate.Id,
	}, nil
}

func (s *CandidateServiceImpl) GetCandidate(ctx context.Context, request *GetCandidateRequest) (*GetCandidateResponse, error) {
	candidate, err := s.candidateService.GetCandidate(request.Id)
	if err != nil {
		return nil, err
	}

	return &GetCandidateResponse{
		Candidate: mapCandidate(candidate),
	}, nil
}

func (s *CandidateServiceImpl) SearchCandidates(ctx context.Context, request *SearchCandidatesRequest) (*SearchCandidatesResponse, error) {
	page, err := s.candidateService.SearchCandidate(domain.Pageable{
		Page:           int(request.Page),
		Size:           int(request.Size),
		OrderBy:        mapCandidateOrderBy(request.OrderBy),
		OrderPredicate: mapOrderDirection(request.OrderDirection),
	})

	if err != nil {
		return nil, err
	}

	return &SearchCandidatesResponse{
		List:  mapCandidates(page.Items),
		Count: int32(page.TotalCount),
	}, nil
}

func (s *CandidateServiceImpl) mustEmbedUnimplementedCandidatesServiceServer() {
	panic("implement me")
}
