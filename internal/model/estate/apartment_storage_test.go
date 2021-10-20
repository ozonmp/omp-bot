package estate

import (
	"reflect"
	"testing"
)

func TestEmptyInMemoryApartmentStorage(t *testing.T) {
	s := NewEmptyInMemoryApartmentStorage()

	apartments, err := s.List(0, 0)
	if err != nil {
		t.Errorf(".List() of empty storage: unexpected error = %v", err)
		return
	}
	if apartments != nil {
		t.Errorf(".List() of empty storage: expected to get nil list of apartments, got = %v", apartments)
		return
	}

	apartments, err = s.List(0, 5)
	if err != nil {
		t.Errorf(".List() of empty storage: unexpected error = %v", err)
		return
	}
	if apartments != nil {
		t.Errorf(".List() of empty storage: expected to get nil list of apartments, got = %v", apartments)
		return
	}

	apartments, err = s.List(10, 5)
	if err != nil {
		t.Errorf(".List() of empty storage: unexpected error = %v", err)
		return
	}
	if apartments != nil {
		t.Errorf(".List() of empty storage: expected to get nil list of apartments, got = %v", apartments)
		return
	}

	apartment, err := s.Describe(3)
	if err == nil {
		t.Errorf(".Describe() on empty storage: error was expected, got = nil")
		return
	}
	if apartment != nil {
		t.Errorf(".Describe() on empty storage: nil apartment was expected, got = %v", apartment)
		return
	}

	err = s.Update(1, Apartment{Title: "ap1", Price: 10})
	if err == nil {
		t.Errorf(".Update() on empty storage: error was expected, got = nil")
		return
	}
	apartment, err = s.Describe(1)
	if err == nil {
		t.Errorf(".Describe() after failed update: error was expected, got = nil")
		return
	}
	if apartment != nil {
		t.Errorf(".Describe() after failed update: nil apartment was expected, got = %v", apartment)
		return
	}

	ok, err := s.Remove(1)
	if err != nil {
		t.Errorf(".Remove() on empty storage: unexpected error = %v", err)
		return
	}
	if ok {
		t.Errorf(".Remove() on empty storage: expected not ok, got = %v", ok)
		return
	}
}

func TestEmptyInMemoryApartmentStorageWithCreate(t *testing.T) {
	s := NewEmptyInMemoryApartmentStorage()

	apartment0 := Apartment{ID: 0, Title: "ap0", Price: 10}
	id, err := s.Create(apartment0)
	if err != nil {
		t.Errorf(".Create() on empty storage: unexpected error = %v", err)
		return
	}
	if id != 0 {
		t.Errorf(".Create() on empty storage: expected new apartment id = %v", id)
		return
	}

	apartment, err := s.Describe(0)
	if err != nil {
		t.Errorf(".Describe() of existing element: unexpected error = %v", err)
		return
	}
	if apartment == nil {
		t.Errorf(".Describe() of existing element: non-nil apartment was expected, got = nil")
		return
	}
	if *apartment != apartment0 {
		t.Errorf(".Describe() of existing element: got wrong value, expected = %v, got = %v", apartment0, *apartment)
		return
	}

	apartment1 := Apartment{ID: 1, Title: "ap1", Price: 13}
	apartment1WithWrongID := apartment1
	apartment1WithWrongID.ID = 29

	id, err = s.Create(apartment1WithWrongID)
	if err != nil {
		t.Errorf(".Create() supply element with wrong id: unexpected error = %v", err)
		return
	}
	if id != 1 {
		t.Errorf(".Create() supply element with wrong id: expected new apartment id = %v", id)
		return
	}

	apartment, err = s.Describe(1)
	if err != nil {
		t.Errorf(".Describe() of existing element: unexpected error = %v", err)
		return
	}
	if apartment == nil {
		t.Errorf(".Describe() of existing element: non-nil apartment was expected, got = nil")
		return
	}
	if *apartment != apartment1 {
		t.Errorf(".Describe() of existing element: got wrong value, expected = %v, got = %v", apartment1, *apartment)
		return
	}

	expectedApartments := []Apartment{
		apartment0,
		apartment1,
	}
	apartments, err := s.List(0, 5)
	if err != nil {
		t.Errorf(".List() two elements: unexpected error = %v", err)
		return
	}
	if !reflect.DeepEqual(apartments, expectedApartments) {
		t.Errorf(".List() two elements: got wrong value, expected = %v, got = %v", expectedApartments, apartments)
		return
	}

	apartment2 := Apartment{ID: 2, Title: "ap2", Price: 16}
	id, err = s.Create(apartment2)
	if err != nil {
		t.Errorf(".Create() third element: unexpected error = %v", err)
		return
	}
	if id != 2 {
		t.Errorf(".Create() third element: expected new apartment id = %v", id)
		return
	}

	expectedApartments = []Apartment{
		apartment0,
		apartment1,
	}
	apartments, err = s.List(0, 2)
	if err != nil {
		t.Errorf(".List() two elements: unexpected error = %v", err)
		return
	}
	if !reflect.DeepEqual(apartments, expectedApartments) {
		t.Errorf(".List() two elements: got wrong value, expected = %v, got = %v", expectedApartments, apartments)
		return
	}

	expectedApartments = []Apartment{
		apartment1,
	}
	apartments, err = s.List(1, 1)
	if err != nil {
		t.Errorf(".List() one element in the middle: unexpected error = %v", err)
		return
	}
	if !reflect.DeepEqual(apartments, expectedApartments) {
		t.Errorf(".List() one element in the middle: got wrong value, expected = %v, got = %v", expectedApartments, apartments)
		return
	}

	updatedApartment1 := Apartment{ID: 1, Title: "another_ap1", Price: 199}
	updatedApartment1WithWrongID := updatedApartment1
	updatedApartment1WithWrongID.ID = 74
	err = s.Update(1, updatedApartment1WithWrongID)
	if err != nil {
		t.Errorf(".Update() existing element: unepected error = %v", err)
		return
	}
	apartment, err = s.Describe(1)
	if err != nil {
		t.Errorf(".Describe() of updated element: unexpected error = %v", err)
		return
	}
	if apartment == nil {
		t.Errorf(".Describe() of updated element: non-nil apartment was expected, got = nil")
		return
	}
	if *apartment != updatedApartment1 {
		t.Errorf(".Describe() of updated element: got wrong value, expected = %v, got = %v", updatedApartment1, *apartment)
		return
	}

	ok, err := s.Remove(1)
	if err != nil {
		t.Errorf(".Remove() existing element: unexpected error = %v", err)
		return
	}
	if !ok {
		t.Errorf(".Remove() existing element: expected ok, got = %v", ok)
		return
	}

	expectedApartments = []Apartment{
		apartment0,
		apartment2,
	}
	apartments, err = s.List(0, 3)
	if err != nil {
		t.Errorf(".List() two elements after remove: unexpected error = %v", err)
		return
	}
	if !reflect.DeepEqual(apartments, expectedApartments) {
		t.Errorf(".List() two elements after remove: got wrong value, expected = %v, got = %v", expectedApartments, apartments)
		return
	}
}

