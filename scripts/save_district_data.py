from pymongo import MongoClient
import json

## One use script, use it to save all the districts data on the utility database
## Only run it using python <script>

MONGO_URI = "mongodb://localhost:27017/"
DATABASE_NAME = "harmony_data"
COLLECTION_NAME = "districts"

paths = [
    {
        "path": "./scripts/data/city_bogota_ne.json",
        "city": "BOGOTA",
    },
    {
        "path": "./scripts/data/city_mexico_ne.json",
        "city": "MEXICO_CITY",
    },
    {
        "path": "./scripts/data/city_miami_ne.json",
        "city": "MIAMI",
    },
]

districts_data = []

for district in paths:
    with open(district["path"], "r", encoding="utf-8") as file:
        try:
            data = json.load(file)
            
            for j in data:
                j["city"] = district["city"]
                districts_data.append(j)
        except json.JSONDecodeError as e:
            print(f"Error decoding JSON: {e}")

try:
    client = MongoClient(MONGO_URI)
    
    print("Connected")

    db = client[DATABASE_NAME]
    collection = db[COLLECTION_NAME]

    result = collection.insert_many(districts_data)
    
    print(f"Data inserted with ID: {result.inserted_ids}")

except Exception as e:
    print("Error connecting to MongoDB:", e)

finally:
    client.close()
    print("Connection closed")