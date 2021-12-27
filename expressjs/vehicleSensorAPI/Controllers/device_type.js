const { pg_client } = require("../Adapter/db");

const deviceTypeGet = (req, res) => {
  let query =
    "Select device_name, device_description, is_active from devices_type"; // query string.
  pg_client.query(query, (err, result) => {
    if (err) {
      console.log(err);
      res.status(500).send(err);
      res.end();
    } else {
      console.log(result.rows);
      res.send(result.rows);
      res.end();
    }
  });
};

const deviceTypePost = (req, res) => {
  let requestObj = req.body;
  let objArray = Object.values(requestObj);
  console.log(objArray);
  let query =
    "insert into devices_type (device_name, device_description, is_active)" +
    "values($1, $2, $3)";
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

const deviceTypeDelete = (req, res) => {
  let obj = req.body;
  console.log(obj);
  let query = "delete from devices where id='" + obj.id + "'";
  pg_client
    .query(query)
    .then(() => {
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
  deviceTypeGet,
  deviceTypePost,
  deviceTypeDelete,
};
