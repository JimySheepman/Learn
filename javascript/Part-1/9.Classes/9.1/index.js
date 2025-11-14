class User {
  constructor(name) {
    this.name = name;
  }
  sayHi() {
    console.log(this.name);
  }
}
let user = new User("John");
user.sayHi();
console.log(typeof User);

class Person {
  constructor(name) {
    this.name = name;
  }
  get name() {
    return this._name;
  }
  set name(value) {
    if (value.length < 4) {
      console.log("Name is too short.");
      return;
    }
    this._name = value;
  }
}

let person = new Person("Merlins");
console.log(person.name);
person = new Person("");

class Button {
  constructor(value) {
    this.value = value;
  }
  click = () => {
    console.log(this.value);
  };
}

let button = new Button("Hello");
setTimeout(button.click, 1000);
