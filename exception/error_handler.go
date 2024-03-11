package exception

import (
	"Ubersnap-middle-backend-programmer-test/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var response utils.WebResponse = utils.WebResponse{
		Code:    500,
		Status:  "INTERNAL_SERVER_ERROR",
		Message: err.Error(),
		Data:    false,
	}

	_, ok := err.(ValidationError)
	if ok {
		response = utils.WebResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
			Data:    false,
		}
	} else if _, ok := err.(ServerError); ok {
		response = utils.WebResponse{
			Code:    500,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: err.Error(),
			Data:    false,
		}
	} else if _, ok := err.(NotFoundError); ok {
		response = utils.WebResponse{
			Code:    404,
			Status:  "NOT_FOUND",
			Message: err.Error(),
			Data:    false,
		}
	} else if _, ok := err.(UnauthenticatedError); ok {
		response = utils.WebResponse{
			Code:    401,
			Status:  "UNAUTHENTICATED",
			Message: err.Error(),
			Data:    false,
		}
	} else if _, ok := err.(ClientError); ok {
		response = utils.WebResponse{
			Code:    400,
			Status:  "CLIENT_ERROR",
			Message: err.Error(),
			Data:    false,
		}
	} else if _, ok := err.(MaintenanceError); ok {
		response = utils.WebResponse{
			Code:    http.StatusServiceUnavailable,
			Status:  "SERVICE_UNAVAILABLE",
			Message: err.Error(),
			Data:    false,
		}
	} else if _, ok := err.(TooManyRequestError); ok {
		response = utils.WebResponse{
			Code:    http.StatusTooManyRequests,
			Status:  "TOO_MANY_REQUEST",
			Message: err.Error(),
			Data:    false,
		}
	}

	headers := ctx.GetReqHeaders()
	utils.Logger.Warn(response.Status, utils.LogAny("error", err), utils.LogAny("data", map[string]interface{}{
		"actor":    headers["X-User"],
		"headers":  headers,
		"time":     time.Now().Format("02-Jan-2006 15:04:05"),
		"method":   ctx.Method(),
		"path":     ctx.Path(),
		"response": response,
	}))

	return ctx.Status(response.Code).JSON(response)
}
