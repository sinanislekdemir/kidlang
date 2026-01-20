package interpreter

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type IDELanguage string

const (
	IDE_LANG_EN IDELanguage = "en"
	IDE_LANG_TR IDELanguage = "tr"
	IDE_LANG_FI IDELanguage = "fi"
	IDE_LANG_DE IDELanguage = "de"
)

var ideLanguage IDELanguage = IDE_LANG_EN
var ideTranslations = createIDETranslations()

type IDEStrings struct {
	// Menu items
	MenuFile     string
	MenuEdit     string
	MenuRun      string
	MenuExamples string
	MenuHelp     string
	MenuLanguage string

	// File menu
	FileNew    string
	FileOpen   string
	FileSave   string
	FileSaveAs string
	FileClose  string
	FileExit   string

	// Edit menu
	EditCut    string
	EditCopy   string
	EditPaste  string
	EditDelete string
	EditClear  string

	// Run menu
	RunRun   string
	RunStop  string
	RunDebug string

	// Examples menu
	ExamplesBrowse string

	// Help menu
	HelpAbout     string
	HelpDocs      string
	HelpShortcuts string
	HelpLangRef   string

	// Language menu
	LangEnglish string
	LangTurkish string
	LangFinnish string
	LangGerman  string

	// Titles
	TitleDocumentation  string
	TitleKeyboardShort  string
	TitleExamplesSelect string
	TitleExamplesPrefix string

	// Status bar
	StatusUntitled string
	StatusLine     string
	StatusCol      string

	// Help text
	HelpRunning     string
	HelpMenuOpen    string
	HelpMenuNav     string
	HelpFileSelect  string
	HelpStopProgram string

	// Messages
	MsgRunning       string
	MsgFinished      string
	MsgInterrupted   string
	MsgPressKey      string
	MsgPressEnter    string
	MsgSaveChanges   string
	MsgYes           string
	MsgNo            string
	MsgEnterFilename string
	MsgFileSaved     string
	MsgFileLoaded    string
	MsgError         string

	// About dialog
	AboutTitle    string
	AboutSubtitle string
	AboutAuthor   string

	// Keyboard shortcuts
	ShortcutRun       string
	ShortcutMenu      string
	ShortcutEsc       string
	ShortcutNew       string
	ShortcutOpen      string
	ShortcutSave      string
	ShortcutArrows    string
	ShortcutEnter     string
	ShortcutBackspace string
	ShortcutDelete    string
	ShortcutInMenu    string
	ShortcutMenuLR    string
	ShortcutMenuUD    string
	ShortcutMenuEnter string
}

