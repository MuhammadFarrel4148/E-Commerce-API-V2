const Joi = require('joi');

const addAddressSchema = Joi.object({
    street: Joi.string().required(),
    city: Joi.string().required(),
    state: Joi.string().required(),
    country: Joi.string().required()
});

const putAddressSchema = Joi.object({
    street: Joi.string().allow(''),
    city: Joi.string().allow(''),
    state: Joi.string().allow(''),
    country: Joi.string().allow('')
}).min(1);

module.exports = {
    addAddressSchema,
    putAddressSchema
};