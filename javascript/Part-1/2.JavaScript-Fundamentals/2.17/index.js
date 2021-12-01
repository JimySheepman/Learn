// let func = (arg1, arg2, ..., argN) => expression
let func = function (arg1, arg2, argN) {
  return expression;
};

let sum = (a, b) => a + b;
console.log(sum(123, 123));

let double = (n) => n * 2;

console.log(double(3));

let sayHi = () => console.log("Hello Word");

sayHi();

let age = prompt("What is your age?", 18);

let welcome = age < 18 ? () => alert("Hello") : () => alert("Greetings!");

welcome();

let sum = (a, b) => {
  let result = a + b;
  return result;
};

console.log(sum(1, 2));
