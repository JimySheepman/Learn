const mongoose = require("mongoose");
const Schema = mongoose.Schema;

const InfoSchema = new Schema({
  title: String,
  description: String,
  icon: String,
  dateCreated: {
    type: Date,
    default: Date.now,
  },
});

const Info = mongoose.model("Info", InfoSchema);

module.exports = Info;
