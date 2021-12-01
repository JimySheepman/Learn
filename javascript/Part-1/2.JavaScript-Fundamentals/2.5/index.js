let message = "hello";
message = 1234;

// Number
let n = 123;
n = 12.45123;

console.log(1 / 0);
console.log(Infinity);
console.log("not a number" / 2);

// BigInt

const bigInt = 12314123141511561616636n;
console.log(bigInt);

// string

let str = "hello";
let phrase = `can embed another ${str}`;
console.log(phrase);

// Boolean
let flag = true;

let isGreater = 4 > 1;
console.log(isGreater);

//nul
let age = null;

// undefined

let sow;
console.log(sow);

// typeof operator
typeof undefined; // "undefined"

typeof 0; // "number"

typeof 10n; // "bigint"

typeof true; // "boolean"

typeof "foo"; // "string"

typeof Symbol("id"); // "symbol"

typeof Math; // "object"  (1)

typeof null; // "object"  (2)

typeof alert; // "function"  (3)
