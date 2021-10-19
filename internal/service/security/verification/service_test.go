package verification

import (
	"runtime/debug"
	"testing"
)

type VerificationServiceChecker struct {
	t *testing.T
}

func (ch VerificationServiceChecker) CheckDescribeOK(s VerificationService, id uint64, expected Verification) {
	entity, err := s.Describe(id)
	if entity == nil {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected not nil, got nil")
	}
	if err != nil {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected nil, got not nil")
	}
	if entity.String() != expected.String() {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected [%s], got [%s]", expected.String(), entity.String())
	}
}

func (ch VerificationServiceChecker) CheckDescribeNotOK(s VerificationService, id uint64, expectedErr error) {
	entity, err := s.Describe(id)
	if entity != nil {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected nil, got not nil")
	}
	if err == nil {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected not nil, got nil")
	}
	if err != expectedErr {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected [%#+v], got [%#+v]", expectedErr, err)
	}
}

type testListResult []Verification

func (t testListResult) String() string {
	res := ""
	for _, v := range t {
		res += "|" + v.String() + "|"
	}
	return res
}

func (ch VerificationServiceChecker) CheckListOK(s VerificationService, cursor uint64, limit uint64, expected testListResult) {
	res, err := s.List(cursor, limit)
	if res == nil {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected not nil, got nil")
	}
	if err != nil {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected nil, got not nil")
	}

	testRes := testListResult(res)

	if testRes.String() != expected.String() {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected [%s], got [%s]", expected, testRes)
	}
}

func (ch VerificationServiceChecker) CheckListNotOK(s VerificationService, cursor uint64, limit uint64, expectedErr error) {
	res, _ := s.List(cursor, limit)
	if res != nil {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected nil, got not nil")
	}
}

func (ch VerificationServiceChecker) CheckRemoveOK(s *VerificationService, verificationID uint64, expected []Verification) {
	ok, err := s.Remove(verificationID)
	if !ok {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Failed to remove element %d", verificationID)
	}
	if err != nil {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected nil, got not nil")
	}
	if testListResult(s.allEntities).String() != testListResult(expected).String() {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected [%s], got [%s]", testListResult(expected), testListResult(s.allEntities))
	}
}

func (ch VerificationServiceChecker) CheckRemoveNotOK(s *VerificationService, verificationID uint64, expectedErr error) {
	ok, err := s.Remove(verificationID)
	if ok {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected false")
	}
	if err != expectedErr {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected [%s], got [%s]", expectedErr, err)
	}
}

func TestVerificationServiceDescribeOK(t *testing.T) {
	var allEntities = []Verification{
		{Title: "one"},
		{Title: "two"},
		{Title: "three"},
		{Title: "four"},
		{Title: "five"},
	}
	ch := VerificationServiceChecker{t: t}
	ch.CheckDescribeOK(VerificationService{allEntities: allEntities}, 0, allEntities[0])
	ch.CheckDescribeOK(VerificationService{allEntities: allEntities}, 1, allEntities[1])
	ch.CheckDescribeOK(VerificationService{allEntities: allEntities}, 4, allEntities[4])
}

func TestVerificationServiceDescribeNotOK(t *testing.T) {
	ch := VerificationServiceChecker{t: t}
	ch.CheckDescribeNotOK(VerificationService{}, 0, ErrEntityNotExists)
	ch.CheckDescribeNotOK(VerificationService{allEntities: []Verification{{Title: "one"}}}, 1, ErrEntityNotExists)
	ch.CheckDescribeNotOK(VerificationService{allEntities: []Verification{{Title: "one"}}}, 2, ErrEntityNotExists)
	ch.CheckDescribeNotOK(VerificationService{allEntities: []Verification{{Title: "one"}}}, 100500, ErrEntityNotExists)
}

func TestVerificationServiceListOK(t *testing.T) {
	var allEntities = []Verification{
		{Title: "one"},   // 0
		{Title: "two"},   // 1
		{Title: "three"}, // 2
		{Title: "four"},  // 3
		{Title: "five"},  // 4
	}
	ch := VerificationServiceChecker{t: t}
	ch.CheckListOK(VerificationService{allEntities: allEntities}, 0, 1, allEntities[0:1])
	ch.CheckListOK(VerificationService{allEntities: allEntities}, 0, 0, allEntities[0:0])
	ch.CheckListOK(VerificationService{allEntities: allEntities}, 1, 1, []Verification{{Title: "two"}})
	ch.CheckListOK(VerificationService{allEntities: allEntities}, 2, 1, []Verification{{Title: "three"}})
	ch.CheckListOK(VerificationService{allEntities: allEntities}, 2, 2, []Verification{{Title: "three"}, {Title: "four"}})
	ch.CheckListOK(VerificationService{allEntities: allEntities}, 0, uint64(len(allEntities)), allEntities)
	ch.CheckListOK(VerificationService{allEntities: allEntities}, 0, 100500, allEntities)
	ch.CheckListOK(VerificationService{allEntities: allEntities}, 0, 100000000000, allEntities)
	ch.CheckListOK(VerificationService{allEntities: allEntities}, 1, 10, allEntities[1:])
	ch.CheckListOK(VerificationService{allEntities: allEntities}, 2, 10, allEntities[2:])
}

func TestVerificationServiceListNotOK(t *testing.T) {
	ch := VerificationServiceChecker{t: t}
	ch.CheckListNotOK(VerificationService{}, 0, 1, ErrEntityNotExists)
	ch.CheckListNotOK(VerificationService{allEntities: []Verification{{Title: "three"}, {Title: "four"}}}, 100500, 1, ErrEntityNotExists)
	ch.CheckListNotOK(VerificationService{allEntities: []Verification{{Title: "three"}, {Title: "four"}}}, 3, 10, ErrEntityNotExists)
}

func TestVerificationServiceCopyResult(t *testing.T) {
	s := VerificationService{
		allEntities: []Verification{
			{Title: "one"},   // 0
			{Title: "two"},   // 1
			{Title: "three"}, // 2
			{Title: "four"},  // 3
			{Title: "five"},  // 4
		},
	}
	res, _ := s.List(0, 1)
	res[0] = Verification{Title: "changed-1"}
	if s.allEntities[0].Title != "one" {
		t.Errorf("List should return a copy")
	}
}

func TestVerificationServiceRemoveOK(t *testing.T) {
	ch := VerificationServiceChecker{t: t}

	//one by one
	s := VerificationService{
		allEntities: []Verification{
			{Title: "one"},   // 0
			{Title: "two"},   // 1
			{Title: "three"}, // 2
			{Title: "four"},  // 3
			{Title: "five"},  // 4
		},
	}
	ch.CheckRemoveOK(&s, 0, []Verification{
		{Title: "two"},   // 1
		{Title: "three"}, // 2
		{Title: "four"},  // 3
		{Title: "five"},  // 4
	})
	ch.CheckRemoveOK(&s, 0, []Verification{
		{Title: "three"}, // 2
		{Title: "four"},  // 3
		{Title: "five"},  // 4
	})
	ch.CheckRemoveOK(&s, 0, []Verification{
		{Title: "four"}, // 3
		{Title: "five"}, // 4
	})
	ch.CheckRemoveOK(&s, 0, []Verification{
		{Title: "five"}, // 4
	})
	ch.CheckRemoveOK(&s, 0, []Verification{})

	//random
	s = VerificationService{
		allEntities: []Verification{
			{Title: "one"},   // 0
			{Title: "two"},   // 1
			{Title: "three"}, // 2
			{Title: "four"},  // 3
			{Title: "five"},  // 4
		},
	}

	ch.CheckRemoveOK(&s, 3, []Verification{
		{Title: "one"},   // 0
		{Title: "two"},   // 1
		{Title: "three"}, // 2
		{Title: "five"},  // 4
	})

	ch.CheckRemoveOK(&s, 2, []Verification{
		{Title: "one"},  // 0
		{Title: "two"},  // 1
		{Title: "five"}, // 4
	})

	ch.CheckRemoveOK(&s, 1, []Verification{
		{Title: "one"},  // 0
		{Title: "five"}, // 4
	})

	ch.CheckRemoveOK(&s, 0, []Verification{
		{Title: "five"}, // 4
	})

	// from end
	s = VerificationService{
		allEntities: []Verification{
			{Title: "one"},   // 0
			{Title: "two"},   // 1
			{Title: "three"}, // 2
			{Title: "four"},  // 3
			{Title: "five"},  // 4
		},
	}

	ch.CheckRemoveOK(&s, 4, []Verification{
		{Title: "one"},   // 0
		{Title: "two"},   // 1
		{Title: "three"}, // 2
		{Title: "four"},  // 3
	})

	ch.CheckRemoveOK(&s, 3, []Verification{
		{Title: "one"},   // 0
		{Title: "two"},   // 1
		{Title: "three"}, // 2
	})

	ch.CheckRemoveOK(&s, 2, []Verification{
		{Title: "one"}, // 0
		{Title: "two"}, // 1
	})

	ch.CheckRemoveOK(&s, 1, []Verification{
		{Title: "one"}, // 0
	})
}

func TestVerificationServiceRemoveNotOK(t *testing.T) {
	ch := VerificationServiceChecker{t: t}

	ch.CheckRemoveNotOK(&VerificationService{}, 0, ErrEntityNotExists)
	ch.CheckRemoveNotOK(&VerificationService{}, 1, ErrEntityNotExists)
	ch.CheckRemoveNotOK(&VerificationService{allEntities: []Verification{}}, 0, ErrEntityNotExists)
	ch.CheckRemoveNotOK(&VerificationService{allEntities: []Verification{}}, 1, ErrEntityNotExists)
	ch.CheckRemoveNotOK(&VerificationService{
		allEntities: []Verification{
			{Title: "one"},   // 0
			{Title: "two"},   // 1
			{Title: "three"}, // 2
			{Title: "four"},  // 3
			{Title: "five"},  // 4
		},
	}, 5, ErrEntityNotExists)
}

func TestVerificationServiceCreate(t *testing.T) {
	s := VerificationService{}
	id, err := s.Create(Verification{Title: "new"})
	if id != 0 {
		t.Errorf("expected 0")
	}
	if err.Error() != "not implemented" {
		t.Errorf("expected not implemented")
	}
}
func TestVerificationServiceUpdate(t *testing.T) {
	s := VerificationService{}
	err := s.Update(0, Verification{"updated"})
	if err.Error() != "not implemented" {
		t.Errorf("expected not implemented")
	}
}
