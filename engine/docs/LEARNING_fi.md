# 🎮 Tervetuloa KidLangiin! Koodausseikkailusi alkaa tästä!

**Hei tuleva ohjelmoija!** 👋

Oletko valmis oppimaan koodaamaan? KidLang on erityinen ohjelmointikieli, joka on tehty juuri sinun kaltaisillesi lapsille (8-13-vuotiaille)! Se on tosi hauskaa ja helppoa oppia. Aloitetaan koodausseikkailusi!

---

## 🌟 Mikä on ohjelmointi?

Ohjelmointi on kuin ohjeiden antamista tietokoneelle. Aivan kuten seuraat reseptiä leipoessasi keksejä, tietokone seuraa koodiasi tehdäkseen hienoja asioita!

---

## 🎯 Ensimmäinen ohjelmasi - Sano terve!

Kirjoitetaan ensimmäinen ohjelmasi! Kirjoita tämä:

```kidlang
tulosta Hei, maailma!
tulosta Opin ohjelmoimaan!
```

**Mitä tapahtuu?** Tietokone tulostaa (näyttää) viestisi ruudulla! 🎉

> **Hauska fakta:** Sinun ei edes tarvitse kirjoittaa `tulosta` yksinkertaisille viesteille. Kirjoita vain tekstiä ja se toimii! Kokeile: `Terve kaikille!`

---

## 📦 Laatikoiden käyttäminen (Muuttujat)

Ajattele **laatikkoa** kuin säiliötä, johon voit laittaa tavaroita. Voit laittaa laatikkoon numeron, sanan tai mitä tahansa haluat!

### Numeron tallentaminen

```kidlang
laatikko ika = 10
tulosta Olen laatikko ika vuotta vanha
```

### Matematiikan tekeminen laatikoilla

```kidlang
laatikko omenoita = 5
laatikko appelsiineja = 3
laatikko yhteensa = laatikko omenoita + laatikko appelsiineja
tulosta Minulla on laatikko yhteensa hedelmää!
```

**Hienoja asioita, joita voit tehdä:**
- **Yhteenlasku:** `laatikko a = 10 + 5` → Tulos: 15
- **Vähennyslasku:** `laatikko b = 20 - 8` → Tulos: 12
- **Kertolasku:** `laatikko c = 4 * 3` → Tulos: 12
- **Jakolasku:** `laatikko d = 15 / 3` → Tulos: 5

---

## 💬 Ohjelmasi kanssa puhuminen (Syötteen saaminen)

Haluatko ohjelmasi kysyvän sinulta kysymyksiä? Käytä **kysy**-komentoa!

```kidlang
kysy Mikä on nimesi?
tulosta Hei laatikko vastaus
tulosta Hauska tavata!
```

**Mitä tapahtuu?** 
1. Ohjelma kysyy nimeäsi
2. Kirjoitat nimesi
3. Vastauksesi tallennetaan erityiseen laatikkoon nimeltä `vastaus`
4. Ohjelma tervehtii sinua!

### Kokeile tätä: Ikälaskuri

```kidlang
kysy Kuinka vanha olet?
laatikko minun_ikani = vastaus
laatikko ensi_vuonna = laatikko minun_ikani + 1
tulosta Ensi vuonna olet laatikko ensi_vuonna vuotta vanha!
```

---

## 🤔 Päätösten tekeminen (Jos-lauseet)

Joskus haluat ohjelmasi tekevän valintoja. Käytä **jos/niin/loppu**!

```kidlang
laatikko pisteet = 85

jos laatikko pisteet > 80 niin
tulosta Hienoa työtä! Sait kympin!
loppu
```

### Numeronarvauspeli

```kidlang
kysy Arvaa numero 1:n ja 10:n väliltä:
laatikko arvaus = vastaus
laatikko salainen = 7

jos laatikko arvaus = laatikko salainen niin
tulosta Voitit! Numero oli laatikko salainen
loppu

jos laatikko arvaus != laatikko salainen niin
tulosta Voi ei! Yritä uudelleen
loppu
```

