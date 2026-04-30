DE
// Mittelerde Händler - Ein Handelsabenteuer
// Kaufe günstig, verkaufe teuer, verdiene 1000 Gold!

funktion marktpreis(kiste basis, kiste spannweite)
zufall() % kiste spannweite + kiste basis
ende

funktion gesamtpreis(kiste preis, kiste menge)
kiste preis * kiste menge
ende

funktion ruestungsplatz(kiste menge)
kiste menge * 2
ende

funktion freier_platz(kiste tasche, kiste belegt)
kiste tasche - kiste belegt
ende

funktion sicheres_banditen_gold(kiste gold)
wenn kiste gold > 30 dann
   rueckgabe kiste gold - 30
ende
0
ende

funktion haendler_titel(kiste gold)
wenn kiste gold >= 1000 dann
   rueckgabe Meisterhaendler
ende

wenn kiste gold >= 500 dann
   rueckgabe Starker Haendler
ende

Lernender Haendler
ende

schreib ====================================
schreib   MITTELERDE HAENDLER
schreib ====================================
schreib
schreib Du bist ein Händler in Mittelerde
schreib Ziel: Verdiene 1000 Gold in 30 Tagen!
schreib

kiste tag = 1
kiste gold = 100
kiste tasche = 50
kiste belegt = 0
kiste traenke = 0
kiste staebe = 0
kiste ruestung = 0

anfang:
schreib
schreib === TAG kiste tag of 30 ===
schreib Gold: kiste gold
schreib Tasche: kiste belegt of kiste tasche Plätzen
schreib Freier Platz: freier_platz(kiste tasche, kiste belegt)
schreib Inventar - Tränke: kiste traenke | Stäbe: kiste staebe | Rüstung: kiste ruestung
schreib

kiste t_preis = marktpreis(20, 30)
kiste s_preis = marktpreis(40, 50)  
kiste r_preis = marktpreis(60, 80)

schreib === MARKTPREISE ===
schreib Tränke: kiste t_preis Gold
schreib Stäbe: kiste s_preis Gold
schreib Rüstung: kiste r_preis Gold (braucht 2 Plätze)
schreib
schreib === AKTIONEN ===
schreib 1. Tränke kaufen
schreib 2. Stäbe kaufen
schreib 3. Rüstung kaufen
schreib 4. Tränke verkaufen
schreib 5. Stäbe verkaufen
schreib 6. Rüstung verkaufen
schreib 7. Zur nächsten Stadt reisen
frag Wähle:

kiste wahl = antwort

wenn kiste wahl = 1 dann geh traenke_kaufen
wenn kiste wahl = 2 dann geh staebe_kaufen
wenn kiste wahl = 3 dann geh ruestung_kaufen
wenn kiste wahl = 4 dann geh traenke_verkaufen
wenn kiste wahl = 5 dann geh staebe_verkaufen
wenn kiste wahl = 6 dann geh ruestung_verkaufen
wenn kiste wahl = 7 dann geh reisen
schreib Ungültige Wahl!
warte 1000
geh anfang

traenke_kaufen:
frag Wie viele Tränke?
kiste menge = antwort
kiste kosten = gesamtpreis(kiste t_preis, kiste menge)
kiste bedarf = kiste belegt + kiste menge
wenn kiste kosten > kiste gold dann
   schreib Nicht genug Gold!
   warte 1000
   geh anfang
ende
wenn kiste bedarf > kiste tasche dann
   schreib Nicht genug Platz!
   warte 1000
   geh anfang
ende
kiste gold = kiste gold - kiste kosten
kiste traenke = kiste traenke + kiste menge
kiste belegt = kiste belegt + kiste menge
schreib kiste menge Tränke für kiste kosten Gold gekauft!
warte 1000
geh anfang

staebe_kaufen:
frag Wie viele Stäbe?
kiste menge = antwort
kiste kosten = gesamtpreis(kiste s_preis, kiste menge)
kiste bedarf = kiste belegt + kiste menge
wenn kiste kosten > kiste gold dann
   schreib Nicht genug Gold!
   warte 1000
   geh anfang
ende
wenn kiste bedarf > kiste tasche dann
   schreib Nicht genug Platz!
   warte 1000
   geh anfang
ende
kiste gold = kiste gold - kiste kosten
kiste staebe = kiste staebe + kiste menge
kiste belegt = kiste belegt + kiste menge
schreib kiste menge Stäbe für kiste kosten Gold gekauft!
warte 1000
geh anfang

