# Vehicle Sensor API

## Scenario

You have a transport company, and there are 10 trucks (long vehicles), there are 2 devices (the case temperature control sensor and the other gps location information) on these vehicles, and these devices need to be registered in the system, steps to be followed while registering to the system:

1. Register vehicles (Truck) to the system with license plates
2. Then you will add 2 device tours to the device_type table (such as temperature sensor, GPS)
3. You will define 2 devices for each TIR registered in the system.
4. Then you send the data read to the log_temperature and log_location route part with the post method.

- Routes:
  - Parent root -> vehicle
  - Get -> vehicle_list
  - Post-> vehicle_add
  - Patch-> vehicle_update
  - Delete-> vehicle_delete
- Parent root -> device
  - Get -> device_list
  - Post-> device_add
  - Patch-> device_update
  - Delete-> device_delete
- Parent root -> device_type
  - Get ->type_list
  - Post-> type_add
  - Delete-> device_delete
- Parent root -> log_temperature
  - Get -> temp_list
  - Post-> temp_add
- Parent root -> log_location
  - Get -> gps_list
  - Post-> gps_add

## How to run?

### Database

```bash
# install images
docker pull postgres
# run contenier
docker run --name postgres -e POSTGRES_PASSWORD=pass -d -p 5432:5432 postgres
# postgresql restore database
cat database.sql | docker exec -i postgres psql -U postgres
```

### Server

```bash
# install dependencies
npm i
# Run server
npm start
```
