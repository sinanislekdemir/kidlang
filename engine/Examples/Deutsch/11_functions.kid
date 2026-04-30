DE
// Funktionen sind kleine Helfer, die du wieder benutzen kannst.

funktion jubel(kiste name)
Hurra fuer kiste name !
ende

funktion extra_sterne(kiste aktuell)
rueckgabe kiste aktuell + 2
ende

frag Wie heisst du?
kiste spieler = antwort

kiste sterne = 4
kiste gesamt = extra_sterne(kiste sterne)

schreib jubel(kiste spieler)
schreib Jetzt hast du kiste gesamt Sterne.
schreib Dein Name hat laenge(kiste spieler) Buchstaben.
