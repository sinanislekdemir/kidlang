FI
lista tuotteet
laatikko maara = 0
tulosta === Ostoslista ===
tulosta
valikko:
tulosta Mitä haluat tehdä?
tulosta 1 Lisää tuote
tulosta 2 Näytä lista
tulosta 3 Lopeta
kysy Valitse:
laatikko valinta = vastaus

jos laatikko valinta = 1 niin
kysy Kirjoita tuotteen nimi:
laatikko nimi = vastaus
laatikko maara = laatikko maara + 1
lista tuotteet[laatikko maara] = laatikko nimi
tulosta Lisätty!
mene valikko
loppu

jos laatikko valinta = 2 niin
tulosta
tulosta Ostoslistasi:
laatikko i = 1
nayta:
jos laatikko i <= laatikko maara niin
tulosta laatikko i . lista tuotteet[laatikko i]
laatikko i = laatikko i + 1
mene nayta
loppu
tulosta
mene valikko
loppu

jos laatikko valinta = 3 niin
tulosta Näkemiin!
loppu
