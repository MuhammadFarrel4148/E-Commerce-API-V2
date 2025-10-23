const { Pool } = require('pg')
const { nanoid } = require('nanoid');
const bcrypt = require('bcrypt');
const InvariantError = require('../exceptions/InvariantError');
const AuthenticationError = require('../exceptions/AuthenticationError');

class UsersService {
    constructor() {
        this._pool = new Pool();
    }

    async verifyUsername(username) {
        const query = {
            text: `SELECT * FROM users WHERE username = $1`,
            values: [username]
        };
        const result = await this._pool.query(query);

        if(result.rows.length > 0) {
            throw new InvariantError('Username telah digunakan');
        };
    };

    async verifyCredential(username, password) {
        const query = {
            text: `SELECT * FROM users WHERE username = $1`,
            values: [username]
        };
        const result = await this._pool.query(query);

        if(!result.rows.length) {
            throw new AuthenticationError('Kredensial yang diberikan salah');
        };

        const { user_id: id, password: hashedPassword } = result.rows[0];

        const match = await bcrypt.compare(password, hashedPassword);

        if(!match) {
            throw new AuthenticationError('Kredensial yang diberikan salah');
        };

        return id;
    };

    async addUsersService(username, password, fullname) {
        await this.verifyUsername(username);

        const user_id = `user-${nanoid(16)}`;
        const hashedPassword = await bcrypt.hash(password, 10);
        
        const query = {
            text: `INSERT INTO users(user_id, username, password, fullname) VALUES($1, $2, $3, $4) RETURNING user_id`,
            values: [user_id, username, hashedPassword, fullname]
        };
        const result = await this._pool.query(query);

        if(!result.rows.length) {
            throw new InvariantError('User gagal ditambahkan');
        };

        return result.rows[0].user_id;
    };
};

module.exports = UsersService;