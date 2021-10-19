package testsubdomain

import (
	"runtime/debug"
	"testing"
)

type TestSubdomainServiceChecker struct {
	t *testing.T
}

func (ch TestSubdomainServiceChecker) CheckDescribeOK(s TestSubdomainService, id uint64, expected TestSubdomain) {
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

func (ch TestSubdomainServiceChecker) CheckDescribeNotOK(s TestSubdomainService, id uint64, expectedErr error) {
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

type testListResult []TestSubdomain

func (t testListResult) String() string {
	res := ""
	for _, v := range t {
		res += "|" + v.String() + "|"
	}
	return res
}

func (ch TestSubdomainServiceChecker) CheckListOK(s TestSubdomainService, cursor uint64, limit uint64, expected testListResult) {
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

func (ch TestSubdomainServiceChecker) CheckListNotOK(s TestSubdomainService, cursor uint64, limit uint64, expectedErr error) {
	res, _ := s.List(cursor, limit)
	if res != nil {
		ch.t.Log(string(debug.Stack()))
		ch.t.Errorf("Expected nil, got not nil")
	}
}

func (ch TestSubdomainServiceChecker) CheckRemoveOK(s *TestSubdomainService, verificationID uint64, expected []TestSubdomain) {
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

func (ch TestSubdomainServiceChecker) CheckRemoveNotOK(s *TestSubdomainService, verificationID uint64, expectedErr error) {
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

func TestSubdomainServiceDescribeOK(t *testing.T) {
	var allEntities = []TestSubdomain{
		{Title: "one"},
		{Title: "two"},
		{Title: "three"},
		{Title: "four"},
		{Title: "five"},
	}
	ch := TestSubdomainServiceChecker{t: t}
	ch.CheckDescribeOK(TestSubdomainService{allEntities: allEntities}, 0, allEntities[0])
	ch.CheckDescribeOK(TestSubdomainService{allEntities: allEntities}, 1, allEntities[1])
	ch.CheckDescribeOK(TestSubdomainService{allEntities: allEntities}, 4, allEntities[4])
}

func TestSubdomainServiceDescribeNotOK(t *testing.T) {
	ch := TestSubdomainServiceChecker{t: t}
	ch.CheckDescribeNotOK(TestSubdomainService{}, 0, ErrEntityNotExists)
	ch.CheckDescribeNotOK(TestSubdomainService{allEntities: []TestSubdomain{{Title: "one"}}}, 1, ErrEntityNotExists)
	ch.CheckDescribeNotOK(TestSubdomainService{allEntities: []TestSubdomain{{Title: "one"}}}, 2, ErrEntityNotExists)
	ch.CheckDescribeNotOK(TestSubdomainService{allEntities: []TestSubdomain{{Title: "one"}}}, 100500, ErrEntityNotExists)
}

func TestSubdomainServiceListOK(t *testing.T) {
	var allEntities = []TestSubdomain{
		{Title: "one"},   // 0
		{Title: "two"},   // 1
		{Title: "three"}, // 2
		{Title: "four"},  // 3
		{Title: "five"},  // 4
	}
	ch := TestSubdomainServiceChecker{t: t}
	ch.CheckListOK(TestSubdomainService{allEntities: allEntities}, 0, 1, allEntities[0:1])
	ch.CheckListOK(TestSubdomainService{allEntities: allEntities}, 0, 0, allEntities[0:0])
	ch.CheckListOK(TestSubdomainService{allEntities: allEntities}, 1, 1, []TestSubdomain{{Title: "two"}})
	ch.CheckListOK(TestSubdomainService{allEntities: allEntities}, 2, 1, []TestSubdomain{{Title: "three"}})
	ch.CheckListOK(TestSubdomainService{allEntities: allEntities}, 2, 2, []TestSubdomain{{Title: "three"}, {Title: "four"}})
	ch.CheckListOK(TestSubdomainService{allEntities: allEntities}, 0, uint64(len(allEntities)), allEntities)
	ch.CheckListOK(TestSubdomainService{allEntities: allEntities}, 0, 100500, allEntities)
	ch.CheckListOK(TestSubdomainService{allEntities: allEntities}, 0, 100000000000, allEntities)
	ch.CheckListOK(TestSubdomainService{allEntities: allEntities}, 1, 10, allEntities[1:])
	ch.CheckListOK(TestSubdomainService{allEntities: allEntities}, 2, 10, allEntities[2:])
}

func TestSubdomainServiceListNotOK(t *testing.T) {
	ch := TestSubdomainServiceChecker{t: t}
	ch.CheckListNotOK(TestSubdomainService{}, 0, 1, ErrEntityNotExists)
	ch.CheckListNotOK(TestSubdomainService{allEntities: []TestSubdomain{{Title: "three"}, {Title: "four"}}}, 100500, 1, ErrEntityNotExists)
	ch.CheckListNotOK(TestSubdomainService{allEntities: []TestSubdomain{{Title: "three"}, {Title: "four"}}}, 3, 10, ErrEntityNotExists)
}

func TestSubdomainServiceCopyResult(t *testing.T) {
	s := TestSubdomainService{
		allEntities: []TestSubdomain{
			{Title: "one"},   // 0
			{Title: "two"},   // 1
			{Title: "three"}, // 2
			{Title: "four"},  // 3
			{Title: "five"},  // 4
		},
	}
	res, _ := s.List(0, 1)
	res[0] = TestSubdomain{Title: "changed-1"}
	if s.allEntities[0].Title != "one" {
		t.Errorf("List should return a copy")
	}
}

func TestSubdomainServiceRemoveOK(t *testing.T) {
	ch := TestSubdomainServiceChecker{t: t}

	//one by one
	s := TestSubdomainService{
		allEntities: []TestSubdomain{
			{Title: "one"},   // 0
			{Title: "two"},   // 1
			{Title: "three"}, // 2
			{Title: "four"},  // 3
			{Title: "five"},  // 4
		},
	}
	ch.CheckRemoveOK(&s, 0, []TestSubdomain{
		{Title: "two"},   // 1
		{Title: "three"}, // 2
		{Title: "four"},  // 3
		{Title: "five"},  // 4
	})
	ch.CheckRemoveOK(&s, 0, []TestSubdomain{
		{Title: "three"}, // 2
		{Title: "four"},  // 3
		{Title: "five"},  // 4
	})
	ch.CheckRemoveOK(&s, 0, []TestSubdomain{
		{Title: "four"}, // 3
		{Title: "five"}, // 4
	})
	ch.CheckRemoveOK(&s, 0, []TestSubdomain{
		{Title: "five"}, // 4
	})
	ch.CheckRemoveOK(&s, 0, []TestSubdomain{})

	//random
	s = TestSubdomainService{
		allEntities: []TestSubdomain{
			{Title: "one"},   // 0
			{Title: "two"},   // 1
			{Title: "three"}, // 2
			{Title: "four"},  // 3
			{Title: "five"},  // 4
		},
	}

	ch.CheckRemoveOK(&s, 3, []TestSubdomain{
		{Title: "one"},   // 0
		{Title: "two"},   // 1
		{Title: "three"}, // 2
		{Title: "five"},  // 4
	})

	ch.CheckRemoveOK(&s, 2, []TestSubdomain{
		{Title: "one"},  // 0
		{Title: "two"},  // 1
		{Title: "five"}, // 4
	})

	ch.CheckRemoveOK(&s, 1, []TestSubdomain{
		{Title: "one"},  // 0
		{Title: "five"}, // 4
	})

	ch.CheckRemoveOK(&s, 0, []TestSubdomain{
		{Title: "five"}, // 4
	})

	// from end
	s = TestSubdomainService{
		allEntities: []TestSubdomain{
			{Title: "one"},   // 0
			{Title: "two"},   // 1
			{Title: "three"}, // 2
			{Title: "four"},  // 3
			{Title: "five"},  // 4
		},
	}

	ch.CheckRemoveOK(&s, 4, []TestSubdomain{
		{Title: "one"},   // 0
		{Title: "two"},   // 1
		{Title: "three"}, // 2
		{Title: "four"},  // 3
	})

	ch.CheckRemoveOK(&s, 3, []TestSubdomain{
		{Title: "one"},   // 0
		{Title: "two"},   // 1
		{Title: "three"}, // 2
	})

	ch.CheckRemoveOK(&s, 2, []TestSubdomain{
		{Title: "one"}, // 0
		{Title: "two"}, // 1
	})

	ch.CheckRemoveOK(&s, 1, []TestSubdomain{
		{Title: "one"}, // 0
	})
}

func TestSubdomainServiceRemoveNotOK(t *testing.T) {
	ch := TestSubdomainServiceChecker{t: t}

	ch.CheckRemoveNotOK(&TestSubdomainService{}, 0, ErrEntityNotExists)
	ch.CheckRemoveNotOK(&TestSubdomainService{}, 1, ErrEntityNotExists)
	ch.CheckRemoveNotOK(&TestSubdomainService{allEntities: []TestSubdomain{}}, 0, ErrEntityNotExists)
	ch.CheckRemoveNotOK(&TestSubdomainService{allEntities: []TestSubdomain{}}, 1, ErrEntityNotExists)
	ch.CheckRemoveNotOK(&TestSubdomainService{
		allEntities: []TestSubdomain{
			{Title: "one"},   // 0
			{Title: "two"},   // 1
			{Title: "three"}, // 2
			{Title: "four"},  // 3
			{Title: "five"},  // 4
		},
	}, 5, ErrEntityNotExists)
}

func TestSubdomainServiceCreate(t *testing.T) {
	s := TestSubdomainService{}
	id, err := s.Create(TestSubdomain{Title: "new"})
	if id != 0 {
		t.Errorf("expected 0")
	}
	if err.Error() != "not implemented" {
		t.Errorf("expected not implemented")
	}
}
func TestSubdomainServiceUpdate(t *testing.T) {
	s := TestSubdomainService{}
	err := s.Update(0, TestSubdomain{"updated"})
	if err.Error() != "not implemented" {
		t.Errorf("expected not implemented")
	}
}
