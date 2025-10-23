const { Pool } = require('pg')
const { nanoid } = require('nanoid');
const bcrypt = require('bcrypt');
const InvariantError = require('../exceptions/InvariantError');

class UsersService {
    constructor() {
        this._pool = new Pool();
    }

    async addUsersService(username, password, fullname) {
        const user_id = `user-${nanoid(16)}`;
        const hashedPassword = bcrypt.hash(password, 10);

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

modoule.exports = UsersService;