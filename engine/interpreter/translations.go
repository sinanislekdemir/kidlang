package interpreter

import (
	"fmt"
	"strings"
)

const (
	LANG_EN = "en"
	LANG_TR = "tr"
	LANG_FI = "fi"
	LANG_DE = "de"
)

var translations = createTranslationMap()
var activeLanguage = LANG_EN

func isReservedKeyword(word string) bool {
	if activeLanguage == LANG_EN {
		// reserved keywords in English, check if word exists in translations keys
		_, exists := translations["tr"][strings.ToUpper(word)]
		return exists
	}

	for keyword := range translations[activeLanguage] {
		if translations[activeLanguage][keyword] == strings.ToUpper(word) {
			return true
		}
	}
	return false

}

func createTranslationMap() map[string]map[string]string {
	nestedLang := make(map[string]map[string]string)

	nestedLang["tr"] = map[string]string{
		"BOX":      "KUTU",
		"FILE":     "DOSYA",
		"STACK":    "LISTE",
		"PORT":     "PORT", // reserved for later use
		"PRINT":    "YAZ",
		"ASK":      "SOR",
		"GOTO":     "GIT",
		"PUT":      "KOY",
		"RANDOM":   "RASTGELE",
		"NOW":      "TARIH",
		"ANSWER":   "CEVAP",
		"IF":       "EGER",
		"THEN":     "ISE",
		"END":      "SON",
		"SQRT":     "KAREKOK",
		"ABS":      "MOD",
		"SQR":      "KARE",
		"ASSIGN":   "ATAMA",
		"EXEC":     "CALISTIR",
		"CLOSE":    "KAPAT",
		"SLEEP":    "BEKLE",
		"OPEN":     "AC",
		"READ":     "OKU",
		"WRITE":    "YAZ",
		"READLINE": "SATIROKU",
		"SEEK":     "SIRA",
		"AND":      "VE",
		"OR":       "VEYA",
		"\\n":      "\\n",
	}

	nestedLang["fi"] = map[string]string{
		"BOX":      "LAATIKKO",
		"FILE":     "TIEDOSTO",
		"STACK":    "LISTA",
		"PORT":     "PORTTI",
		"PRINT":    "TULOSTA",
		"ASK":      "KYSY",
		"GOTO":     "MENE",
		"PUT":      "LAITA",
		"RANDOM":   "SATUNNAINEN",
		"NOW":      "AIKA",
		"ANSWER":   "VASTAUS",
		"IF":       "JOS",
		"THEN":     "NIIN",
		"END":      "LOPPU",
		"SQRT":     "NELIOJUURI",
		"ABS":      "ITSEISARVO",
		"SQR":      "NELIO",
		"ASSIGN":   "ASETA",
		"EXEC":     "SUORITA",
		"CLOSE":    "SULJE",
		"SLEEP":    "ODOTA",
		"OPEN":     "AVAA",
		"READ":     "LUE",
		"WRITE":    "KIRJOITA",
		"READLINE": "LUERIVI",
		"SEEK":     "HAE",
		"AND":      "JA",
		"OR":       "TAI",
		"\\n":      "\\n",
	}

	nestedLang["de"] = map[string]string{
		"BOX":      "KISTE",
		"FILE":     "DATEI",
		"STACK":    "LISTE",
		"PORT":     "PORT",
		"PRINT":    "SCHREIB",
		"ASK":      "FRAG",
		"GOTO":     "GEH",
		"PUT":      "TU",
		"RANDOM":   "ZUFALL",
		"NOW":      "ZEIT",
		"ANSWER":   "ANTWORT",
		"IF":       "WENN",
		"THEN":     "DANN",
		"END":      "ENDE",
		"SQRT":     "WURZEL",
		"ABS":      "BETRAG",
		"SQR":      "QUADRAT",
		"ASSIGN":   "SETZE",
		"EXEC":     "FUEHREAUS",
		"CLOSE":    "SCHLIESSE",
		"SLEEP":    "WARTE",
		"OPEN":     "OEFFNE",
		"READ":     "LIES",
		"WRITE":    "SCHREIB",
		"READLINE": "LIESZEILE",
		"SEEK":     "SUCHE",
		"AND":      "UND",
		"OR":       "ODER",
		"\\n":      "\\n",
	}

	return nestedLang
}

func getTranslation(word string) string {
	if activeLanguage == LANG_EN {
		return word
	}
	if val, exists := translations[activeLanguage][word]; !exists {
		fmt.Printf("translation is missing for %s!", word)
		panic("translation is missing!") // This case is a programming error and must be fixed!
	} else {
		return val
	}
}
