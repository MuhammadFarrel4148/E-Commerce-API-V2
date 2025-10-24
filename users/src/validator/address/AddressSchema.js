const Joi = require('joi');

const addAddressSchema = Joi.object({
    street: Joi.string().required(),
    city: Joi.string().required(),
    state: Joi.string().required(),
    country: Joi.string().required()
});

module.exports = addAddressSchema;