const { vehicleGet, vehiclePost,vehiclePatch,vehicleDelete } = require("../Controllers/vehicle");
const { deviceGet, devicePost,devicePatch,deviceDelete } = require("../Controllers/device");
const { deviceTypeGet, deviceTypePost,deviceTypeDelete} =require("../Controllers/device_type");
const {logTemperatureGet, logTemperaturePost}=require("../Controllers/log_temperature");
const {logLocationGet, logLocationPost}=require("../Controllers/log_location");
const express = require('express');
const router = express.Router();


router.route("/vehicle")
    .get(vehicleGet)
    .post(vehiclePost)
    .patch(vehiclePatch)
    .delete(vehicleDelete);

router.route("/device")
    .get(deviceGet)
    .post(devicePost)
    .patch(devicePatch)
    .delete(deviceDelete);

router.route("/device_type")
    .get(deviceTypeGet)
    .post(deviceTypePost)
    .delete(deviceTypeDelete);

router.route("/log_temperature")
    .get(logTemperatureGet)
    .post(logTemperaturePost);

router.route("/log_location")
    .get(logLocationGet)
    .post(logLocationPost);

module.exports = {
    router,
};