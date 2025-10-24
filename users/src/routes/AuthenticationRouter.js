const express = require('express');
const AuthenticationsController = require('../controller/AuthenticationController');
const AuthenticationsService = require('../service/AuthenticationService');
const UserService = require('../service/UsersService');
const TokenManager = require('../tokenize/TokenManager');
const AuthenticationsValidator = require('../validator/authentications/AuthenticationsIndex');

const authenticationRouter = express.Router();
const authenticationsService = new AuthenticationsService();
const usersService = new UserService();
const authenticationsController = new AuthenticationsController(authenticationsService, usersService, TokenManager, AuthenticationsValidator);

authenticationRouter.post('/add', authenticationsController.addAuthenticationsController);
authenticationRouter.put('/put', authenticationsController.putAuthenticationsController);

module.exports = authenticationRouter;