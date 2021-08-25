/*
alert('Merhaba'); alert('Dünya');

alert('Merhaba')
alert('Dünya')

alert("Bu mesajdan sonra hata verecek")
[1, 2].forEach(alert)

function f() {
  // fonksiyon tanımından sonra noktalı virgül yazılmaz
}

for(;;) {
  // döngüden sonra noktalı virgül yazılmaz
}

'use strict';

...

let
const (sabit, değiştirilemez)
var (eski tip)

let x = 5;
x = "Ahmet";

7 çeşit veri tipi bulunmaktadır:

number ( sayı ) floating-point ve doğal sayılar için kullanılır.
string (karakter dizileri),
boolean Mantıksal değerler için dogru/yanlis,
null – sadece null değerini tutar ve bu da “boş” veya “varolmayan” anlamına gelir,
undefined – sadece undefined değerine sahiptir. Bu da “değer atanmamış” demektir,
object ve symbol – karmaşık veri yapıları için ve tek tanıtıcı(unique identifier) için kullanılabilir. Bu konular henüz anlatılmadı.


typeof null == "object" // hata verir
typeof function(){} == "function" // fonksiyonlara özel davranılır.

prompt(soru[, varsayılan])
soru sor ve kullanıcının girdiği değeri dönder. Eğer kullanıcı “iptal” tuşuna bakarsa null dönder.
confirm(soru)
soru sor ve “Tamam” mı yoksa “İptal” mi diye seçenekler sun. Sonuçta seçilene göre true/false dönder.
alert(mesaj)
Mesajın çıktısını ekrana uyarı olarak ver.

let ziyaretci = prompt("Adınız?", "İbrahim");
let cayIstermi = confirm("Biraz çay ister misiniz?");
alert( "Ziyaretçi: " + ziyaretci ); // İbrahim
alert( "Çay isteriyor mu?: " + cayIstermi ); // true

Normal işlemler: * + - /, mod alma % ve ** üs alma için bu operatörler kullanılır.
alert( '1' + 2 ); // '12', karakter dizisi
alert( 1 + '2' ); // '12', karakter dizisi


Değer atama
Basit bir şekilde a = b şeklinde kullanılabilir. Veya birleşik olarak a *= 2 gibi de kullanıma sahiptir.

Bit seviyesi işlemler
Bit seviye operatörleri şu şekilde kullanılabilir: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators

Üçlü operatör
Üç tane paremetreden oluşur: koşul ? sonucA : sonucB. Eğer koşul doğru ise sonucA döndürür, yanlış ise sonucB

Mantıksal operatörler:
Mantıksal VE &&, VEYA || operatörleri ile bu işlemler yapılabilir.

Karşılaştırma
Eşitlik kontrolü ==, farklı tipteki verileri sayıya çevirip kontrol eder. null ve undefined hariç, bu ikisi de birbirine eşittir.

alert( 0 == false ); // true
alert( 0 == '' ); // true

sıkı eşitlik operatörü === bu çeviriyi yapmamaktadır: farklı tipler her zaman farklı değerler ifade eder, öyleyse:
null ve undefined değerleri özeldir: == şeklinde birbirlerine eşittirler. Fakat başka hiç bir değere eşit değildirler. Büyüktür/Küçüktür karşılaştırmasında karakter dizileri karakter karakter karşılaştırılır. Diğer tipler sayıya çevrilir.

// 1
while (koşul) {
  ...
}

// 2
do {
  ...
} while (koşul);

// 3
for(let i = 0; i < 10; i++) {
  ...
}


let age = prompt('Kaç yaşındasın?', 18);

switch (age) {
  case 18:
    alert("Çalışmaz"); // `prompt` ile tutulan değer sayı değil karakterdir!!!

  case "18":
    alert("Çalışır!");
    break;

  default:
    alert("Değer yukarıda bulunan koşullara uymamakta");
}


function toplam(a, b) {
  let sonuc = a + b;

  return sonuc;
}


let toplam = function(a, b) {
  let sonuc = a + b;

  return sonuc;
}


// ifada sağ tarafta
let toplam = (a, b) => a + b;

// Çoklu satır için {..} kullanılmalı ve `return` ile değerin dönderilmesi gerekmektedir:
let toplam = (a, b) => {
  // ...
  return a + b;
}

// argümansız
let selamVer = () => alert("Merhaba");

// tek argümanlı
let ikiyeKatla = n => n * 2;


çalışmasını görmek için lütfen geliştirici konsolunu açınız.
for (let i = 0; i < 5; i++) {
  console.log("deger", i);
}

describe("us", function() {

  it("n. kuvvetini alir", function() {
    assert.equal(us(2, 3), 8);
  });

});
Bu özelliğin 3 ana bölümü vardır:

describe("baslik", function() { ... })
Fonksiyonun neyi tanımladığı yazılır. Bizim durumumuzda bu üs’tür.

it("baslik", function() { ... })
it bölümünde ise daha okunaklı bir şekilde, hangi koşulda ne yaptığı açıklanır. ikinci argüman ise bunu test eder.

assert.equal(value1, value2)
it bloğunun içindeki ko eğer doğru ise hatasız döner.

assert* fonksiyonu us'ün beklendiği gibi çalışıp çalışmadığını kontrol eder. Burada assert.equal'ı kullanılmaktadır. Argümanları karşılaştırarak eşitlik olmadığı durumda hata verir. Burada us(2,3), 8 e eşit mi diye bakılır

İleriki dönemlerde farklı karşılaştırmaları göreceksiniz.

*/