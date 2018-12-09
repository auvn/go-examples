use strips;
db.createCollection("trips");
db.trips.createIndex({completed: 1, driver_id: 1}, {unique:true, partialFilterExpression:{completed:false}})
db.trips.createIndex({completed: 1, rider_id: 1}, {unique:true, partialFilterExpression:{completed:false}})

use shistory;
db.createCollection("history");
db.history.createIndex({created_at: 1});

use stracking;
db.createCollection("drivers");
db.drivers.createIndex({updated_at: 1}, {expireAfterSeconds: 10});
db.drivers.createIndex({lock_deadline: 1});