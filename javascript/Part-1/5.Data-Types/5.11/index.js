let now = new Date();
console.log(now);

let x = new Date(0);
console.log(x);
let y = new Date(24 * 3600 * 1000);
console.log(y);
let z = new Date(-24 * 3600 * 1000);
console.log(z);
let date = new Date("2017-01-26");
console.log(date);
date = new Date(2011, 0, 1, 2, 3, 4, 567);
console.log(date);
date = new Date();
console.log(date.getHours(), date.getUTCHours(), date.getTimezoneOffset());

let today = new Date();

today.setHours(0);
console.log(today);
today.setHours(0, 0, 0, 0);
console.log(today);
date = new Date(2013, 0, 32);
console.log(date);

date = new Date();
console.log(+date);

let start = Date.now();
for (let i = 0; i < 100000; i++) {
  let doSomething = i * i * i;
}
let end = Date.now();

console.log(`The loop took ${end - start} ms`);

function diffSubtract(date1, date2) {
  return date2 - date1;
}

function diffGetTime(date1, date2) {
  return date2.getTime() - date1.getTime();
}
