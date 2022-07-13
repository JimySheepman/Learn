const express = require('express')
const mongoose = require('mongoose')
const mongodb = require('mongodb')
const bodyParser = require('body-parser')
const app = express()
const cors = require('cors')
require('dotenv').config();

// Basic Configuration
const uri = process.env.MONGO_URI;
const listener = app.listen(process.env.PORT || 3000, () => {
  console.log('Your app is listening on port ' + listener.address().port)
});
app.use(bodyParser.urlencoded({ extended: false }));
app.use(cors());
app.use(express.json());

mongoose.connect(uri, {
  useNewUrlParser: true,
  useUnifiedTopology: true
});

// Render View
app.use(express.static('public'));
app.get('/', (req, res) => {
  res.sendFile(process.cwd() + '/views/index.html');
});

// DB Schema
const exerciseSchema = new mongoose.Schema({
  description: {type: String, required: true},
  duration: {type: Number, required: true},
  date: {type: String, required: true}
});
const userSchema = new mongoose.Schema({
  username: {type:String, required: [true, 'username required']},
  log: [exerciseSchema]
});

// DB Model
const Exercise = mongoose.model('Exercise', exerciseSchema);
const User = mongoose.model('User', userSchema);

app.post('/api/users', bodyParser.urlencoded({extended:false}), (req, res) => {
  if(req.body.username != '') {
    const addUser = new User({
      username:req.body.username
    });

  //Add user - return json Schema
    addUser.save((err, savedUser) => {
      if (err) {
        return console.error(err);
      }
      res.json({
        username: savedUser.username,
        _id: savedUser._id
      })
    });
  }
  else {
    res.send("Username is required. Reload form and enter username in appropriate field.");
  }

  app.get('/api/users', (req, res) => {
    User.find({}, (err, array) => {
      if (err) {
        return console.error(err);
      }
      res.json(array);
    })
  })

  app.post('/api/users/:_id/exercises', bodyParser.urlencoded({extended:false}), (req, res) => {
    if(req.params._id != '') {
      let session = new Exercise({
        description: req.body.description,
        duration: parseInt(req.body.duration),
        date: new Date(req.body.date).toDateString()
      })
      if(session.date === '') {
        session.date = new Date().toISOString().substring(0, 10)
        console.log(session.date)
      }
      User.findByIdAndUpdate(req.params._id,{$push : {log: session}},{new: true},(err, updatedUser) => {
        console.log(req.params);
        console.log(updatedUser);
        if(updatedUser != undefined) {
          if(err) {
          return console.log(error);
          }
          console.log('User updated; session added');
          console.log(updatedUser.username);
          res.json({
            _id: req.params._id,
            username: updatedUser.username,
            description: session.description,
            duration: session.duration,
            date: session.date
          })
        }
        else {
          res.send("There is no user with this ID in the database.")
        }
      })
    }
    else {
      res.send("Id is required.");
    }
  })

  app.get('/api/users/:_id/logs', (req, res) => {
    console.log("line 127 inside get(/api/users/:_id/logs)")
    User.findById(req.params._id, (err, result) => {
      if(err) {
        return console.error(error)
      }
      const exerciseCount = result.log.length;
      for(each in result.log) {
        if(result.log[each].date == 'Invalid Date') {
          result.log[each].date = new Date().toDateString();
        }
      }
      if(Object.keys(req.query).length != 0) {
        const from = new Date(req.query['from']);
        const to = new Date(req.query['to']);
        let fromLog = new Date(result.log[0].date);
        let resLog = [];
        if(req.query.hasOwnProperty('limit')) {
          for(each in result.log) {
            if(each < req.query['limit']) {
              resLog.push(result.log[each]);
            }
          }
        }
        if(req.query.hasOwnProperty('from')) {
          for(each in result.log) {
            console.log(result.log.length);
            console.log(each)
            fromLog = new Date(result.log[each].date);
            if(from < fromLog && to > fromLog) {
              console.log("include this log");
              resLog.push(result.log[each]);
            }
          }
        }
        result.log = resLog;
      }
      res.json({"username": result.username, "count": exerciseCount, "_id": result._id, "log": result.log})
    })
  })
});

