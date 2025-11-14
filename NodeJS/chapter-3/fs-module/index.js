const fs = require("fs");

fs.readFile("password.txt", "utf-8", (err, data) => {
  if (err) console.log(err);
  console.log(data);
  console.log("Dosya okundu");
});

fs.writeFile("example.txt", "Hello, nodejs world", "utf-8", (err, data) => {
  if (err) console.log(err);
  console.log("dosyaya yazma işlemi tamamlandı.");
});

fs.writeFile(
  "example.json",
  '{"name":"merlins","age":25}',
  "utf-8",
  (err, data) => {
    if (err) console.log(err);
    console.log("json dosyasına yazma işlemi tamamlandı.");
  }
);

fs.appendFile("example.txt", "\n is magic", "utf-8", (err, data) => {
  if (err) console.log(err);
  console.log("dosyasına append  işlemi tamamlandı.");
});

fs.unlink("example.json", (err) => {
  if (err) console.log(err);
  console.log("dosyasına silindi.");
});

fs.mkdir('upload/img',{recursive:true},(err)=>{
    if (err) console.log(err);
    console.log("klasorler olusturuldu");
})

fs.rmdir('upload',{recursive:true},(err)=>{
    if (err) console.log(err);
    console.log("klasorler silindi");
})

