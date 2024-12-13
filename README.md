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
