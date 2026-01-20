TR
// Basit Macera Oyunu
// Seçimler yap ve keşfet!

yaz ================================
yaz    HAZINE AVI MACERASI
yaz ================================
yaz

kutu saglik = 100
kutu altin = 0

basla:
yaz Gizemli bir ormanda uyanıyorsun.
yaz Sağlığın: kutu saglik
yaz Altınların: kutu altin
yaz
yaz Ne yapmak istersin?
yaz 1. Ormanı keşfet
yaz 2. Çantanı kontrol et
yaz 3. Dinlen
yaz 4. Oyundan çık
sor Seç:

eger cevap = 1 ise git kesfet
eger cevap = 2 ise git canta_kontrol
eger cevap = 3 ise git dinlen
eger cevap = 4 ise git cikis

yaz Geçersiz seçim!
git basla

kesfet:
yaz
yaz Ormanın derinliklerine doğru ilerliyorsun...
bekle 1000

kutu olay = rastgele % 3

eger kutu olay = 0 ise git altin_bul
eger kutu olay = 1 ise git canavar_bul
git hicbisey_bul

altin_bul:
kutu buldu = rastgele % 20 + 10
kutu altin = kutu altin + kutu buldu
yaz kutu buldu altın buldun! ✨
yaz Toplam altın: kutu altin
bekle 2000
git basla

canavar_bul:
yaz Bir canavar belirdi! 👹
kutu hasar = rastgele % 30 + 10
kutu saglik = kutu saglik - kutu hasar
yaz Sana kutu hasar hasar verdi!
yaz Kalan sağlık: kutu saglik

eger kutu saglik <= 0 ise
   git oyun_bitti
son

bekle 2000
git basla

hicbisey_bul:
yaz Burada ilginç bir şey bulamadın.
bekle 1000
git basla

canta_kontrol:
yaz
yaz === DURUMUN ===
yaz Sağlık: kutu saglik
yaz Altın: kutu altin
yaz
eger kutu altin >= 100 ise
   yaz Kazanmak için yeterli altının var!
   yaz Tebrikler! 🏆
   git cikis
son
yaz Kazanmak için 100 altına ihtiyacın var.
yaz Keşfetmeye devam et!
bekle 2000
git basla

dinlen:
yaz
yaz Bir süre dinleniyorsun...
bekle 1500
kutu iyilesme = 20
kutu saglik = kutu saglik + kutu iyilesme
eger kutu saglik > 100 ise
   kutu saglik = 100
son
yaz kutu iyilesme sağlık kazandın!
yaz Sağlık: kutu saglik
bekle 1500
git basla

oyun_bitti:
yaz
yaz ================================
yaz       OYUN BİTTİ
yaz ================================
yaz Yenildin!
yaz Son altın: kutu altin
yaz
git cikis

cikis:
yaz
yaz Oynadığın için teşekkürler!
yaz Hoşçakal! 👋
