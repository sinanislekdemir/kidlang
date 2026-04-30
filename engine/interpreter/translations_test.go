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
	if got := getTranslation("CONTAINS"); got != "ICINDEVAR" {
		t.Errorf("getTranslation('CONTAINS') = %s; want 'ICINDEVAR'", got)
	}

	activeLanguage = LANG_DE
	if got := getTranslation("REPLACE"); got != "ERSETZE" {
		t.Errorf("getTranslation('REPLACE') = %s; want 'ERSETZE'", got)
	}

	activeLanguage = LANG_FI
	if got := getTranslation("STACKLEN"); got != "LISTANPITUUS" {
		t.Errorf("getTranslation('STACKLEN') = %s; want 'LISTANPITUUS'", got)
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
	if got := isReservedKeyword("ICINDEVAR"); !got {
		t.Errorf("isReservedKeyword('ICINDEVAR') = %v; want true", got)
	}

	activeLanguage = LANG_DE
	if got := isReservedKeyword("ERSETZE"); !got {
		t.Errorf("isReservedKeyword('ERSETZE') = %v; want true", got)
	}

	activeLanguage = LANG_FI
	if got := isReservedKeyword("LISTANPITUUS"); !got {
		t.Errorf("isReservedKeyword('LISTANPITUUS') = %v; want true", got)
	}

	if got := isReservedKeyword("UNKNOWN"); got {
		t.Errorf("isReservedKeyword('UNKNOWN') = %v; want false", got)
	}

	// Reset activeLanguage to English
	activeLanguage = LANG_EN
}
