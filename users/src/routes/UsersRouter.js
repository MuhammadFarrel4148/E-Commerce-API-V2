const express = require('express');
const UsersController = require('../controller/UsersController');
const UsersService = require('../service/UsersService');
const UsersValidator = require('../validator/users/UsersIndex');

const userRouter = express.Router();
const usersService = new UsersService();
const usersController = new UsersController(usersService, UsersValidator);

userRouter.post('/add', usersController.addUsersController);

module.exports = userRouter;