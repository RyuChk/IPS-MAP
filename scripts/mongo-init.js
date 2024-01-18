rs.status();
use admin; 
db.createUser({user: 'admin', pwd: 'admin', roles: [ 
    { role: 'userAdminAnyDatabase', db: 'admin' },
    { role: "dbAdminAnyDatabase", db: "admin" }, 
    { role: "readWriteAnyDatabase", db: "admin" },
    { role: 'readWrite', db: 'ips-rssi-beta' }
]});
use ips-map;
db.createCollection("map-image-collection");
