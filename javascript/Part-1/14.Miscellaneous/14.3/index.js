function curry(f) {
  return function (a) {
    return function (b) {
      return f(a, b);
    };
  };
}
function sum(a, b) {
  return a + b;
}

let curriedSum = curry(sum);

alert(curriedSum(1)(2));
