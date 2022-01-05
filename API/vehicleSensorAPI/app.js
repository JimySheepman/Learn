const { routers } = require("./Routers/routers");
const express = require("express");
const app = express();

app.use(express.json());
app.use(routers);

app.get("/api", (req, res) => {
  console.log(req.body);
  res.status(200).send("api services");
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => `Sunucu ${PORT} portunda çalışmaktadır!`);
