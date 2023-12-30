package middlewares

import (
	"fmt"
	"net/http"
	"numenv_subscription_api/errors/logs"
	"numenv_subscription_api/models/responses"
	"numenv_subscription_api/repositories"

	"github.com/labstack/echo/v4"
)

// Verify that the queried session is not already full
func IsSessionFull(next echo.HandlerFunc) echo.HandlerFunc {
  return func(ctx echo.Context) error {

    sess, err := repositories.GetSessionBySpeaker(
      ctx.Request().Context(),
      ctx.Param("speaker"),
    )
    count, err := repositories.GetSessionNumberSubscribersBySpeaker(ctx)
    if err != nil || count == nil {
      return err
    }

    if *count >= sess.NumSubscribers {
      errMsg := fmt.Sprintf(
        "Max number of subscribers reached for session : %s - \"%s\".", 
        sess.Speaker,
        sess.Name,
      ) 
      logs.Output(
        logs.WARNING,
        errMsg,
      )
      return ctx.JSON(
        http.StatusNotAcceptable,
        responses.ErrorResponse{
          Message: errMsg,
        },
      )
    }

    return next(ctx)
  }
}
