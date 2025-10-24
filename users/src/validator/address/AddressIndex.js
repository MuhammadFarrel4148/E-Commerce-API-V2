const addAddressSchema = require('./AddressSchema');
const InvariantError = require('../../exceptions/InvariantError');

const AddressValidator = {    
    validateAddAddressPayload: (payload) => {
        const validationResult = addAddressSchema.validate(payload);

        if(validationResult.error) {
            throw new InvariantError(validationResult.error);
        };
    }
};

module.exports = AddressValidator;