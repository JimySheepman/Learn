function sayHi(){
    console.log('Hello');
}


setTimeout(sayHi,1000);

setTimeout(() => {
    console.log("2");
}, 4000);


setTimeout(() =>  console.log('hi'), 1000);



let timerId = setTimeout(() => console.log("never happens"), 1000);
console.log(timerId);

clearTimeout(timerId);
console.log(timerId);


let timerId_1 = setInterval(() => console.log('tick'), 2000);


setTimeout(() => { clearInterval(timerId_1); console.log('stop'); }, 5000);


setTimeout(() => console.log("World"));

console.log("Hello");


let start = Date.now();
let times = [];

setTimeout(function run() {
  times.push(Date.now() - start);

  if (start + 100 < Date.now()) console.log(times); 
  else setTimeout(run); 
});