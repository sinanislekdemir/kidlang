TR
liste notlar
liste isimler
yaz === Not Defteri ===
kutu adet = 0

menu:
yaz
yaz 1 Öğrenci ekle
yaz 2 Tüm öğrencileri göster
yaz 3 Çık
sor Seç:
kutu secim = cevap

eger kutu secim = 1 ise
sor Öğrenci adı:
kutu isim = cevap
sor Öğrenci notu:
kutu not = cevap
kutu adet = kutu adet + 1
liste isimler[kutu adet] = kutu isim
liste notlar[kutu adet] = kutu not
yaz Eklendi!
git menu
son

eger kutu secim = 2 ise
yaz
yaz === Tüm Öğrenciler ===
kutu i = 1
goster:
eger kutu i <= kutu adet ise
yaz liste isimler[kutu i]: liste notlar[kutu i]
kutu i = kutu i + 1
git goster
son
git menu
son

eger kutu secim = 3 ise
yaz Hoşçakal!
son
