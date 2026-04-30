DE
// Eingebaute Helfer koennen mit Text, Dateien und Listen arbeiten.

kiste haustier = regenbogen roboter
schreib enthaelt(kiste haustier, bogen)
schreib ersetze(kiste haustier, roboter, drache)
schreib teil(kiste haustier, 1, 10)

datei heft = notizen
oeffne heft helpernotes.txt
schreib gibtsdatei(heft)
schreib dateigroesse(heft)
schliesse heft

liste spielzeug
kiste erstes = listesetz(spielzeug, 1, roboter)
kiste zweites = listesetz(spielzeug, 2, drache)
kiste schluessel = listeschluessel(spielzeug)
schreib verbinde(kiste schluessel, :)
schreib listenlaenge(spielzeug)
