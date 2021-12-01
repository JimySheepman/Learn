alert("Hello");
alert("World");

alert("Hello");
alert("World");

alert("There will be an error after this message")[(1, 2)].forEach(alert);

function f() {
  // no semicolon needed after function declaration
}

for (;;) {
  // no semicolon needed after the loop
}

("use strict");

let x = 5;
x = "John";

typeof null == "object"; // error in the language
typeof function () {} == "function"; // functions are treated specially

let userName = prompt("Your name?", "Alice");
let isTeaWanted = confirm("Do you want some tea?");

alert("Visitor: " + userName); // Alice
alert("Tea wanted: " + isTeaWanted); // true

alert("1" + 2); // '12', string
alert(1 + "2"); // '12', string

alert(0 == false); // true
alert(0 == ""); // true

// 1
while (condition) {
  continue;
}

// 2
do {
  continue;
} while (condition);

// 3
for (let i = 0; i < 10; i++) {
  continue;
}

let age = prompt("Your age?", 18);

switch (age) {
  case 18:
    alert("Won't work"); // the result of prompt is a string, not a number
    break;

  case "18":
    alert("This works!");
    break;

  default:
    alert("Any value not equal to one above");
}

function sum(a, b) {
  let result = a + b;

  return result;
}

let sum = function (a, b) {
  let result = a + b;

  return result;
};

// expression at the right side
let sum = (a, b) => a + b;

// or multi-line syntax with { ... }, need return here:
let sum = (a, b) => {
  // ...
  return a + b;
};

// without arguments
let sayHi = () => alert("Hello");

// with a single argument
let double = (n) => n * 2;
