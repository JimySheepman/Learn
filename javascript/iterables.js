/*
Özet
Objeler for..of ile kullanılırsa sıralı erişim objesi adını alır.

Teknik olarak, sıralı erişim objelerinin Symbol.iterator metodunu uygulamış olması gerekir.
obj[Symbol.iterator]'ün sonucunda bu objeye sıralı erişim objesi denir ve for..of içerisinde tekrarlanabilir.
Bir sıralı erişim objesi next() metoduna kesinlikle sahip olmalıdır. Bu metod { done: Boolean, value:any} döndürmelidir. Burada done:true olur ise bu döngü bitti anlamına gelir. Diğer türlü value bir sonraki değerdir.
Symbol.iterator metodu for..of tarafından otomatik olarak çağrılmaktadır. Elbette doğrudan da çağırılabilir.
Var olan sıralı erişilebilir objeler, yani karakterler ve diziler de Symbol.iterator metodunu yapmışlardır.
Karakter döngüsü vekil ikilileri anlayabilir.
İndekslenmiş özelliklere ve length özelliğine sahip objelere dizi-benzeri denir. Böyle objeler başka özellik ve metodlara da sahip olabilir. Fakat dizilerin sahip olduğu metodlardan yoksundurlar.

Eğer şartnameye bakılacak olursa – Varolan çoğu metodun iterables veya dizi-benzeri ile çalışabileceği vurgulanmıştır. Gerçek diziler daha soyut kalmaktadır bundan dolayı pek bahsedilmez.

Array.from(obj[, mapFn, thisArg]) metodu iterable veya dizi-benzeri'inden gerçek Array üretirler, sonrasında bunu herhangi bir dizi metoduyla kullanılabilir.
 mapFn ve thisArg gibi isteğe bağlı metodlar dizinin her bir elemanın istenilen fonksiyona uygular.
*/
let aralik = {
  baslangic: 1,
  bitis: 5,

  [Symbol.iterator]() {
    this.current = this.baslangic;
    return this;
  },

  next() {
    if (this.current <= this.bitis) {
      return { done: false, value: this.current++ };
    } else {
      return { done: true };
    }
  },
};

for (let num of aralik) {
  console.log(num); // 1, then 2, 3, 4, 5
}

for (let char of "test") {
  console.log(char); // t, sonra e, sonra s, sonra t
}

let str = "𝒳😂";
for (let char of str) {
  console.log(char); // 𝒳, sonra 😂
}

let str1 = "Hello";

// for (let char of str) alert(char);
// ile aynı şekilde çalışır

let iterator = str1[Symbol.iterator]();

while (true) {
  let result = iterator.next();
  if (result.done) break;
  console.log(result.value); //karakterlerin bir bir çıktısını verir.
}

let diziBenzeri = {
  0: "Merhaba",
  1: "Dünya",
  length: 2,
};

let arr = Array.from(diziBenzeri); // (*)
console.log(arr.pop()); // Dünya (metod çalışmakta)

// aralik'in yukarıdan alındığı varsayılırsa

// her sayının karesinin alınması.
let arr2 = Array.from(aralik, (num) => num * num);

console.log(arr2); // 1,4,9,16,25

function slice(str2, start, end) {
  return Array.from(str2).slice(start, end).join("");
}

let str2 = "𝒳😂𩷶";

console.log(slice(str2, 1, 3)); // 😂𩷶

// Varolan metodlar vekil çiftleri desteklemez.
console.log(str2.slice(1, 3)); // çöp
