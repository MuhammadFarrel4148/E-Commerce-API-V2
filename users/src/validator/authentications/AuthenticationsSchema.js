const Joi = require('joi');

const AuthenticationsSchema = Joi.object({
    username: Joi.string.required(),
    password: Joi.string.required()
});

module.exports = AuthenticationsSchema;