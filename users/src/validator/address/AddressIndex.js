const {
    addAddressSchema,
    putAddressSchema
} = require('./AddressSchema');
const InvariantError = require('../../exceptions/InvariantError');

const AddressValidator = {    
    validateAddAddressPayload: (payload) => {
        const validationResult = addAddressSchema.validate(payload);

        if(validationResult.error) {
            throw new InvariantError(validationResult.error);
        };
    },
    validatePutAddressPayload: (payload) => {
        const validationResult = putAddressSchema.validate(payload);

        if(validationResult.error) {
            throw new InvariantError(validationResult.error.message);
        };
    }
};

module.exports = AddressValidator;