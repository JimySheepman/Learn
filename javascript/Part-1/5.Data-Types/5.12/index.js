let user = {
  name: "John",
  age: 30,

  toString() {
    return `name:"${this.name}",age:${this.age}`;
  },
};
console.log(user);

let student = {
  name: "John",
  age: 30,
  isAdmin: false,
  courses: ["html", "css", "js"],
  wife: null,
};

let json = JSON.stringify(student);

console.log(typeof json); // we've got a string!

console.log(json);
console.log(JSON.stringify(1));
console.log(JSON.stringify("test"));
console.log(JSON.stringify(true));
console.log(JSON.stringify([1, 2, 3]));

user = {
  sayHi() {
    console.log("hello");
  },
  [Symbol("id")]: 123,
  something: undefined,
};
console.log(JSON.stringify(user));

let meetup = {
  title: "Conference",
  room: {
    number: 123,
    participants: ["john", "ann"],
  },
};
console.log(JSON.stringify(meetup));

let room = {
  number: 23,
};

meetup = {
  title: "Conference",
  participants: [{ name: "John" }, { name: "Alice" }],
  place: room, // meetup references room
};

room.occupiedBy = meetup; // room references meetup

console.log(JSON.stringify(meetup, ["title", "participants"]));

room = {
  number: 23,
};

meetup = {
  title: "Conference",
  participants: [{ name: "John" }, { name: "Alice" }],
  place: room, // meetup references room
};

room.occupiedBy = meetup; // room references meetup

console.log(
  JSON.stringify(meetup, function replacer(key, value) {
    console.log(`${key}: ${value}`);
    return key == "occupiedBy" ? undefined : value;
  })
);

user = {
  name: "John",
  age: 25,
  roles: {
    isAdmin: false,
    isEditor: true,
  },
};

console.log(JSON.stringify(user, null, 2));

room = {
  number: 23,
};

meetup = {
  title: "Conference",
  date: new Date(Date.UTC(2017, 0, 1)),
  room,
};
console.log(JSON.stringify(meetup));

let numbers = "[0, 1, 2, 3]";
numbers = JSON.parse(numbers);
console.log(numbers[1]);

let userData =
  '{ "name": "John", "age": 35, "isAdmin": false, "friends": [0,1,2,3] }';
user = JSON.parse(userData);
console.log(user.friends[1]);

str = '{"title":"Conference","date":"2017-11-30T12:00:00.000Z"}';

meetup = JSON.parse(str, function (key, value) {
  if (key == "date") return new Date(value);
  return value;
});

console.log(meetup.date.getDate());
