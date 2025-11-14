function sum(a, b) {
  return a + b;
}

console.log(sum(1, 2, 3, 4, 5));

function sumAll(...args) {
  let sum = 0;
  for (let arg of args) {
    sum += arg;
  }
  return sum;
}

console.log(sumAll(1));
console.log(sumAll(1, 2));
console.log(sumAll(1, 2, 3));

function showName(firstName, lastName, ...titles) {
  console.log(firstName + " " + lastName);
  console.log(titles[0]);
  console.log(titles[1]);
  console.log(titles.length);
}
showName("Jimy", "Sheepman", "Programmer", "Software Engineer");

let arr = [3, 4, 6, 12];
console.log("M " + Math.max(3, 4, 6, 12));
console.log("M " + Math.max(arr));
console.log("Arg " + Math.max(...arr));

let arr1 = [3, 5, 1];
let arr2 = [8, 9, 15];
let merged = [0, ...arr, 2, ...arr2];
console.log(merged);

let str = "Hello";

console.log([...str]);

let arr3 = [1, 2, 3];
let arrCopy = [...arr3];

console.log(JSON.stringify(arr3) === JSON.stringify(arrCopy));
console.log(arr3 === arrCopy);
arr3.push(4);
console.log(arr3);
console.log(arrCopy);

let obj = { a: 1, b: 2, c: 3 };
let objCopy = { ...obj };

console.log(JSON.stringify(obj) === JSON.stringify(objCopy));
console.log(obj === objCopy);
obj.d = 4;
console.log(JSON.stringify(obj));
console.log(JSON.stringify(objCopy));
