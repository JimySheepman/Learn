const { pg_client } = require("../Adapter/db");

const logTemperatureGet = (req, res) => {
  let query =
    "Select vehicle_id, device_id, read_data, created_at from log_temperature";
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

const logTemperaturePost = (req, res) => {
  let requestObj = req.body;
  let objArray = Object.values(requestObj);
  console.log(objArray);
  let query =
    "insert into log_temperature (vehicle_id, device_id, read_data, created_at)" +
    "values($1, $2, $3, $4)";

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
  logTemperatureGet,
  logTemperaturePost,
};
