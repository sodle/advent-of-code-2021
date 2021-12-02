package shared

import "testing"

func TestReadVectorsFromFile(t *testing.T) {
	want := []Vector{
		{5, 0},
		{0, 5},
		{8, 0},
		{0, -3},
		{0, 8},
		{2, 0},
	}
	got := ReadVectorsFromFile("testdata/vectors.txt")

	if len(got) != len(want) {
		t.Errorf("vectors.txt length incorrect, got: %d, want: %d", len(got), len(want))
	}
	for i, s := range want {
		if s != got[i] {
			t.Errorf("vectors.txt unexpected value at line %d, got: (%d, %d), want: (%d, %d)", i, got[i].X, got[i].Y, s.X, s.Y)
		}
	}
}

func BenchmarkReadVectorsFromFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadVectorsFromFile("testdata/vectors.txt")
	}
}
