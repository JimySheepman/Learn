function showMessage() {
  console.log("hello word!");
}

function showMessage2() {
  let message = "Hello, I'm JavaScript!"; // local variable
  console.log(message);
}

let userName = "Jhon";
function showMessage3() {
  let message = "Hello, " + userName;
  console.log(message);
}

showMessage();
showMessage2();
showMessage3();

function showMessage4(from, text) {
  from = "*" + from + "*";
  console.log(from + ": " + text);
}

let from = "ann";
showMessage4(from, "hello ! ! ");

function showMessage5(from, text = "Default text") {
  console.log(from + ": " + text);
}
showMessage5("ann");

function showMessage6(from, text = anotherFunction()) {
  // anotherFunction() only executed if no text given
  // its result becomes the value of text
}

function showMessage7(from, text = "Default text") {
  return console.log(from + ": " + text);
}
