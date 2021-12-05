let user = new Object();
let user_2 = {};

let user_3 = {
  name: "John",
  age: 30,
};

console.log(user_3.name);
console.log(user_3.age);
user_3.isAdmin = true;
delete user_3.age;

let user_4 = {
  name: "john",
  age: 25,
  "likes birds": true,
};
console.log(user_4["likes birds"]);
let key = "likes birds";
console.log(user_4[key]);

let fruit = prompt("Which furit to buy?", "apple");

let bag = {
  [fruit]: 5,
};
console.log(bag.fruit);

function makeUser(name, age) {
  return {
    name: name,
    age: age,
    // ...other properties
  };
}

let user_5 = makeUser("merlins", 31);
console.log(user_5.name);

let obj = {
  for: 1,
  let: 2,
  return: 3,
};
console.log(obj.for + obj.let + obj.return);

let user_6 = {};
console.log(user_6.noSuchProperty === undefined);
// "key" in Object

let user_7 = { age: 30 };
let key = "age";
console.log(key in user);

let user_8 = {
  name: "John",
  age: 30,
  isAdmin: true,
};

for (let key in user) {
  console.log(key);
  console.log(user[key]);
}

let codes = {
    "49": "Germany",
    "41": "Switzerland",
    "44": "Great Britain",
    // ..,
    "1": "USA"
  };

for (let code in codes) {
  alert(code);
}
