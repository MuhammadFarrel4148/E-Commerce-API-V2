const {
    AuthenticationsSchema,
    putAuthenticationsSchema
} = require('./AuthenticationsSchema');
const InvariantError = require('../../exceptions/InvariantError');

const AuthenticationsValidator = {
    validateAuthenticationsPayload: (payload) => {
        const validationResult = AuthenticationsSchema.validate(payload);

        if(validationResult.error) {
            throw new InvariantError(validationResult.error);
        };
    },
    validatePutAuthenticationsPayload: (payload) => {
        const validationResult = putAuthenticationsSchema.validate(payload);

        if(validationResult.error) {
            throw new InvariantError(validationResult.error);
        };
    }
};

module.exports = AuthenticationsValidator;