package grpc_service

import "github.com/mayerkv/go-candidates/domain"

func mapFromContacts(contacts []*Contact) []domain.Contact {
	var res []domain.Contact

	for _, c := range contacts {
		res = append(res, mapContact(*c))
	}

	return res
}

func mapContact(c Contact) domain.Contact {
	var t domain.ContactType
	switch c.Type {
	case Contact_EMAIL:
		t = domain.ContactTypeEmail
	case Contact_PHONE:
		t = domain.ContactTypePhone
	}

	return domain.Contact{
		Type:  t,
		Value: c.Value,
	}
}

func mapCandidate(candidate *domain.Candidate) *Candidate {
	return &Candidate{
		Id:       candidate.Id,
		Name:     candidate.Name,
		Surname:  candidate.Surname,
		Contacts: mapToContacts(candidate.Contacts),
	}
}

func mapToContacts(contacts []domain.Contact) []*Contact {
	var res []*Contact

	for _, c := range contacts {
		res = append(res, mapToContact(c))
	}
	return res
}

func mapToContact(contact domain.Contact) *Contact {
	var t Contact_Type
	switch contact.Type {
	case domain.ContactTypeEmail:
		t = Contact_EMAIL
	case domain.ContactTypePhone:
		t = Contact_PHONE
	}

	return &Contact{
		Type:  t,
		Value: contact.Value,
	}
}

func mapCandidates(candidates []domain.Candidate) []*Candidate {
	var res []*Candidate

	for _, c := range candidates {
		res = append(res, mapCandidate(&c))
	}

	return res
}

func mapOrderDirection(direction OrderDirection) domain.Predicate {
	if direction == OrderDirection_DESC {
		return domain.PredicateDesc
	}

	return domain.PredicateAsc
}

func mapCandidateOrderBy(by Candidate_Order) string {
	if by == Candidate_FULL_NAME {
		return "fullName"
	}

	return ""
}
