require('dotenv').config();

const express = require('express');

// Define app express and port
const app = express();
const port = process.env.PORT;

// Automate parse to JSON
app.use(express.json());

// Define Routing
app.use('/', (req, res) => {
    res.send('Hello World');
});

// Start Server
app.listen(port, () => {
    console.log(`Listening on port ${port}`); 
});