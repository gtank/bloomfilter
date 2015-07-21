package bloomfilter

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestFalseNegatives(t *testing.T) {
	var trials = 1000000
	var rate = 16384
	var counter = 0

	set := NewBloomFilter(trials, rate)

	// Add random data to set
	for i := 0; i < trials; i++ {
		c := 8
		b := make([]byte, c)
		_, err := rand.Read(b)
		if err != nil {
			t.Fatalf("Could not generate random string")
			return
		}
		set.Add(b)
		if set.Check(b) {
			counter++
		}
	}

	if trials != counter {
		t.Fatalf("Lost some random set entries!")
	}

}

func TestFalsePositives(t *testing.T) {
	var trials = 1000000
	var rate = 16384
	var counter = 0

	set := NewBloomFilter(trials, rate)

	// Add random data to set
	for i := 0; i < trials; i++ {
		c := 8
		b := make([]byte, c)
		_, err := rand.Read(b)
		if err != nil {
			t.Fatalf("Could not generate random string")
			return
		}
		set.Add(b)
	}

	// Check for different random data
	for i := 0; i < trials; i++ {
		c := 8
		b := make([]byte, c)
		_, err := rand.Read(b)
		if err != nil {
			t.Fatalf("Could not generate random string")
			return
		}
		if set.Check(b) {
			counter++
		}
	}

	expected := float64(1) / float64(rate)
	actual := float64(counter) / float64(trials)

	fmt.Println("expected error rate: ", expected)
	fmt.Println("actual error rate: ", actual)
}
