let start = Date.now();
let times = [];

setTimeout(function run() {
  times.push(Date.now() - start); // remember delay from the previous call

  if (start + 100 < Date.now()) console.log(times); // show the delays after 100ms
  else setTimeout(run); // else re-schedule
});

function sayHi(phrase, who) {
    console.log( phrase + ', ' + who );
  }
  
  setTimeout(sayHi, 1000, "Hello", "John"); 