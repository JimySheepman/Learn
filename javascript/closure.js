let name = "Ahmet";


function getHello() {
    console.log("Hello, "+ name);    
}
name = "Mehmet";

getHello();


function createEmployee() {
    let name="Mehmet";

    return function(){
        console.log(name)
    };    
}

name = "Zafer";

let work = createEmployee()
work()


function employee(name,surname){

    function fullName() {
        return name + " " + surname; 
    }

    console.log("Hi, " + fullName() );
    console.log("Bye, " + fullName());
}

function User(name){
    this.user = function(){
        console.log(name);
    };
}

let user = new User('Ahmet');
user.employee


function sayacUret() {
    let sayac = 0;
  
    return function() {
      return sayac++;
    };
  }
  
  let sayac = sayacUret();
  
  console.log( sayac() );
  console.log( sayac() );
  console.log( sayac() );


  function sayacUret() {
    let sayac = 0;
  
    return function() {
      return sayac++; 
    };
  }
  
  let sayac1 = sayacUret();
  let sayac2 = sayacUret();
  
  console.log( sayac1() );
  console.log( sayac1() );
  
  console.log( sayac2() );




  function f() {
    let deger = Math.random();
  
    function g() {
      debugger;
    }
  
    return g;
  }
  
  let g = f();
  g();