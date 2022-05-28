package echo

import (
	"github.com/danielcosme/curious-ape/internal/core/application"
	"github.com/danielcosme/curious-ape/internal/transport/echo/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func Routes(a *application.App) http.Handler {
	h := Handler{App: a}
	e := echo.New()
	e.HTTPErrorHandler = createErrHandler(a)
	e.Use(middleware.Logger(a))
	e.Use(middleware.Recover(a))

	e.GET("/ping", h.Ping)

	days := e.Group("/days")
	{
		days.GET("", h.DaysGetAll)
		daysByDate := days.Group("/:date", middleware.SetDay(a))
		{ // /days/:date/habits
			daysByDate.POST("/habits", h.HabitCreate)
		}
	}

	habits := e.Group("/habits")
	{
		habits.GET("", h.HabitsGetAll)
		habits.POST("", h.HabitCreate, middleware.SetDay(a))
		habits.GET("/categories", h.HabitsGetAllCategories)
		habitsByID := habits.Group("/:id", middleware.SetHabit(a))
		{
			habitsByID.GET("", h.HabitGetByID)
			habitsByID.PUT("", h.HabitFullUpdate)
			habitsByID.DELETE("", h.HabitDelete)
		}
	}

	sleepRecords := e.Group("/sleep")
	{
		sleepRecords.GET("/debug", h.FitbitDebug)
	}

	oauths := e.Group("/oauth2")
	{
		oauths.GET("/:provider/connect", h.Oauth2Connect)
		oauths.GET("/:provider/success", h.Oauth2Success)
	}

	return e
}


func createErrHandler(a *application.App) echo.HTTPErrorHandler {
	return func (err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		props := map[string]string{}
		props["Code"] = strconv.Itoa(c.Response().Status)
		props["Method"] = c.Request().Method
		a.Error(err, props)
	}
}
