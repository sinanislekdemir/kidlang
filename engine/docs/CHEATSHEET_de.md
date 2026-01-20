# Kidlang Spickzettel (Deutsch)

## Sprachauswahl
```
DE
```

## Kommentare
```
// Dies ist ein Kommentar
```

## Variablen

### Kiste (einfache Variable)
```
kiste name = 42
kiste name = 3.14
kiste name = hallo
kiste name = kiste andere
```

### Liste (Wörterbuch/Map)
```
liste spielzeuge
liste spielzeuge[1] = auto     // oder spielzeuge(1) = auto
liste spielzeuge[2] = ball     // oder spielzeuge(2) = ball
spielzeuge[1] = lastwagen      // oder spielzeuge(1) = lastwagen
schreib spielzeuge[1]          // oder schreib spielzeuge(1)
```

### Datei
```
datei meinedatei
```

## Ausgabe
```
schreib Hallo Welt             // Anführungszeichen optional
schreib "Hallo Welt"           // Anführungszeichen bewahren Text
schreib kiste name             // Gibt Wert der Variable 'name' aus
schreib "kiste name"           // Gibt Text "kiste name" aus
schreib 1 + 2
Jeder Text ohne Befehl         // Implizite Ausgabe
```

## Eingabe
```
frag Wie heißt du?             // Fordert Benutzereingabe an
schreib antwort                // Ergebnis wird in 'antwort' Variable gespeichert
frag "Wert eingeben:"          // Anführungszeichen optional
kiste x = antwort              // Verwende die Antwort
```

## Mathematische Operationen
```
kiste ergebnis = 5 + 3
kiste ergebnis = 10 - 2
kiste ergebnis = 4 * 3
kiste ergebnis = 10 / 2
kiste ergebnis = 10 % 3
kiste ergebnis = 2 ^ 3  // XOR für Zahlen, Verschlüsselung für Strings
```

## Mathematische Funktionen
```
wurzel 16     // Quadratwurzel
betrag -5     // Absolutwert
quadrat 4     // Quadrat (4*4)
sin 1.57      // Sinus
cos 0         // Kosinus
tan 0.785     // Tangens
log 2.718     // Natürlicher Logarithmus
asin 0.5      // Arkussinus
acos 0.5      // Arkuskosinus
```

## String-Operationen
```
kiste text = hallo + welt      // Verkettung (Anführungszeichen optional)
kiste text = hallo - l         // Entfernt alle 'l' Zeichen
kiste text = abc * 3           // String wiederholen
kiste zeichen = hallo / 2      // Zeichen am Index 2 holen
```

## Bedingungen
```
wenn kiste x = 5 dann schreib x ist 5
ende

wenn kiste x > 10 dann
schreib x ist größer als 10
ende

wenn kiste x < 5 dann geh ueberspringe
```

### Vergleichsoperatoren
```
=    // Gleich
!=   // Ungleich
>    // Größer als
<    // Kleiner als
>=   // Größer oder gleich
<=   // Kleiner oder gleich
```

### Logische Operatoren
```
wenn kiste x > 5 und kiste y < 10 dann schreib beide wahr
ende

wenn kiste a = 1 oder kiste b = 2 dann schreib eines wahr
ende
```

## Labels und Sprünge
```
start:
schreib "Hallo"
geh start

ueberspringe:
schreib "Hierher gesprungen"
```

## Datei-Operationen
```
datei meinedatei
oeffne meinedatei daten.txt    // Öffnet/erstellt Datei
lies meinedatei kiste inhalt   // Liest gesamte Datei
lieszeile meinedatei kiste zeile // Liest eine Zeile
schreib meinedatei etwas text  // Schreibt in Datei
suche meinedatei 5             // Springt zu Zeile 5
schliesse meinedatei           // Schließt Datei
```

## Spezialwerte
```
zufall  // Zufällige Ganzzahl
zeit    // Aktuelles Datum/Uhrzeit
\n      // Zeilenumbruch
```

## Systembefehle
```
fuehreaus ls -la       // Shell-Befehl ausführen (Anführungszeichen optional)
warte 1000             // Warte 1000 Millisekunden
```

## Ablaufsteuerung
- Programme werden Zeile für Zeile von oben nach unten ausgeführt
- `geh` springt zu einem Label
- `wenn...dann...ende` erstellt bedingte Blöcke
- `ende` schließt einen `wenn` Block
