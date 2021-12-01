package shared

import "testing"

func TestReadLinesFromFile(t *testing.T) {
	want := []string{
		"hello",
		"world",
		"blank space",
	}
	got := ReadLinesFromFile("testdata/lines.txt")

	if len(got) != len(want) {
		t.Errorf("lines.txt length incorrect, got: %d, want: %d", len(got), len(want))
	}
	for i, s := range want {
		if s != got[i] {
			t.Errorf("lines.txt unexpected number at line %d, got: %s, want: %s", i, got[i], s)
		}
	}
}

func TestReadNumberFile(t *testing.T) {
	want := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	got := ReadNumberFile("testdata/numbers.txt")
	gotTrailingNewline := ReadNumberFile("testdata/numbers_trailing_newline.txt")

	if len(got) != len(want) {
		t.Errorf("numbers.txt length incorrect, got: %d, want: %d", len(got), len(want))
	}
	for i, n := range want {
		if n != got[i] {
			t.Errorf("numbers.txt unexpected number at line %d, got: %d, want: %d", i, got[i], n)
		}
	}

	if len(gotTrailingNewline) != len(want) {
		t.Errorf("numbers_trailing_newline.txt length incorrect, got: %d, want: %d", len(gotTrailingNewline), len(want))
	}
	for i, n := range want {
		if n != gotTrailingNewline[i] {
			t.Errorf("numbers_trailing_newline.txt unexpected number at line %d, got: %d, want: %d", i, gotTrailingNewline[i], n)
		}
	}
}

func BenchmarkReadLinesFromFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadLinesFromFile("testdata/numbers.txt")
	}
}

func BenchmarkReadNumberFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadNumberFile("testdata/numbers.txt")
	}
}
