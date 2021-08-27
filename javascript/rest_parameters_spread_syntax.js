function sum(a, b) {
  return a + b;
}
console.log(sum(1, 2, 3, 4, 5));

function allSum(...args) {
  let sum = 0;
  for (let arg of args) sum += arg;
  return sum;
}

console.log(allSum(1));
console.log(allSum(1, 2));
console.log(allSum(4, 3, 2, 2, 1, 2));

function showName(name, surname, ...title) {
  console.log(name + " " + surname);
  console.log(title[0]);
  console.log(title[1]);
  console.log(title.length);
}

showName("Julius", "Caesar", "Konsil", "İmparator", "Salvador");

function argumanNames() {
  console.log(arguments.length);
  console.log(arguments[0]);
  console.log(arguments[1]);
}

argumanNames("Julius", "Caesar");
argumanNames("Ahmet");

function f() {
  let showArg = () => console.log(arguments[0]);
  showArg();
}

f(1);

console.log(Math.max(3, 5, 1));
let arr1 = [1, -8, 3, 4];
let arr2 = [8, 3, -2, 1];

console.log(Math.max(...arr1, ...arr2));
console.log(Math.max(...arr1, ...arr2,25));
console.log(Math.min(...arr1, ...arr2));

let str = "Hello";
console.log(str);
console.log( [...str] );