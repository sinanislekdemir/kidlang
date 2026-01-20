fi

// Matematiikkaohjelma suomeksi
tulosta === Laskuharjoitus ===
kysy Anna ensimmäinen numero: 
laatikko luku1 = vastaus
kysy Anna toinen numero: 
laatikko luku2 = vastaus

laatikko summa = laatikko luku1 + laatikko luku2
laatikko erotus = laatikko luku1 - laatikko luku2
laatikko tulo = laatikko luku1 * laatikko luku2

tulosta 
tulosta Tulokset:
tulosta laatikko luku1 + laatikko luku2 = laatikko summa
tulosta laatikko luku1 - laatikko luku2 = laatikko erotus
tulosta laatikko luku1 * laatikko luku2 = laatikko tulo

if laatikko luku2 != 0 then
   goto jako
end
tulosta Ei voi jakaa nollalla!
goto loppu

jako:
laatikko jakolasku = laatikko luku1 / laatikko luku2
tulosta laatikko luku1 / laatikko luku2 = laatikko jakolasku

loppu:
tulosta 
tulosta Kiitos käytöstä!
