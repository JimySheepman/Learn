const promise1 = new Promise((resolve, reject) => {
  resolve("Veriler alındı"); // olumlu dönüş yapıyor.
  // reject("Bağlantı hatası"); // olumsuz yani hata dönüyor
});

const promise2 = new Promise((resolve, reject) => {
  // resolve("Veriler alındı"); // olumlu dönüş yapıyor.
  reject("Bağlantı hatası"); // olumsuz yani hata dönüyor
});
const promise3 = new Promise((resolve, reject) => {
  resolve("Veriler alındı promise3"); // olumlu dönüş yapıyor.
  reject("Bağlantı hatası"); // olumsuz yani hata dönüyor
});

promise1.then((value) => {
  console.log(value);
});

promise2.catch((error) => {
  console.log(error);
});

promise3
  .then((value) => {
    console.log(value);
  })
  .catch((error) => {
    console.log(error);
  });

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

const addbook = (newBook) => {
  const promise4 = new Promise((resolve, reject) => {
    books.push(newBook);
    resolve(books);
    //reject("Hata oluştu")
  });
  return promise4;
};

addbook({ name: "kitap 4", author: "yazar 4" })
  .then(() => {
    console.log("Yeni Liste");
    listbook();
  })
  .catch((error) => {
    console.log(error);
  });
