const fs = require("fs");
// create

fs.writeFileSync(
  "employees.json",
  '{"name": "Employee 1 Name", "salary": 2000}',
  "utf-8",
  (err) => {
    if (err) console.log(err);
    console.log("File is create.");
  }
);
// read
fs.readFileSync("employees.json", "utf-8", (err, data) => {
  if (err) console.log(err);
  console.log("File is read:");
  console.log(data);
});
// update
let data = fs.readFileSync("employees.json", (err, data) => {
  if (err) console.log(err);
  console.log(data);
});

let parseData = JSON.parse(data);
parseData["city"] = "Ankara";

fs.writeFileSync(
  "employees.json",
  JSON.stringify(parseData),
  "utf-8",
  (err) => {
    if (err) console.log(err);
    console.log("File is create.");
  }
);

// Delete(if you want to work this function, delete to comment lines'/* and */')
/* 
fs.unlinkSync("employees.json", (err) => {
  if (err) console.log(err);
  console.log('File is delete')
});
 */
