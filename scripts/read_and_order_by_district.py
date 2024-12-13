from pymongo import MongoClient
import json
import math
from unidecode import unidecode

MONGO_URI = "mongodb://localhost:27017/"
DATABASE_NAME = "harmony_data"
COLLECTION_NAME = "sellers"

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

districts_data = {}

# IMPORTANT Only for the test purposes, Im using a regular districts definition but normally it must be irregular polygons
# IMPORTANT It is not the best aproach so it cant be used on production


def degrees_to_radians(degrees: float) -> float:
    return degrees * (math.pi / 180)


def distance_between_points(lat1: float, lon1: float, lat2: float, lon2: float) -> float:
    earth_radius_km = 6371

    d_lat = degrees_to_radians(lat2 - lat1)
    d_lon = degrees_to_radians(lon2 - lon1)

    r_lat1 = degrees_to_radians(lat1)
    r_lat2 = degrees_to_radians(lat2)

    a = (math.sin(d_lat / 2) ** 2 + math.sin(d_lon / 2) ** 2 * math.cos(r_lat1) * math.cos(r_lat2))
    c = 2 * math.atan2(math.sqrt(a), math.sqrt(1 - a))

    distance = earth_radius_km * c
    return distance


for district in paths:
    with open(district["path"], "r", encoding="utf-8") as file:
        try:
            data = json.load(file)

            for j in data:
                if district["city"] not in districts_data:
                    districts_data[district["city"]] = []
                districts_data[district["city"]].append(j)
        except json.JSONDecodeError as e:
            print(f"Error decoding JSON: {e}")

with open("./scripts/data/data.json", "r", encoding="utf-8") as file:
    try:
        data = json.load(file)

        sellers_ordered = []
        mexico_sellers = []
        bogota_sellers = []
        miami_sellers = []

        for seller in data:
            if seller["city"] == "Mexico City":
                mexico_sellers.append(seller)
            elif seller["city"] == "Bogotá":
                bogota_sellers.append(seller)
            elif seller["city"] == "Miami":
                miami_sellers.append(seller)

        formated = [
            {
                "sellers": mexico_sellers,
                "city": "MEXICO_CITY"
            },
            {
                "sellers": bogota_sellers,
                "city": "BOGOTA"
            },
            {
                "sellers": miami_sellers,
                "city": "MIAMI"
            }
        ]

        for parsing_data in formated:
            for seller in parsing_data["sellers"]:
                for district in districts_data[parsing_data["city"]]:
                    distance = distance_between_points(float(district["latitude"]), float(district["longitude"]), seller["latitude"], seller["longitude"])
                    if distance * 1000 < district["radius"]:
                        seller["district"] = unidecode(district["name"]).upper()
                        sellers_ordered.append(seller)

        try:
            client = MongoClient(MONGO_URI)

            print("Connected")

            db = client[DATABASE_NAME]
            collection = db[COLLECTION_NAME]

            result = collection.insert_many(sellers_ordered, ordered=False)
        except Exception as e:
            print("Error connecting to MongoDB:", e)

        finally:
            client.close()
            print("Connection closed")

    except json.JSONDecodeError as e:
        print(f"Error decoding JSON: {e}")
