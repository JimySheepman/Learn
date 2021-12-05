let user = {
  name: "john",
  age: 30,
};
user.sayHi = function () {
  console.log("hello world!");
};
user.sayHi();

let user_2 = {
  // ... some code
};
function sayHii() {
  console.log("hellooo");
}

user_2.sayHii = sayHii;
user_2.sayHii();

user_3 = {
  sayHi: function () {
    console.log("hello");
  },
};
user_3 = {
  sayHi() {
    // same as "sayHi: function(){...}"
    alert("Hello");
  },
};

let user_4 = { name: "John" };
let admin = { name: "Admin" };

function sayHi() {
  alert(this.name);
}

// use the same function in two objects
user_4.f = sayHi;
admin.f = sayHi;

// these calls have different this
// "this" inside the function is the object "before the dot"
user_4.f(); // John  (this == user)
admin.f(); // Admin  (this == admin)

admin["f"](); // Admin (dot or square brackets access the method – doesn't matter)

let user_5 = {
  firstName: "Ilya",
  sayHi() {
    let arrow = () => alert(this.firstName);
    arrow();
  },
};

user_5.sayHi(); // Ilya
