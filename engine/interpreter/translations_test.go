package interpreter

import (
	"testing"
)

func TestGetTranslation(t *testing.T) {
	// Test case for English language
	activeLanguage = LANG_EN
	if got := getTranslation("BOX"); got != "BOX" {
		t.Errorf("getTranslation('BOX') = %s; want 'BOX'", got)
	}

	// Test case for Turkish language
	activeLanguage = LANG_TR
	if got := getTranslation("BOX"); got != "KUTU" {
		t.Errorf("getTranslation('BOX') = %s; want 'KUTU'", got)
	}
	// activeLanguage = LANG_EN

	// Test case for missing translation
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("getTranslation('MISSING') did not panic")
		}
		activeLanguage = LANG_EN
	}()
	getTranslation("MISSING")
}

func TestIsReservedKeyword(t *testing.T) {
	// Test case for English language
	activeLanguage = LANG_EN
	if got := isReservedKeyword("BOX"); !got {
		t.Errorf("isReservedKeyword('BOX') = %v; want true", got)
	}
	if got := isReservedKeyword("UNKNOWN"); got {
		t.Errorf("isReservedKeyword('UNKNOWN') = %v; want false", got)
	}

	// Test case for Turkish language
	activeLanguage = LANG_TR
	if got := isReservedKeyword("KUTU"); !got {
		t.Errorf("isReservedKeyword('KUTU') = %v; want true", got)
	}

	if got := isReservedKeyword("UNKNOWN"); got {
		t.Errorf("isReservedKeyword('UNKNOWN') = %v; want false", got)
	}

	// Reset activeLanguage to English
	activeLanguage = LANG_EN
}
