package screen

import (
	"testing"
)

func TestScreenGetTextReturnsEmptyByDefault(t *testing.T) {
	// Reset state
	mu.Lock()
	displayText = ""
	mu.Unlock()

	if got := ScreenGetText(); got != "" {
		t.Fatalf("ScreenGetText() = %q, want empty", got)
	}
}

func TestScreenUpdateTextSetsText(t *testing.T) {
	ScreenUpdateText("hello world")
	defer ScreenUpdateText("")

	if got := ScreenGetText(); got != "hello world" {
		t.Fatalf("ScreenGetText() = %q, want hello world", got)
	}
}

func TestScreenUpdateTextOverwritesPrevious(t *testing.T) {
	ScreenUpdateText("first")
	ScreenUpdateText("second")
	defer ScreenUpdateText("")

	if got := ScreenGetText(); got != "second" {
		t.Fatalf("ScreenGetText() = %q, want second", got)
	}
}

func TestScreenUpdateTextWithEmptyString(t *testing.T) {
	ScreenUpdateText("something")
	ScreenUpdateText("")

	if got := ScreenGetText(); got != "" {
		t.Fatalf("ScreenGetText() = %q, want empty", got)
	}
}

func TestScreenUpdateTextWithUnicode(t *testing.T) {
	ScreenUpdateText("你好世界")
	defer ScreenUpdateText("")

	if got := ScreenGetText(); got != "你好世界" {
		t.Fatalf("ScreenGetText() = %q, want 你好世界", got)
	}
}

func TestScreenGetTextIsThreadSafe(t *testing.T) {
	ScreenUpdateText("concurrent")
	defer ScreenUpdateText("")

	done := make(chan string, 10)
	for i := 0; i < 10; i++ {
		go func() {
			done <- ScreenGetText()
		}()
	}
	for i := 0; i < 10; i++ {
		got := <-done
		if got != "concurrent" {
			t.Fatalf("goroutine got %q, want concurrent", got)
		}
	}
}
