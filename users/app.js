require('dotenv').config();

const express = require('express');
const userRouter = require('./src/routes/UsersRouter');
const authenticationRouter = require('./src/routes/AuthenticationRouter');
const addressRouter = require('./src/routes/AddressUsersRouter');
const ClientError = require('./src/exceptions/ClientError');

const app = express();
const port = process.env.PORT;

app.use(express.json());
app.use('/users', userRouter);
app.use('/authentications', authenticationRouter);
app.use('/address', addressRouter);
app.use((err, req, res, next) => {
    if(err instanceof ClientError) {
        res.status(err.statusCode).json({
            status: 'fail',
            message: err.message
        });
    };
});

app.listen(port, () => {
    console.log(`Listening on port ${port}`); 
});