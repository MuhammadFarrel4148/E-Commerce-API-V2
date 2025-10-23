require('dotenv').config();

const jwt = require('jsonwebtoken');

const TokenManager = {
    generateAccessToken: (payload) => jwt.sign(payload, process.env.ACCESS_TOKEN, { expiresIn: '1h' }),
    generateRefreshToken: (payload) => jwt.sign(payload, process.env.REFRESH_TOKEN)
};

module.exports = TokenManager;