**Vertailusymbolit:**
- `=` tarkoittaa "on yhtä suuri kuin"
- `!=` tarkoittaa "EI ole yhtä suuri kuin"
- `>` tarkoittaa "suurempi kuin"
- `<` tarkoittaa "pienempi kuin"
- `>=` tarkoittaa "suurempi tai yhtä suuri kuin"
- `<=` tarkoittaa "pienempi tai yhtä suuri kuin"

---

## 🔄 Asioiden toistaminen (Silmukat merkinnöillä)

Haluatko tehdä jotain yhä uudelleen ja uudelleen? Käytä **merkintöjä** ja **mene**!

```kidlang
laatikko laskuri = 1

alku:
tulosta Lasketaan: laatikko laskuri
laatikko laskuri = laatikko laskuri + 1

jos laatikko laskuri < 6 niin mene alku

tulosta Laskeminen valmis!
```

**Mitä tapahtuu?** Tämä laskee 1:stä 5:een!

### Lähtölaskenta

```kidlang
laatikko aika = 10

laskenta:
tulosta laatikko aika
nuku 1
laatikko aika = laatikko aika - 1

jos laatikko aika > 0 niin mene laskenta

tulosta Laukaisu!
```

---

## 🎲 Hauskaa matematiikkafunktioilla

KidLangilla on erityisiä matematiikkavoimia!

```kidlang
// Neliöjuuri (mikä luku kerrottuna itsellään antaa tämän?)
laatikko a = neliojuuri(16)
tulosta laatikko a
// Tulos: 4 (koska 4 × 4 = 16)

// Neliö (kerro luku itsellään)
laatikko b = nelio(5)
tulosta laatikko b
// Tulos: 25 (koska 5 × 5 = 25)

// Itseisarvo (poista miinusmerkki)
laatikko c = itseisarvo(-10)
tulosta laatikko c
// Tulos: 10

// Satunnaisluku 0:n ja 1:n väliltä
laatikko d = satunnainen()
tulosta Sait: laatikko d
```

---

## 📝 Sanojen kanssa työskentely (Merkkijonot)

Voit tehdä hienoja asioita myös sanoilla!

### Sanojen yhdistäminen

```kidlang
laatikko eka = Hei
laatikko toka = Maailma
laatikko yhdessa = laatikko eka + laatikko toka
tulosta laatikko yhdessa
// Tulos: HeiMaailma
```

### Sanojen toistaminen

```kidlang
laatikko nauru = Ha * 5
tulosta laatikko nauru
// Tulos: HaHaHaHaHa
```

### Yhden kirjaimen saaminen

```kidlang
laatikko sana = Pizza
laatikko kirjain = laatikko sana / 1
tulosta laatikko kirjain
// Tulos: P (ensimmäinen kirjain!)
```

---

## 📚 Listojen käyttäminen (Pinot)

**Pino** on kuin laatikko, joka voi pitää sisällään monia asioita, joilla jokaisella on numero tai nimi!

```kidlang
pino lelut
pino lelut[1] = Robotti
pino lelut[2] = Pallo
pino lelut[3] = Palapeli

tulosta Ensimmäinen leluni on: pino lelut[1]
tulosta Toinen leluni on: pino lelut[2]
tulosta Kolmas leluni on: pino lelut[3]
```

Voit käyttää myös sanoja merkintöinä:

```kidlang
pino kaveri
pino kaveri[nimi] = Matti
pino kaveri[ika] = 10
pino kaveri[harrastus] = Jalkapallo

tulosta Nimi: pino kaveri[nimi]
tulosta Ikä: pino kaveri[ika]
tulosta Harrastus: pino kaveri[harrastus]
```

---

## 🎮 Miniprojekti: Kertolaskuvisa

Yhdistetään opitut asiat hauskaksi visapeliksi!

```kidlang
tulosta === KERTOLASKUVISA ===

kysy Mikä on 7 kertaa 8?
laatikko vastaus1 = vastaus

jos laatikko vastaus1 = 56 niin
tulosta Oikein! Hienoa työtä!
loppu

jos laatikko vastaus1 != 56 niin
tulosta Ei aivan! Vastaus on 56
loppu

kysy Mikä on 9 kertaa 6?
laatikko vastaus2 = vastaus

jos laatikko vastaus2 = 54 niin
tulosta Täydellistä! Olet matematiikkatähti!
loppu

jos laatikko vastaus2 != 54 niin
tulosta Oikea vastaus on 54
loppu

tulosta Kiitos pelaamisesta!
```

