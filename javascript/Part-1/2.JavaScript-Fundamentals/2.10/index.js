let year = prompt(
  "In which year was ECMAScrpit-2015 specification published? ",
  ""
);
if (year == 2015) console.log(true);
else console.log(false);

if (1) {
  console.log(true);
} else if (0) {
  console.log(false);
} else {
  console.log(null);
}

// let result = condition ? value1 : value2;
let age = prompt("age?", 18);

let message =
  age < 3
    ? "Hi, baby!"
    : age < 18
    ? "Hello!"
    : age < 100
    ? "Greetings!"
    : "What an unusual age!";

if (age < 3) {
  message = "Hi, baby!";
} else if (age < 18) {
  message = "Hello!";
} else if (age < 100) {
  message = "Greetings!";
} else {
  message = "What an unusual age!";
}
