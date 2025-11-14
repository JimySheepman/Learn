alert("Hello_1");
window.alert("Hello_2");
var gVar = 5;
console.log(window.gVar);

let gLet = 5;

console.log(window.gLet);

window.currentUser = {
  name: "John",
};

console.log(currentUser.name);
console.log(window.currentUser.name);

if (!window.Promise) {
    alert("Your browser is really old!");
  }