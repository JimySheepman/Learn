let a = 1;

function f() {
  let a = 2;
  eval("alert(a)");
}

f();

let x = 1;
{
  let x = 5;
  window.eval("alert(x)");
}
