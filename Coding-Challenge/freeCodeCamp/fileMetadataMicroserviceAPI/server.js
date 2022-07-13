const express = require('express')
const bodyParser = require('body-parser')
const app = express()
const cors = require('cors')
var multer  = require('multer');

const listener = app.listen(process.env.PORT || 3000, () => {
  console.log('Your app is listening on port ' + listener.address().port)
});
app.use(bodyParser.urlencoded({ extended: false }));
app.use(cors());
app.use(express.json());

// Render View
app.use('/public', express.static(process.cwd() + '/public'));
app.get('/', function(req, res) {
  res.sendFile(process.cwd() + '/views/index.html');
});

app.post("/api/fileanalyse",
  multer({ dest: 'uploadedFiles/' }).single('upfile'),
  (req, res) => {
    if(res == 0) {
      console.log("Error")
    } else {
      res.json({
      name: req.file.originalname,
      type: req.file.mimetype,
      size: req.file.size + ' ' + 'bytes',
    });
    }

});

app.use((req, res, next) => {
  return next({status: 404, message: 'not found'})
})
