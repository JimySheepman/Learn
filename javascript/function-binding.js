let user = {
  firstName: "John",
  sayHi() {
    console.log(`Hello, ${this.firstName}!`);
  },
};

setTimeout(() => user.sayHi(), 1000);

// ...within 1 second
user = {
  sayHi() {
    console.log("Another user in setTimeout!");
  },
};

// Another user in setTimeout?!?

let user_1 = {
  firstName: "John",
};

function func() {
  console.log(this.firstName);
}

let funcUser = func.bind(user_1);
funcUser(); // John

let user_2 = {
  firstName: "John",
};

function func(phrase) {
  console.log(phrase + ", " + this.firstName);
}

// this'i user'a bağla.
let funcUser_1 = func.bind(user_2);

funcUser_1("Hello"); // Hello, John ("Hello" iletildi ve this=user oldu)

function mul(a, b) {
  return a * b;
}

let double = mul.bind(null, 2);

console.log(double(3)); // = mul(2, 3) = 6
console.log(double(4)); // = mul(2, 4) = 8
console.log(double(5)); // = mul(2, 5) = 10

let triple = mul.bind(null, 3);

console.log(triple(3)); // = mul(3, 3) = 9
console.log(triple(4)); // = mul(3, 4) = 12
console.log(triple(5)); // = mul(3, 5) = 15

function curry(func) {
  return function (a) {
    return function (b) {
      return func(a, b);
    };
  };
}
function sum(a, b) {
  return a + b;
}

let carriedSum = curry(sum);

console.log(carriedSum(1)(2)); // 3
