const { pg_client } = require("../Adapter/db");

const deviceGet = (req, res) => {
  let query =
    "Select  devices.id, vehicles.vehicle_plate, devices.device_name, devices.is_online, devices.is_active" +
    " from devices  inner join vehicles on devices.vehicle_id=vehicles.id order by devices.id asc";
  pg_client.query(query, (err, result) => {
    if (err) {
      console.log(err);
      res.status(500).send(err);
      res.end();
    } else {
      console.log(result.rows);
      res.status(200).send(result.rows);
      res.end();
    }
  });
};

const devicePost = (req, res) => {
  let requestObj = req.body;
  let objArray = Object.values(requestObj);
  console.log(objArray);
  let query =
    "insert into devices (vehicle_id, device_type_id, device_name, is_online, is_active) " +
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

const devicePatch = (req, res) => {
  let oldObj = req.body[0];
  let newObj = req.body[1];
  let valArray = Object.values(newObj);
  valArray.push(oldObj.vehicle_id);
  console.log(valArray);
  let query =
    "update devices set vehicle_id=$1, device_type_id=$2, device_name=$3, is_online=$4, is_active=$5 " +
    "where vehicle_id=$6;";
  pg_client
    .query(query, valArray)
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

const deviceDelete = (req, res) => {
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
      res.end();
    });
};

module.exports = {
  deviceGet,
  devicePost,
  devicePatch,
  deviceDelete,
};
