export let months = [
  "Jan",
  "Feb",
  "Mar",
  "Apr",
  "Aug",
  "Sep",
  "Oct",
  "Nov",
  "Dec",
];
export const MODULES_BECAME_STANDARD_YEAR = 2015;
export class User {
  constructor(name) {
    this.name = name;
  }
}

function sayHi(user) {
  alert(`Hello, ${user}!`);
}

function sayBye(user) {
  alert(`Bye, ${user}!`);
}

export { sayHi, sayBye };

import * as say from "./say.js";

say.sayHi("John");
say.sayBye("John");


export function sayHi() { "..." }
export function sayBye() { "..." }
export function becomeSilent() { "..." }
export {sayHi as hi, sayBye as bye};

export default class User { // just add "default"
    constructor(name) {
      this.name = name;
    }
  }

  export default class User {
    constructor(name) {
      this.name = name;
    }
  }
  
  export function sayHi(user) {
    alert(`Hello, ${user}!`);
  }