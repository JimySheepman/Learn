let single = "sigle-quoted";
let double = "double-quoted";
let backticks = `backticks`;
console.log(single, double, backticks);

function sum(a, b) {
  return a + b;
}
console.log(`1 + 2 = ${sum(1, 2)}.`);
let guestList = `Guests:
* John
* Pete
* Mary
`;
console.log(guestList);
console.log("Guests:\n * John\n * Pete\n * Mary");

let str1 = "Hello\nWorld";
let str2 = `Hello
World`;
console.log(str1 == str2);

console.log(
  "Lorem,\n ipsum dolor sit\r amet consectetur adipisicing\\ elit.\t Minus \b omnis\f voluptatibus\v neque velit, corrupti sed quasi quidem quisquam autem similique? Numquam minus labore rem soluta blanditiis molestias neque similique commodi!"
);

console.log("\u00A9");
console.log("\u{20331}");
console.log("\u{1F60D}");
console.log("I'm the walrus!");
console.log(`The backslash: \\`);
console.log(`my\n`.length);

let str3 = `Hello`;
console.log(str3[0]);
console.log(str3.charAt(0));
console.log(str3[str3.length - 1]);

for (let char of "Hello") {
  console.log(char);
}

console.log("Interface".toLowerCase());
console.log("Interface".toUpperCase());

let str4 = "Widget with id";
console.log(str4.indexOf("Widget"));
console.log(str4.indexOf("id"));
console.log(str4.indexOf("id", 2));

/* let str5 = "As sly as a fox, as strong as an ox";
let target = "as";

let pos = 0;
while (true) {
  let foundPos = str5.indexOf(target, pos);
  if (foundPos == 1) break;
  console.log(`Found at  ${foundPos}`);
  pos = foundPos + 1;
  break;
} */

console.log(~2);
console.log(~1);
console.log(~0);
console.log(~-1);

console.log("Widget with id".includes("Widget"));
console.log("Hello".includes("Bye"));

let str6 = "stringify";
console.log(str6.slice(0, 5));
console.log(str6.slice(0, 1));
console.log(str6.slice(2));
console.log(str6.slice(-4, -1));
console.log(str6.substring(6, 2));

console.log("a" < "Z");
console.log("z".codePointAt(0));
console.log(String.fromCodePoint(122));

let str7 = "";
for (let i = 65; i < 220; i++) {
  str7 += String.fromCodePoint(i);
}
console.log(str7);

let s1 = "S\u0307\u0323";
let s2 = "S\u0323\u0307";

console.log(`s1: ${s1}, s2: ${s2}`);

console.log(s1 == s2);
