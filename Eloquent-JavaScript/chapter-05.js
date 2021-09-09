let total = 0,
  count = 1;
while (count <= 10) {
  total += count;
  count += 1;
}

console.log(total);

for (let i = 0; i < 10; i++) {
  console.log(i);
}

function repeatLog(n) {
  for (let i = 0; i < n; i++) {
    console.log(i);
  }
}

repeatLog(10);

function repeat(n, action) {
  for (let i = 0; i < n; i++) {
    action(i);
  }
}
repeat(3, console.log);

let labels = [];
repeat(5, (i) => {
  labels.push(`Unit ${i + 1}`);
});
console.log(labels);

function greaterThan(n) {
  return (m) => m > n;
}

let greaterThan10 = greaterThan(10);
console.log(greaterThan10(11));

function noisy(f) {
  return (...args) => {
    console.log("calling with", args);
    let result = f(...args);
    console.log("called with", args, "returned", result);
    return result;
  };
}
noisy(Math.min)(3, 2, 1, -41243, 4121);

function unless(test, then) {
  if (!test) then();
}

repeat(3, (n) => {
  unless(n % 2 == 1, () => {
    console.log(n, "is even");
  });
});

["A", "B"].forEach((l) => console.log(l));

function countBy(items, groupName) {
    let counts = [];
    for (let item of items) {
      let name = groupName(item);
      let known = counts.findIndex(c => c.name == name);
      if (known == -1) {
        counts.push({name, count: 1});
      } else {
        counts[known].count++;
      }
    }
    return counts;
  }
  
  console.log(countBy([1, 2, 3, 4, 5], n => n > 2));