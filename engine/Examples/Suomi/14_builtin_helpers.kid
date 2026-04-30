FI
// Valmiit apurit toimivat tekstin, tiedostojen ja listojen kanssa.

laatikko lemmikki = sateenkaari robotti
tulosta sisaltaa(laatikko lemmikki, kaari)
tulosta korvaa(laatikko lemmikki, robotti, leija)
tulosta pala(laatikko lemmikki, 1, 11)

tiedosto vihko = muistiinpanot
avaa vihko helpernotes.txt
kirjoita vihko omena|banaani|porkkana
tulosta onkotiedosto(vihko)
laatikko vihko_teksti = luetiedosto(vihko)
tulosta laatikko vihko_teksti
sulje vihko

lista lelut
laatikko eka = laitalistaan(lelut, 1, robotti)
laatikko toka = laitalistaan(lelut, 2, leija)
laatikko avaimet = listanavaimet(lelut)
tulosta yhdista(laatikko avaimet, :)
tulosta listanpituus(lelut)
