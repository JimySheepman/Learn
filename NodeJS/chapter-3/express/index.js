const express = require("express");

const app = express();

app.get("/", function (req, res) {
  res.status(200).send("Index");
});

app.get("/about", function (req, res) {
  res.status(200).send("About");
});

app.get("/contact", function (req, res) {
  res.status(200).send("contact");
});

app.get("*", (req, res) => {
  res.status(404).send("404 sayfa bulunamadı.");
});
app.listen(3000);
