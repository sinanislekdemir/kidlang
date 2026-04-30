# 🎮 KidLang'e Hoş Geldin! Kodlama Maceran Burada Başlıyor!

**Merhaba, geleceğin programcısı!** 👋

Kod yazmayı öğrenmeye hazır mısın? KidLang, senin gibi çocuklar için (8-13 yaş) özel olarak yapılmış bir programlama dilidir! Süper eğlenceli ve öğrenmesi kolay. Hadi kodlama yolculuğuna başlayalım!

---

## 🌟 Programlama Nedir?

Programlama, bilgisayara talimatlar vermek gibidir. Tıpkı kurabiye yapmak için tarif takip ettiğin gibi, bilgisayar da senin kodunu takip ederek harika şeyler yapar!

---

## 🎯 İlk Programın - Merhaba De!

Hadi ilk programını yazalım! Bunu yaz:

```kidlang
yaz Merhaba, Dünya!
yaz Kod yazmayı öğreniyorum!
```

**Ne olur?** Bilgisayar mesajını ekrana yazar (gösterir)! 🎉

> **Eğlenceli Bilgi:** Basit mesajlar için `yaz` yazmana bile gerek yok. Sadece metni yaz ve çalışır! Dene: `Selam!`

---

## 📦 Kutu Kullanmak (Değişkenler)

Bir **kutu**yu, içine bir şeyler koyabileceğin bir kap gibi düşün. Kutuya sayı, kelime veya istediğin herhangi bir şey koyabilirsin!

### Sayı Saklamak

```kidlang
kutu yas = 10
yaz Ben kutu yas yasindayim
```

### Kutularla Matematik Yapmak

```kidlang
kutu elmalar = 5
kutu portakallar = 3
kutu toplam = kutu elmalar + kutu portakallar
yaz Bende kutu toplam meyve var!
```

**Yapabileceğin Harika Şeyler:**
- **Toplama:** `kutu a = 10 + 5` → Sonuç: 15
- **Çıkarma:** `kutu b = 20 - 8` → Sonuç: 12
- **Çarpma:** `kutu c = 4 * 3` → Sonuç: 12
- **Bölme:** `kutu d = 15 / 3` → Sonuç: 5

---

## 💬 Programınla Konuşmak (Girdi Almak)

Programının sana sorular sormasını ister misin? **sor** kullan!

```kidlang
sor Adin ne?
yaz Merhaba kutu cevap
yaz Tanistigimiza memnun oldum!
```

**Ne olur?** 
1. Program adını sorar
2. Adını yazarsın
3. Cevabın `cevap` adında özel bir kutuya kaydedilir
4. Program sana merhaba der!

### Bunu Dene: Yaş Hesaplayıcı

```kidlang
sor Kac yasindasin?
kutu benim_yasim = cevap
kutu gelecek_yil = kutu benim_yasim + 1
yaz Gelecek yil kutu gelecek_yil yasinda olacaksin!
```

---

## 🤔 Karar Vermek (If İfadeleri)

Bazen programının seçim yapmasını istersin. **eğer/ise/son** kullan!

```kidlang
kutu puan = 85

eğer kutu puan > 80 ise
yaz Harika is! A aldin!
son
```

### Sayı Tahmin Oyunu

```kidlang
sor 1 ile 10 arasinda bir sayi tahmin et:
kutu tahmin = cevap
kutu gizli = 7

eğer kutu tahmin = kutu gizli ise
yaz Kazandin! Gizli sayi kutu gizli idi
son

eğer kutu tahmin != kutu gizli ise
yaz Uzgunum! Tekrar dene
son
```

**Karşılaştırma Sembolleri:**
- `=` "eşittir" demek
- `!=` "eşit DEĞİLDİR" demek
- `>` "büyüktür" demek
- `<` "küçüktür" demek
- `>=` "büyük veya eşittir" demek
- `<=` "küçük veya eşittir" demek

---

## 🔄 Şeyleri Tekrarlamak (Etiketlerle Döngüler)

Bir şeyi tekrar tekrar yapmak ister misin? **etiketler** ve **git** kullan!

```kidlang
kutu sayac = 1

basla:
yaz Sayma: kutu sayac
kutu sayac = kutu sayac + 1

eğer kutu sayac < 6 ise git basla

yaz Saymak bitti!
```

**Ne olur?** Bu 1'den 5'e kadar sayar!

### Geri Sayım

```kidlang
kutu zaman = 10

geri_say:
yaz kutu zaman
uyku 1
kutu zaman = kutu zaman - 1

eğer kutu zaman > 0 ise git geri_say

yaz Firlatildi!
```

---

## 🎲 Matematik Fonksiyonlarıyla Eğlence

KidLang'in özel matematik güçleri var!

```kidlang
// Karekök (hangi sayı kendisiyle çarpılınca bunu verir?)
kutu a = karekok(16)
yaz kutu a
// Sonuç: 4 (çünkü 4 × 4 = 16)

// Kare (bir sayıyı kendisiyle çarp)
kutu b = kare(5)
yaz kutu b
// Sonuç: 25 (çünkü 5 × 5 = 25)

// Mutlak değer (eksi işaretini kaldır)
kutu c = mod(-10)
yaz kutu c
// Sonuç: 10

// 0 ile 1 arasında rastgele sayı
kutu d = rastgele()
yaz Sonuc: kutu d
```

---

## 📝 Kelimelerle Çalışmak (Metinler)

Kelimelerle de harika şeyler yapabilirsin!

### Kelimeleri Birleştirmek

```kidlang
kutu ilk = Merhaba
kutu ikinci = Dunya
kutu birlikte = kutu ilk + kutu ikinci
yaz kutu birlikte
// Sonuç: MerhabaDunya
```

