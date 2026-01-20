DE
// Zahlenratespiel
// Versuche die geheime Zahl zu erraten!

schreib ================================
schreib    ZAHLENRATESPIEL
schreib ================================
schreib

kiste geheimnis = zufall % 50 + 1
kiste versuche = 0
kiste max_versuche = 7

schreib Ich denke an eine Zahl zwischen 1 und 50
schreib Du hast kiste max_versuche Versuche!
schreib

spiel:
kiste versuche = kiste versuche + 1

schreib Versuch kiste versuche von kiste max_versuche
frag Deine Vermutung:

wenn antwort = kiste geheimnis dann
   geh gewonnen
ende

wenn antwort > kiste geheimnis dann
   schreib Zu hoch! Versuche eine kleinere Zahl.
ende

wenn antwort < kiste geheimnis dann
   schreib Zu niedrig! Versuche eine größere Zahl.
ende

wenn kiste versuche >= kiste max_versuche dann
   geh verloren
ende

schreib
geh spiel

gewonnen:
schreib
schreib ================================
schreib    GLÜCKWUNSCH!
schreib ================================
schreib Du hast es erraten! Die Zahl war kiste geheimnis
schreib Du hast kiste versuche Versuche gebraucht!
geh schluss

verloren:
schreib
schreib ================================
schreib    SPIEL VORBEI
schreib ================================
schreib Leider sind deine Versuche aufgebraucht!
schreib Die geheime Zahl war kiste geheimnis
geh schluss

schluss:
schreib
schreib Danke fürs Spielen!
