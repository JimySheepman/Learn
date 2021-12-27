const { Pool } = require('pg');

const connectionInformation = {
    user: "posgres",
    host: "localhost",
    database: "posgres",
    password: "pass",
    port: 5432,
}

const pg_client = new Pool(connectionInformation)
try {
    pg_client.connect();
    console.log("::> PostgreSQL Server is Ready");
} catch (err) {
    console.log(err.stack);
}

exports.pg_client = pg_client