FI
// Keski-Maa Kauppias - Kauppaseikkailu
// Osta halvalla, myy kalliilla, ansaitse 1000 kultaa!

funktio markkinahinta(laatikko perus, laatikko vaihtelu)
satunnainen() % laatikko vaihtelu + laatikko perus
loppu

funktio kokonaishinta(laatikko hinta, laatikko maara)
laatikko hinta * laatikko maara
loppu

funktio haarniska_tila(laatikko maara)
laatikko maara * 2
loppu

funktio vapaa_tila(laatikko laukku, laatikko kaytetty)
laatikko laukku - laatikko kaytetty
loppu

funktio turvallinen_rosvo_kulta(laatikko kulta)
jos laatikko kulta > 30 niin
   palauta laatikko kulta - 30
loppu
0
loppu

funktio kauppias_titteli(laatikko kulta)
jos laatikko kulta >= 1000 niin
   palauta Mestarikauppias
loppu

jos laatikko kulta >= 500 niin
   palauta Vahva Kauppias
loppu

Oppiva Kauppias
loppu

tulosta ====================================
tulosta   KESKI-MAA KAUPPIAS
tulosta ====================================
tulosta
tulosta Olet kauppias Keski-Maassa
tulosta Tavoite: Ansaitse 1000 kultaa 30 päivässä!
tulosta

laatikko paiva = 1
laatikko kulta = 100
laatikko laukku = 50
laatikko kaytetty = 0
laatikko juomat = 0
laatikko sauvat = 0
laatikko haarniska = 0

alku:
tulosta
tulosta === PAIVA laatikko paiva of 30 ===
tulosta Kulta: laatikko kulta
tulosta Laukku: laatikko kaytetty of laatikko laukku tilaa
tulosta Vapaa tila: vapaa_tila(laatikko laukku, laatikko kaytetty)
tulosta Varasto - Juomat: laatikko juomat | Sauvat: laatikko sauvat | Haarniska: laatikko haarniska
tulosta

laatikko j_hinta = markkinahinta(20, 30)
laatikko s_hinta = markkinahinta(40, 50)  
laatikko h_hinta = markkinahinta(60, 80)

tulosta === MARKKINAHINNAT ===
tulosta Juomat: laatikko j_hinta kultaa
tulosta Sauvat: laatikko s_hinta kultaa
tulosta Haarniska: laatikko h_hinta kultaa (vie 2 tilaa)
tulosta
tulosta === TOIMINNOT ===
tulosta 1. Osta juomia
tulosta 2. Osta sauvoja
tulosta 3. Osta haarniskaa
tulosta 4. Myy juomia
tulosta 5. Myy sauvoja
tulosta 6. Myy haarniskaa
tulosta 7. Matkusta seuraavaan kaupunkiin
kysy Valitse:

laatikko valinta = vastaus

jos laatikko valinta = 1 niin mene osta_juomat
jos laatikko valinta = 2 niin mene osta_sauvat
jos laatikko valinta = 3 niin mene osta_haarniska
jos laatikko valinta = 4 niin mene myy_juomat
jos laatikko valinta = 5 niin mene myy_sauvat
jos laatikko valinta = 6 niin mene myy_haarniska
jos laatikko valinta = 7 niin mene matkusta
tulosta Virheellinen valinta!
odota 1000
mene alku

osta_juomat:
kysy Kuinka monta juomaa?
laatikko maara = vastaus
laatikko hinta = kokonaishinta(laatikko j_hinta, laatikko maara)
laatikko tarve = laatikko kaytetty + laatikko maara
jos laatikko hinta > laatikko kulta niin
   tulosta Ei tarpeeksi kultaa!
   odota 1000
   mene alku
loppu
jos laatikko tarve > laatikko laukku niin
   tulosta Ei tarpeeksi laukkutilaa!
   odota 1000
   mene alku
loppu
laatikko kulta = laatikko kulta - laatikko hinta
laatikko juomat = laatikko juomat + laatikko maara
laatikko kaytetty = laatikko kaytetty + laatikko maara
tulosta Ostit laatikko maara juomaa laatikko hinta kullalla!
odota 1000
mene alku

osta_sauvat:
kysy Kuinka monta sauvaa?
laatikko maara = vastaus
laatikko hinta = kokonaishinta(laatikko s_hinta, laatikko maara)
laatikko tarve = laatikko kaytetty + laatikko maara
jos laatikko hinta > laatikko kulta niin
   tulosta Ei tarpeeksi kultaa!
   odota 1000
   mene alku
loppu
jos laatikko tarve > laatikko laukku niin
   tulosta Ei tarpeeksi laukkutilaa!
   odota 1000
   mene alku
loppu
laatikko kulta = laatikko kulta - laatikko hinta
laatikko sauvat = laatikko sauvat + laatikko maara
laatikko kaytetty = laatikko kaytetty + laatikko maara
tulosta Ostit laatikko maara sauvaa laatikko hinta kullalla!
odota 1000
mene alku

