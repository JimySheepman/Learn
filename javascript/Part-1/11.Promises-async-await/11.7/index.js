setTimeout(() => alert("timeout"));

Promise.resolve().then(() => alert("promise"));

alert("code");

Promise.resolve()
  .then(() => {
    setTimeout(() => alert("timeout"), 0);
  })
  .then(() => {
    alert("promise");
  });

let promise = Promise.reject(new Error("Promise Failed!"));
promise.catch((err) => alert("caught"));
window.addEventListener("unhandledrejection", (event) => alert(event.reason));
