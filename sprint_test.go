package main

import "testing"

func TestNameWithoutPoints(t *testing.T) {
	test1 := NameWithoutPoints("Este es un texto sin puntos de historia")
	assertEquals(t, "Este es un texto sin puntos de historia", test1)

	test2 := NameWithoutPoints("Este es un texto con puntos de historia al final (5)")
	assertEquals(t, "Este es un texto con puntos de historia al final", test2)

	test3 := NameWithoutPoints("(3) Este es un texto con puntos de historia al comienzo")
	assertEquals(t, "Este es un texto con puntos de historia al comienzo", test3)
}
