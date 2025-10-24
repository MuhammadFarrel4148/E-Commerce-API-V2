const { Pool } = require('pg');
const AuthenticationError = require('../exceptions/AuthenticationError');

class AuthenticationsService {
    constructor() {
        this._pool = new Pool();
    };

    async verifyToken(token) {
        const query = {
            text: `SELECT * FROM token WHERE token = $1`,
            values: [token]
        };
        const result = await this._pool.query(query);

        if(!result.rows.length) {
            throw new AuthenticationError('Kredensial yang diberikan salah');
        };
    };

    async addTokenService(token) {
        const query = {
            text: `INSERT INTO token(token) VALUES($1)`,
            values: [token]
        };
        await this._pool.query(query);
    };

    async deleteTokenService(token) {
        const query = {
            text: `DELETE FROM token WHERE token = $1`,
            values: [token]
        };
        await this._pool.query(query);
    };
};

module.exports = AuthenticationsService;