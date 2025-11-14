let sum = new Function("a", "b", "return a+b");
console.log(sum(3, 6));

let sayHi = new Function('console.log("Hello")');
sayHi();
