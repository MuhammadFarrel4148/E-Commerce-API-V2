class AuthenticationsController {
    constructor(authenticationsService, usersService, tokenManager, validator) {
        this._authenticationsService = authenticationsService;
        this._usersService = usersService;
        this._tokenManager = tokenManager;
        this._validator = validator;

        this.addAuthenticationsController = this.addAuthenticationsController.bind(this);
        this.putAuthenticationsController = this.putAuthenticationsController.bind(this);
        this.deleteAuthenticationsController = this.deleteAuthenticationsController.bind(this);
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

    async putAuthenticationsController(req, res) {
        await this._validator.validatePutAuthenticationsPayload(req.body);

        const { refreshToken } = req.body;

        await this._authenticationsService.verifyToken(refreshToken);

        const userId = await this._tokenManager.updateAccessToken(refreshToken);
        const accessToken = await this._tokenManager.generateAccessToken({ userId });

        res.status(201).json({
            status: 'success',
            data: {
                accessToken
            }
        });
    };

    async deleteAuthenticationsController(req, res) {
        await this._validator.validatePutAuthenticationsPayload(req.body);

        const { refreshToken } = req.body;

        await this._authenticationsService.verifyToken(refreshToken);
        
        await this._authenticationsService.deleteTokenService(refreshToken);

        res.status(201).json({
            status: 'success',
            message: 'Token berhasil dihapus'
        });
    };
};

module.exports = AuthenticationsController;