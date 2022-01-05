const express = require("express");
const mongoose = require("mongoose");
const fileUpload = require("express-fileupload");
const methodOverride = require("method-override");
const fileController = require("./controllers/fileController");
const pageController = require("./controllers/pageController");

// create app
const app = express();

// connect DB
mongoose
  .connect("mongodb://localhost:27017/agency")
  .then(() => {
    console.log("DB CONNECTED!");
  })
  .catch((err) => {
    console.log(err);
  });

// middlewares
app.use(express.urlencoded({ extended: true }));
app.use(express.json());
app.use(fileUpload());
app.use(
  methodOverride("_method", {
    methods: ["POST", "GET"],
  })
);

//ROUTES
app.get("/", fileController.getAllFiles);
app.get("/file/:id", fileController.getFile);
app.post("/file", fileController.createFile);
app.put("/file/:id", fileController.updateFile);
app.delete("/file/:id", fileController.deleteFile);

app.get("/about", pageController.getAboutPage);
app.get("/add", pageController.getAddPage);
app.get("/file/edit/:id", pageController.getEditPage);

const port = process.env.PORT || 3000;
app.listen(port, () => {
  console.log(`Server start to port: ${port}`);
});
