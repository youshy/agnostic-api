from config import client
from app import app
from bson.json_util import dumps
from flask import request, jsonify
import json
import ast
import imp

helper_module = imp.load_source('*', './app/helpers.py')

db = client.restfulapi
collection = db.users

@app.route("/users", methods=['GET'])
def get_all_users():
    try:
        query_params = helper_module.parse_query_params(request.query_string)
        if query_params:
            query = {k: int(v) if isinstance(v, str) and v.isdigit() else v for k, v in query_params.items()}

            records_fetched = collection.find(query)

            if records_retched.count() > 0:
                return dumps(records_fetched)
            else:
                return "", 404

        else:
            if collection.find().count > 0:
                return dumps(collection.find())
            else:
                return jsonify([])
    except:
        return "", 500
