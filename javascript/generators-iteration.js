function* generateSequence() {
  yield 1;
  yield 2;
  return 3;
}

// "generator function" creates "generator object"
let generator = generateSequence();
alert(generator); // [object Generator]

function* generateSequence() {
  yield 1;
  yield 2;
  return 3;
}

let generator = generateSequence();

let one = generator.next();

alert(JSON.stringify(one)); // {value: 1, done: false}

function* generateSequence() {
  yield 1;
  yield 2;
  return 3;
}

let generator = generateSequence();

for (let value of generator) {
  alert(value); // 1, then 2
}

function* generateSequence() {
  yield 1;
  yield 2;
  yield 3;
}

let generator = generateSequence();

for (let value of generator) {
  alert(value); // 1, then 2, then 3
}

function* generateSequence() {
  yield 1;
  yield 2;
  yield 3;
}

let sequence = [0, ...generateSequence()];

alert(sequence); // 0, 1, 2, 3

let range = {
  from: 1,
  to: 5,

  // for..of range calls this method once in the very beginning
  [Symbol.iterator]() {
    // ...it returns the iterator object:
    // onward, for..of works only with that object, asking it for next values
    return {
      current: this.from,
      last: this.to,

      // next() is called on each iteration by the for..of loop
      next() {
        // it should return the value as an object {done:.., value :...}
        if (this.current <= this.last) {
          return { done: false, value: this.current++ };
        } else {
          return { done: true };
        }
      },
    };
  },
};

// iteration over range returns numbers from range.from to range.to
alert([...range]); // 1,2,3,4,5

let range = {
  from: 1,
  to: 5,

  *[Symbol.iterator]() {
    // a shorthand for [Symbol.iterator]: function*()
    for (let value = this.from; value <= this.to; value++) {
      yield value;
    }
  },
};

alert([...range]); // 1,2,3,4,5

function* generateSequence(start, end) {
  for (let i = start; i <= end; i++) yield i;
}

function* generatePasswordCodes() {
  // 0..9
  yield* generateSequence(48, 57);

  // A..Z
  yield* generateSequence(65, 90);

  // a..z
  yield* generateSequence(97, 122);
}

let str = "";

for (let code of generatePasswordCodes()) {
  str += String.fromCharCode(code);
}

alert(str); // 0..9A..Za..z

function* generateSequence(start, end) {
  for (let i = start; i <= end; i++) yield i;
}

function* generateAlphaNum() {
  // yield* generateSequence(48, 57);
  for (let i = 48; i <= 57; i++) yield i;

  // yield* generateSequence(65, 90);
  for (let i = 65; i <= 90; i++) yield i;

  // yield* generateSequence(97, 122);
  for (let i = 97; i <= 122; i++) yield i;
}

let str = "";

for (let code of generateAlphaNum()) {
  str += String.fromCharCode(code);
}

alert(str); // 0..9A..Za..z

function* gen() {
  // Pass a question to the outer code and wait for an answer
  let result = yield "2 + 2 = ?"; // (*)

  alert(result);
}

let generator = gen();

let question = generator.next().value; // <-- yield returns the value

generator.next(4); // --> pass the result into the generator

function* gen() {
  let ask1 = yield "2 + 2 = ?";

  alert(ask1); // 4

  let ask2 = yield "3 * 3 = ?";

  alert(ask2); // 9
}

let generator = gen();

alert(generator.next().value); // "2 + 2 = ?"

alert(generator.next(4).value); // "3 * 3 = ?"

alert(generator.next(9).done); // true

function* gen() {
  try {
    let result = yield "2 + 2 = ?"; // (1)

    alert(
      "The execution does not reach here, because the exception is thrown above"
    );
  } catch (e) {
    alert(e); // shows the error
  }
}

let generator = gen();

let question = generator.next().value;

generator.throw(new Error("The answer is not found in my database")); // (2)

function* generate() {
  let result = yield "2 + 2 = ?"; // Error in this line
}

let generator = generate();

let question = generator.next().value;

try {
  generator.throw(new Error("The answer is not found in my database"));
} catch (e) {
  alert(e); // shows the error
}

// Async iteration and generators

let range = {
  from: 1,
  to: 5,

  [Symbol.iterator]() {
    // called once, in the beginning of for..of
    return {
      current: this.from,
      last: this.to,

      next() {
        // called every iteration, to get the next value
        if (this.current <= this.last) {
          return { done: false, value: this.current++ };
        } else {
          return { done: true };
        }
      },
    };
  },
};

for (let value of range) {
  alert(value); // 1 then 2, then 3, then 4, then 5
}

let range = {
  from: 1,
  to: 5,

  [Symbol.asyncIterator]() {
    // (1)
    return {
      current: this.from,
      last: this.to,

      async next() {
        // (2)

        // note: we can use "await" inside the async next:
        await new Promise((resolve) => setTimeout(resolve, 1000)); // (3)

        if (this.current <= this.last) {
          return { done: false, value: this.current++ };
        } else {
          return { done: true };
        }
      },
    };
  },
};

(async () => {
  for await (let value of range) {
    // (4)
    alert(value); // 1,2,3,4,5
  }
})();

function* generateSequence(start, end) {
  for (let i = start; i <= end; i++) {
    yield i;
  }
}

for (let value of generateSequence(1, 5)) {
  alert(value); // 1, then 2, then 3, then 4, then 5
}

let range = {
  from: 1,
  to: 5,

  *[Symbol.iterator]() {
    // a shorthand for [Symbol.iterator]: function*()
    for (let value = this.from; value <= this.to; value++) {
      yield value;
    }
  },
};

for (let value of range) {
  alert(value); // 1, then 2, then 3, then 4, then 5
}

async function* generateSequence(start, end) {
  for (let i = start; i <= end; i++) {
    // Wow, can use await!
    await new Promise((resolve) => setTimeout(resolve, 1000));

    yield i;
  }
}

(async () => {
  let generator = generateSequence(1, 5);
  for await (let value of generator) {
    alert(value); // 1, then 2, then 3, then 4, then 5 (with delay between)
  }
})();

let range = {
    from: 1,
    to: 5,
  
    // this line is same as [Symbol.asyncIterator]: async function*() {
    async *[Symbol.asyncIterator]() {
      for(let value = this.from; value <= this.to; value++) {
  
        // make a pause between values, wait for something
        await new Promise(resolve => setTimeout(resolve, 1000));
  
        yield value;
      }
    }
  };
  
  (async () => {
  
    for await (let value of range) {
      alert(value); // 1, then 2, then 3, then 4, then 5
    }
  
  })();


  (async () => {

    let count = 0;
  
    for await (const commit of fetchCommits('javascript-tutorial/en.javascript.info')) {
  
      console.log(commit.author.login);
  
      if (++count == 100) { // let's stop at 100 commits
        break;
      }
    }
  
  })();
  
  // Note: If you are running this in an external sandbox, you'll need to paste here the function fetchCommits described above