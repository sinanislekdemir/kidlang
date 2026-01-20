TR
liste urunler
kutu adet = 0
yaz === Alışveriş Listesi ===
yaz
menu:
yaz Ne yapmak istersin?
yaz 1 Ürün ekle
yaz 2 Listeyi göster
yaz 3 Çık
sor Seç:
kutu secim = cevap

eger kutu secim = 1 ise
sor Ürün adını gir:
kutu isim = cevap
kutu adet = kutu adet + 1
liste urunler[kutu adet] = kutu isim
yaz Eklendi!
git menu
son

eger kutu secim = 2 ise
yaz
yaz Alışveriş listesi:
kutu i = 1
goster:
eger kutu i <= kutu adet ise
yaz kutu i . liste urunler[kutu i]
kutu i = kutu i + 1
git goster
son
yaz
git menu
son

eger kutu secim = 3 ise
yaz Hoşçakal!
son
