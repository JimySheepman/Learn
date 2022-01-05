// server.js
// where your node app starts
function isValidDate(d) {
  return d instanceof Date && !isNaN(d);
}
// init project
var express = require("express");
var app = express();

// enable CORS (https://en.wikipedia.org/wiki/Cross-origin_resource_sharing)
// so that your API is remotely testable by FCC
var cors = require("cors");
app.use(cors({ optionSuccessStatus: 200 })); // some legacy browsers choke on 204

// http://expressjs.com/en/starter/static-files.html
app.use(express.static("public"));

// http://expressjs.com/en/starter/basic-routing.html
app.get("/", function (req, res) {
  res.sendFile(__dirname + "/views/index.html");
});

// your first API endpoint...
app.get("/api/hello", function (req, res) {
  res.json({ greeting: "hello API" });
});

app.get("/api/", function (req, res) {
  let dateRes = new Date();
  res.json({ unix: dateRes.getTime(), utc: dateRes.toUTCString() });
});

app.get("/api/:date", function (req, res) {
  let dateStr = req.params.date;
  console.log(dateStr);
  if (dateStr.indexOf("-") === -1) {
    dateStr = new Number(dateStr);
  }
  let dateRes = new Date(dateStr);
  if (isValidDate(dateRes)) {
    res.json({ unix: dateRes.getTime(), utc: dateRes.toUTCString() });
  } else {
    res.json({ unix: null, utc: "Invalid Date" });
  }
});

// listen for requests :)
var listener = app.listen(3000, function () {
  console.log("Your app is listening on port " + listener.address().port);
});
