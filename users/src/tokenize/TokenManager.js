require('dotenv').config();

const jwt = require('jsonwebtoken');
const AuthenticationError = require('../exceptions/AuthenticationError');

const TokenManager = {
    generateAccessToken: (payload) => jwt.sign(payload, process.env.ACCESS_TOKEN, { expiresIn: '1h' }),
    generateRefreshToken: (payload) => jwt.sign(payload, process.env.REFRESH_TOKEN),
    updateAccessToken: (payload) => {
        try {
            const decoded = jwt.verify(payload, process.env.REFRESH_TOKEN);
            const { userId } = decoded;
            return userId;
        } catch(error) {
            throw new AuthenticationError('Kredensial yang diberikan salah')
        };
    }
};

module.exports = TokenManager;