let listOfNumber = [2, 3, 4, 5, 7];
console.log(listOfNumber[2]);
for (let i = 0; i < listOfNumber.length; i++) {
  console.log(listOfNumber[i]);
}

let doh = "Doh";
console.log(typeof doh.toUpperCase);
console.log(doh.toUpperCase());

let sequence = [1, 2, 3];

sequence.push(4);
sequence.push(5);
console.log(sequence);
sequence.pop();
console.log(sequence);

let day1 = {
  squirrel: false,
  events: ["work", "touched tree", "pizza", "running"],
};

console.log(day1.squirrel);
console.log(day1.wolf);
day1.wolf = false;
console.log(day1.wolf);

let anObject = { left: 1, right: 2 };
console.log(anObject.left);
delete anObject.left;
console.log(anObject.left);
console.log("left" in anObject);
console.log("right" in anObject);

console.log(Object.keys({ x: 0, y: 0, z: 2 }));

let objectA = {
  a: 1,
  b: 2,
};
Object.assign(objectA, { b: 3, c: 4 });
console.log(objectA);

let object1 = { value: 10 };
let object2 = object1;
let object3 = { value: 10 };

console.log(object1 == object2);
console.log(object1 == object3);
object1.value = 15;
console.log(object2.value);
console.log(object3.value);

/* const score = {visitors: 0, home: 0};
// This is okay
score.visitors = 1;
// This isn't allowed
score = {visitors: 1, home: 1};
 */
let journal = [];

function addEntry(events, squirrel) {
  journal.push({ events, squirrel });
}

addEntry(["work", "touched tree", "pizza", "running", "television"], false);
addEntry(
  [
    "work",
    "ice cream",
    "cauliflower",
    "lasagna",
    "touched tree",
    "brushed teeth",
  ],
  false
);
addEntry(["weekend", "cycling", "break", "peanuts", "beer"], true);

function phi(table) {
  return (
    (table[3] * table[0] - table[2] * table[1]) /
    Math.sqrt(
      (table[2] + table[3]) *
        (table[0] + table[1]) *
        (table[1] + table[3]) *
        (table[0] + table[2])
    )
  );
}

console.log(phi([76, 9, 4, 1]));

function tableFor(event, journal) {
  let table = [0, 0, 0, 0];
  for (let i = 0; i < journal.length; i++) {
    let entry = journal[i],
      index = 0;
    if (entry.events.includes(event)) index += 1;
    if (entry.squirrel) index += 2;
    table[index] += 1;
  }
  return table;
}

console.log(tableFor("pizza", journal));

for (let i = 0; i < journal.length; i++) {
  let entry = journal[i];
}

for (let entry of journal) {
  console.log(`${entry.events.length} events.`);
}

function journalEvents(journal) {
  let events = [];
  for (let entry of journal) {
    for (let event of entry.events) {
      if (!events.includes(event)) {
        events.push(event);
      }
    }
  }
  return events;
}

console.log(journalEvents(journal));

for (let event of journalEvents(journal)) {
  console.log(event + ":", phi(tableFor(event, journal)));
}

let todolist = [];
function remember(task) {
  todolist.push(task);
}
function getTask() {
  return todolist.shift();
}
function rememberUrgently(task) {
  todolist.unshift(task);
}

remember("groceries");
getTask();
console.log(todolist);

function remove(array, index) {
  return array.slice(0, index).concat(array.slice(index + 1));
}
console.log(remove(["a", "b", "c", "d", "e"], 2));


console.log("coconuts".slice(4, 7));
console.log("coconut".indexOf("u"));
console.log("one two three".indexOf("ee"));
console.log("  okay \n ".trim());
console.log(String(6).padStart(3, "0"));
let sentence = "Secretarybirds specialize in stomping";
let words = sentence.split(" ");
console.log(words);
console.log(words.join(". "));
console.log("LA".repeat(3));
let string = "abc";
console.log(string.length);
console.log(string[1]);


function max(...numbers){
  let result = -Infinity;
  for (let number of numbers){
    if(number > result){
      result = number;
    }
  }
  return result;
}

console.log(max(4,2,1,5,6,78,8,1));

let numbers = [12,5,6,123,6,7,1234,1];
console.log(max(...numbers));


function randomPointOnCircle(radius) {
  let angle = Math.random() * 2 *Math.PI;
  return{
    x: radius * Math.cos(angle),
    y : radius * Math.sin(angle)};
}

console.log(randomPointOnCircle(4));
console.log(Math.random());
console.log(Math.random());
console.log(Math.random());
console.log(Math.floor(Math.random()*10));


function phi(table) {
  return (table[3] * table[0] - table[2] * table[1]) /
    Math.sqrt((table[2] + table[3]) *
              (table[0] + table[1]) *
              (table[1] + table[3]) *
              (table[0] + table[2]));
}


let _string = JSON.stringify({squirrel: false,events: ["weekend"]});
console.log(_string);
console.log(JSON.parse(_string).events);