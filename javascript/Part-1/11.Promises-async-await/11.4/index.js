fetch("/article/promise-chaining/user.json")
  .then((response) => response.json())
  .then((user) => fetch(`https://api.github.com/users/${user.name}`))
  .then((response) => response.json())
  .then(
    (githubUser) =>
      new Promise((resolve, reject) => {
        let img = document.createElement("img");
        img.src = githubUser.avatar_url;
        img.className = "promise-avatar-example";
        document.body.append(img);

        setTimeout(() => {
          img.remove();
          resolve(githubUser);
        }, 3000);
      })
  )
  .catch((error) => alert(error.message));

new Promise((resolve, reject) => {
  resolve("ok");
})
  .then((result) => {
    throw new Error("Whoops!");
  })
  .catch(alert);

new Promise((resolve, reject) => {
  throw new Error("Whoops!");
})
  .catch(function (error) {
    if (error instanceof URIError) {
    } else {
      alert("Can't handle such error");
      throw error;
    }
  })
  .then(function () {})
  .catch((error) => {
    alert(`The unknown error has occurred: ${error}`);
  });

class HttpError extends Error {
  constructor(response) {
    super(`${response.status} for ${response.url}`);
    this.name = "HttpError";
    this.response = response;
  }
}

function loadJson(url) {
  return fetch(url).then((response) => {
    if (response.status == 200) {
      return response.json();
    } else {
      throw new HttpError(response);
    }
  });
}

loadJson("no-such-user.json").catch(alert);

window.addEventListener("unhandledrejection", function (event) {
  alert(event.promise);
  alert(event.reason);
});

new Promise(function () {
  throw new Error("Whoops!");
});
