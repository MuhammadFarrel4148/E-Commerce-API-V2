class AuthenticationsController {
    constructor(authenticationsService, usersService, tokenManager, validator) {
        this._authenticationsService = authenticationsService;
        this._usersService = usersService;
        this._tokenManager = tokenManager;
        this._validator = validator;

        this.addAuthenticationsController = this.addAuthenticationsController.bind(this);
    };

    async addAuthenticationsController(req, res) {
        await this._validator.validateAuthenticationsPayload(req.body);

        const { username, password } = req.body;

        const userId = await this._usersService.verifyCredential(username, password);

        const accessToken = await this._tokenManager.generateAccessToken({ userId });
        const refreshToken = await this._tokenManager.generateRefreshToken({ userId });
        await this._authenticationsService.addTokenService(refreshToken);

        res.status(201).json({
            status: 'success',
            data: {
                accessToken,
                refreshToken
            }
        });
    };
};

module.exports = AuthenticationsController;