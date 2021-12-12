// Promise.resolve
function loadCached(url) {
  let cache = loadCached.cache || (loadCached.cache = new Map());

  if (cache.has(url)) {
    return Promise.resolve(cache.get(url)); // (*)
  }

  return fetch(url)
    .then((response) => response.text())
    .then((text) => {
      cache.set(url, text);
      return text;
    });
}

// Promise.reject
let promise = new Promise((resolve, reject) => reject(error));

// Promise.all
let urls = [
  "https://api.github.com/users/iliakan",
  "https://api.github.com/users/remy",
  "https://api.github.com/users/jeresig",
];
let requests = urls.map((url) => fetch(url));
Promise.all(requests).then((responses) =>
  responses.forEach((response) => alert(`${response.url}: ${response.status}`))
);

let names = ["iliakan", "remy", "jeresig"];

requests = names.map((name) => fetch(`https://api.github.com/users/${name}`));

Promise.all(requests)
  .then((responses) => {
    for (let response of responses) {
      alert(`${response.url}: ${response.status}`);
    }

    return responses;
  })
  .then((responses) => Promise.all(responses.map((r) => r.json())))
  .then((users) => users.forEach((user) => alert(user.name)));
Promise.all([
  new Promise((resolve, reject) => setTimeout(() => resolve(1), 1000)),
  new Promise((resolve, reject) =>
    setTimeout(() => reject(new Error("Whoops!")), 2000)
  ),
  new Promise((resolve, reject) => setTimeout(() => resolve(3), 3000)),
]).catch(alert);

// Promise.allSettled
urls = [
  "https://api.github.com/users/iliakan",
  "https://api.github.com/users/remy",
  "https://no-such-url",
];

Promise.allSettled(urls.map((url) => fetch(url))).then((results) => {
  results.forEach((result, num) => {
    if (result.status == "fulfilled") {
      alert(`${urls[num]}: ${result.value.status}`);
    }
    if (result.status == "rejected") {
      alert(`${urls[num]}: ${result.reason}`);
    }
  });
});

// Polyfill
if (!Promise.allSettled) {
  Promise.allSettled = function (promises) {
    return Promise.all(
      promises.map((p) =>
        Promise.resolve(p).then(
          (v) => ({
            state: "fulfilled",
            value: v,
          }),
          (r) => ({
            state: "rejected",
            reason: r,
          })
        )
      )
    );
  };
}
