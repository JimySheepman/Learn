let map = new Map();

map.set("1", "str1");
map.set(1, "num1");
map.set(true, "bool1");
console.log(map.get(1));
console.log(map.get("1"));
console.log(map.size);

let recipeMap = new Map([
  ["cucumber", 500],
  ["tomatoes", 350],
  ["onion", 50],
]);

for (let vegetable of recipeMap.keys()) {
  console.log(vegetable);
}

for (let amount of recipeMap.values()) {
  console.log(amount);
}

for (let entry of recipeMap) {
  console.log(entry);
}

let set = new Set();

let john = { name: "John" };
let pete = { name: "Pete" };
let mary = { name: "Mary" };

// visits, some users come multiple times
set.add(john);
set.add(pete);
set.add(mary);
set.add(john);
set.add(mary);

// set keeps only unique values
console.log(set.size); // 3

for (let user of set) {
  console.log(user.name); // John (then Pete and Mary)
}
