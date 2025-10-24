require('dotenv').config();

const AuthorizationError = require('../exceptions/AuthorizationError');
const jwt = require('jsonwebtoken');

class AuthorizationsService {
    constructor() {

    };

    AuthMiddleware(req, res, next) {
        const authHeader = req.headers.authorization;

        if(!authHeader) {
            throw new AuthorizationError('Unauthorized!');
        };

        const token = authHeader.split(' ')[1];
        const decoded = jwt.verify(token, process.env.ACCESS_TOKEN);
        req.user = decoded;
        next();
    };
};

module.exports = AuthorizationsService;