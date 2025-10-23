const AuthenticationsSchema = require('./AuthenticationsSchema');
const InvariantError = require('../../exceptions/InvariantError');

const AuthenticationsValidator = {
    validateAuthenticationsPayload: (payload) => {
        const validationResult = AuthenticationsSchema.validate(payload);

        if(validationResult.error) {
            throw new InvariantError(validationResult.error);
        };
    }
};

module.exports = AuthenticationsValidator;