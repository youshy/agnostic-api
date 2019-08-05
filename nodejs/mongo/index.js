const express = require("express")
const mongoose = require("mongoose")
const bodyParser = require("body-parser")

const app = express()

const mongoURI = "mongodb://localhost:27017/mydb"

const db = mongoose
  .connect(
    mongoURI,
    { useNewUrlParser: true }
  )
  .then(() => console.log("MongoDB connected!"))
  .catch(err => console.log("Problem connecting to the db", err))

require("./model")
const User = mongoose.model("users")

app.use(bodyParser.json())
app.use(bodyParser.urlencoded({
  extended: false
}))

app.get('/users', async (req, res) => {
  const users = await User.find()

  res.status(200).json(users)
})

app.get('/user', async (req, res) => {
  const user = await User.findOne({
    firstname: req.query.firstname,
    lastname: req.query.lastname
  })

  res.status(200).json(user)
})

app.post('/user', async (req, res) => {
  const {firstname, lastname, age} = req.body

  const newUser = {
    firstname,
    lastname,
    age
  }

  await new User(newUser).save()

  res.status(200).send("new user has been saved")
})

app.put('/user', (req, res) => {
  User.findOne({
    firstname: req.query.firstname,
    lastname: req.query.lastname
  }).then(user => {
    user.firstname = req.body.firstname
    user.lastname = req.body.lastname
    user.age = req.body.age

    user.save().then(() => {
      res.status(200).send("user data has been changed")
    })
  })
})

app.delete('/user', (req, res) => {
  User.deleteOne({
    firstname: req.query.firstname,
    lastname: req.query.lastname
  }).then(() => {
    res.status(200).send("user has been removed from the db")
  })
})

const PORT = 3000

app.listen(PORT, () => {
  console.log(`Server is listening on port ${PORT}`)
})
