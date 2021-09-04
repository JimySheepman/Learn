let target = {};
let proxy = new Proxy(target, {}); // empty handler

proxy.test = 5; // writing to proxy (1)
alert(target.test); // 5, the property appeared in target!

alert(proxy.test); // 5, we can read it from proxy too (2)

for (let key in proxy) alert(key); // test, iteration works (3)

let numbers = [0, 1, 2];

numbers = new Proxy(numbers, {
  get(target, prop) {
    if (prop in target) {
      return target[prop];
    } else {
      return 0; // default value
    }
  },
});

alert(numbers[1]); // 1
alert(numbers[123]); // 0 (no such item)

let dictionary = {
  Hello: "Hola",
  Bye: "Adiós",
};

alert(dictionary["Hello"]); // Hola
alert(dictionary["Welcome"]); // undefined

let dictionary = {
  Hello: "Hola",
  Bye: "Adiós",
};

dictionary = new Proxy(dictionary, {
  get(target, phrase) {
    // intercept reading a property from dictionary
    if (phrase in target) {
      // if we have it in the dictionary
      return target[phrase]; // return the translation
    } else {
      // otherwise, return the non-translated phrase
      return phrase;
    }
  },
});

// Look up arbitrary phrases in the dictionary!
// At worst, they're not translated.
alert(dictionary["Hello"]); // Hola
alert(dictionary["Welcome to Proxy"]); // Welcome to Proxy (no translation)

let numbers = [];

numbers = new Proxy(numbers, {
  // (*)
  set(target, prop, val) {
    // to intercept property writing
    if (typeof val == "number") {
      target[prop] = val;
      return true;
    } else {
      return false;
    }
  },
});

numbers.push(1); // added successfully
numbers.push(2); // added successfully
alert("Length is: " + numbers.length); // 2

numbers.push("test"); // TypeError ('set' on proxy returned false)

alert("This line is never reached (error in the line above)");

let user = {
  name: "John",
  age: 30,
  _password: "***",
};

user = new Proxy(user, {
  ownKeys(target) {
    return Object.keys(target).filter((key) => !key.startsWith("_"));
  },
});

// "ownKeys" filters out _password
for (let key in user) alert(key); // name, then: age

// same effect on these methods:
alert(Object.keys(user)); // name,age
alert(Object.values(user)); // John,30

let user = {};

user = new Proxy(user, {
  ownKeys(target) {
    return ["a", "b", "c"];
  },
});

alert(Object.keys(user)); // <empty>

let user = {};

user = new Proxy(user, {
  ownKeys(target) {
    // called once to get a list of properties
    return ["a", "b", "c"];
  },

  getOwnPropertyDescriptor(target, prop) {
    // called for every property
    return {
      enumerable: true,
      configurable: true,
      /* ...other flags, probable "value:..." */
    };
  },
});

alert(Object.keys(user)); // a, b, c

let user = {
  name: "John",
  _password: "secret",
};

alert(user._password); // secret

let user = {
  name: "John",
  _password: "***",
};

user = new Proxy(user, {
  get(target, prop) {
    if (prop.startsWith("_")) {
      throw new Error("Access denied");
    }
    let value = target[prop];
    return typeof value === "function" ? value.bind(target) : value; // (*)
  },
  set(target, prop, val) {
    // to intercept property writing
    if (prop.startsWith("_")) {
      throw new Error("Access denied");
    } else {
      target[prop] = val;
      return true;
    }
  },
  deleteProperty(target, prop) {
    // to intercept property deletion
    if (prop.startsWith("_")) {
      throw new Error("Access denied");
    } else {
      delete target[prop];
      return true;
    }
  },
  ownKeys(target) {
    // to intercept property list
    return Object.keys(target).filter((key) => !key.startsWith("_"));
  },
});

// "get" doesn't allow to read _password
try {
  alert(user._password); // Error: Access denied
} catch (e) {
  alert(e.message);
}

// "set" doesn't allow to write _password
try {
  user._password = "test"; // Error: Access denied
} catch (e) {
  alert(e.message);
}

