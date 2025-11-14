/* let user = {
  name: "John",
  surname: "Smith",
  get fullName() {
    return `${this.name} ${this.surname}`;
  },
  set fullName(value) {
    [this.name, this.surname] = value.split(" ");
  },
};
console.log(user.fullName);

user.fullName = "Jimy Sheepman";
console.log(user.name);
console.log(user.surname);
 */
user = {
  get name() {
    return this._name;
  },
  set name(value) {
    if (value.length < 4) {
      console.log("Name is too short, min char = 5");
      return;
    }
    this._name  = value;
  },
};

user.name = "Pete"
console.log(user.name);
user.name =""
console.log(user.name)