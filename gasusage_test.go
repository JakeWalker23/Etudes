package main

import ("testing")

func TestAdd(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}



