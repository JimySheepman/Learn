const express = require("express");
const Joi = require("@hapi/joi");
const app = express();

app.use(express.json());

const student = [
  { id: 1, name: "Ahmet" },
  { id: 2, name: "Ali" },
  { id: 3, name: "Veli" },
  { id: 4, name: "Ayşe" },
  { id: 5, name: "Fatma" },
];
app.get("/", (req, res) => {
  res.send("index page.");
});

app.get("/api/products/:id", (req, res) => {
  res.send(req.params.id);
});
app.get("/api/shipments/:year/:month", (req, res) => {
  res.send(req.params);
});
app.get("/api/student", (req, res) => {
  res.send(student);
});

app.get("/api/student/:id", (req, res) => {
  const stud = student.find((std) => std.id === parseInt(req.params.id));
  if (!stud) res.status(404).send("Student is not founded...");
  else {
    res.send(stud);
  }
});

// sent
app.post("/api/student", (req, res) => {
  const schema = Joi.object({
    name_con: Joi.string().min(3).required(),
  });
  const result = schema.validate({ name_con: req.body.name });
  console.log(result);
  if (result.error) {
    res.status(400).send(result.error.details[0].message);
    return;
  }
  const stu = {
    id: student.length + 1,
    name: req.body.name,
  };
  student.push(stu);
  res.send(student);
});

app.put("/api/student/:id", (req, res) => {
  // id check
  const stud = student.find((std) => std.id === parseInt(req.params.id));
  if (!stud) res.status(404).send("Student is not founded...");
  // validation
  const schema = Joi.object({
    name_con: Joi.string().min(3).required(),
  });
  const result = schema.validate({ name_con: req.body.name });
  console.log(result);
  if (result.error) {
    res.status(400).send(result.error.details[0].message);
    return;
  }
  // update
  stud.name = req.body.name;
  res.send(stud);
});

app.delete("/api/student/:id", (req, res) => {
  // id check
  const stud = student.find((std) => std.id === parseInt(req.params.id));
  if (!stud) res.status(404).send("Student is not founded...");
  // delete
  const index = student.indexOf(stud);
  student.splice(index, 1);
  res.send(stud);
});

const port = process.env.PORT || 3000;

app.listen(port, () => {
  console.log(`App start ${port} port.`);
});
