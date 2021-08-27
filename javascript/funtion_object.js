function getHi() {
    console.log("Hello");    
}

console.log(getHi.name)

let getHello = function() {
    console.log("hi");
}

console.log(getHello.name);

function f1(a) {}
function f2(a, b) {}
function many(a, b, ...more) {}

console.log(f1.length); // 1
console.log(f2.length); // 2
console.log(many.length); // 2

function sor(soru, ...isleyiciler) {
    let dogruMu = confirm(soru);
  
    for(let isleyici of isleyiciler) {
      if (isleyici.length == 0) {
        if (dogruMu) isleyici();
      } else {
        isleyici(dogruMu);
      }
    }
  
  }

  sor("Soru?", () => alert('Evet dedin'), sonuc => alert(sonuc));
  
  function makeCounter() {
    // let count = 0 yazmak yerine

  function counter() {
    return counter.count++;
  };

  counter.count = 0;

  return counter;
}

let counter = makeCounter();
console.log( counter() ); // 0
console.log( counter() ); // 1



  let selamVer = function func(kim) {
    console.log(`Selam, ${kim}`);
  };
  
  selamVer("Ahmet");


  let sum = new Function('arg1', 'arg2', 'return arg1 + arg2');

console.log( sum(1, 2) );


function FonkAl() {
    let deger = "test";
  
    let func = function() { alert(deger); };
  
    return func;
  }
  
  FonkAl()();


  let group = {
    title: "Our Group",
    students: ["John", "Pete", "Alice"],
  
    showList() {
      this.students.forEach(
        student => alert(this.title + ': ' + student)
      );
    }
  };
  
  group.showList();



  function defer(f, ms) {
    return function() {
      setTimeout(() => f.apply(this, arguments), ms)
    };
  }
  
  function sayHi(who) {
    alert('Hello, ' + who);
  }
  
  let sayHiDeferred = defer(sayHi, 2000);
  sayHiDeferred("John"); 