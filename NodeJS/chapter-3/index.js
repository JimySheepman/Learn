const http = require("http");

const server = http.createServer((req, res) => {
  const url = req.url;
  if (url === "/") {
    res.writeHead(200, { "Content-Type": "text/html" });
    res.write("<h1>index sayfası</h1>");
  } else if (url === "/about") {
    res.writeHead(200, { "Content-Type": "text/html" });
    res.write("<h1></h1>about sayfası</h1>");
  } else if ("/contact") {
    res.writeHead(200, { "Content-Type": "text/html" });
    res.write("<h1>contact sayfası</h1>");
  } else {
    res.writeHead(404, { "Content-Type": "text/html" });
    res.write("<h1>404 sayfa bulunamadı</h1>");
  }

  res.end();
});

const port = 3000;

server.listen(port, () => {
  console.log(`Server port:${port}başlatıldı..`);
});
