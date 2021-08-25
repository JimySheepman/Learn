/*
JSON.stringify objeyi JSON’a çevirir.
JSON.parse JSON’dan objeye çevirmeye yarar.
Örneğin, aşağıda JSON.stringify 
*/

let kullanici = {
    adi: "Ahmet",
    yasi: 30,
  
    toString() {
      return `{adi: "${this.adi}", yasi: ${this.yasi}}`;
    }
  };
  
  alert(kullanici); // {adi: "Ahmet", yasi: 30}


  let ogrenci = {
    adi: 'Ahmet',
    yasi: 30,
    adminMi: false,
    dersler: ['html', 'css', 'js'],
    esi: null
  };
  
  let json = JSON.stringify(ogrenci);
  
  alert(typeof json); // string dönecektir.!
  
  alert(json);
  /* JSON'a çevirilmiş obje:
  {
    "adi": 'Ahmet',
    "yasi": 30,
    "adminMi": false,
    "dersler": ['html', 'css', 'js'],
    "esi": null
  }
  */

  let json = JSON.stringify(deger[ degistirici, bosluk])

  let oda = {
    sayi: 23
  };
  
  let tanisma = {
    baslik: "Konferans",
    katilimcilar: [{adi: "Ahmet"}, {adi: "Mehmet"}],
    yer: oda // tanışma odayı referans gösteriyor.
  };
  
  oda.dolduruldu = tanisma; // oda tanışmayı referans gösteriyor.
  
  alert( JSON.stringify(tanisma, ['baslik', 'katilimcilar']) );
  // {"baslik":"Konferans","katilimcilar":[{},{}]}


  let oda = {
    sayi: 23
  };
  
  let tanisma = {
    baslik: "Konferans",
    katilimcilar: [{adi: "Ahmet"}, {adi: "Mehmet"}],
    yer: oda // tanışma odayı referans gösteriyor.
  };
  
  oda.dolduruldu = tanisma; // oda tanışmayı referans gösteriyor.
  
  alert( JSON.stringify(tanisma, ['baslik', 'katilimcilar', 'yer', 'adi', 'sayi']) );
  
  /*
  {
    "baslik":"Konferans",
    "katilimcilar":[{"adi":"Ahmet"},{"adi":"Mehmet"}],
    "yer":{"sayi":23}
  }
  */


  let oda = {
    sayi: 23
  };
  
  let tanisma = {
    baslik: "Konferans",
    katilimcilar: [{adi: "Ahmet"}, {adi: "Mehmet"}],
    yer: oda // tanışma odayı referans gösteriyor.
  };
  
  oda.dolduruldu = tanisma; // oda tanışmayı referans gösteriyor
  
  alert( JSON.stringify(tanisma, function degistirici(anahtar, deger) {
    alert(`${anahtar}: ${deger}`); // degistiriciye gelen
    return (anahtar == 'dolduruldu') ? undefined : deger;
  }));
  
  /* degistiriciye gelen anahtar:deger çifti:
  :             [object Object]
  baslik:        Conference
  katilimci:    [object Object],[object Object]
  0:            [object Object]
  adi:         Ahmet
  1:            [object Object]
  adi:         Mehmet
  yer:        [object Object]
  sayi:       23
  */


  let kullanici = {
    adi: "Ahmet",
    yas: 25,
    roller: {
      admin: false,
      editor: true
    }
  };
  
  alert(JSON.stringify(kullanici, null, 2));
  /* iki boşluk:
  {
    "adi": "Ahmet",
    "yasi": 25,
    "roller": {
      "admin": false,
      "editor": true
    }
  }
  */
  
  /*  JSON.stringify(user, null, 4) için ise çıktı aşağıdaki gibi olur:
  {
      "adi": "Ahmet",
      "yasi": 25,
      "roller": {
          "admin": false,
          "editor": true
      }
  }
  */


  let oda = {
    sayi: 23
  };
  
  let toplanti = {
    baslik: "Konferans",
    tarih: new Date(Date.UTC(2017, 0, 1)),
    oda
  };
  
  alert( JSON.stringify(toplanti) );
  /*
    {
      "baslik":"Konferans",
      "tarih":"2017-01-01T00:00:00.000Z",  // (1)
      "oda": {"sayi":23}               // (2)
    }
  */

    let deger = JSON.parse(str[ alıcı]);

    // metne çevrilmiş dizi
let sayilar = "[0, 1, 2, 3]";

sayilar = JSON.parse(sayilar);

alert( sayilar[1] ); // 1

let str = '{"baslik":"Konferans","tarih":"2017-11-30T12:00:00.000Z"}';

let tanisma = JSON.parse(str, function(anahtar, deger) {
  if (anahtar == 'tarih') return new Date(deger);
  return deger;
});

alert( tanisma.tarih.getDate() ); // Şimdi çalışıyor!


let program = `{
    "tanismalar": [
      {"baslik":"Konferans","tarih":"2017-11-30T12:00:00.000Z"},
      {"baslik":"DogumGunu","tarih":"2017-04-18T12:00:00.000Z"}
    ]
  }`;
  
  program = JSON.parse(program, function(anahtar, deger) {
    if (anahtar == 'tarih') return new Date(deger);
    return deger;
  });
  
  alert( program.tanismalar[1].tarih.getDate() ); // çalışır!