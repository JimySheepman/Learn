{
  let massage = "Hello";
  console.log(massage);
}
// console.log(massage); // Scope error
{
  let message = "Hello";
  console.log(message);
}

{
  let message = "Goodbye";
  console.log(message);
}

function sayHiBye(firstName, lastName) {
  function getFullName() {
    return firstName + " " + lastName;
  }

  console.log("Hello, " + getFullName());
  console.log("Bye, " + getFullName());
}
