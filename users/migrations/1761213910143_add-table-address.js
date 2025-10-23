/**
 * @type {import('node-pg-migrate').ColumnDefinitions | undefined}
 */
export const shorthands = undefined;

/**
 * @param pgm {import('node-pg-migrate').MigrationBuilder}
 * @param run {() => void | undefined}
 * @returns {Promise<void> | void}
 */
export const up = (pgm) => {
    pgm.createTable('address', {
        address_id: {
            type: 'VARCHAR(50)',
            primaryKey: true
        },
        user_id: {
            type: 'VARCHAR(50)',
            notNull: true
        },
        street: {
            type: 'TEXT',
            notNull: true
        },
        city: {
            type: 'TEXT',
            notNull: true
        },
        state: {
            type: 'TEXT',
            notNull: true
        },
        country: {
            type: 'TEXT',
            notNull: true
        }
    });
    pgm.addConstraint('address', 'fk_constraint.address_userid.users_userid', 'FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE');
};

/**
 * @param pgm {import('node-pg-migrate').MigrationBuilder}
 * @param run {() => void | undefined}
 * @returns {Promise<void> | void}
 */
export const down = (pgm) => {
    pgm.dropTable('address');
    pgm.dropConstraint('address', 'fk_constraint.address_userid.users_userid');
};
