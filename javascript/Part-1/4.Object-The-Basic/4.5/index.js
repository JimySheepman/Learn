function User(name) {
  this.name = name;
  this.isAdmin = false;
}

let user = new User("jack");

console.log(user.name);
console.log(user.isAdmin);

function User_2(name) {
  if (!new.target) {
    return new User_2(name);
  }
  this.name = name;
}

let john = User_2("John");
console.log(john.name);

function BigUser() {
  this.name = "John";

  return { name: "Godzilla" }; // <-- returns this object
}

alert(new BigUser().name); // Godzilla, got that object

function SmallUser() {
  this.name = "John";

  return; // <-- returns this
}

alert(new SmallUser().name); // John
