DE
liste artikel
kiste anzahl = 0
schreib === Einkaufsliste ===
schreib
menue:
schreib Was möchtest du tun?
schreib 1 Artikel hinzufügen
schreib 2 Liste anzeigen
schreib 3 Beenden
frag Wähle:
kiste wahl = antwort

wenn kiste wahl = 1 dann
frag Gib Artikelname ein:
kiste name = antwort
kiste anzahl = kiste anzahl + 1
liste artikel[kiste anzahl] = kiste name
schreib Hinzugefügt!
geh menue
ende

wenn kiste wahl = 2 dann
schreib
schreib Deine Einkaufsliste:
kiste i = 1
zeigen:
wenn kiste i <= kiste anzahl dann
schreib kiste i . liste artikel[kiste i]
kiste i = kiste i + 1
geh zeigen
ende
schreib
geh menue
ende

wenn kiste wahl = 3 dann
schreib Auf Wiedersehen!
ende
