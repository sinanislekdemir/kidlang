TR
// Sayı Tahmin Oyunu
// Gizli sayıyı tahmin etmeye çalış!

yaz ================================
yaz    SAYI TAHMİN OYUNU
yaz ================================
yaz

kutu gizli = rastgele % 50 + 1
kutu denemeler = 0
kutu maks_deneme = 7

yaz 1 ile 50 arasında bir sayı düşünüyorum
yaz kutu maks_deneme denemede tahmin etmen gerek!
yaz

oyun:
kutu denemeler = kutu denemeler + 1

yaz Deneme kutu denemeler : kutu maks_deneme
sor Tahminin nedir:

eger cevap = kutu gizli ise
   git kazandın
son

eger cevap > kutu gizli ise
   yaz Çok yüksek! Daha küçük bir sayı dene.
son

eger cevap < kutu gizli ise
   yaz Çok düşük! Daha büyük bir sayı dene.
son

eger kutu denemeler >= kutu maks_deneme ise
   git kaybettin
son

yaz
git oyun

kazandın:
yaz
yaz ================================
yaz    TEBRİKLER!
yaz ================================
yaz Buldun! Sayı kutu gizli idi
yaz kutu denemeler denemede buldun!
git son

kaybettin:
yaz
yaz ================================
yaz    OYUN BİTTİ
yaz ================================
yaz Maalesef denemen bitti!
yaz Gizli sayı kutu gizli idi
git son

son:
yaz
yaz Oynadığın için teşekkürler!
