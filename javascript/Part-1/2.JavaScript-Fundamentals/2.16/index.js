function sayHi() {
  alert("Hello");
}

let func = sayHi;
func();
sayHi();

function ask(question, yes, no) {
  if (confirm(question)) yes();
  else no();
}

function showOk() {
  alert("You agreed.");
}

function showCancel() {
  alert("You canceled the execution.");
}

ask("Do you agree?", showOk, showCancel);

let sum = function (a, b) {
  return a + b;
};
