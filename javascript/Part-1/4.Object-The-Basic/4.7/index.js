let id1 = Symbol("id");
let id2 = Symbol("id");

alert(id1 == id2); // false

let user = {
  // belongs to another code
  name: "John",
};

let id = Symbol("id");

user[id] = 1;

alert(user[id]);

// read from the global registry
let id = Symbol.for("id"); // if the symbol did not exist, it is created

// read it again (maybe from another part of the code)
let idAgain = Symbol.for("id");

// the same symbol
alert(id === idAgain); // true

// get symbol by name
let sym = Symbol.for("name");
let sym2 = Symbol.for("id");

// get name by symbol
alert(Symbol.keyFor(sym)); // name
alert(Symbol.keyFor(sym2)); // id
