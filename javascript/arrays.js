let arr = new Array();
let arr = [];
let meyveler = ["Elma", "Portakal", "Erik"];

let meyveler = ["Elma", "Portakal", "Erik"];

alert(meyveler[0]); // Elma
alert(meyveler[1]); // Portakal
alert(meyveler[2]); // Erik

meyveler[2] = "Armut"; // Şimdi ["Elma", "Portakal", "Armut"]
meyveler[3] = "Limon"; // Şimdi ["Elma", "Portakal", "Armut", "Limon"]
meyveler[3] = "Limon"; // Şimdi ["Elma", "Portakal", "Armut", "Limon"]
let meyveler = ["Elma", "Portakal", "Erik"];

alert(meyveler.length); // 3

let meyveler = ["Elma", "Portakal", "Erik"];

alert(meyveler); // Elma,Portakal,Erik

// Karmaşık tipler
let arr = [
  "Elma",
  { isim: "Ahmet" },
  true,
  function () {
    alert("merhaba");
  },
];

// Birinci indeksteki değeri al ve "isim" özelliğini görüntüle
alert(arr[1].name); // John

// 3. indeksteki fonksiyonu al ve çalıştır.
arr[3](); // merhaba

let meyveler = ["Elma", "Portakal", "Erik"];

let meyveler = ["Elma", "Portakal", "Armut"];

alert(meyveler.pop()); // Armutu sil ve bunu ekranda bildir.

alert(meyveler); // Elma, Portakal

let meyveler = ["Elma", "Portakal"];

meyveler.push("Armut");

alert(meyveler); // Elma, Portakal, Armut

let meyveler = ["Elma", "Portakal", "Armut"];

alert(meyveler.shift()); // Elmayı sil ve bunu ekranda bildir.

alert(meyveler); // Portakal, Armut

let meyveler = ["Portakal", "Armut"];

meyveler.unshift("Elma");

alert(meyveler); // Elma, Portakal, Armut

let meyveler = ["Elma"];

meyveler.push("Portakal", "Armut");
meyveler.unshift("Ananas", "Limon");

// ["Ananas", "Limon", "Elma", "Portakal", "Armut"]
alert(meyveler);

let meyveler = ["Muz"];

let arr = meyveler; // iki değişken aynı diziye referans verir. ( referans ile kopyalama )

alert(arr === fruits); // true

arr.push("Armut"); // diziyi referans ile düzenleme

alert(meyveler); // Muz, Armut - 2 eleman

let meyveler = []; // Dizi yap

meyveler[99999] = 5; // boyutundan çokça büyük bir özelliğe veri ata.

meyveler.yas = 25; // doğrudan özelliğe isim vererek atama yap.

meyveler.shift(); // Başlangıçtan bir eleman al

fruits.pop(); // Sondan bir eleman al

let dizi = ["Elma", "Portakal", "Armut"];

for (let i = 0; i < arr.length; i++) {
  alert(arr[i]);
}

let meyveler = ["Elma", "Portakal", "Erik"];

// Dizi elemanları üzerinden döner.
for (let meyve of meyveler) {
  alert(meyve);
}

let arr = ["Elma", "Portakal", "Erik"];

for (let key in arr) {
  alert(arr[key]); // Elma, Portakal, Erik
}

let meyveler = [];
meyveler[123] = "Elma";

alert(meyveler.length); // 124

let arr = [1, 2, 3, 4, 5];

arr.length = 2; // 2 elemana düşür
alert(arr); // [1, 2]

arr.length = 5; // uzunluğu geri al
alert(arr[3]); // undefined: değer dönmez

let arr = new Array("Elma", "Armut", "vs");

let arr = new Array(2); //  [2] diye bir dizi mi oluşturacak ?

alert(arr[0]); // undefined! böyle bir eleman yok

alert(arr.length); // length 2

let matrix = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9],
];

alert(matrix[1][1]); // merkez eleman

let arr = [1, 2, 3];

alert(arr); // 1,2,3
alert(String(arr) === "1,2,3"); // true

alert([] + 1); // "1"
alert([1] + 1); // "11"
alert([1, 2] + 1); // "1,21"

alert("" + 1); // "1"
alert("1" + 1); // "11"
alert("1,2" + 1); // "1,21"

// köşeli parantez ile (genel kullanılan)
let arr = [item1, item2,];

// new Array (Çok nadir kullanım)
let arr = new Array(item1, item2,);

/*
Dizi özel bir çeşit objedir, verilerin sıralı bir şekilde saklanması için uygun bir tiptir.

Tanım: new Array(number) verilen boyutlarda yeni bir dizi yaratır, fakat eleman olmadan.

length özelliği dizinin boyu ve daha net olmak gerekirse son index sayısı + 1 şeklindedir. Dizi metodları tarafından otomatik olarak ayarlanır.

Eğer length'i elle küçültürseniz dizi de kısalacaktır, tabi bu veri kayıplarına neden olabilir.

Dizi üzerinde aşağıdaki işlemler yapılabilir:

push(...items) items'ı sona ekler.
pop() sondan bir eleman siler ve dönderir.
shift() başlangıçtan eleman siler ve bunu dönderir.
unshift(...items) başlangıca items ekler.
Dizinin elemanlarını for döngüsü ile dönme:

for(let i=0; i<arr.length; i++) – hızlı çalışır ve eski tarayıcılara uyarlıdır.
for(let item of arr) – sadece elemanların yazımı için modern yazım sağlar.
for(let i in arr) – kullanılamaz.
*/