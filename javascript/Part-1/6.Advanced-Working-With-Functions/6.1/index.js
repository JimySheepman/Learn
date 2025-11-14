function pow(x, n) {
  let result = 1;

  for (let i = 0; i < n; i++) {
    result *= x;
  }
  return result;
}
console.log("Iterative thinking:", pow(2, 3));

function pow_2(x, n) {
  if (n == 1) {
    return x;
  } else {
    return x * pow_2(x, n - 1);
  }
}
console.log("Recursive thinking:", pow_2(2, 3));

let company = {
  sales: [
    { name: "John", salary: 1000 },
    { name: "Alice", salary: 1600 },
  ],
  development: {
    sites: [
      { name: "Peter", salary: 2000 },
      { name: "Alex", salary: 1800 },
    ],
    internals: [{ name: "Jack", salary: 1300 }],
  },
};

function sumSalary(department) {
  if (Array.isArray(department)) {
    return department.reduce((prev, current) => prev + current.salary, 0);
  } else {
    let sum = 0;
    for (let subdep of Object.values(department)) {
      sum += sumSalary(subdep);
    }
    return sum;
  }
}

console.log(sumSalary(company));