TR
// Orta Dünya Tüccarı - Bir ticaret macerası
// Ucuza al, pahalıya sat ve 1000 altın kazan!

yaz ====================================
yaz   ORTA DUNYA TUCCARI
yaz ====================================
yaz
yaz Sen Orta Dünya'da bir tüccarsın
yaz Hedef: 30 günde 1000 altın kazan!
yaz

kutu gun = 1
kutu altin = 100
kutu canta = 50
kutu dolu = 0
kutu iksir = 0
kutu asa = 0
kutu zirh = 0

basla:
yaz
yaz === GUN kutu gun of 30 ===
yaz Altın: kutu altin
yaz Çanta: kutu dolu of kutu canta alan
yaz Envanter iksir: kutu iksir | Asa: kutu asa | Zırh: kutu zirh
yaz

kutu i_fiyat = rastgele % 30 + 20
kutu a_fiyat = rastgele % 50 + 40  
kutu z_fiyat = rastgele % 80 + 60

yaz === PAZAR FIYATLARI ===
yaz İksir: kutu i_fiyat altın
yaz Asa: kutu a_fiyat altın
yaz Zırh: kutu z_fiyat altın (2 alan kaplar)
yaz
yaz === EYLEMLER ===
yaz 1. İksir Al
yaz 2. Asa Al
yaz 3. Zırh Al
yaz 4. İksir Sat
yaz 5. Asa Sat
yaz 6. Zırh Sat
yaz 7. Sonraki şehre seyahat et
sor Seç:

kutu secim = cevap

eger kutu secim = 1 ise git iksir_al
eger kutu secim = 2 ise git asa_al
eger kutu secim = 3 ise git zirh_al
eger kutu secim = 4 ise git iksir_sat
eger kutu secim = 5 ise git asa_sat
eger kutu secim = 6 ise git zirh_sat
eger kutu secim = 7 ise git seyahat
yaz Geçersiz seçim!
bekle 1000
git basla

iksir_al:
sor Kaç iksir?
kutu miktar = cevap
kutu maliyet = kutu i_fiyat * kutu miktar
kutu gerek = kutu dolu + kutu miktar
eger kutu maliyet > kutu altin ise
   yaz Yeterli altın yok!
   bekle 1000
   git basla
son
eger kutu gerek > kutu canta ise
   yaz Yeterli çanta alanı yok!
   bekle 1000
   git basla
son
kutu altin = kutu altin - kutu maliyet
kutu iksir = kutu iksir + kutu miktar
kutu dolu = kutu dolu + kutu miktar
yaz kutu miktar iksir alındı! kutu maliyet altın ödendi!
bekle 1000
git basla

asa_al:
sor Kaç asa?
kutu miktar = cevap
kutu maliyet = kutu a_fiyat * kutu miktar
kutu gerek = kutu dolu + kutu miktar
eger kutu maliyet > kutu altin ise
   yaz Yeterli altın yok!
   bekle 1000
   git basla
son
eger kutu gerek > kutu canta ise
   yaz Yeterli çanta alanı yok!
   bekle 1000
   git basla
son
kutu altin = kutu altin - kutu maliyet
kutu asa = kutu asa + kutu miktar
kutu dolu = kutu dolu + kutu miktar
yaz kutu miktar asa alındı! kutu maliyet altın ödendi!
bekle 1000
git basla

zirh_al:
sor Kaç zırh?
kutu miktar = cevap
kutu maliyet = kutu z_fiyat * kutu miktar
kutu gerek = kutu dolu + kutu miktar * 2
eger kutu maliyet > kutu altin ise
   yaz Yeterli altın yok!
   bekle 1000
   git basla
son
eger kutu gerek > kutu canta ise
   yaz Yeterli çanta alanı yok!
   bekle 1000
   git basla
son
kutu altin = kutu altin - kutu maliyet
kutu zirh = kutu zirh + kutu miktar
kutu dolu = kutu dolu + kutu miktar * 2
yaz kutu miktar zırh alındı! kutu maliyet altın ödendi!
bekle 1000
git basla

iksir_sat:
sor Kaç iksir?
kutu miktar = cevap
eger kutu miktar > kutu iksir ise
   yaz O kadar yok!
   bekle 1000
   git basla
son
kutu kazanc = kutu i_fiyat * kutu miktar
kutu altin = kutu altin + kutu kazanc
kutu iksir = kutu iksir - kutu miktar
kutu dolu = kutu dolu - kutu miktar
yaz kutu miktar iksir satıldı! kutu kazanc altın kazanıldı!
bekle 1000
git basla

asa_sat:
sor Kaç asa?
kutu miktar = cevap
eger kutu miktar > kutu asa ise
   yaz O kadar yok!
   bekle 1000
   git basla
son
kutu kazanc = kutu a_fiyat * kutu miktar
kutu altin = kutu altin + kutu kazanc
kutu asa = kutu asa - kutu miktar
kutu dolu = kutu dolu - kutu miktar
yaz kutu miktar asa satıldı! kutu kazanc altın kazanıldı!
bekle 1000
git basla

zirh_sat:
sor Kaç zırh?
kutu miktar = cevap
eger kutu miktar > kutu zirh ise
   yaz O kadar yok!
   bekle 1000
   git basla
son
kutu kazanc = kutu z_fiyat * kutu miktar
kutu altin = kutu altin + kutu kazanc
kutu zirh = kutu zirh - kutu miktar
kutu dolu = kutu dolu - kutu miktar * 2
yaz kutu miktar zırh satıldı! kutu kazanc altın kazanıldı!
bekle 1000
git basla

seyahat:
yaz
yaz Sonraki şehre seyahat ediliyor...
bekle 1000
kutu gun = kutu gun + 1

kutu olay = rastgele % 8

eger kutu olay = 0 ise
   yaz Dostane bir büyücü sana 50 altın verdi! ✨
   kutu altin = kutu altin + 50
   bekle 1500
son

eger kutu olay = 1 ise
   yaz Haydutlar saldırdı! 30 altın kaybedildi! ⚔️
   eger kutu altin > 30 ise
      kutu altin = kutu altin - 30
   son
   eger kutu altin <= 30 ise
      kutu altin = 0
   son
   bekle 1500
son

eger kutu olay = 2 ise
   yaz Yolda bir iksir buldun! 🧪
   eger kutu dolu < kutu canta ise
      kutu iksir = kutu iksir + 1
      kutu dolu = kutu dolu + 1
   son
   bekle 1500
son

eger kutu altin >= 1000 ise git kazan
eger kutu gun > 30 ise git kaybet

git basla

kazan:
yaz
yaz ====================================
yaz   TEBRİKLER!
yaz ====================================
yaz kutu gun günde 1000 altın kazandın!
yaz Usta Tüccarsın! 🏆
yaz
git son

kaybet:
yaz
yaz ====================================
yaz   SÜRE BİTTİ!
yaz ====================================
yaz 30 gün geçti...
yaz Son altın: kutu altin
yaz
eger kutu altin >= 500 ise
   yaz Bir tüccar için fena değil!
son
eger kutu altin < 500 ise
   yaz Ticaretini geliştirmeye devam et!
son
yaz
git son

son:
yaz Orta Dünya Tüccarı oynadığın için teşekkürler!
yaz Hoşçakal! 👋
