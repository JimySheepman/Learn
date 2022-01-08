const express = require("express");
const mongoose = require("mongoose");
const pageController = require("./controllers/pageController");

// create app
const app = express();

// connect DB
mongoose
  .connect("mongodb://localhost:27017/freelancer")
  .then(() => {
    console.log("DB CONNECTED!");
  })
  .catch((err) => {
    console.log(err);
  });

// middlewares
app.use(express.urlencoded({ extended: true }));
app.use(express.json());

//ROUTES
app.get("/", pageController.getPage);
app.post("/add", pageController.getAddPage);
app.put("/edit", pageController.getEditPage);
app.post("/contact", pageController.sendEmail);

const port = process.env.PORT || 3000;
app.listen(port, () => {
  console.log(`Server start to port: ${port}`);
});
