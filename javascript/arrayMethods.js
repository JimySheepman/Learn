/**
 Özet
Dizi metodlarının kısa açıklamaları:

Eleman ekleme/silme metodları:

push(...items) – elemanları sona ekler,
pop() – en sondaki elemanı alır,
shift() – başlangıçtan eleman alır,
unshift(...items) – başlangıça eleman ekler
splice(pos, deleteCount, ...items) – pos indeksinde deleteCount sayısı kadar elemanı siler ve bunları items'a ekler.
slice(start, end) – start ile end pozisyonları arasındaki (end dahil değil) elemanları yeni bir diziye kopyalar.
concat(...items) – yeni bir dizi döndürür: var olan dizideki tüm elemanları kopyalar ve items'ı ekler. Eğer items dizi ise bunun elemanları da alınır.
Elemanları aramaya yönelik metodlar:

indexOf/lastIndexOf(item, pos) – pos'tan başlayarak item'ı arar. Bulursa indeksini döndürür, bulamaz ise -1 döndürür.
includes(value) – eğer dizi value'ya sahipse true döndürür. Diğer türlü false döndürür.
find/filter(func) – Elemanları fonksiyonlar ile filtreler. Buna göre fonksiyonu true yapan ilk veya tamamını döner.
findIndex aynı find gibidir fakat bir değer yerine index döner.
Diziler üzerinde dönüşümler:

map(func) – her eleman için func çağrılır ve bunların sonuçlarından bir dizi üretilerek döndürülür.
sort(func) – diziyi olduğu yerde sıralar ve döndürür.
reverse() – diziyi terse çevirir ve döndürür.
split/join – karakterleri diziye çevirir veya dizileri karaktere çevirir.
reduce(func, initial) – dizide bulunan elemanlar sıra ile func fonksiyonu üzerinden hesaplanır ve son değer döndürülür.
Elemanlar üzerinden dönme:

forEach(func) – dizide bulunan her eleman için func çağrılır. Hiç birşey döndürmez.
Ek olarak:

Array.isArray(arr) arr'in dizi olup olmadığını kontrol eder.
Bu metodların içinden sadece sort, reverse ve splice doğrudan dizinin kendisi üzerinden işlem yapar. Diğerleri değer döndürür.

Yukarıdaki metodlar projelerin çoğundaki kullanılan dizi fonksiyonlarının %99’unu kapsar. Fakat bunun yanında farklı metodlar da bulunmaktadır:

arr.some(fn)/arr.every(fn) diziyi kontrol eder.

Dizinin her elemanı için fn çağırılır. Map'e çok benzer fakat herhangi biri/hepsi true ise true döndürür. Diğer türlü false döndürür.

arr.fill(value, start, end) – diziyi tekrar eden value değeri ile start ile index arasına doldurur.

arr.copyWithin(target, start, end) – start tan end'e kadar olan elemanları target'tan itibaren var olanların üzerine yazarak yapıştırır.

Tüm liste için kullanım talimatları sayfasına bakabilirsiniz.

Görünürde çok fazla metod varmış gibi ve ezberlemesi zormuş gibi görünse de aslında göründüğünden çok daha kolaydır.

Sadece tanımların bulunduğu sayfaya bakmanız yeterlid. Ardından bu bölümdeki örnekleri çözerek pratik yaparsanız metodlar ile ilgili yeteri kadar bilgi sahibi olmuş olursunuz.

Daha sonrasında metodlar ile ilgili birşey yapmak istediğinizde, nasıl yapıldığını bilmiyorsanız, buraya tekrar gelip doğru metodu bulabilirsiniz. 
Buradaki örnekler doğru bir şekilde yazmanıza yardımcı olacaktır. Sonrasında metodları hiç bir özel çaba harcamadan hatırlayacak duruma gelebilirsiniz.
 */

let arr = ["eve", "gitmek", "istiyorum"];
delete arr[1];
console.log(arr[1]);
console.log(arr.length);

