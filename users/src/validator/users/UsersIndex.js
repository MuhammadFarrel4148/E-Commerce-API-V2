const UsersSchema = require('./UsersSchema');
const InvariantError = require('../../exceptions/InvariantError');

const UsersValidator = {
    validateUsersPayload: (payload) => {
        const validationResult = UsersSchema.validate(payload);
        console.log(validationResult);

        if(validationResult.error) {
            throw new InvariantError(validationResult.error);
        };
    }
};

module.exports = UsersValidator;