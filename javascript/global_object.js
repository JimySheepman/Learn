alert("Hello");
window.alert("Hello1")


var hi = "hello"

function getHi(){
    alert(hi)
}

alert(window.hi );
alert(window.getHi);

window.test = 5;

alert(test)
if (window.XMLHttpRequest) {
    alert('XMLHttpRequest tanımlı!')
  }


  function f() {
      alert(this);
  }

  f();

  if (!window.Promise) {
    alert("Your browser is really old!");
  }

  if (!window.Promise) {
    window.Promise = s
  }