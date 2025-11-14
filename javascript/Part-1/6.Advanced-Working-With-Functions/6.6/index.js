function sayHi() {
  console.log("hi");
}
console.log(sayHi.name);

let sayHi_2 = function () {
  console.log("hi");
};
console.log(sayHi_2.name);

function f(sayHi_3 = function () {}) {
  console.log(sayHi_3.name);
}
f();
function f1(a) {}
function f2(a, b) {}
console.log(f1.length); // 1
console.log(f2.length); // 2

function ask(question, ...handlers) {
  let isYes = confirm(question);

  for (let handler of handlers) {
    if (handler.length == 0) {
      if (isYes) handler();
    } else {
      handler(isYes);
    }
  }
}

ask(
  "Question?",
  () => alert("You said yes"),
  (result) => alert(result)
);

function makeCounter() {
  function counter() {
    return counter.count++;
  }
  counter.count = 0;
  return counter;
}

let counter = makeCounter();
console.log(counter());
console.log(counter());
