### HARMONY DATA API Microservice

The purpose of this service is maintain the sellers information, order all the information and provide endpoints to collect the information

## Pre installation

Order all information to improve the data getters of database efficiency

### Run Save districts script

    python ./scripts/save_district_data.py

It will load all the inforamtion on preloaded files and the save it on the database

### Run Read and order sellers data

    python ./scripts/read_and_order_by_district.py

It will save all the sellers information on the database, then you will be able to request it and fragment it using districts (BOGOTA, MIAMI, MEXICO_CITY), it will improve the data reading for the frontend

The script was created using a circular distribution for all the districts so it cant be used on proudction, I created thsi only for the test puposes. It uses a circular distribution to detect where the sellers is located on the map so the script cant reach all the possible sellers on each city.

This script was create to reduce the load on the frontend.

# REST API

Use the data aspi to collect information from the database, the information was ordered by pre installation scripts so you must use the scripts to initialize the databases

## Status

### Request

`GET /status`

Get the state of the application, it will return true when it is working normally

### Response

    {
        "state": true
    }

## Get all sellers

### Request

`GET /sellers`

Get the sellers list on the database, it will return all the sellers without filtering

### Response

    {
        "data": [
            {
                "pdv": 15836,
                "vendor_code": "",
                "address": "Calle Miguel Laurent #1700-5, Apt 48",
                "sales_liters": 0,
                "sales_usd": 0,
                "sales_units": 0,
                "latitude": 19.373766,
                "longitude": -99.156176,
                "city": "Mexico City",
                "product_name": "",
                "vendor_name": "",
                "country": "Mexico",
                "continent": "Central America",
                "category": "Soda",
                "brand": "Coca Cola",
                "sub_brand": "",
                "item": "Coca Cola Zero Returnable 250ml Glass",
                "year": 2024,
                "month": "January",
                "quarter": "Q1",
                "NSE": 4
            },
        ]
        "error": false,
    }

## Get all sellers by district

### Request

`GET /sellers-by-district/:city/:district`

Get the sellers list on the database, it will return only the sellers who are realted with the city and the district. React the read_and_order script to know what is the possible order of the seller for each district

### Response

    {
        "data": [
            {
                "pdv": 15836,
                "vendor_code": "",
                "address": "Calle Miguel Laurent #1700-5, Apt 48",
                "sales_liters": 0,
                "sales_usd": 0,
                "sales_units": 0,
                "latitude": 19.373766,
                "longitude": -99.156176,
                "city": "Mexico City",
                "product_name": "",
                "vendor_name": "",
                "country": "Mexico",
                "continent": "Central America",
                "category": "Soda",
                "brand": "Coca Cola",
                "sub_brand": "",
                "item": "Coca Cola Zero Returnable 250ml Glass",
                "year": 2024,
                "month": "January",
                "quarter": "Q1",
                "NSE": 4
            },
        ]
        "error": false,
    }

## Get all sellers by city

### Request

`GET /sellers-by-city/:city`

Get the sellers list on the database, it will return only the sellers who are realted with the city

### Response

    {
        "data": [
            {
                "pdv": 15836,
                "vendor_code": "",
                "address": "Calle Miguel Laurent #1700-5, Apt 48",
                "sales_liters": 0,
                "sales_usd": 0,
                "sales_units": 0,
                "latitude": 19.373766,
                "longitude": -99.156176,
                "city": "Mexico City",
                "product_name": "",
                "vendor_name": "",
                "country": "Mexico",
                "continent": "Central America",
                "category": "Soda",
                "brand": "Coca Cola",
                "sub_brand": "",
                "item": "Coca Cola Zero Returnable 250ml Glass",
                "year": 2024,
                "month": "January",
                "quarter": "Q1",
                "NSE": 4
            },
        ]
        "error": false,
    }

## Get all districts

### Request

`GET /districts`

The api returns all districts on the database, go to the distrcts scripts to know what it means, it was ordered to reach a big part of the cities but not all

### Response

    {
        "data": [
            {
                "name": "Usaquén",
                "city": "BOGOTA",
                "uuid": "USAQUEN",
                "latitude": "4.700960",
                "longitude": "-74.030115",
                "radius": 4000
            },
            .
            .
            .
        ]
        "error": false,
    }

## Get all districts by city

### Request

`GET /districts`

The api returns all districts with the city relation

### Response

    {
        "data": [
            {
                "name": "Usaquén",
                "city": "BOGOTA",
                "uuid": "USAQUEN",
                "latitude": "4.700960",
                "longitude": "-74.030115",
                "radius": 4000
            },
            .
            .
            .
        ]
        "error": false,
    }
