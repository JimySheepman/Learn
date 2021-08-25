/*
Map

new Map() – map yaratır.
map.set(key, value) – Anahtara değer atar.
map.get(key) – Anahtarın değerini döndürür. Eğer öyle bir anahtar yoksa undefined döndürür.
map.has(key) – Eğer öyle bir anahtar varsa true yoksa false döndürür.
map.delete(key) – Verilen anahtara ait değeri siler.
map.clear() – Mapin içini temizler.
map.size – anlık eleman sayısını döndürür.


Set

new Set(iterable) – set oluşturur, isteğe bağlı olarak değerler içeren diziden de oluşturulabilir.
set.add(value) – bir değer ekler, set’in kendisini döndürür
set.delete(value) – değeri siler. Eğer öyle bir değer varsa true yoksa false döndürür.
set.has(value) – Eğer öyle bir değer varsa true yoksa false döndürür.
set.clear() – set’in içindeki her şeyi siler.
set.size – eleman sayısını döndürür.

*/

let map = new Map();

map.set("1", "str1"); // String tipinde anahtar
map.set(1, "num1"); // Sayı tipinde anahtar
map.set(true, "bool1"); // boolean tipinde anahtar

// sıradan Objeleri hatırlıyorsunuzdur. Anahtar değerleri stringe dönüşürdü
// Map anahtar tipini de korur, tıpkı şu 2 farklı şekilde olduğu gibi:
alert(map.get(1)); // 'num1'
alert(map.get("1")); // 'str1'

alert(map.size); // 3

// id değeri ekledik
let john = { name: "John", id: 1 };

let ziyaretSayisi = {};

// şimdi id kullanarak veriyi tuttuk
ziyaretSayisi[john.id] = 123;

alert(ziyaretSayisi[john.id]); // 123

// [key, value] çiftlerinden oluşan dizi
let map1 = new Map([
  ["1", "str1"],
  [1, "num1"],
  [true, "bool1"],
]);

let map2 = new Map(
  Object.entries({
    name: "John",
    age: 30,
  })
);

let yemekMap = new Map([
  ["salatalik", 500],
  ["domates", 350],
  ["sogan", 50],
]);

// anahtarlar üzerinde yineleme (sebzeler)
for (let vegetable of yemekMap.keys()) {
  alert(vegetable); // salatalik, domates, sogan
}

// değerler üzerinde yineleme (miktarlar)
for (let amount of yemekMap.values()) {
  alert(amount); // 500, 350, 50
}

// [anahtar, değer] üzerinde yineleme
for (let entry of yemekMap) {
  // yemekMap.entries() ile aynı
  alert(entry); // salatalik,500 (vb.)
}


let set = new Set();

let john = { name: "John" };
let pete = { name: "Pete" };
let mary = { name: "Mary" };

// ziyaretler, bazıları birden çok kez gelmiş
set.add(john);
set.add(pete);
set.add(mary);
set.add(john);
set.add(mary);

// set sadece eşsiz değerler tutar
alert( set.size ); // 3

for(let user of set) {
  alert(user.name); // John (sonraki döngülerde Pete and Mary)
}

let set = new Set(["portakal", "elma", "muz"]);

for(let value of set) alert(value);

// forEach ile de aynı şekilde:
set.forEach((value, valueAgain, set) => {
  alert(value);
});

let john = { name: "John" };

// Obje erişebilir, john da onun referansı

// referansın üzerine yazalım (overwrite)
john = null;

// obje daha fazla erişebilir olmadığı için bellekten silinir.


let john = { name: "John" };

let map = new Map();
map.set(john, "...");

john = null; // referansın üzerine yazalım (overwrite)

// john map içinde tutuldu
// map.keys() kullanarak ona ulaşabiliriz

let weakMap = new WeakMap();

let obj = {};

weakMap.set(obj, "ok"); // düzgün çalışır (anahtar bir obje)

weakMap.set("test", "Whoops"); // Hata verir, çünkü "test" ilkel bir tipte

let john = { name: "John" };

let weakMap = new WeakMap();
weakMap.set(john, "...");

john = null; // referansın üzerine yazalım (overwrite)

// john bellekten silindi!