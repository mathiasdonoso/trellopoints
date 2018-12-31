package main

import (
	"testing"
)

func assertEquals(t *testing.T, spected interface{}, returned interface{}) {
	if spected != returned {
		t.Fatal("It should be", spected, "but instead it returns", returned)
	}
}

func TestHasStoryPoint(t *testing.T) {
	test1 := HasStoryPoint("Este es un texto sin puntos de historia")
	assertEquals(t, false, test1)

	test2 := HasStoryPoint("Este es un texto (esto no son puntos de historia) sin puntos de historia")
	assertEquals(t, false, test2)

	// test3 := HasStoryPoint("Este es un texto (13) sin puntos de historia")
	// assertEquals(t, false, test3)

	test4 := HasStoryPoint("Este es un texto con puntos de historia al final (5)")
	assertEquals(t, true, test4)

	test5 := HasStoryPoint("(5) Este es un texto con puntos de historia al comienzo")
	assertEquals(t, true, test5)
}

func TestGetStoryPoint(t *testing.T) {
	test1, err := GetStoryPoint("Texto sin puntos de historia")
	assertEquals(t, float64(0), test1)
	assertEquals(t, err, ErrDoesntHavePoints)

	test2, err := GetStoryPoint("Esta es una historia pivote (5)")
	assertEquals(t, float64(5), test2)
	assertEquals(t, err, nil)

	test3, err := GetStoryPoint("(3) Esta es otra historia")
	assertEquals(t, float64(3), test3)
	assertEquals(t, err, nil)

	test4, err := GetStoryPoint("(3) Esta es otra historia con puntos al principio (5)")
	assertEquals(t, float64(3), test4)
	assertEquals(t, err, nil)

	// test5, err := GetStoryPoint("Esta es otra historia (3) con puntos al final (5)")
	// assertEquals(t, float64(5), test5)
	// assertEquals(t, err, nil)
}
