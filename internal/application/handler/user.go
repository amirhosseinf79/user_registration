package handler

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/interfaces"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) interfaces.UserHandler {
	return &userHandler{
		userService: userService,
	}
}

// @Summary Get user details
// @Description Get User Details by ID
// @Tags User
// @Accept json
// @Produce json
// @Param userID path int true "UserID"
// @Success 200 {object} user.ResponseDetails
// @Failure 404 {object} shared.ResponseOneMessage
// @Router /user/{userID} [get]
func (uh *userHandler) GetUserByID(ctx *fiber.Ctx) error {
	userID, _ := ctx.ParamsInt("userID", 0)
	userDetails, err := uh.userService.GetUserDetailsByID(uint(userID))
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(userDetails)
}

// @Summary Get all users
// @Description Get list of all users
// @Tags User
// @Accept json
// @Produce json
// @Param filters query user.FilterUser false "Filters"
// @Success 200 {object} user.ResponseList
// @Router /user/all [get]
func (uh *userHandler) GetUsersList(ctx *fiber.Ctx) error {
	var filter user.FilterUser
	ctx.QueryParser(&filter)
	userDetails, err := uh.userService.GetUserList(filter)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(userDetails)
}

// @Summary Get user profile
// @Description Get User profile
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} user.ResponseDetails
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /profile [get]
func (uh *userHandler) GetUserProfile(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)
	userDetails, err := uh.userService.GetUserDetailsByID(userID)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(userDetails)
}

// @Summary Update user Profile
// @Description Update User Profile
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body user.FieldUpdateDetails true "user"
// @Success 200 {object} user.ResponseDetails
// @Failure 400 {object} shared.ResponseOneMessage
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /profile/update/info [patch]
func (uh *userHandler) UpdateProfileInfo(ctx *fiber.Ctx) error {
	var fields user.FieldUpdateDetails
	ctx.BodyParser(&fields)
	userID := ctx.Locals("userID").(uint)
	userDetails, err := uh.userService.UpdateUserProfile(userID, fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(userDetails)
}

// @Summary Update user password
// @Description Update User Password
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param fields body user.FieldUpdatePassword true "Fields"
// @Success 200 {object} user.ResponseDetails
// @Failure 400 {object} shared.ResponseOneMessage
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /profile/update/password [put]
func (uh *userHandler) UpdateUserPassword(ctx *fiber.Ctx) error {
	var fields user.FieldUpdatePassword
	ctx.BodyParser(&fields)
	userID := ctx.Locals("userID").(uint)
	userDetails, err := uh.userService.UpdateUserPassword(userID, fields)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(userDetails)
}

// @Summary Verify Mobile
// @Description verify Mobile
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param fields body user.FieldVerifyOTP true "Fields"
// @Success 200 {object} user.ResponseDetails
// @Failure 400 {object} shared.ResponseOneMessage
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /profile/verify/mobile [post]
func (ah *userHandler) VerifyUserOTP(ctx *fiber.Ctx) error {
	var fields user.FieldVerifyOTP
	ctx.BodyParser(&fields)
	userID := ctx.Locals("userID").(uint)
	response, err := ah.userService.VerifyUserMobile(userID, fields.Code)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}

// @Summary Send Verify Code
// @Description Send Verify Code
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} auth.OTPOkMock
// @Failure 401 {object} shared.ResponseOneMessage
// @Router /profile/send/otp [post]
func (ah *userHandler) SendUserOTP(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(uint)
	response, err := ah.userService.SendVerifyOTP(userID)
	if err != nil {
		return ctx.Status(err.Code).JSON(err)
	}
	return ctx.JSON(response)
}
