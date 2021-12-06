let arr = ["I", "go", "home"];
delete arr[2];
console.log(arr.length);
console.log(arr);

arr.splice(2, 1); // form index 1 remove 1 element
console.log(arr);
arr = ["I", "study", "JavaScript", "right", "now"];
arr.splice(0, 3, "Let's", "dance"); // remove 3 first elements and replace them with another
console.log(arr);

arr = ["t", "e", "s", "t"];
console.log(arr.slice(1, 3));
console.log(arr.slice(-2));

arr = [1, 2];
console.log(arr.concat([3, 4]));
console.log(arr);
let arrayLike = {
  0: "something",
  1: "else",
  [Symbol.isConcatSpreadable]: true,
  length: 2,
};
console.log(arr.concat(arrayLike));
console.log(arr);

arr = [1, 0, false];

console.log(arr.indexOf(false));
console.log(arr.indexOf(0));
console.log(arr.indexOf(null));
console.log(arr.includes(1));

let users = [
  { id: 1, name: "John" },
  { id: 2, name: "Pete" },
  { id: 3, name: "Mary" },
];

let user = users.find((item) => item.id == 1);

console.log(user.name);

let someUsers = users.filter((item) => item.id < 3);
console.log(someUsers.length);

let lenghts = ["Bilbo", "Gandalf", "Nazgul"].map((item) => item.length);
console.log(lenghts);

arr = [1, 2, 13, 24, 5, 6];
arr.sort();
console.log(arr);

function compareNumeric(a, b) {
  if (a > b) return 1;
  if (a == b) return 0;
  if (a < b) return -1;
}
arr.sort(compareNumeric);
console.log(arr);

[1, -2, 15, 2, 0, 8].sort(function (a, b) {
  console.log(a + " <> " + b);
  return a - b;
});

arr.reverse();
console.log(arr);

let names = "Bilbo, Gandalf, Nazgul";
arr = names.split(", ");
for (let name of arr) {
  console.log(name);
}
let newStr = arr.join(";");
console.log(newStr);

arr = [1, 2, 3, 4, 5];
let result = arr.reduce((sum, current) => sum + current, 0);
console.log(result);

let army = {
  minAge: 18,
  maxAge: 27,
  canJoin(user) {
    return user.age >= this.minAge && user.age < this.maxAge;
  },
};

users = [{ age: 16 }, { age: 20 }, { age: 23 }, { age: 30 }];

// find users, for who army.canJoin returns true
let soldiers = users.filter(army.canJoin, army);

console.log(soldiers[0].age);
console.log(soldiers.length);
console.log(soldiers[1].age);
