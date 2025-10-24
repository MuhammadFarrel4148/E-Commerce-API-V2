class AddressUsersController {
    constructor(addressUsersService, validator) {
        this._addressUsersService = addressUsersService;
        this._validator = validator;

        this.addAddressUsersController = this.addAddressUsersController.bind(this);
        this.putAddressUsersController = this.putAddressUsersController.bind(this);
    };

    async addAddressUsersController(req, res) {
        await this._validator.validateAddAddressPayload(req.body);

        const { userId } = req.user;
        const { street, city, state, country } = req.body;

        const addressId = await this._addressUsersService.addAddressUsersService(street, city, state, country, userId);

        res.status(201).json({
            status: 'success',
            data: {
                addressId
            }
        });
    };

    async putAddressUsersController(req, res) {
        // ambil 
    };  
};

module.exports = AddressUsersController;