package candgo

import "testing"

func TestTrue(t *testing.T) {
	if !True() {
		t.Fatalf("True() : expected %v got %v", true, True())
	}
}

func TestMax(t *testing.T) {
	a, b := 56, 108
	if Max(a, b) != b {
		t.Fatalf("Max(%v,%v) : expected %v got %v", a, b, b, Max(a, b))
	}
}
func TestInc(t *testing.T) {
	v := 201
	if Inc(&v); v != 202 {
		t.Fatalf("Inc(201): expected 202,got %v", v)
	}
}
