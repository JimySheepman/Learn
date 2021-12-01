let user;
console.log(user ?? "Anonymos");

let user2 = "Jhon";
console.log(user ?? "Anonymos");

let firstName = null,
  lastName = null,
  nickName = "merlins";

console.log(firstName ?? lastName ?? nickName ?? "Anonymos");
console.log(firstName || lastName || nickName || "Anonymos");

let height = 0;
console.log(height || 100);
console.log(height ?? 100);

height = null;
let width = null;
let area = (height ?? 100) * (width ?? 50);
console.log(area);

let area2 = height ?? 100 * width ?? 50;
console.log(area2);

let area3 = height ?? 100 * width ?? 50;
console.log(area3);

let x = (1 && 2) ?? 3;
console.log(x);
