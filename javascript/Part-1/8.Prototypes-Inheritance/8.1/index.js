let animal = {
  eats: true,
  walk() {
    console.log("Animal walk");
  },
};
let rabbit = {
  jumps: true,
  __proto__: animal,
};
let longEar = {
  earLength: 10,
  __proto__: rabbit,
};
rabbit.__proto__ = animal;
console.log(rabbit.eats);
console.log(rabbit.jumps);
rabbit.walk();
longEar.walk();
console.log(longEar.jumps);

let user = {
  name: "john",
  surname: "Smith",
  set fullName(value) {
    [this.name, this.surname] = value.split(" ");
  },
  get fullName() {
    return `${this.name} ${this.surname}`;
  },
};
let admin = {
  __proto__: user,
  isAdmin: true,
};

console.log(admin.fullName);
admin.fullName = "Alice Cooper";

console.log(admin.fullName);
console.log(user.fullName);
