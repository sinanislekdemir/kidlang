FI
lista pisteet
lista nimet
tulosta === Arviointikirja ===
laatikko maara = 0

valikko:
tulosta
tulosta 1 Lisää opiskelija
tulosta 2 Näytä kaikki opiskelijat
tulosta 3 Lopeta
kysy Valitse:
laatikko valinta = vastaus

jos laatikko valinta = 1 niin
kysy Opiskelijan nimi:
laatikko nimi = vastaus
kysy Opiskelijan pisteet:
laatikko pisteet_arvo = vastaus
laatikko maara = laatikko maara + 1
lista nimet[laatikko maara] = laatikko nimi
lista pisteet[laatikko maara] = laatikko pisteet_arvo
tulosta Lisätty!
mene valikko
loppu

jos laatikko valinta = 2 niin
tulosta
tulosta === Kaikki opiskelijat ===
laatikko i = 1
nayta:
jos laatikko i <= laatikko maara niin
tulosta lista nimet[laatikko i]: lista pisteet[laatikko i]
laatikko i = laatikko i + 1
mene nayta
loppu
mene valikko
loppu

jos laatikko valinta = 3 niin
tulosta Näkemiin!
loppu
