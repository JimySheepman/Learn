function slow(x) {
  console.log(`Called with ${x}`);
  return x;
}

function cachingDecorator(func) {
  let cache = new Map();
  return function (x) {
    if (cache.has(x)) {
      return cache.get(x);
    }
    let result = func(x);
    cache.set(x, result);
    return result;
  };
}

slow = cachingDecorator(slow);
console.log(slow(1));
console.log("Again: " + slow(1));
console.log(slow(2));
console.log("Again: " + slow(2));

let worker = {
  someMethod() {
    return 1;
  },

  slow(x) {
    alert("Called with " + x);
    return x * this.someMethod();
  },
};

function cachingDecorator(func) {
  let cache = new Map();
  return function (x) {
    if (cache.has(x)) {
      return cache.get(x);
    }
    let result = func(x);
    cache.set(x, result);
    return result;
  };
}

alert(worker.slow(1));
worker.slow = cachingDecorator(worker.slow);
alert(worker.slow(2));
