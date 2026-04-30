FI
// Yksinkertainen seikkailupeli
// Tee valintoja ja tutkiskele!

tulosta ================================
tulosta    AARTEENETSINTÄSEIKKAILU
tulosta ================================
tulosta

laatikko terveys = 100
laatikko kulta = 0

aloita:
tulosta Heräät salaperäisessä metsässä.
tulosta Terveytesi: laatikko terveys
tulosta Kultasi: laatikko kulta
tulosta
tulosta Mitä haluat tehdä?
tulosta 1. Tutki metsää
tulosta 2. Tarkista laukkusi
tulosta 3. Lepää
tulosta 4. Lopeta peli
kysy Valitse:

jos vastaus = 1 niin mene tutki
jos vastaus = 2 niin mene tarkista_laukku
jos vastaus = 3 niin mene lepaa
jos vastaus = 4 niin mene lopeta

tulosta Virheellinen valinta!
mene aloita

tutki:
tulosta
tulosta Uskallaudut syvemmälle metsään...
odota 1000

laatikko tapahtuma = satunnainen() % 3

jos laatikko tapahtuma = 0 niin mene loyda_kulta
jos laatikko tapahtuma = 1 niin mene loyda_hirvio
mene loyda_ei_mitaan

loyda_kulta:
laatikko loytyi = satunnainen() % 20 + 10
laatikko kulta = laatikko kulta + laatikko loytyi
tulosta Löysit laatikko loytyi kultakolikkoa! ✨
tulosta Kulta yhteensä: laatikko kulta
odota 2000
mene aloita

loyda_hirvio:
tulosta Hirviö ilmestyy! 👹
laatikko vahinko = satunnainen() % 30 + 10
laatikko terveys = laatikko terveys - laatikko vahinko
tulosta Se hyökkää sinua vastaan laatikko vahinko vahingolla!
tulosta Jäljellä oleva terveys: laatikko terveys

jos laatikko terveys <= 0 niin
   mene peli_ohi
loppu

odota 2000
mene aloita

loyda_ei_mitaan:
tulosta Et löytänyt mitään mielenkiintoista täältä.
odota 1000
mene aloita

tarkista_laukku:
tulosta
tulosta === TILASI ===
tulosta Terveys: laatikko terveys
tulosta Kulta: laatikko kulta
tulosta
jos laatikko kulta >= 100 niin
   tulosta Sinulla on tarpeeksi kultaa voittoon!
   tulosta Onnittelut! 🏆
   mene lopeta
loppu
tulosta Tarvitset 100 kultaa voittaaksesi.
tulosta Jatka tutkimista!
odota 2000
mene aloita

lepaa:
tulosta
tulosta Lepäät hetken...
odota 1500
laatikko parannus = 20
laatikko terveys = laatikko terveys + laatikko parannus
jos laatikko terveys > 100 niin
   laatikko terveys = 100
loppu
tulosta Sait takaisin laatikko parannus terveyttä!
tulosta Terveys: laatikko terveys
odota 1500
mene aloita

peli_ohi:
tulosta
tulosta ================================
tulosta       PELI OHI
tulosta ================================
tulosta Sinut voitettiin!
tulosta Lopullinen kulta: laatikko kulta
tulosta
mene lopeta

lopeta:
tulosta
tulosta Kiitos pelaamisesta!
tulosta Näkemiin! 👋