func createIDETranslations() map[IDELanguage]IDEStrings {
	translations := make(map[IDELanguage]IDEStrings)

	// English
	translations[IDE_LANG_EN] = IDEStrings{
		MenuFile:     "File",
		MenuEdit:     "Edit",
		MenuRun:      "Run",
		MenuExamples: "Examples",
		MenuHelp:     "Help",
		MenuLanguage: "Language",

		FileNew:    "New",
		FileOpen:   "Open",
		FileSave:   "Save",
		FileSaveAs: "Save As",
		FileClose:  "Close",
		FileExit:   "Exit",

		EditCut:    "Cut",
		EditCopy:   "Copy",
		EditPaste:  "Paste",
		EditDelete: "Delete",
		EditClear:  "Clear",

		RunRun:   "Run (F5)",
		RunStop:  "Stop",
		RunDebug: "Debug",

		ExamplesBrowse: "Browse Examples",

		HelpAbout:     "About",
		HelpDocs:      "Documentation",
		HelpShortcuts: "Keyboard Shortcuts",
		HelpLangRef:   "Language Reference",

		LangEnglish: "English",
		LangTurkish: "Turkish",
		LangFinnish: "Finnish",
		LangGerman:  "German",

		TitleDocumentation:  "Documentation",
		TitleKeyboardShort:  "Keyboard Shortcuts",
		TitleExamplesSelect: "Examples - Select Category",
		TitleExamplesPrefix: "Examples - ",

		StatusUntitled: "Untitled",
		StatusLine:     "Line",
		StatusCol:      "Col",

		HelpRunning:     "F5=Run  F10=Menu  ^S=Save  ^N=New  ^O=Open",
		HelpMenuOpen:    "Up/Down=Select  Left/Right=Menu  Enter=OK  ESC=Cancel",
		HelpMenuNav:     "Left/Right=Menu  Up/Down=Select  Enter=OK  ESC=Cancel",
		HelpFileSelect:  "Up/Down=Select  Enter=Load  ESC=Back",
		HelpStopProgram: "Press Ctrl+C to stop the program",

		MsgRunning:       "=== Running Program ===",
		MsgFinished:      "=== Program Finished ===",
		MsgInterrupted:   "=== Program Interrupted ===",
		MsgPressKey:      "Press any key to continue...",
		MsgPressEnter:    "Press Enter to continue...",
		MsgSaveChanges:   "Save changes before closing?",
		MsgYes:           "Yes",
		MsgNo:            "No",
		MsgEnterFilename: "Enter filename:",
		MsgFileSaved:     "File saved successfully",
		MsgFileLoaded:    "File loaded successfully",
		MsgError:         "Error",

		AboutTitle:    "KidLang IDE",
		AboutSubtitle: "Batman Edition",
		AboutAuthor:   "Developed by Sinan Islekdemir",

		ShortcutRun:       "F5           - Run program",
		ShortcutMenu:      "F10          - Open menu",
		ShortcutEsc:       "ESC          - Close menu/dialog",
		ShortcutNew:       "Ctrl+N       - New file",
		ShortcutOpen:      "Ctrl+O       - Open file",
		ShortcutSave:      "Ctrl+S       - Save file",
		ShortcutArrows:    "Arrow Keys   - Navigate editor",
		ShortcutEnter:     "Enter        - New line",
		ShortcutBackspace: "Backspace    - Delete character before cursor",
		ShortcutDelete:    "Delete       - Delete character at cursor",
		ShortcutInMenu:    "In Menu:",
		ShortcutMenuLR:    "  Left/Right - Switch menus",
		ShortcutMenuUD:    "  Up/Down    - Select items",
		ShortcutMenuEnter: "  Enter      - Confirm selection",
	}

	// Turkish
	translations[IDE_LANG_TR] = IDEStrings{
		MenuFile:     "Dosya",
		MenuEdit:     "Düzenle",
		MenuRun:      "Çalıştır",
		MenuExamples: "Örnekler",
		MenuHelp:     "Yardım",
		MenuLanguage: "Dil",

		FileNew:    "Yeni",
		FileOpen:   "Aç",
		FileSave:   "Kaydet",
		FileSaveAs: "Farklı Kaydet",
		FileClose:  "Kapat",
		FileExit:   "Çıkış",

		EditCut:    "Kes",
		EditCopy:   "Kopyala",
		EditPaste:  "Yapıştır",
		EditDelete: "Sil",
		EditClear:  "Temizle",

		RunRun:   "Çalıştır (F5)",
		RunStop:  "Durdur",
		RunDebug: "Hata Ayıkla",

		ExamplesBrowse: "Örneklere Gözat",

		HelpAbout:     "Hakkında",
		HelpDocs:      "Dökümanlar",
		HelpShortcuts: "Klavye Kısayolları",
		HelpLangRef:   "Dil Referansı",

		LangEnglish: "İngilizce",
		LangTurkish: "Türkçe",
		LangFinnish: "Fince",
		LangGerman:  "Almanca",

		TitleDocumentation:  "Dökümanlar",
		TitleKeyboardShort:  "Klavye Kısayolları",
		TitleExamplesSelect: "Örnekler - Kategori Seçin",
		TitleExamplesPrefix: "Örnekler - ",

		StatusUntitled: "Adsız",
		StatusLine:     "Satır",
		StatusCol:      "Sütun",

		HelpRunning:     "F5=Çalıştır  F10=Menü  ^S=Kaydet  ^N=Yeni  ^O=Aç",
		HelpMenuOpen:    "Yukarı/Aşağı=Seç  Sol/Sağ=Menü  Enter=Tamam  ESC=İptal",
		HelpMenuNav:     "Sol/Sağ=Menü  Yukarı/Aşağı=Seç  Enter=Tamam  ESC=İptal",
		HelpFileSelect:  "Yukarı/Aşağı=Seç  Enter=Yükle  ESC=Geri",
		HelpStopProgram: "Programı durdurmak için Ctrl+C'ye basın",

		MsgRunning:       "=== Program Çalışıyor ===",
		MsgFinished:      "=== Program Tamamlandı ===",
		MsgInterrupted:   "=== Program Kesildi ===",
		MsgPressKey:      "Devam için herhangi bir tuşa basın...",
		MsgPressEnter:    "Devam için Enter'a basın...",
		MsgSaveChanges:   "Kapatmadan önce değişiklikleri kaydet?",
		MsgYes:           "Evet",
		MsgNo:            "Hayır",
		MsgEnterFilename: "Dosya adı girin:",
		MsgFileSaved:     "Dosya başarıyla kaydedildi",
		MsgFileLoaded:    "Dosya başarıyla yüklendi",
		MsgError:         "Hata",

		AboutTitle:    "KidLang IDE",
		AboutSubtitle: "Batman Versiyonu",
		AboutAuthor:   "Geliştiren: Sinan Islekdemir",

		ShortcutRun:       "F5           - Programı çalıştır",
		ShortcutMenu:      "F10          - Menüyü aç",
		ShortcutEsc:       "ESC          - Menü/diyalogu kapat",
		ShortcutNew:       "Ctrl+N       - Yeni dosya",
		ShortcutOpen:      "Ctrl+O       - Dosya aç",
		ShortcutSave:      "Ctrl+S       - Dosyayı kaydet",
		ShortcutArrows:    "Yön Tuşları  - Editörde gezin",
		ShortcutEnter:     "Enter        - Yeni satır",
		ShortcutBackspace: "Backspace    - İmleçten önceki karakteri sil",
		ShortcutDelete:    "Delete       - İmleçteki karakteri sil",
		ShortcutInMenu:    "Menüde:",
		ShortcutMenuLR:    "  Sol/Sağ    - Menüler arası geç",
		ShortcutMenuUD:    "  Yukarı/Aşağı- Öğe seç",
		ShortcutMenuEnter: "  Enter      - Seçimi onayla",
	}

	// Finnish
	translations[IDE_LANG_FI] = IDEStrings{
		MenuFile:     "Tiedosto",
		MenuEdit:     "Muokkaa",
		MenuRun:      "Suorita",
		MenuExamples: "Esimerkit",
		MenuHelp:     "Ohje",
		MenuLanguage: "Kieli",

		FileNew:    "Uusi",
		FileOpen:   "Avaa",
		FileSave:   "Tallenna",
		FileSaveAs: "Tallenna nimellä",
		FileClose:  "Sulje",
		FileExit:   "Poistu",

		EditCut:    "Leikkaa",
		EditCopy:   "Kopioi",
		EditPaste:  "Liitä",
		EditDelete: "Poista",
		EditClear:  "Tyhjennä",

		RunRun:   "Suorita (F5)",
		RunStop:  "Pysäytä",
		RunDebug: "Virheenkorjaus",

		ExamplesBrowse: "Selaa esimerkkejä",

		HelpAbout:     "Tietoja",
		HelpDocs:      "Dokumentaatio",
		HelpShortcuts: "Pikanäppäimet",
		HelpLangRef:   "Kieli Viite",

		LangEnglish: "Englanti",
		LangTurkish: "Turkki",
		LangFinnish: "Suomi",
		LangGerman:  "Saksa",

		TitleDocumentation:  "Dokumentaatio",
		TitleKeyboardShort:  "Pikanäppäimet",
		TitleExamplesSelect: "Esimerkit - Valitse kategoria",
		TitleExamplesPrefix: "Esimerkit - ",

		StatusUntitled: "Nimetön",
		StatusLine:     "Rivi",
		StatusCol:      "Sarake",

		HelpRunning:     "F5=Suorita  F10=Valikko  ^S=Tallenna  ^N=Uusi  ^O=Avaa",
		HelpMenuOpen:    "Ylös/Alas=Valitse  Vasen/Oikea=Valikko  Enter=OK  ESC=Peruuta",
		HelpMenuNav:     "Vasen/Oikea=Valikko  Ylös/Alas=Valitse  Enter=OK  ESC=Peruuta",
		HelpFileSelect:  "Ylös/Alas=Valitse  Enter=Lataa  ESC=Takaisin",
		HelpStopProgram: "Pysäytä ohjelma painamalla Ctrl+C",

		MsgRunning:       "=== Ohjelma suoritetaan ===",
		MsgFinished:      "=== Ohjelma valmis ===",
		MsgInterrupted:   "=== Ohjelma keskeytetty ===",
		MsgPressKey:      "Paina mitä tahansa näppäintä jatkaaksesi...",
		MsgPressEnter:    "Paina Enter jatkaaksesi...",
		MsgSaveChanges:   "Tallenna muutokset ennen sulkemista?",
		MsgYes:           "Kyllä",
		MsgNo:            "Ei",
		MsgEnterFilename: "Anna tiedostonimi:",
		MsgFileSaved:     "Tiedosto tallennettu onnistuneesti",
		MsgFileLoaded:    "Tiedosto ladattu onnistuneesti",
		MsgError:         "Virhe",

		AboutTitle:    "KidLang IDE",
		AboutSubtitle: "Batman-painos",
		AboutAuthor:   "Kehittäjä: Sinan Islekdemir",

		ShortcutRun:       "F5           - Suorita ohjelma",
		ShortcutMenu:      "F10          - Avaa valikko",
		ShortcutEsc:       "ESC          - Sulje valikko/dialogi",
		ShortcutNew:       "Ctrl+N       - Uusi tiedosto",
		ShortcutOpen:      "Ctrl+O       - Avaa tiedosto",
		ShortcutSave:      "Ctrl+S       - Tallenna tiedosto",
		ShortcutArrows:    "Nuolinäppäimet - Navigoi editorissa",
		ShortcutEnter:     "Enter        - Uusi rivi",
		ShortcutBackspace: "Backspace    - Poista merkki ennen kursoria",
		ShortcutDelete:    "Delete       - Poista merkki kursorilla",
		ShortcutInMenu:    "Valikossa:",
		ShortcutMenuLR:    "  Vasen/Oikea- Vaihda valikkoja",
		ShortcutMenuUD:    "  Ylös/Alas  - Valitse kohteita",
		ShortcutMenuEnter: "  Enter      - Vahvista valinta",
	}

	// German
	translations[IDE_LANG_DE] = IDEStrings{
		MenuFile:     "Datei",
		MenuEdit:     "Bearbeiten",
		MenuRun:      "Ausführen",
		MenuExamples: "Beispiele",
		MenuHelp:     "Hilfe",
		MenuLanguage: "Sprache",

		FileNew:    "Neu",
		FileOpen:   "Öffnen",
		FileSave:   "Speichern",
		FileSaveAs: "Speichern unter",
		FileClose:  "Schließen",
		FileExit:   "Beenden",

		EditCut:    "Ausschneiden",
		EditCopy:   "Kopieren",
		EditPaste:  "Einfügen",
		EditDelete: "Löschen",
		EditClear:  "Löschen",

		RunRun:   "Ausführen (F5)",
		RunStop:  "Stoppen",
		RunDebug: "Debuggen",

		ExamplesBrowse: "Beispiele durchsuchen",

		HelpAbout:     "Über",
		HelpDocs:      "Dokumentation",
		HelpShortcuts: "Tastaturkürzel",
		HelpLangRef:   "Sprachreferenz",

		LangEnglish: "Englisch",
		LangTurkish: "Türkisch",
		LangFinnish: "Finnisch",
		LangGerman:  "Deutsch",

		TitleDocumentation:  "Dokumentation",
		TitleKeyboardShort:  "Tastaturkürzel",
		TitleExamplesSelect: "Beispiele - Kategorie wählen",
		TitleExamplesPrefix: "Beispiele - ",

		StatusUntitled: "Unbenannt",
		StatusLine:     "Zeile",
		StatusCol:      "Spalte",

		HelpRunning:     "F5=Ausführen  F10=Menü  ^S=Speichern  ^N=Neu  ^O=Öffnen",
		HelpMenuOpen:    "Auf/Ab=Wählen  Links/Rechts=Menü  Enter=OK  ESC=Abbrechen",
		HelpMenuNav:     "Links/Rechts=Menü  Auf/Ab=Wählen  Enter=OK  ESC=Abbrechen",
		HelpFileSelect:  "Auf/Ab=Wählen  Enter=Laden  ESC=Zurück",
		HelpStopProgram: "Drücken Sie Ctrl+C um das Programm zu stoppen",

		MsgRunning:       "=== Programm läuft ===",
		MsgFinished:      "=== Programm beendet ===",
		MsgInterrupted:   "=== Programm unterbrochen ===",
		MsgPressKey:      "Drücken Sie eine beliebige Taste um fortzufahren...",
		MsgPressEnter:    "Drücken Sie Enter um fortzufahren...",
		MsgSaveChanges:   "Änderungen vor dem Schließen speichern?",
		MsgYes:           "Ja",
		MsgNo:            "Nein",
		MsgEnterFilename: "Dateiname eingeben:",
		MsgFileSaved:     "Datei erfolgreich gespeichert",
		MsgFileLoaded:    "Datei erfolgreich geladen",
		MsgError:         "Fehler",

		AboutTitle:    "KidLang IDE",
		AboutSubtitle: "Batman Edition",
		AboutAuthor:   "Entwickelt von Sinan Islekdemir",

		ShortcutRun:       "F5           - Programm ausführen",
		ShortcutMenu:      "F10          - Menü öffnen",
		ShortcutEsc:       "ESC          - Menü/Dialog schließen",
		ShortcutNew:       "Ctrl+N       - Neue Datei",
		ShortcutOpen:      "Ctrl+O       - Datei öffnen",
		ShortcutSave:      "Ctrl+S       - Datei speichern",
		ShortcutArrows:    "Pfeiltasten  - Im Editor navigieren",
		ShortcutEnter:     "Enter        - Neue Zeile",
		ShortcutBackspace: "Backspace    - Zeichen vor Cursor löschen",
		ShortcutDelete:    "Delete       - Zeichen am Cursor löschen",
		ShortcutInMenu:    "Im Menü:",
		ShortcutMenuLR:    "  Links/Rechts - Menüs wechseln",
		ShortcutMenuUD:    "  Auf/Ab      - Elemente wählen",
		ShortcutMenuEnter: "  Enter       - Auswahl bestätigen",
	}

	return translations
}

