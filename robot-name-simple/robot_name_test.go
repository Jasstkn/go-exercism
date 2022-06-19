package robotname

import (
	"regexp"
	"testing"
)

var namePat = regexp.MustCompile(`^[A-Z]{2}\d{3}$`)
var seen = map[string]int{}

func New() *Robot { return new(Robot) }

// getName is a test helper function to facilitate optionally checking for seen
// robot names.
func (r *Robot) getName(t testing.TB, expectSeen bool) string {
	t.Helper()
	newName, err := r.Name()
	if err != nil {
		t.Fatalf("Name() returned unexpected error: %v", err)
	}
	if len(newName) != 5 {
		t.Fatalf("names should have 5 characters: name '%s' has %d character(s)", newName, len(newName))
	}

	_, chk := seen[newName]
	if !expectSeen && chk {
		t.Fatalf("Name %s reissued after %d robots.", newName, len(seen))
	}
	seen[newName] = 0
	return newName
}

func TestNameValid(t *testing.T) {
	n := New().getName(t, false)
	if !namePat.MatchString(n) {
		t.Errorf(`Invalid robot name %q, want form "AA###".`, n)
	}
}

func TestNameSticks(t *testing.T) {
	r := New()
	n1 := r.getName(t, false)
	n2 := r.getName(t, true)
	if n2 != n1 {
		t.Errorf(`Robot name changed.  Now %s, was %s.`, n2, n1)
	}
}

func TestSuccessiveRobotsHaveDifferentNames(t *testing.T) {
	n1 := New().getName(t, false)
	n2 := New().getName(t, false)
	if n1 == n2 {
		t.Errorf(`Robots with same name.  Two %s's.`, n1)
	}
}

func TestResetName(t *testing.T) {
	r := New()
	n1 := r.getName(t, false)
	r.Reset()
	if r.getName(t, false) == n1 {
		t.Errorf(`Robot name not cleared on reset.  Still %s.`, n1)
	}
}

// Test 1000 names are unique - this should run reasonably quickly even with a sub-optimal solution
// (e.g. pick a random name, then pick a new name if it's been seen before)
func TestMultipleNames(t *testing.T) {
	// Test uniqueness for new robots.
	for i := len(seen); i <= 1000; i++ {
		New().getName(t, false)
	}
}

var maxNames = 26 * 26 * 10 * 10 * 10

const lotsOfNames = 76000

// TestCollisions tests if unique names are generated by creating new robots until all names are used.
func TestCollisions(t *testing.T) {
	// Remove the next line to make this test run
	// t.Skip("skipping test as it can take a long time to run if solution is sub-optimal.")

	// Test uniqueness for new robots.
	for i := len(seen); i <= lotsOfNames; i++ {
		New().getName(t, false)
	}

	// Test that names aren't recycled either.
	// Note that this runs till names are exhausted.
	r := New()
	for i := len(seen); i < maxNames; i++ {
		r.Reset()
		r.getName(t, false)
	}

	// Test that name exhaustion is handled more or less correctly.
	_, err := New().Name()
	if err == nil {
		t.Fatalf("should return error if namespace is exhausted")
	}
}
