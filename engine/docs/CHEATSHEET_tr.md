# Kidlang Kopya Kağıdı (Türkçe)

## Dil Seçimi
```
TR
```

## Yorumlar
```
// Bu bir yorum satırı
```

## Değişkenler

### Kutu (basit değişken)
```
kutu isim = 42
kutu isim = 3.14
kutu isim = merhaba
kutu isim = kutu diger
```

### Liste (sözlük/harita)
```
liste oyuncaklar
liste oyuncaklar[1] = araba    // veya oyuncaklar(1) = araba
liste oyuncaklar[2] = top      // veya oyuncaklar(2) = top
oyuncaklar[1] = kamyon         // veya oyuncaklar(1) = kamyon
yaz oyuncaklar[1]              // veya yaz oyuncaklar(1)
```

### Dosya
```
dosya dosyam
```

## Çıktı
```
yaz Merhaba Dünya          // Tırnak işareti isteğe bağlı
yaz "Merhaba Dünya"        // Tırnaklar metin değerini korur
yaz kutu isim              // 'isim' değişkeninin değerini yazar
yaz "kutu isim"            // "kutu isim" metnini yazar
yaz 1 + 2
Komut olmayan metin        // Örtük yazdırma
```

## Girdi
```
sor Adın ne?               // Kullanıcıdan girdi ister
yaz cevap                  // Sonuç 'cevap' değişkeninde saklanır
sor "Değer gir:"           // Tırnak işareti isteğe bağlı
kutu x = cevap             // Cevabı kullan
```

## Matematiksel İşlemler
```
kutu sonuc = 5 + 3
kutu sonuc = 10 - 2
kutu sonuc = 4 * 3
kutu sonuc = 10 / 2
kutu sonuc = 10 % 3
kutu sonuc = 2 ^ 3  // Sayılar için XOR, metinler için şifreleme
```

## Matematik Fonksiyonları
```
karekok 16    // Karekök
mod -5        // Mutlak değer
kare 4        // Kare (4*4)
sin 1.57      // Sinüs
cos 0         // Kosinüs
tan 0.785     // Tanjant
log 2.718     // Doğal logaritma
asin 0.5      // Arksinüs
acos 0.5      // Arkkosinüs
```

## Metin İşlemleri
```
kutu metin = merhaba + dünya   // Birleştirme (tırnak isteğe bağlı)
kutu metin = merhaba - a       // Tüm 'a' karakterlerini çıkar
kutu metin = abc * 3           // Metni tekrarla
kutu harf = merhaba / 2        // 2. indeksteki karakteri al
```

## Koşullar
```
eger kutu x = 5 ise yaz x beştir
son

eger kutu x > 10 ise
yaz x ondan büyüktür
son

eger kutu x < 5 ise git atla
```

### Karşılaştırma Operatörleri
```
=    // Eşit
!=   // Eşit değil
>    // Büyüktür
<    // Küçüktür
>=   // Büyük eşit
<=   // Küçük eşit
```

### Mantıksal Operatörler
```
eger kutu x > 5 ve kutu y < 10 ise yaz ikisi de doğru
son

eger kutu a = 1 veya kutu b = 2 ise yaz biri doğru
son
```

## Etiketler ve Atlamalar
```
basla:
yaz "Merhaba"
git basla

atla:
yaz "Buraya atlandı"
```

## Dosya İşlemleri
```
dosya dosyam
ac dosyam veri.txt         // Dosyayı açar/oluşturur
oku dosyam kutu icerik     // Tüm dosyayı oku
satiroku dosyam kutu satir // Bir satır oku
yaz dosyam bir metin       // Dosyaya yaz
sira dosyam 5              // 5. satıra git
kapat dosyam               // Dosyayı kapat
```

## Özel Değerler
```
rastgele  // Rastgele tamsayı
tarih     // Güncel tarih/saat
\n        // Yeni satır karakteri
```

## Sistem Komutları
```
calistir ls -la        // Kabuk komutu çalıştır (tırnak isteğe bağlı)
bekle 1000             // 1000 milisaniye bekle
```

## Çalışma Akışı
- Programlar yukarıdan aşağıya satır satır çalışır
- `git` bir etikete atlar
- `eger...ise...son` koşullu bloklar oluşturur
- `son` bir `eger` bloğunu kapatır
