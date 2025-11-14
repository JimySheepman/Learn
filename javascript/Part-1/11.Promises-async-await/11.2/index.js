let promise = new Promise(function (resolve, reject) {
  setTimeout(() => reject(new Error("Whoops!")), 1000);
})
  .finally(() => alert("Promise ready"))
  .then((result) => alert(result));

promise.then(
  (result) => alert(result),
  (error) => alert(error)
);
promise.catch(alert);
