DE
// Einfaches Abenteuerspiel
// Triff Entscheidungen und erkunde!

schreib ================================
schreib    SCHATZSUCHE ABENTEUER
schreib ================================
schreib

kiste gesundheit = 100
kiste gold = 0

anfang:
schreib Du wachst in einem geheimnisvollen Wald auf.
schreib Deine Gesundheit: kiste gesundheit
schreib Dein Gold: kiste gold
schreib
schreib Was möchtest du tun?
schreib 1. Den Wald erkunden
schreib 2. Deine Tasche prüfen
schreib 3. Ausruhen
schreib 4. Spiel beenden
frag Wähle:

wenn antwort = 1 dann geh erkunden
wenn antwort = 2 dann geh tasche_pruefen
wenn antwort = 3 dann geh ausruhen
wenn antwort = 4 dann geh beenden

schreib Ungültige Wahl!
geh anfang

erkunden:
schreib
schreib Du wagst dich tiefer in den Wald...
warte 1000

kiste ereignis = zufall % 3

wenn kiste ereignis = 0 dann geh gold_finden
wenn kiste ereignis = 1 dann geh monster_finden
geh nichts_finden

gold_finden:
kiste gefunden = zufall % 20 + 10
kiste gold = kiste gold + kiste gefunden
schreib Du hast kiste gefunden Goldmünzen gefunden! ✨
schreib Gesamt Gold: kiste gold
warte 2000
geh anfang

monster_finden:
schreib Ein Monster erscheint! 👹
kiste schaden = zufall % 30 + 10
kiste gesundheit = kiste gesundheit - kiste schaden
schreib Es greift dich mit kiste schaden Schaden an!
schreib Verbleibende Gesundheit: kiste gesundheit

wenn kiste gesundheit <= 0 dann
   geh spiel_vorbei
ende

warte 2000
geh anfang

nichts_finden:
schreib Du hast hier nichts Interessantes gefunden.
warte 1000
geh anfang

tasche_pruefen:
schreib
schreib === DEIN STATUS ===
schreib Gesundheit: kiste gesundheit
schreib Gold: kiste gold
schreib
wenn kiste gold >= 100 dann
   schreib Du hast genug Gold zum Gewinnen!
   schreib Glückwunsch! 🏆
   geh beenden
ende
schreib Du brauchst 100 Gold zum Gewinnen.
schreib Erkunde weiter!
warte 2000
geh anfang

ausruhen:
schreib
schreib Du ruhst dich eine Weile aus...
warte 1500
kiste heilung = 20
kiste gesundheit = kiste gesundheit + kiste heilung
wenn kiste gesundheit > 100 dann
   kiste gesundheit = 100
ende
schreib Du hast kiste heilung Gesundheit wiederhergestellt!
schreib Gesundheit: kiste gesundheit
warte 1500
geh anfang

spiel_vorbei:
schreib
schreib ================================
schreib       SPIEL VORBEI
schreib ================================
schreib Du wurdest besiegt!
schreib Finales Gold: kiste gold
schreib
geh beenden

beenden:
schreib
schreib Danke fürs Spielen!
schreib Auf Wiedersehen! 👋
