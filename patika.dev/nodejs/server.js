const http = require("http");

const server = http.createServer((req, res) => {
  const url = req.url;
  if (url === "/") {
    res.writeHead(200, { "Content-Type": "text/html" });
    res.write("<h2>Index sayfasına hoşgeldinz</h2>");
  } else if (url === "/about") {
    res.writeHead(200, { "Content-Type": "text/html" });
    res.write("<h2>About sayfasına hoşgeldinz</h2>");
  } else if (url === "/contact") {
    res.writeHead(200, { "Content-Type": "text/html" });
    res.write("<h2>Contact sayfasına hoşgeldinz</h2>");
  } else {
    res.writeHead(404, { "Content-Type": "text/html" });
    res.write("<h2>status:404 -> Not Found</h2>");
  }
  res.end();
});

const port = 5000;
server.listen(port, () => {
  console.log(`Server is start. Port:${port}`);
});
