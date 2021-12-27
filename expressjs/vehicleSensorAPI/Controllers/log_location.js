const { pg_client } = require("../Adapter/db");

const logLocationGet = (req, res) => {
  let query =
    "Select vehicle_id, device_id, latitude, longtitude, created_at from log_location";
  pg_client.query(query, (err, result) => {
    if (err) {
      console.log(err);
      res.status(500).send(err);
      res.end();
    }
    console.log(result.rows);

    res.send(result.rows);
    res.end();
  });
};

const logLocationPost = (req, res) => {
  let requestObj = req.body;
  let objArray = Object.values(requestObj);
  console.log(objArray);
  let query =
    "insert into log_location (vehicle_id, device_id, latitude, longtitude, created_at)" +
    "values($1, $2, $3, $4, $5)";
  pg_client
    .query(query, objArray)
    .then((result) => {
      res.status(200).send("OK");
      res.end();
    })
    .catch((err) => {
      console.log(err);
      res.status(500).send(err);
      res.end();
    });
};

module.exports = {
  logLocationGet,
  logLocationPost,
};
