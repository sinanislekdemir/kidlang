TR
// Yerlesik yardimcilar yazi, dosya ve liste ile calisabilir.

kutu evcil = gokkusagi robot
yaz icindevar(kutu evcil, kus)
yaz degistir(kutu evcil, robot, ucurtma)
yaz parcaal(kutu evcil, 1, 9)

dosya defter = notlar
ac defter helpernotes.txt
yaz dosyavarmi(defter)
yaz dosyaboyu(defter)
kapat defter

liste oyuncaklar
kutu ilk = listeyekoy(oyuncaklar, 1, robot)
kutu ikinci = listeyekoy(oyuncaklar, 2, ucurtma)
kutu anahtarlar = listeanahtarlari(oyuncaklar)
yaz birlestir(kutu anahtarlar, :)
yaz listeuzunluk(oyuncaklar)