osta_haarniska:
kysy Kuinka monta haarniskaa?
laatikko maara = vastaus
laatikko hinta = kokonaishinta(laatikko h_hinta, laatikko maara)
laatikko tarve = laatikko kaytetty + haarniska_tila(laatikko maara)
jos laatikko hinta > laatikko kulta niin
   tulosta Ei tarpeeksi kultaa!
   odota 1000
   mene alku
loppu
jos laatikko tarve > laatikko laukku niin
   tulosta Ei tarpeeksi laukkutilaa!
   odota 1000
   mene alku
loppu
laatikko kulta = laatikko kulta - laatikko hinta
laatikko haarniska = laatikko haarniska + laatikko maara
laatikko kaytetty = laatikko kaytetty + haarniska_tila(laatikko maara)
tulosta Ostit laatikko maara haarniskaa laatikko hinta kullalla!
odota 1000
mene alku

myy_juomat:
kysy Kuinka monta juomaa?
laatikko maara = vastaus
jos laatikko maara > laatikko juomat niin
   tulosta Sinulla ei ole niin montaa!
   odota 1000
   mene alku
loppu
laatikko ansio = kokonaishinta(laatikko j_hinta, laatikko maara)
laatikko kulta = laatikko kulta + laatikko ansio
laatikko juomat = laatikko juomat - laatikko maara
laatikko kaytetty = laatikko kaytetty - laatikko maara
tulosta Myit laatikko maara juomaa laatikko ansio kullalla!
odota 1000
mene alku

myy_sauvat:
kysy Kuinka monta sauvaa?
laatikko maara = vastaus
jos laatikko maara > laatikko sauvat niin
   tulosta Sinulla ei ole niin montaa!
   odota 1000
   mene alku
loppu
laatikko ansio = kokonaishinta(laatikko s_hinta, laatikko maara)
laatikko kulta = laatikko kulta + laatikko ansio
laatikko sauvat = laatikko sauvat - laatikko maara
laatikko kaytetty = laatikko kaytetty - laatikko maara
tulosta Myit laatikko maara sauvaa laatikko ansio kullalla!
odota 1000
mene alku

myy_haarniska:
kysy Kuinka monta haarniskaa?
laatikko maara = vastaus
jos laatikko maara > laatikko haarniska niin
   tulosta Sinulla ei ole niin montaa!
   odota 1000
   mene alku
loppu
laatikko ansio = kokonaishinta(laatikko h_hinta, laatikko maara)
laatikko kulta = laatikko kulta + laatikko ansio
laatikko haarniska = laatikko haarniska - laatikko maara
laatikko kaytetty = laatikko kaytetty - haarniska_tila(laatikko maara)
tulosta Myit laatikko maara haarniskaa laatikko ansio kullalla!
odota 1000
mene alku

matkusta:
tulosta
tulosta Matkustetaan seuraavaan kaupunkiin...
odota 1000
laatikko paiva = laatikko paiva + 1

laatikko tapahtuma = satunnainen() % 8

jos laatikko tapahtuma = 0 niin
   tulosta Ystävällinen velho antaa sinulle 50 kultaa! ✨
   laatikko kulta = laatikko kulta + 50
   odota 1500
loppu

jos laatikko tapahtuma = 1 niin
   tulosta Rosvot hyökkäävät! Menetit 30 kultaa! ⚔️
   laatikko kulta = turvallinen_rosvo_kulta(laatikko kulta)
   odota 1500
loppu

jos laatikko tapahtuma = 2 niin
   tulosta Löysit juoman tieltä! 🧪
   jos laatikko kaytetty < laatikko laukku niin
      laatikko juomat = laatikko juomat + 1
      laatikko kaytetty = laatikko kaytetty + 1
   loppu
   odota 1500
loppu

jos laatikko kulta >= 1000 niin mene voitto
jos laatikko paiva > 30 niin mene havisi

mene alku

voitto:
tulosta
tulosta ====================================
tulosta   ONNITTELUT!
tulosta ====================================
tulosta Ansaitsit 1000 kultaa laatikko paiva päivässä!
tulosta Titteli: kauppias_titteli(laatikko kulta)
tulosta Olet valmis seuraaville markkinoille! 🏆
tulosta
mene loppu

havisi:
tulosta
tulosta ====================================
tulosta   AIKA LOPPUI!
tulosta ====================================
tulosta 30 päivää on kulunut...
tulosta Lopullinen kulta: laatikko kulta
tulosta Titteli: kauppias_titteli(laatikko kulta)
tulosta
jos laatikko kulta >= 500 niin
   tulosta Ei hullumpaa kauppiaalle!
loppu
jos laatikko kulta < 500 niin
   tulosta Jatka kaupankäynnin harjoittelua!
loppu
tulosta
mene loppu

loppu:
tulosta Kiitos Keski-Maa Kauppias pelin pelaamisesta!
tulosta Näkemiin! 👋