---

## 🎨 Miniprojekti: Tarinankertoja

```kidlang
tulosta Luodaan hauska tarina!

kysy Mikä on lempieläimesi?
laatikko elain = vastaus

kysy Mikä on lempiruokasi?
laatikko ruoka = vastaus

kysy Mikä on lempivärisi?
laatikko vari = vastaus

tulosta ================
tulosta TARINASI:
tulosta ================
tulosta Olipa kerran laatikko vari -värinen laatikko elain
tulosta Tämä laatikko elain rakasti syödä laatikko ruoka joka päivä!
tulosta Eräänä päivänä laatikko elain löysi maagisen laatikko ruoka
tulosta Ja he elivät onnellisina elämänsä loppuun!
tulosta ================
```

---

## 🏆 Haasteita sinulle!

Nyt kun tiedät perusteet, kokeile rakentaa nämä hauskat projektit:

### 1. 🎯 Yksinkertainen laskin
Tee ohjelma, joka kysyy kaksi numeroa ja laskee ne yhteen!

### 2. 🌡️ Lämpötilamuunnin
Muunna lämpötiloja Fahrenheitista Celsiuksiksi!
(Vihje: Celsius = (Fahrenheit - 32) × 5 / 9)

### 3. 🎲 Nopanheittäjä
Käytä `satunnainen()`-funktiota nopanheiton simulointiin!

### 4. 📊 Arvosanalaskuri
Kysy koetuloksia ja laske keskiarvo!

### 5. 🎪 Tivolipeli
Luo numeronarvauspeli, jossa on useita yrityksiä!

---

## 💡 Vinkkejä nuorille ohjelmoijille

1. **Älä pelkää virheitä!** Kaikki tekevät niitä. Korjaa ja opi!
2. **Kokeile!** Vaihda numeroita ja sanoja nähdäksesi mitä tapahtuu
3. **Aloita pienestä!** Tee ensin yksinkertaisia ohjelmia, lisää sitten ominaisuuksia
4. **Pidä hauskaa!** Ohjelmoinnin pitäisi olla nautinnollista, kuten pulmien ratkaiseminen
5. **Tallenna työsi!** Anna ohjelmillesi nimet kuten `minunpelini.kid`

---

## 🎓 Mitä olet oppinut!

✅ Kuinka tulostaa viestejä  
✅ Kuinka käyttää laatikoita (muuttujia) asioiden tallentamiseen  
✅ Kuinka tehdä matematiikkaa (+, -, ×, ÷)  
✅ Kuinka kysyä kysymyksiä ja saada vastauksia  
✅ Kuinka tehdä päätöksiä jos/niin-lauseilla  
✅ Kuinka toistaa asioita merkinnöillä ja mene-komennolla  
✅ Kuinka käyttää listoja (pinoja)  
✅ Kuinka rakentaa hauskoja projekteja!  

---

## 🚀 Mitä seuraavaksi?

Haluatko oppia lisää? Tutustu näihin tiedostoihin:

- **TUTORIAL_BEGINNER.md** - Lisää aloittelijan oppitunteja
- **TUTORIAL_ALGORITHMS.md** - Opi lajittelua ja hakua
- **TUTORIAL_PROJECTS.md** - Rakenna täydellisiä projekteja
- **examples/** -kansio - Katso yli 20 esimerkkiohjelmaa!

---

## 🎉 Olet nyt ohjelmoija!

Onnittelut! Olet oppinut ohjelmoimaan KidLangilla! Jatka harjoittelua, jatka luomista ja tärkeintä—**pidä hauskaa koodatessasi!** 🌟

Muista: Jokainen asiantuntija-ohjelmoija aloitti täsmälleen siitä, missä sinä nyt olet. Teet mainiosti! 💪

---

**Hauskaa koodausta! 🎮✨**