func TestInMemoryApartmentStorageWithStartingItems(t *testing.T) {
	apartment0 := Apartment{Title: "ap0", Price: 5}
	apartment1 := Apartment{Title: "ap1", Price: 13}
	apartment2 := Apartment{Title: "ap2", Price: 6}
	apartment3 := Apartment{Title: "ap3", Price: 17}
	apartment4 := Apartment{Title: "ap4", Price: 10}
	apartment5 := Apartment{Title: "ap5", Price: 10}
	apartment6 := Apartment{Title: "ap6", Price: 154}

	startingItems := []Apartment{
		apartment0,
		apartment1,
		apartment2,
		apartment3,
		apartment4,
		apartment5,
		apartment6,
	}

	apartment0.ID = 0
	apartment1.ID = 1
	apartment2.ID = 2
	apartment3.ID = 3
	apartment4.ID = 4
	apartment5.ID = 5
	apartment6.ID = 6

	s := NewInMemoryApartmentStorage(startingItems)

	apartments, err := s.List(10, 5)
	if err != nil {
		t.Errorf(".List() outside of range: unexpected error = %v", err)
		return
	}
	if apartments != nil {
		t.Errorf(".List() outside of range: expected to get nil list of apartments, got = %v", apartments)
		return
	}

	expectedApartments := []Apartment{
		apartment0,
		apartment1,
		apartment2,
		apartment3,
		apartment4,
		apartment5,
		apartment6,
	}
	apartments, err = s.List(0, 10)
	if err != nil {
		t.Errorf(".List() all elements by exceeding range: unexpected error = %v", err)
		return
	}
	if !reflect.DeepEqual(apartments, expectedApartments) {
		t.Errorf(".List() all elements by exceeding range: got wrong value, expected = %v, got = %v", expectedApartments, apartments)
		return
	}

	apartments, err = s.List(0, 7)
	if err != nil {
		t.Errorf(".List() exactly all elements: unexpected error = %v", err)
		return
	}
	if !reflect.DeepEqual(apartments, expectedApartments) {
		t.Errorf(".List() exactly all elements: got wrong value, expected = %v, got = %v", expectedApartments, apartments)
		return
	}

	expectedApartments = []Apartment{
		apartment2,
		apartment3,
		apartment4,
	}
	apartments, err = s.List(2, 3)
	if err != nil {
		t.Errorf(".List() elements from the middle: unexpected error = %v", err)
		return
	}
	if !reflect.DeepEqual(apartments, expectedApartments) {
		t.Errorf(".List() elements from the middle: got wrong value, expected = %v, got = %v", expectedApartments, apartments)
		return
	}

	ok, err := s.Remove(3)
	if err != nil {
		t.Errorf(".Remove() existing element: unexpected error = %v", err)
		return
	}
	if !ok {
		t.Errorf(".Remove() existing element: expected ok, got = %v", ok)
		return
	}

	ok, err = s.Remove(3)
	if err != nil {
		t.Errorf(".Remove() after removing before: unexpected error = %v", err)
		return
	}
	if ok {
		t.Errorf(".Remove() after removing before: expected not ok, got = %v", ok)
		return
	}

	ok, err = s.Remove(4)
	if err != nil {
		t.Errorf(".Remove() existing element: unexpected error = %v", err)
		return
	}
	if !ok {
		t.Errorf(".Remove() existing element: expected ok, got = %v", ok)
		return
	}

	expectedApartments = []Apartment{
		apartment5,
		apartment6,
	}
	apartments, err = s.List(3, 3)
	if err != nil {
		t.Errorf(".List() starting from non-existing element: unexpected error = %v", err)
		return
	}
	if !reflect.DeepEqual(apartments, expectedApartments) {
		t.Errorf(".List() starting from non-existing element: got wrong value, expected = %v, got = %v", expectedApartments, apartments)
		return
	}
}
