const { Pool } = require('pg');
const { nanoid } = require('nanoid');
const InvariantError = require('../exceptions/InvariantError');

class AddressUsersService {
    constructor() {
        this._pool = new Pool();
    };

    async addAddressUsersService(street, city, state, country, userId) {
        const addressId = `address-${nanoid(16)}`;

        const query = {
            text: `INSERT INTO address(address_id, user_id, street, city, state, country) VALUES($1, $2, $3, $4, $5, $6) RETURNING address_id`,
            values: [addressId, userId, street, city, state, country]
        };
        const result = await this._pool.query(query);

        if(!result.rows.length) {
            throw new InvariantError('Address gagal ditambahkan');
        };

        return result.rows[0].address_id;
    };

    async putAddressUsersService(street, city, state, country, userId) {
        const fieldsUpdate = [];
        const valuesUpdate = [];
        let paramIndex = 1;

        if(street !== undefined) {
            fieldsUpdate.push(`street = $${paramIndex++}`);
            valuesUpdate.push(street);
        };

        if(city !== undefined) {
            fieldsUpdate.push(`city = $${paramIndex++}`);
            valuesUpdate.push(city);
        };

        if(state !== undefined) {
            fieldsUpdate.push(`state = $${paramIndex++}`);
            valuesUpdate.push(state);
        };

        if(country !== undefined) {
            fieldsUpdate.push(`country = $${paramIndex++}`);
            valuesUpdate.push(country);
        };

        const condition = `WHERE user_id = $${paramIndex++}`;
        valuesUpdate.push(userId);

        const query = {
            text: `UPDATE Address SET ${fieldsUpdate.join(', ')} ${condition}`,
            values: valuesUpdate
        };
        await this._pool.query(query);
    };
};

module.exports = AddressUsersService;