### Kelimeleri Tekrarlamak

```kidlang
kutu gulme = Ha * 5
yaz kutu gulme
// Sonuç: HaHaHaHaHa
```

### Bir Harf Almak

```kidlang
kutu kelime = Pizza
kutu harf = kutu kelime / 1
yaz kutu harf
// Sonuç: P (ilk harf!)
```

---

## 📚 Liste Kullanmak (Yığınlar)

Bir **liste**, her biri numara veya isimle işaretlenmiş bir sürü şey tutabilen bir kutu gibidir!

```kidlang
liste oyuncaklar
liste oyuncaklar[1] = Robot
liste oyuncaklar[2] = Top
liste oyuncaklar[3] = Yapboz

yaz Ilk oyuncagim: liste oyuncaklar[1]
yaz Ikinci oyuncagim: liste oyuncaklar[2]
yaz Ucuncu oyuncagim: liste oyuncaklar[3]
```

Etiket olarak kelimeler de kullanabilirsin:

```kidlang
liste arkadas
liste arkadas[isim] = Ahmet
liste arkadas[yas] = 10
liste arkadas[hobi] = Futbol

yaz Isim: liste arkadas[isim]
yaz Yas: liste arkadas[yas]
yaz Hobi: liste arkadas[hobi]
```

---

## 🎮 Mini Proje: Çarpım Testi

Öğrendiklerini eğlenceli bir test oyununda birleştirelim!

```kidlang
yaz === CARPIM TESTI ===

sor 7 kere 8 kac eder?
kutu cevap1 = cevap

eğer kutu cevap1 = 56 ise
yaz Dogru! Harika is!
son

eğer kutu cevap1 != 56 ise
yaz Tam degil! Cevap 56
son

sor 9 kere 6 kac eder?
kutu cevap2 = cevap

eğer kutu cevap2 = 54 ise
yaz Mukemmel! Matematik yildizi olacaksin!
son

eğer kutu cevap2 != 54 ise
yaz Dogru cevap 54
son

yaz Oynadigin icin tesekkurler!
```

---

## 🎨 Mini Proje: Hikaye Yaratıcı

```kidlang
yaz Haydi komik bir hikaye yaratalim!

sor En sevdigin hayvan nedir?
kutu hayvan = cevap

sor En sevdigin yemek nedir?
kutu yemek = cevap

sor En sevdigin renk nedir?
kutu renk = cevap

yaz ================
yaz HIKAYEN:
yaz ================
yaz Bir zamanlar kutu renk renkli bir kutu hayvan varmis
yaz Bu kutu hayvan her gun kutu yemek yemeyi cok severmis!
yaz Bir gun kutu hayvan sihirli bir kutu yemek bulmus
yaz Ve sonsuza kadar mutlu yasamis!
yaz ================
```

---

## 🏆 Senin İçin Meydan Okuma Projeleri!

Artık temelleri biliyorsun, bu eğlenceli projeleri yapmayı dene:

### 1. 🎯 Basit Hesap Makinesi
İki sayı soran ve bunları toplayan bir program yap!

### 2. 🌡️ Sıcaklık Dönüştürücü
Fahrenheit'ı Celsius'a çevir!
(İpucu: Celsius = (Fahrenheit - 32) × 5 / 9)

### 3. 🎲 Zar Atıcı
Zar atmayı simüle etmek için `rastgele()` kullan!

### 4. 📊 Not Hesaplayıcı
Sınav notlarını sor ve ortalamayı hesapla!

### 5. 🎪 Lunapark Oyunu
Birden fazla şansı olan bir sayı tahmin oyunu yarat!

---

## 💡 Genç Programcılar İçin İpuçları

1. **Hatalardan korkma!** Herkes hata yapar. Düzelt ve öğren!
2. **Deneyler yap!** Sayıları ve kelimeleri değiştir, ne olduğunu gör
3. **Küçük başla!** Önce basit programlar yap, sonra özellikler ekle
4. **Eğlen!** Programlama bulmaca çözmek gibi keyifli olmalı
5. **Çalışmanı kaydet!** Programlarına `oyunum.kid` gibi isimler ver

---

## 🎓 Neler Öğrendin!

✅ Mesajları nasıl yazdıracağını  
✅ Şeyleri saklamak için kutuları (değişkenleri) nasıl kullanacağını  
✅ Nasıl matematik yapılacağını (+, -, ×, ÷)  
✅ Nasıl soru sorulup cevap alınacağını  
✅ Eğer/ise ile nasıl karar verileceğini  
✅ Etiketler ve git ile nasıl tekrar yapılacağını  
✅ Listeleri (yığınları) nasıl kullanacağını  
✅ Eğlenceli projeler nasıl yapılacağını!  

---

## 🚀 Sırada Ne Var?

Daha fazla öğrenmek ister misin? Bu dosyalara göz at:

- **TUTORIAL_BEGINNER.md** - Daha fazla başlangıç dersi
- **TUTORIAL_ALGORITHMS.md** - Sıralama ve arama öğren
- **TUTORIAL_PROJECTS.md** - Eksiksiz projeler yap
- **examples/** klasörü - 20+ örnek programı gör!

---

## 🎉 Artık Bir Programcısın!

Tebrikler! KidLang'de kod yazmayı öğrendin! Pratik yapmaya devam et, yaratmaya devam et ve en önemlisi—**kodlama yaparken eğlen!** 🌟

Unutma: Her uzman programcı tam olarak senin şu anda olduğun yerden başladı. Harika gidiyorsun! 💪

---

**Mutlu Kodlamalar! 🎮✨**
