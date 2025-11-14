function getData(data) {
  return new Promise((resolve, reject) => {
    console.log("Veriler alınmaya çalışılıyor...");
    if (data) {
      resolve("veriler alındı");
    } else {
      reject("Veriler alınmadı..");
    }
  });
}

function cleanData(receivedData) {
  return new Promise((resolve, reject) => {
    console.log("Veriler düzenleniyor..");
    if (receivedData) {
      resolve("Veriler düzenlendi");
    } else {
      reject("Veriler düzenlenemedi...");
    }
  });
}
/* Promise kullanımı -> asengrondur. parallen işlemlerde yapılır.
getData(true).then((result) => {
  console.log(result);
  return cleanData(true)
    .then((result) => {
      console.log(result);
    })
    .catch((error) => {
      console.log(error);
    });
});
*/
// async kullanımı -> sekrondur.
async function processData() {
  try {
    const receivedData = await getData(false);
    console.log(receivedData);
    const cleanedData = await cleanData(true);
    console.log(cleanedData);
  } catch (error) {
    console.log(error);
  }
}

processData();

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
    // reject("Hata oluştu");
  });
  return promise4;
};

async function showBooks() {
  try {
    await addbook({ name: "kitap 12", author: "yazar 4" });
    listbook();
  } catch (error) {
    console.log(error);
  }
}
showBooks();
