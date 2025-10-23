const { Pool } = require('pg');

class AuthenticationsService {
    constructor() {
        this._pool = new Pool();
    };

    async addTokenService(token) {
        const query = {
            text: `INSERT INTO token(token) VALUES($1)`,
            values: [token]
        };
        await this._pool.query(query);
    };
};

module.exports = AuthenticationsService;