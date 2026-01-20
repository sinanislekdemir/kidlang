DE
liste noten
liste namen
schreib === Notenbuch ===
kiste anzahl = 0

menue:
schreib
schreib 1 Schüler hinzufügen
schreib 2 Alle Schüler anzeigen
schreib 3 Beenden
frag Wähle:
kiste wahl = antwort

wenn kiste wahl = 1 dann
frag Schülername:
kiste name = antwort
frag Schülernote:
kiste note = antwort
kiste anzahl = kiste anzahl + 1
liste namen[kiste anzahl] = kiste name
liste noten[kiste anzahl] = kiste note
schreib Hinzugefügt!
geh menue
ende

wenn kiste wahl = 2 dann
schreib
schreib === Alle Schüler ===
kiste i = 1
zeigen:
wenn kiste i <= kiste anzahl dann
schreib liste namen[kiste i]: liste noten[kiste i]
kiste i = kiste i + 1
geh zeigen
ende
geh menue
ende

wenn kiste wahl = 3 dann
schreib Auf Wiedersehen!
ende
