class UsersController {
    constructor(usersService, validator) {
        this._usersService = usersService;
        this._validator = validator; 
        
        this.addUsersController = this.addUsersController.bind(this);
    };

    async addUsersController(req, res) {
        await this._validator.validateUsersPayload(req.body);

        const { username, password, fullname } = req.body;

        const userId = await this._usersService.addUsersService(username, password, fullname)

        res.status(201).json({
            status: 'success',
            data: {
                userId
            }
        });
    };
};

module.exports = UsersController;