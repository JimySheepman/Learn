const File = require("../models/File");

exports.getAboutPage = (req, res) => {
  res.status(200).json({
    status: "succes",
    req,
  });
};

exports.getAddPage = (req, res) => {
  res.status(200).json({
    status: "succes",
    req,
  });
};

exports.getEditPage = async (req, res) => {
  const photo = await Photo.findOne({ _id: req.params.id });
  res.status(200).json({
    status: "succes",
    msg: "edit",
    photo,
  });
};
