const redis = require("redis");
const express = require("express");
const fetch = require("node-fetch");

const USER_NAME = "username";
const PORT = 3000;
const REDIS_PORT = 6379;

const client = redis.createClient(REDIS_PORT);
const app = express();

function formatOutput(username, numOfRepos) {
  return `${username} has ${numOfRepos} repos`;
}

async function getRepo(req, res) {
  try {
    const username = req.params[USER_NAME];
    const response = await fetch(`https://api.github.com/user/${username}`);
    const { public_repos } = await response.json();
    client.set(username, public_repos);
    res.send(formatOutput(username, public_repos));
  } catch (error) {
    console.log(error);
    res.status(404);
  }
}

// cashe middleware
function cache(req, res, next) {
  const username = req.params[USER_NAME];

  client.get(username, (err, data) => {
    if (err) throw err;
    if (data != null) {
      res.send(formatOutput(username, data));
    } else {
      next();
    }
  });
}

app.listen(PORT, () => {
  console.log(`App listening on port ${PORT}`);
});
