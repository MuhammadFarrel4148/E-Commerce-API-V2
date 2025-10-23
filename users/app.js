require('dotenv').config();

const express = require('express');
const userRouter = require('./src/routes/UsersRouter');

// Define app express and port
const app = express();
const port = process.env.PORT;

// Automate parse to JSON
app.use(express.json());

// Define Routing
app.use('/users', userRouter);

// Start Server
app.listen(port, () => {
    console.log(`Listening on port ${port}`); 
});