func getIDETranslation() IDEStrings {
	return ideTranslations[ideLanguage]
}

func setIDELanguage(lang IDELanguage) {
	ideLanguage = lang
}

func getIDELanguage() IDELanguage {
	return ideLanguage
}

// Exported functions for testing
func GetIDETranslation() IDEStrings {
	return getIDETranslation()
}

func SetIDELanguage(lang IDELanguage) {
	setIDELanguage(lang)
}

func GetIDELanguage() IDELanguage {
	return getIDELanguage()
}

// Load language from config file
func loadIDELanguageFromConfig() {
	configPath := getConfigPath()
	if configPath == "" {
		return
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "language=") {
			lang := strings.TrimPrefix(line, "language=")
			lang = strings.TrimSpace(lang)
			switch IDELanguage(lang) {
			case IDE_LANG_EN, IDE_LANG_TR, IDE_LANG_FI, IDE_LANG_DE:
				ideLanguage = IDELanguage(lang)
			}
			return
		}
	}
}

// Save language to config file
func saveIDELanguageToConfig() error {
	configPath := getConfigPath()
	if configPath == "" {
		return fmt.Errorf("cannot determine config path")
	}

	// Create directory if needed
	configDir := filepath.Dir(configPath)
	if configDir != "" {
		os.MkdirAll(configDir, 0755)
	}

	// Read existing config if it exists
	var lines []string
	data, err := os.ReadFile(configPath)
	if err == nil {
		lines = strings.Split(string(data), "\n")
	}

	// Update or add language setting
	found := false
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "language=") {
			lines[i] = fmt.Sprintf("language=%s", ideLanguage)
			found = true
			break
		}
	}

	if !found {
		lines = append(lines, fmt.Sprintf("language=%s", ideLanguage))
	}

	// Write back
	content := strings.Join(lines, "\n")
	return os.WriteFile(configPath, []byte(content), 0644)
}

// Get config file path based on platform
func getConfigPath() string {
	if runtime.GOOS == "windows" {
		// On Windows, we just fail to save (as requested)
		return ""
	}

	// Try to save in current directory first
	_, err := os.Stat("kidlang.ini")
	if err == nil {
		// File exists, try to use it
		f, err := os.OpenFile("kidlang.ini", os.O_RDWR, 0644)
		if err == nil {
			f.Close()
			return "kidlang.ini"
		}
	} else {
		// File doesn't exist, try to create it
		f, err := os.Create("kidlang.ini")
		if err == nil {
			f.Close()
			os.Remove("kidlang.ini") // Remove it, we'll recreate when saving
			return "kidlang.ini"
		}
	}

	// Can't write to current directory, use home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	return filepath.Join(home, ".kidlang.ini")
}
