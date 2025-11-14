function* generateSequence() {
  yield 1;
  yield 2;
  return 3;
}

let generator = generateSequence();
console.log(generator);
let one = generator.next();
console.log(JSON.stringify(one));
let two = generator.next();
console.log(JSON.stringify(two));
let three = generator.next();
console.log(JSON.stringify(three));

for (let value of generator) {
  alert(value);
}

let sequence = [0, ...generateSequence()];
console.log(sequence);

let range = {
  from: 1,
  to: 5,
  [Symbol.iterator]() {
    return {
      current: this.from,
      last: this.to,
      next() {
        if (this.current <= this.last) {
          return { done: false, value: this.current++ };
        } else {
          return { done: true };
        }
      },
    };
  },
};
console.log([...range]);

function* generateSequence2(start, end) {
  for (let i = start; i <= end; i++) yield i;
}

function* generatePasswordCodes() {
  yield* generateSequence2(48, 57);
  yield* generateSequence2(65, 90);
  yield* generateSequence2(97, 122);
}

let str = "";

for (let code of generatePasswordCodes()) {
  str += String.fromCharCode(code);
}
console.log(str);

function* gen() {
  try {
    let result = yield "2 + 2 = ?";
    console.log(
      "The execution does not reach here, because the exception is thrown above"
    );
  } catch (e) {
    console.log(e);
  }
}
generator = gen();
let question = generator.next().value;
generator.throw(new Error("The answer is not found in my database"));
