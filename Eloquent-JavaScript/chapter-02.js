let caught = 5 * 5;
console.log(caught);

let ten = 10;

console.log(ten * ten);

let mode = "light";
console.log(mode);
mode = "dark";
console.log(mode);

let luigisDebt = 140;
luigisDebt = luigisDebt - 35;
console.log(luigisDebt);

let one = 1,
  two = 2;
console.log(one + two);

var name = "merlins";
const greeting = "Hello ";
console.log(greeting + name);

// prompt("Enter passcode");

let x = 30;
console.log("the value of x is", x);

console.log(Math.max(2, 4));

/* let thNumber = Number(prompt("Pick a number"));
console.log("Your number is the square root of "+ thNumber * thNumber);
 */

/* let theNumber = Number(prompt("Pich a number"));
if(!Number.isNaN(theNumber)){
    console.log("Your number is the square root of " + theNumber * theNumber);
} */

if (1 + 1 == 2) console.log("It's true");

/* let theNumber=Number(prompt("Pick a number"));
if(!Number.isNaN(theNumber)){
    console.log("Your number is the square root of " + theNumber * theNumber);
}else{
    console.log("Hey. Why didn't you give me a number? ");
} */

/* 
let num = Number(prompt("Pick a number"))

if( num <10){
    console.log("Small");
}else if(num<100){
    console.log("Medium");
}else{
    console.log("Large");
} */

console.log(0);
console.log(2);
console.log(4);
console.log(6);
console.log(8);
console.log(10);
console.log(12);

let number = 0;
while (number <= 12) {
  console.log(number);
  number += 2;
}

/* let result = 1;
let counter=0;
while (counter < 10) {
    result *= 2;
    counter += 1;
}
console.log(result);
 */

/* let yourName;
do{
    yourName = prompt("Who I am?");
}while(!yourName);
console.log(yourName); */

if (false != true) {
  console.log("That makes sense.");
  if (1 < 2) {
    console.log("No surprise there.");
  }
}

for (let number = 0; number <= 12; number = number + 2) {
  console.log(number);
}

let result = 1;
for (let counter = 0; counter < 10; counter = counter + 1) {
  result = result * 2;
}
console.log(result);

for (let current = 20; ; current = current + 1) {
  if (current % 7 == 0) {
    console.log(current);
    break;
  }
}

switch (prompt("What is the weather like?")) {
  case "rainy":
    console.log("Remember to bring an umbrella.");
    break;
  case "sunny":
    console.log("Dress lightly.");
  case "cloudy":
    console.log("Go outside.");
    break;
  default:
    console.log("Unknown weather type!");
    break;
}
