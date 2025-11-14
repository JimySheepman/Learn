let arr = new Array();
let arr2 = [];

let fruits = ["Apple", "Orange", "Plum"];
console.log(fruits[0]);
console.log(fruits[1]);
console.log(fruits[2]);

fruits[2] = "Pear";
fruits[3] = "Lemon";

console.log(fruits);

let arr3 = [
  "Apple",
  { name: "John" },
  true,
  function () {
    console.log("hello world");
  },
];

console.log(arr3[1].name);
console.log(arr3);

console.log(fruits.pop());
console.log(fruits);
console.log(fruits.push("Lemon"));
console.log(fruits);
console.log(fruits.shift());
console.log(fruits);
console.log(fruits.unshift("Apple"));
console.log(fruits);

fruits = ["Banana"];
let arr4 = fruits;
console.log(arr4 === fruits);
arr4.push("Pear");
console.log(fruits);

for (let key in arr4) {
  console.log(arr4[key]);
}

let matrix = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9],
];

console.log(matrix[1][1]);
console.log(matrix);

arr = [1, 2, 3];
console.log(arr);
console.log(String(arr) === "1,2,3");
