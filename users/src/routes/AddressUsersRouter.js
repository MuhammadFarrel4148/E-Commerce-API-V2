const express = require('express');
const AuthorizationsService = require('../middleware/AuthorizationService');
const AddressUsersController = require('../controller/AddressUsersController');
const AddressUsersService = require('../service/AddressUsersService');
const AddressValidator = require('../validator/address/AddressIndex');

const addressRouter = express.Router();
const authorizationsService = new AuthorizationsService();
const addressUsersService = new AddressUsersService();
const addressUsersController = new AddressUsersController(addressUsersService, AddressValidator);

addressRouter.post('/add', authorizationsService.AuthMiddleware, addressUsersController.addAddressUsersController);
addressRouter.patch('/put', authorizationsService.AuthMiddleware, addressUsersController.putAddressUsersController);

module.exports = addressRouter;