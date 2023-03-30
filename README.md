# go-aircraft-api

# Overview
This project is an API for looking up aircraft data by icao24 hex code. All aircraft have a unique hex code identifier, allowing lookup of any individual aircraft through publicly available data, such as the regularly updated FAA aircraft registry.

Curious about aircraft flying overhead? Start up the API and look up the icao24 hex codes of nearby aircraft through a website like ADS-B Exchange or FlightRadar24. 

# Usage
Add a hex code to the URL like this: localhost:8080/icao24/hex_code_here

Then, view information such as city of registration, engine make and model, year of manufacture, registered owner name, and more in JSON format.

![image](https://user-images.githubusercontent.com/71681977/228971043-4701726f-8b5d-4d09-bb00-776559269bd3.png)
Figure 1: Example output using Postman.

# Structure
This project uses Go and Gin with a sqlite database. Main calls the controller’s get function, which parses the entry and queries the database. The result is then returned as JSON.

![Go-Aircraft-API drawio](https://user-images.githubusercontent.com/71681977/228971643-8fa50212-ba73-4543-8dbc-6c40140fcf2d.png)

Figure 2: Basic structure of the API
