from pymongo import MongoClient
import json

MONGO_URI = "mongodb://localhost:27017/"
DATABASE_NAME = "harmony_data"
COLLECTION_NAME = "districts"

# IMPORTANT Only for the test purposes, Im using a regular districts definition but normally it must be irregular polygons
# IMPORTANT It is not the best aproach so it cant be used on production


def isOnDistrict():
    return True


with open("./scripts/data/data.json", "r", encoding="utf-8") as file:
    try:
        data = json.load(file)
    except json.JSONDecodeError as e:
        print(f"Error decoding JSON: {e}")
