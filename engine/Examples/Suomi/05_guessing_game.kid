FI
// Numeronarvauspeli
// Yritä arvata salainen numero!

tulosta ================================
tulosta    NUMERONARVAUSPELI
tulosta ================================
tulosta

laatikko salaisuus = satunnainen % 50 + 1
laatikko yritykset = 0
laatikko max_yritykset = 7

tulosta Ajattelen numeroa 1 ja 50 välillä
tulosta Sinulla on laatikko max_yritykset yritystä arvataksesi!
tulosta

peli:
laatikko yritykset = laatikko yritykset + 1

tulosta Yritys laatikko yritykset : laatikko max_yritykset
kysy Anna arvauksesi:

jos vastaus = laatikko salaisuus niin
   mene voitto
loppu

jos vastaus > laatikko salaisuus niin
   tulosta Liian korkea! Kokeile pienempää numeroa.
loppu

jos vastaus < laatikko salaisuus niin
   tulosta Liian matala! Kokeile suurempaa numeroa.
loppu

jos laatikko yritykset >= laatikko max_yritykset niin
   mene häviö
loppu

tulosta
mene peli

voitto:
tulosta
tulosta ================================
tulosta    ONNITTELUT!
tulosta ================================
tulosta Arvasit sen! Numero oli laatikko salaisuus
tulosta Se vei laatikko yritykset yritystä!
mene lopetus

häviö:
tulosta
tulosta ================================
tulosta    PELI LOPPUI
tulosta ================================
tulosta Valitettavasti yritykset loppuivat!
tulosta Salainen numero oli laatikko salaisuus
mene lopetus

lopetus:
tulosta
tulosta Kiitos pelaamisesta!
