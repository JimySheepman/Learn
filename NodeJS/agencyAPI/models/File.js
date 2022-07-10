const mongoose = require("mongoose");
const Schema = mongoose.Schema;

const FileSchema = new Schema({
  title: String,
  description: String,
  file: String,
  dateCreated: {
    type: Date,
    default: Date.now,
  },
});

const File = mongoose.model("File", FileSchema);

module.exports = File;