// "deleteProperty" doesn't allow to delete _password
try {
  delete user._password; // Error: Access denied
} catch (e) {
  alert(e.message);
}

// "ownKeys" filters out _password
for (let key in user) alert(key); // name

let range = {
  start: 1,
  end: 10,
};

range = new Proxy(range, {
  has(target, prop) {
    return prop >= target.start && prop <= target.end;
  },
});

alert(5 in range); // true
alert(50 in range); // false

function delay(f, ms) {
  // return a wrapper that passes the call to f after the timeout
  return function () {
    // (*)
    setTimeout(() => f.apply(this, arguments), ms);
  };
}

function sayHi(user) {
  alert(`Hello, ${user}!`);
}

// after this wrapping, calls to sayHi will be delayed for 3 seconds
sayHi = delay(sayHi, 3000);

sayHi("John"); // Hello, John! (after 3 seconds)

function delay(f, ms) {
  return function () {
    setTimeout(() => f.apply(this, arguments), ms);
  };
}

function sayHi(user) {
  alert(`Hello, ${user}!`);
}

alert(sayHi.length); // 1 (function length is the arguments count in its declaration)

sayHi = delay(sayHi, 3000);

alert(sayHi.length); // 0 (in the wrapper declaration, there are zero arguments)

function delay(f, ms) {
  return new Proxy(f, {
    apply(target, thisArg, args) {
      setTimeout(() => target.apply(thisArg, args), ms);
    },
  });
}

function sayHi(user) {
  alert(`Hello, ${user}!`);
}

sayHi = delay(sayHi, 3000);

alert(sayHi.length); // 1 (*) proxy forwards "get length" operation to the target

sayHi("John"); // Hello, John! (after 3 seconds)

let user = {};

Reflect.set(user, "name", "John");

alert(user.name); // John

let user = {
  name: "John",
};

user = new Proxy(user, {
  get(target, prop, receiver) {
    alert(`GET ${prop}`);
    return Reflect.get(target, prop, receiver); // (1)
  },
  set(target, prop, val, receiver) {
    alert(`SET ${prop}=${val}`);
    return Reflect.set(target, prop, val, receiver); // (2)
  },
});

let name = user.name; // shows "GET name"
user.name = "Pete"; // shows "SET name=Pete"

let user = {
  _name: "Guest",
  get name() {
    return this._name;
  },
};

let userProxy = new Proxy(user, {
  get(target, prop, receiver) {
    return target[prop];
  },
});

alert(userProxy.name); // Guest

let user = {
  _name: "Guest",
  get name() {
    return this._name;
  },
};

let userProxy = new Proxy(user, {
  get(target, prop, receiver) {
    return target[prop]; // (*) target = user
  },
});

let admin = {
  __proto__: userProxy,
  _name: "Admin",
};

// Expected: Admin
alert(admin.name); // outputs: Guest (?!?)

let user = {
  _name: "Guest",
  get name() {
    return this._name;
  },
};

let userProxy = new Proxy(user, {
  get(target, prop, receiver) {
    // receiver = admin
    return Reflect.get(target, prop, receiver); // (*)
  },
});

let admin = {
  __proto__: userProxy,
  _name: "Admin",
};

alert(admin.name); // Admin

let map = new Map();

let proxy = new Proxy(map, {});

proxy.set("test", 1); // Error

let map = new Map();

let proxy = new Proxy(map, {
  get(target, prop, receiver) {
    let value = Reflect.get(...arguments);
    return typeof value == "function" ? value.bind(target) : value;
  },
});

proxy.set("test", 1);
alert(proxy.get("test")); // 1 (works!)

class User {
  #name = "Guest";

  getName() {
    return this.#name;
  }
}

let user = new User();

user = new Proxy(user, {});

alert(user.getName()); // Error

class User {
  #name = "Guest";

  getName() {
    return this.#name;
  }
}

let user = new User();

user = new Proxy(user, {
  get(target, prop, receiver) {
    let value = Reflect.get(...arguments);
    return typeof value == "function" ? value.bind(target) : value;
  },
});

alert(user.getName()); // Guest

let allUsers = new Set();

class User {
  constructor(name) {
    this.name = name;
    allUsers.add(this);
  }
}

