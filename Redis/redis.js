const { createClient } = require("redis");

rd_client = createClient({ url: "redis://" + "127.0.0.1" });

rd_client.on("connect", () => console.log("::> Redis Server is Ready"));
rd_client.on("error", (err) => console.log("<:: Redis Client Error", err));

(async () => {
  await rd_client.connect();
})();

exports.rd_client = rd_client;
