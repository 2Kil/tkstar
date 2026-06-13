package tkstar

import (
	"errors"
	"log"
	"os"
	"strings"
	"testing"
)

func TestRandAtomicReturnsZeroForNonPositiveMax(t *testing.T) {
	if got := RandAtomic(0); got != 0 {
		t.Fatalf("RandAtomic(0) = %d, want 0", got)
	}
	if got := RandAtomic(-5); got != 0 {
		t.Fatalf("RandAtomic(-5) = %d, want 0", got)
	}
}

func TestRandAtomicReturnsValueInRange(t *testing.T) {
	for i := 0; i < 100; i++ {
		got := RandAtomic(10)
		if got < 0 || got >= 10 {
			t.Fatalf("RandAtomic(10) = %d, out of range [0, 10)", got)
		}
	}
}

func TestRandAtomicRadiusReturnsMinWhenMaxLessThanMin(t *testing.T) {
	if got := RandAtomicRadius(5, 3); got != 5 {
		t.Fatalf("RandAtomicRadius(5, 3) = %d, want 5", got)
	}
}

func TestRandAtomicRadiusReturnsValueInRange(t *testing.T) {
	for i := 0; i < 100; i++ {
		got := RandAtomicRadius(3, 7)
		if got < 3 || got > 7 {
			t.Fatalf("RandAtomicRadius(3, 7) = %d, out of range [3, 7]", got)
		}
	}
}

func TestRandAtomicRadiusHandlesEqualMinMax(t *testing.T) {
	if got := RandAtomicRadius(5, 5); got != 5 {
		t.Fatalf("RandAtomicRadius(5, 5) = %d, want 5", got)
	}
}

func TestCheckErrReturnsTrueForNil(t *testing.T) {
	if !CheckErr(nil) {
		t.Fatal("CheckErr(nil) = false, want true")
	}
}

func TestCheckErrReturnsFalseForError(t *testing.T) {
	if CheckErr(errors.New("test error")) {
		t.Fatal("CheckErr(error) = true, want false")
	}
}

func TestCheckErrLogsCustomPrefix(t *testing.T) {
	var buf strings.Builder
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	CheckErr(errors.New("fail"), "CustomPrefix")
	if !strings.Contains(buf.String(), "CustomPrefix") {
		t.Fatalf("expected log to contain 'CustomPrefix', got: %s", buf.String())
	}
}

func TestHelperRemoveDuplicatesRemovesDuplicates(t *testing.T) {
	input := []string{"a", "b", "a", "c", "b"}
	got := HelperRemoveDuplicates(input)
	expected := []string{"a", "b", "c"}
	if len(got) != len(expected) {
		t.Fatalf("len = %d, want %d; got %v", len(got), len(expected), got)
	}
	for i, v := range got {
		if v != expected[i] {
			t.Fatalf("got[%d] = %q, want %q", i, v, expected[i])
		}
	}
}

func TestHelperRemoveDuplicatesRemovesEmptyStrings(t *testing.T) {
	input := []string{"a", "", "b", "", ""}
	got := HelperRemoveDuplicates(input)
	expected := []string{"a", "b"}
	if len(got) != len(expected) {
		t.Fatalf("len = %d, want %d; got %v", len(got), len(expected), got)
	}
}

func TestHelperRemoveDuplicatesHandlesEmptySlice(t *testing.T) {
	got := HelperRemoveDuplicates([]string{})
	if len(got) != 0 {
		t.Fatalf("len = %d, want 0", len(got))
	}
}

func TestHelperRemoveDuplicatesHandlesAllEmpty(t *testing.T) {
	got := HelperRemoveDuplicates([]string{"", "", ""})
	if len(got) != 0 {
		t.Fatalf("len = %d, want 0; got %v", len(got), got)
	}
}

func TestBuildTimeReturnsNonEmptyString(t *testing.T) {
	got := BuildTime()
	if got == "" {
		t.Fatal("BuildTime() returned empty string")
	}
}
