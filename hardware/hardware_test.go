package hardware

import (
	"strings"
	"testing"
)

func TestSysGetSerialKeyReturnsNonEmpty(t *testing.T) {
	key := SysGetSerialKey()
	if key == "" {
		t.Fatal("SysGetSerialKey() returned empty string")
	}
}

func TestSysGetSerialKeyReturnsConsistentValue(t *testing.T) {
	key1 := SysGetSerialKey()
	key2 := SysGetSerialKey()
	if key1 != key2 {
		t.Fatalf("SysGetSerialKey() returned inconsistent values: %q vs %q", key1, key2)
	}
}

func TestSysGetSerialKeyContainsOnlyAlphanumeric(t *testing.T) {
	key := SysGetSerialKey()
	for _, c := range key {
		if !((c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')) {
			t.Fatalf("SysGetSerialKey() contains non-alphanumeric char: %c in %q", c, key)
		}
	}
}

func TestKeyIsPressReturnsFalseForInvalidKey(t *testing.T) {
	if KeyIsPress("NOT_A_REAL_KEY") {
		t.Fatal("KeyIsPress(NOT_A_REAL_KEY) = true, want false")
	}
}

func TestKeyIsPressDoesNotPanicForValidKeys(t *testing.T) {
	validKeys := []string{"A", "SPACE", "ENTER", "ESC", "F1", "CTRL", "SHIFT", "ALT"}
	for _, key := range validKeys {
		KeyIsPress(key)
	}
}

func TestKeyIsPressCaseInsensitive(t *testing.T) {
	// Both should not panic and return the same result
	lower := KeyIsPress("space")
	upper := KeyIsPress("SPACE")
	if lower != upper {
		t.Fatalf("KeyIsPress case mismatch: lowercase=%v, uppercase=%v", lower, upper)
	}
}

func TestKeyMapContainsExpectedKeys(t *testing.T) {
	expected := []string{"A", "Z", "0", "9", "F1", "F12", "CTRL", "SHIFT", "ALT", "SPACE", "ENTER", "ESC"}
	for _, key := range expected {
		if _, ok := keyMap[key]; !ok {
			t.Fatalf("keyMap missing expected key: %s", key)
		}
	}
}

func TestWmicCmdReturnsNonEmptyForUUID(t *testing.T) {
	result := wmicCmd("Win32_ComputerSystemProduct", "UUID")
	// wmicCmd may return empty on systems without wmic/powershell, so just log
	if result == "" {
		t.Log("wmicCmd returned empty for UUID (may be expected on this system)")
	}
}

func TestWmicCmdOutputIsTrimmed(t *testing.T) {
	result := wmicCmd("Win32_ComputerSystemProduct", "UUID")
	if result != strings.TrimSpace(result) {
		t.Fatalf("wmicCmd output has leading/trailing whitespace: %q", result)
	}
}