let arr1 = ["ben", "javascript", "çalışıyorum"];
arr1.splice(1, 1); // index 1'den 1 elaman sil
console.log(arr1);

let arr2 = ["Ben", "şu", "an", "JavaScript", "çalışıyorum"];
arr2.splice(0, 4, "Ders"); // İlk 4 elamanı sil ve öncesine yeni eleman ekle.
console.log(arr2);

let arr3 = ["Ben", "şu", "an", "JavaScript", "çalışıyorum"];
let removed = arr3.splice(0, 2); // ilk iki elemanı sil.
console.log(removed);

let arr4 = [1, 2, 5];

// indeks -1 ( sondan birinci )
// 0 eleman sil,
// 3 vs 4 ekle
arr4.splice(-1, 0, 3, 4);

console.log(arr4); // 1,2,3,4,5

let str = "test";
let ary = ["t", "e", "s", "t"];

console.log(str.slice(1, 3)); // es
console.log(ary.slice(1, 3)); // e,s

console.log(str.slice(-2)); // st
console.log(ary.slice(-2)); // s,t

let arr5 = [1, 2];

// diziyi [3,4] ile birleştir
console.log(arr5.concat([3, 4])); // 1,2,3,4

// diziyi [3,4] ve [5,6] ile birleştir
console.log(arr5.concat([3, 4], [5, 6])); // 1,2,3,4,5,6

// diziyi [3,4] ile birleştir ve ardından 5, 6 ekle
console.log(arr5.concat([3, 4], 5, 6)); // 1,2,3,4,5,6

let arr6 = [1, 0, false];

console.log(arr6.indexOf(0)); // 1
console.log(arr6.indexOf(false)); // 2
console.log(arr6.indexOf(null)); // -1

console.log(arr6.includes(1)); // true

let kullanicilar = [
  { id: 1, isim: "Ahmet" },
  { id: 2, isim: "Muzaffer" },
  { id: 3, isim: "Emine" },
];

let kullanici = kullanicilar.find((eleman) => eleman.id == 1);

console.log(kullanici.isim); // Ahmet

let kullanicilar1 = [
  { id: 1, isim: "Ahmet" },
  { id: 2, isim: "Muzaffer" },
  { id: 3, isim: "Emine" },
];

// ilk iki kullaniciyi döndürür.
let baziKullanicilar = kullanicilar1.filter((eleman) => eleman.id < 3);

console.log(baziKullanicilar.length); // 2

let uzunluklar = ["Bilbo", "Gandalf", "Nazgul"].map((eleman) => eleman.length);
console.log(uzunluklar); // 5,7,6

let arr8 = [1, 2, 15];

// metod dizinin içeriğini sıralar ve döndürür.
arr8.sort();

console.log(arr8); // 1, 15, 2

let arr9 = [1, 2, 3, 4, 5];
arr9.reverse();

console.log(arr9); // 5,4,3,2,1

let isimler = "Bilbo, Gandalf, Nazgul";

let arr10 = isimler.split(", ");

for (let isim of arr10) {
  console.log(` ${name}'e mesaj.`); // Bilbo'e mesaj ve diğerleri.
}

let arr11 = [1, 2, 3, 4, 5];

let result = arr11.reduce((sum, current) => sum + current, 0);

console.log(result); // 15

["Bilbo", "Gandalf", "Nazgul"].forEach((item, index, array) => {
  console.log(`${item} ${array}'in ${index}. indeksinde`);
});

let kullanici2 = {
  yas: 18,
  dahaGenc(digerKullanici) {
    return digerKullanici.yas < this.yas;
  },
};

let kullanicilar3 = [{ yas: 12 }, { yas: 16 }, { yas: 32 }];

// kullanıcıdan daha genç kullanıcıları bulunuz
let dahaGencKullanicilar = kullanicilar3.filter(
  kullanici2.dahaGenc,
  kullanici2
);

console.log(dahaGencKullanicilar.length); // 2
