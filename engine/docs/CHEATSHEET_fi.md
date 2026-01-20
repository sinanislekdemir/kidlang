# Kidlang Muistilista (Suomi)

## Kielen Valinta
```
FI
```

## Kommentit
```
// Tämä on kommentti
```

## Muuttujat

### Laatikko (yksinkertainen muuttuja)
```
laatikko nimi = 42
laatikko nimi = 3.14
laatikko nimi = terve
laatikko nimi = laatikko toinen
```

### Lista (sanakirja/kartta)
```
lista lelut
lista lelut[1] = auto          // tai lelut(1) = auto
lista lelut[2] = pallo         // tai lelut(2) = pallo
lelut[1] = kuorma-auto         // tai lelut(1) = kuorma-auto
tulosta lelut[1]               // tai tulosta lelut(1)
```

### Tiedosto
```
tiedosto tiedostoni
```

## Tuloste
```
tulosta Hei Maailma            // Lainausmerkit valinnaisia
tulosta "Hei Maailma"          // Lainausmerkit säilyttävät tekstin
tulosta laatikko nimi          // Tulostaa muuttujan 'nimi' arvon
tulosta "laatikko nimi"        // Tulostaa tekstin "laatikko nimi"
tulosta 1 + 2
Mikä tahansa teksti            // Implisiittinen tulostus
```

## Syöte
```
kysy Mikä on nimesi?           // Pyytää käyttäjän syötettä
tulosta vastaus                // Tulos tallennetaan 'vastaus' muuttujaan
kysy "Anna arvo:"              // Lainausmerkit valinnaisia
laatikko x = vastaus           // Käytä vastausta
```

## Matemaattiset Toiminnot
```
laatikko tulos = 5 + 3
laatikko tulos = 10 - 2
laatikko tulos = 4 * 3
laatikko tulos = 10 / 2
laatikko tulos = 10 % 3
laatikko tulos = 2 ^ 3  // XOR numeroille, salaus merkkijonoille
```

## Matemaattiset Funktiot
```
neliojuuri 16     // Neliöjuuri
itseisarvo -5     // Itseisarvo
nelio 4           // Neliö (4*4)
sin 1.57          // Sini
cos 0             // Kosini
tan 0.785         // Tangentti
log 2.718         // Luonnollinen logaritmi
asin 0.5          // Arkussini
acos 0.5          // Arkuskosini
```

## Merkkijono-operaatiot
```
laatikko teksti = hei + maailma    // Yhdistäminen (lainausmerkit valinnaisia)
laatikko teksti = hei - e          // Poista kaikki 'e' merkit
laatikko teksti = abc * 3          // Toista merkkijono
laatikko merkki = hei / 2          // Hae merkki indeksistä 2
```

## Ehdolliset Lauseet
```
jos laatikko x = 5 niin tulosta x on 5
loppu

jos laatikko x > 10 niin
tulosta x on suurempi kuin 10
loppu

jos laatikko x < 5 niin mene ohita
```

### Vertailuoperaattorit
```
=    // Yhtä suuri
!=   // Eri suuri
>    // Suurempi kuin
<    // Pienempi kuin
>=   // Suurempi tai yhtä suuri
<=   // Pienempi tai yhtä suuri
```

### Loogiset Operaattorit
```
jos laatikko x > 5 ja laatikko y < 10 niin tulosta molemmat totta
loppu

jos laatikko a = 1 tai laatikko b = 2 niin tulosta yksi totta
loppu
```

## Nimiöt ja Hyppäykset
```
alku:
tulosta "Hei"
mene alku

ohita:
tulosta "Hypättiin tänne"
```

## Tiedosto-operaatiot
```
tiedosto tiedostoni
avaa tiedostoni data.txt       // Avaa/luo tiedosto
lue tiedostoni laatikko sisalto // Lue koko tiedosto
luerivi tiedostoni laatikko rivi // Lue yksi rivi
kirjoita tiedostoni tekstiä    // Kirjoita tiedostoon
hae tiedostoni 5               // Siirry riville 5
sulje tiedostoni               // Sulje tiedosto
```

## Erikoisarvot
```
satunnainen  // Satunnainen kokonaisluku
aika         // Nykyinen päivämäärä/aika
\n           // Rivinvaihto
```

## Järjestelmäkomennot
```
suorita ls -la         // Suorita kuorikomento (lainausmerkit valinnaisia)
odota 1000             // Odota 1000 millisekuntia
```

## Suorituksen Kulku
- Ohjelmat suorittavat riviltä riville ylhäältä alas
- `mene` hyppää nimiöön
- `jos...niin...loppu` luo ehdollisia lohkoja
- `loppu` sulkee `jos` lohkon
