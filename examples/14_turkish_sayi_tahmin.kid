tr

// Türkçe Sayı Tahmin Oyunu
// Gizli sayıyı bulmaya çalış!

yaz ================================
yaz    SAYI TAHMİN OYUNU
yaz ================================
yaz

kutu gizli = rastgele % 50 + 1
kutu deneme = 0

yaz 1 ile 50 arasında bir sayı tuttum!
yaz Bulabilir misin?
yaz

oyun:
kutu deneme = kutu deneme + 1

yaz Deneme kutu deneme
sor Tahminin nedir?

eger cevap = kutu gizli ise
   git kazandin
son

eger cevap > kutu gizli ise
   yaz Daha küçük bir sayı söyle!
son

eger cevap < kutu gizli ise
   yaz Daha büyük bir sayı söyle!
son

yaz
git oyun

kazandin:
yaz
yaz ================================
yaz    TEBRİKLER!
yaz ================================
yaz Doğru bildin! Sayı kutu gizli idi.
yaz kutu deneme denemede buldun!
yaz
yaz Teşekkürler! Hoşça kal!
