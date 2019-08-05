from pymongo import MongoClient

DATABASE = MongoClient()['mydb']
DEBUG = True
client = MongoClient('localhost', 27017)
