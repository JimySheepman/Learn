const File = require("../models/File");
const fs = require("fs");

exports.getAllFiles = async (req, res) => {
  const page = req.query.page || 1;
  const filesPerPage = 3;
  const totalFiles = await File.find().countDocuments();
  const files = await File.find({})
    .sort("-dateCreated")
    .skip((page - 1) * filesPerPage)
    .limit(filesPerPage);

  res.status(200).json({
    status: "succes",
    page,
    filesPerPage,
    totalFiles,
    files,
  });
};

exports.getFile = async (req, res) => {
  const file = await File.findById(req.params.id);
  res.status().json({
    status: "succes",
    file,
  });
};

exports.createFile = async (req, res) => {
  const uploadDir = "public/uploads";

  if (!fs.existsSync(uploadDir)) {
    fs.mkdirSync(uploadDir);
  }

  let uploadeFile = req.files.image;
  let uploadPath = __dirname + "/../public/uploads/" + uploadeFile.name;

  uploadeFile.mv(uploadPath, async () => {
    await File.create({
      ...req.body,
      image: "/uploads/" + uploadeFile.name,
    });
    res.status(200).json({
      status: "succes",
      uploadDir,
      uploadeFile,
      uploadPath,
    });
  });
};

exports.updateFile = async (req, res) => {
  const file = await File.findOne({ _id: req.params.id });
  file.title = req.body.title;
  file.description = req.body.description;
  file.save();

  res.status(200).json({
    status: "succes",
    file,
  });
};

exports.deleteFile = async (req, res) => {
  const file = await File.findOne({ _id: req.params.id });
  let deletedFile = __dirname + "/../public" + file.image;
  fs.unlinkSync(deletedFile);
  await File.findByIdAndRemove(req.params.id);
  res.status(200).json({
    status: "succes",
    file,
    deletedFile,
  });
};
