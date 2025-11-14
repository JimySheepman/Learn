// ! fonksiyon içerisinde fanksiyon çağırma işlemidir.
// ! aynı anda başka bir işem yapılmasını istmek callback demek
const func1 = () => {
  console.log("Func 1 tamamlandı");
  func2();
};

const func2 = () => {
  console.log("Func 2 tamamlandı");
  func3();
};

const func3 = () => {
  console.log("Func 3 tamamlandı");
};

func1();

let x = 5;
console.log("1. gelen veri: ", x);

setTimeout(() => {
  x = x + 5;
}, 1000); // veri kaybı yaşanıyor.

console.log("2. gelen veri: ", x);

x = x + 5;
console.log("3. gelen veri: ", x);

const books = [
  { name: "kitap 1", author: "yazar 1" },
  { name: "kitap 2", author: "yazar 2" },
  { name: "kitap 3", author: "yazar 3" },
];

const listbook = () => {
  books.map((book) => {
    console.log(book.name);
  });
};

const addbook = (newBook, callback) => {
  books.push(newBook);
  callback();// diğer fonsiyonu cağrıyoruz.(callback işemidir bu)
};

addbook({ name: "kitap 4", author: "yazar 4" }, listbook);
