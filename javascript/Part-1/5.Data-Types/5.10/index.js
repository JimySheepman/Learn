let arr = ["John", "Smith"];
let [firstName, surname] = arr;

alert(firstName);
alert(surname);
let [firstName, surname] = "John Smith".split(" ");
alert(firstName);
alert(surname);

let firstName = arr[0];
let surname = arr[1];

let [firstName, , title] = [
  "Julius",
  "Caesar",
  "Consul",
  "of the Roman Republic",
];

alert(title);

let [a, b, c] = "abc";
let [one, two, three] = new Set([1, 2, 3]);

let user = {};
[user.name, user.surname] = "John Smith".split(" ");

alert(user.name);
alert(user.surname);

let user = {
  name: "John",
  age: 30,
};

for (let [key, value] of Object.entries(user)) {
  alert(`${key}:${value}`);
}

let user = new Map();
user.set("name", "John");
user.set("age", "30");

for (let [key, value] of user) {
  alert(`${key}:${value}`);
}

let guest = "Jane";
let admin = "Pete";

[guest, admin] = [admin, guest];

alert(`${guest} ${admin}`);

let [name1, name2] = ["Julius", "Caesar", "Consul", "of the Roman Republic"];

alert(name1);
alert(name2);

let [firstName, surname] = [];

alert(firstName);
alert(surname);

let options = {
  title: "Menu",
  width: 100,
  height: 200,
};

let { title, width, height } = options;

alert(title); // Menu
alert(width); // 100
alert(height); // 200

let options = {
  title: "Menu",
  height: 200,
  width: 100,
};

let { title, ...rest } = options;

alert(rest.height); // 200
alert(rest.width); // 100

let options = {
  size: {
    width: 100,
    height: 200,
  },
  items: ["Cake", "Donut"],
  extra: true,
};

// destructuring assignment split in multiple lines for clarity
let {
  size: {
    // put size here
    width,
    height,
  },
  items: [item1, item2], // assign items here
  title = "Menu", // not present in the object (default value is used)
} = options;

alert(title); // Menu
alert(width); // 100
alert(height); // 200
alert(item1); // Cake
alert(item2); // Donut

// we pass object to function
let options = {
  title: "My menu",
  items: ["Item1", "Item2"],
};

// ...and it immediately expands it to variables
function showMenu({
  title = "Untitled",
  width = 200,
  height = 100,
  items = [],
}) {
  // title, items – taken from options,
  // width, height – defaults used
  alert(`${title} ${width} ${height}`); // My Menu 200 100
  alert(items); // Item1, Item2
}

showMenu(options);
