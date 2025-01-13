package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		finish = []string{}
		dine()
		if len(finish) != 5 {
			t.Errorf("incorrect length of slice; expected 5 but got %d", len(finish))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		finish = []string{}

		eatTime = e.delay
		thinkTime = e.delay
		dine()
		if len(finish) != 5 {
			t.Errorf("%s: incorrect length of slice; expected 5 but got %d", e.name, len(finish))
		}
	}
}
