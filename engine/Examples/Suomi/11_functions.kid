FI
// Funktiot ovat pieniä apureita, joita voi käyttää uudelleen.

funktio hurraa(laatikko nimi)
Hurraa laatikko nimi !
loppu

funktio lisaa_tahdet(laatikko nykyinen)
palauta laatikko nykyinen + 2
loppu

kysy Mikä sinun nimesi on?
laatikko pelaaja = vastaus

laatikko tahdet = 4
laatikko yhteensa = lisaa_tahdet(laatikko tahdet)

tulosta hurraa(laatikko pelaaja)
tulosta Nyt sinulla on laatikko yhteensa tähteä.
tulosta Nimesi pituus on pituus(laatikko pelaaja).