ruestung_kaufen:
frag Wie viele Rüstungen?
kiste menge = antwort
kiste kosten = gesamtpreis(kiste r_preis, kiste menge)
kiste bedarf = kiste belegt + ruestungsplatz(kiste menge)
wenn kiste kosten > kiste gold dann
   schreib Nicht genug Gold!
   warte 1000
   geh anfang
ende
wenn kiste bedarf > kiste tasche dann
   schreib Nicht genug Platz!
   warte 1000
   geh anfang
ende
kiste gold = kiste gold - kiste kosten
kiste ruestung = kiste ruestung + kiste menge
kiste belegt = kiste belegt + ruestungsplatz(kiste menge)
schreib kiste menge Rüstungen für kiste kosten Gold gekauft!
warte 1000
geh anfang

traenke_verkaufen:
frag Wie viele Tränke?
kiste menge = antwort
wenn kiste menge > kiste traenke dann
   schreib Du hast nicht so viele!
   warte 1000
   geh anfang
ende
kiste verdienst = gesamtpreis(kiste t_preis, kiste menge)
kiste gold = kiste gold + kiste verdienst
kiste traenke = kiste traenke - kiste menge
kiste belegt = kiste belegt - kiste menge
schreib kiste menge Tränke für kiste verdienst Gold verkauft!
warte 1000
geh anfang

staebe_verkaufen:
frag Wie viele Stäbe?
kiste menge = antwort
wenn kiste menge > kiste staebe dann
   schreib Du hast nicht so viele!
   warte 1000
   geh anfang
ende
kiste verdienst = gesamtpreis(kiste s_preis, kiste menge)
kiste gold = kiste gold + kiste verdienst
kiste staebe = kiste staebe - kiste menge
kiste belegt = kiste belegt - kiste menge
schreib kiste menge Stäbe für kiste verdienst Gold verkauft!
warte 1000
geh anfang

ruestung_verkaufen:
frag Wie viele Rüstungen?
kiste menge = antwort
wenn kiste menge > kiste ruestung dann
   schreib Du hast nicht so viele!
   warte 1000
   geh anfang
ende
kiste verdienst = gesamtpreis(kiste r_preis, kiste menge)
kiste gold = kiste gold + kiste verdienst
kiste ruestung = kiste ruestung - kiste menge
kiste belegt = kiste belegt - ruestungsplatz(kiste menge)
schreib kiste menge Rüstungen für kiste verdienst Gold verkauft!
warte 1000
geh anfang

reisen:
schreib
schreib Reise zur nächsten Stadt...
warte 1000
kiste tag = kiste tag + 1

kiste ereignis = zufall() % 8

wenn kiste ereignis = 0 dann
   schreib Ein freundlicher Zauberer gibt dir 50 Gold! ✨
   kiste gold = kiste gold + 50
   warte 1500
ende

wenn kiste ereignis = 1 dann
   schreib Banditen greifen an! 30 Gold verloren! ⚔️
   kiste gold = sicheres_banditen_gold(kiste gold)
   warte 1500
ende

wenn kiste ereignis = 2 dann
   schreib Einen Trank auf der Straße gefunden! 🧪
   wenn kiste belegt < kiste tasche dann
      kiste traenke = kiste traenke + 1
      kiste belegt = kiste belegt + 1
   ende
   warte 1500
ende

wenn kiste gold >= 1000 dann geh gewonnen
wenn kiste tag > 30 dann geh verloren

geh anfang

gewonnen:
schreib
schreib ====================================
schreib   GLÜCKWUNSCH!
schreib ====================================
schreib Du hast 1000 Gold in kiste tag Tagen verdient!
schreib Titel: haendler_titel(kiste gold)
schreib Du bist bereit fuer den naechsten Markt! 🏆
schreib
geh ende

verloren:
schreib
schreib ====================================
schreib   ZEIT IST UM!
schreib ====================================
schreib 30 Tage sind vergangen...
schreib Finales Gold: kiste gold
schreib Titel: haendler_titel(kiste gold)
schreib
wenn kiste gold >= 500 dann
   schreib Nicht schlecht für einen Händler!
ende
wenn kiste gold < 500 dann
   schreib Übe weiter deinen Handel!
ende
schreib
geh ende

ende:
schreib Danke fürs Spielen von Mittelerde Händler!
schreib Auf Wiedersehen! 👋
