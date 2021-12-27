const { pg_client } = require("../Adapter/db");

const vehicleGet = (req, res) => {
  let query =
    "Select id, vehicle_plate, current_status, is_active from vehicles";
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

const vehiclePost = (req, res) => {
  let requestObj = req.body;
  let objArray = Object.values(requestObj);
  console.log(objArray);
  let query =
    "insert into vehicles (vehicle_plate, current_status, is_active) " +
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

const vehiclePatch = (req, res) => {
  let oldObj = req.body[0];
  let newObj = req.body[1];
  let valArray = Object.values(newObj);
  valArray.push(oldObj.vehicle_plate);

  console.log(valArray);
  let query =
    "update vehicles set vehicle_plate=$1, current_status=$2, is_active=$3 " +
    "where vehicle_plate=$4;";
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

const vehicleDelete = (req, res) => {
  let obj = req.body;
  console.log(obj);
  let queryToDo =
    "delete from vehicles where vehicle_plate='" + obj.vehicle_plate + "'";
  pg_client
    .query(queryToDo)
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
  vehicleGet,
  vehiclePost,
  vehicleDelete,
  vehiclePatch,
};
