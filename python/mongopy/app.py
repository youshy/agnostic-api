from flask import Flask, jsonify, request, Response
from flask_pymongo import PyMongo
import json
from bson.objectid import ObjectId

app = Flask(__name__)

#MongoDB setup
app.config['MONGO_URI'] = 'mongodb://localhost:27017/mydb'
mongo = PyMongo(app)

class JSONEncoder(json.JSONEncoder):
    ''' extend json-encoder class'''

    def default(self, o):
        if isinstance(o, ObjectId):
            return str(o)
        if isinstance(o, datetime.datetime):
            return str(o)
        return json.JSONEncoder.default(self, o)

app.json_encoder = JSONEncoder

#Get all users
@app.route('/users')
def get_users():
    resp = [doc for doc in mongo.db.users.find()]
    return jsonify(resp)

#Get particular user
@app.route('/user')
def get_one_user():
    firstname = request.args.get('firstname')
    lastname = request.args.get('lastname')
    user = mongo.db.users.find_one({
        'firstname': firstname,
        'lastname': lastname
    })
    return jsonify(user)

app.run(port=3000)
