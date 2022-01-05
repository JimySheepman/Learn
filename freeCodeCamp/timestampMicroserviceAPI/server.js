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

app.get("/api/:date", (req, res) => {
  let dateString = req.params.date;

  if (!isNaN(Date.parse(dateString))) {
    let dateObject = new Date(dateString);
    res.json({ unix: dateObject.valueOf(), utc: dateObject.toUTCString() });
  } else if (/\d{5,}/.test(dateString)) {
    let dateInt = parseInt(dateString);
    res.json({ unix: dateInt, utc: new Date(dateInt).toUTCString() });
  } else {
    res.json({ error: "Invalid Date" });
  }
});

// listen for requests :)
var listener = app.listen(3000, function () {
  console.log("Your app is listening on port " + listener.address().port);
});