let user = new User("John");

alert(allUsers.has(user)); // true

user = new Proxy(user, {});

alert(allUsers.has(user)); // false

let object = {
  data: "Valuable data",
};

let { proxy, revoke } = Proxy.revocable(object, {});

// pass the proxy somewhere instead of object...
alert(proxy.data); // Valuable data

// later in our code
revoke();

// the proxy isn't working any more (revoked)
alert(proxy.data); // Error

let revokes = new WeakMap();

let object = {
  data: "Valuable data",
};

let { proxy, revoke } = Proxy.revocable(object, {});

revokes.set(proxy, revoke);

// ..somewhere else in our code..
revoke = revokes.get(proxy);
revoke();

alert(proxy.data); // Error (revoked)

// Eval: run a code string

let code = 'alert("Hello")';
eval(code); // Hello

let value = eval("1+1");
alert(value); // 2

let value = eval("let i = 0; ++i");
alert(value); // 1

let a = 1;

function f() {
  let a = 2;

  eval("alert(a)"); // 2
}

f();

let x = 5;
eval("x = 10");
alert(x); // 10, value modified

// reminder: 'use strict' is enabled in runnable examples by default

eval("let x = 5; function f() {}");

alert(typeof x); // undefined (no such variable)
// function f is also not visible

let x = 1;
{
  let x = 5;
  window.eval("alert(x)"); // 1 (global variable)
}

let f = new Function("a", "alert(a)");

f(5); // 5

// Currying

function curry(f) {
  // curry(f) does the currying transform
  return function (a) {
    return function (b) {
      return f(a, b);
    };
  };
}

// usage
function sum(a, b) {
  return a + b;
}

let curriedSum = curry(sum);

alert(curriedSum(1)(2)); // 3

function sum(a, b) {
  return a + b;
}

let curriedSum = _.curry(sum); // using _.curry from lodash library

alert(curriedSum(1, 2)); // 3, still callable normally
alert(curriedSum(1)(2)); // 3, called partially

function curry(func) {
  return function curried(...args) {
    if (args.length >= func.length) {
      return func.apply(this, args);
    } else {
      return function (...args2) {
        return curried.apply(this, args.concat(args2));
      };
    }
  };
}

function sum(a, b, c) {
  return a + b + c;
}

let curriedSum = curry(sum);

alert(curriedSum(1, 2, 3)); // 6, still callable normally
alert(curriedSum(1)(2, 3)); // 6, currying of 1st arg
alert(curriedSum(1)(2)(3)); // 6, full currying

// func is the function to transform
function curried(...args) {
  if (args.length >= func.length) {
    // (1)
    return func.apply(this, args);
  } else {
    return function (...args2) {
      // (2)
      return curried.apply(this, args.concat(args2));
    };
  }
}

// Reference Type

let user = {
  name: "John",
  hi() {
    alert(this.name);
  },
  bye() {
    alert("Bye");
  },
};

user.hi(); // works

// now let's call user.hi or user.bye depending on the name
(user.name == "John" ? user.hi : user.bye)(); // Error!

let user = {
  name: "John",
  hi() {
    alert(this.name);
  },
};

// split getting and calling the method in two lines
let hi = user.hi;
hi(); // Error, because this is undefined

// BigInt

const bigint = 1234567890123456789012345678901234567890n;

const sameBigint = BigInt("1234567890123456789012345678901234567890");

const bigintFromNumber = BigInt(10); // same as 10n

alert(1n + 2n); // 3

alert(5n / 2n); // 2

alert(1n + 2); // Error: Cannot mix BigInt and other types

let bigint = 1n;
let number = 2;

// number to bigint
alert(bigint + BigInt(number)); // 3

// bigint to number
alert(Number(bigint) + number); // 3

let bigint = 1n;

alert(+bigint); // error

alert(2n > 1n); // true

alert(2n > 1); // true

alert(1 == 1n); // true

alert(1 === 1n); // false

if (0n) {
  // never executes
}

alert(1n || 2); // 1 (1n is considered truthy)

alert(0n || 2); // 2 (0n is considered falsy)
