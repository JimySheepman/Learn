let message = "hello!";
let phrase = message;

let user = {
  name: "John",
};
let admin = user;

admin.name = "merlins";
console.log(user.name); // referance => merlins

let a = {};
let b = a;
console.log(a == b);
console.log(a === b);

let a = {};
let b = {};
console.log(a == b);

let user_2 = {
  name: "john",
  age: 30,
};
let clone = {};

for (let key in user) {
  clone[key] = user[key];
}

clone.name = "pette";
console.log(user.name); // clone is a not referance => john

let user_3 = { name: "John" };
let permissions1 = { canView: true };
let permissions2 = { canEdit: true };
Object.assign(user, permissions1, permissions2);
// now user = { name: "John", canView: true, canEdit: true }

let user_4 = {
  name: "john",
  sizes: {
    height: 182,
    width: 50,
  },
};

console.log(user_4.sizes.height);

let clone = Object.assign({},user);
console.log(user.sizes === clone.sizes)// true
user.sizes.width++;
console.log(clone.sizes.width) // 51