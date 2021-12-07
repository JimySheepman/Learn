let user = {
  name: "John",
};

let descriptor = Object.getOwnPropertyDescriptor(user, "name");

console.log(JSON.stringify(descriptor, null, 2));

user = {};
Object.defineProperty(user, "name", {
  value: "Martta",
});

descriptor = Object.getOwnPropertyDescriptor(user, "name");
console.log(JSON.stringify(descriptor, null, 2));

user = {
  name: "Samata",
};
Object.defineProperty(user, "name", {
  writable: false,
});
user.name = "Pete"; // Error

Object.defineProperty(user, "name", {
  value: "Jimy",
  enumerable: true,
  configurable: true,
});
console.log(user.name);
user.name = "Pete"; // Error

user = {
  name: "John",
  toString() {
    return this.name;
  },
};

for (let key in user) console.log(key);
console.log(Object.keys(user));

descriptor = Object.getOwnPropertyDescriptor(Math, "PI");
console.log(JSON.stringify(descriptor, null, 2));

user = {
  name: "John",
};
Object.defineProperty(user, "name", {
  configurable: false,
});

user.name = "Pete";
delete user.name; // Error

