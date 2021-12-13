function sayHi() {
  alert("Hello");
}
window.sayHi();
alert(window.innerHeight);

document.body.style.background = "red"

setTimeout(()=>document.body.style.background="",1000);

console.log(location.href)
if(confirm("Go to Wikipedia?")){
    location.href="https://wikipedia.org"
}