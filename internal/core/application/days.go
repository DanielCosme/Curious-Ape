package application

import (
	"fmt"
	"github.com/danielcosme/curious-ape/internal/core/database"
	"github.com/danielcosme/curious-ape/internal/core/entity"
	"github.com/danielcosme/curious-ape/internal/integrations/toggl"
	"github.com/danielcosme/curious-ape/sdk/errors"
	"github.com/danielcosme/curious-ape/sdk/log"
	"time"
)

func (a *App) DayCreate(d *entity.Day) (*entity.Day, error) {
	d.Date = time.Date(d.Date.Year(), d.Date.Month(), d.Date.Day(), 0, 0, 0, 0, time.UTC)
	if err := a.db.Days.Create(d); err != nil {
		return nil, err
	}

	return a.DayGetByDate(d.Date)
}

func (a *App) DaysGetAll() ([]*entity.Day, error) {
	return a.db.Days.Find(entity.DayFilter{}, database.DaysPipeline(a.db)...)
}

func (a *App) DayGetByDate(date time.Time) (*entity.Day, error) {
	d, err := a.db.Days.Get(entity.DayFilter{Dates: []time.Time{date}})
	if err != nil && !errors.Is(err, database.ErrNotFound) {
		return nil, err
	}
	if d == nil {
		// if it does not exist, create new and return.
		d, err = a.DayCreate(&entity.Day{Date: date})
		if err != nil {
			return nil, err
		}
	}

	if err = database.ExecuteDaysPipeline([]*entity.Day{d}, database.DaysPipeline(a.db)...); err != nil {
		return nil, err
	}
	return d, nil
}

func (a *App) SyncDeepWorkByDateRange(start, end time.Time) error {
	togglAPI, o, err := a.TogglAPI()
	if err != nil {
		return err
	}
	days, err := a.daysGetByDateRange(start, end)
	if err != nil {
		return err
	}

	for _, d := range days {
		summary, err := togglAPI.Reports.GetDaySummaryForProjectIDs(d.Date, o.ToogglProjectIDs, o.ToogglWorkSpaceID)
		if err != nil {
			return err
		}
		if _, err := a.dayUpdate(d, workLogFromToggl(summary)); err != nil {
			return nil
		}

		if err := a.createDeepWorkLog(d, entity.Toggl); err != nil {
			return err
		}
		togglSleep()
	}

	return nil
}

func (a *App) SyncDeepWorkLog(date time.Time) error {
	togglAPI, o, err := a.TogglAPI()
	if err != nil {
		return err
	}
	d, err := a.DayGetByDate(date)
	if err != nil {
		return err
	}

	summary, err := togglAPI.Reports.GetDaySummaryForProjectIDs(d.Date, o.ToogglProjectIDs, o.ToogglWorkSpaceID)
	if err != nil {
		return err
	}
	if _, err := a.dayUpdate(d, workLogFromToggl(summary)); err != nil {
		return nil
	}

	return a.createDeepWorkLog(d, entity.Toggl)
}

func (a *App) TogglAPI() (*toggl.API, *entity.Oauth2, error) {
	o, err := a.db.Oauths.Get(entity.Oauth2Filter{Provider: []entity.IntegrationProvider{entity.ProviderToggl}})
	if err != nil {
		return nil, nil, err
	}
	return a.sync.TogglClient(o.AccessToken), o, nil
}

func (a *App) SyncDeepWork() error {
	days, err := a.db.Days.Find(entity.DayFilter{}, database.DaysJoinSleepLogs(a.db))
	if err != nil {
		return err
	}
	togglAPI, o, err := a.TogglAPI()
	if err != nil {
		return err
	}

	for _, d := range days {
		summary, err := togglAPI.Reports.GetDaySummaryForProjectIDs(d.Date, o.ToogglProjectIDs, o.ToogglWorkSpaceID)
		if err != nil {
			return err
		}
		if _, err := a.dayUpdate(d, workLogFromToggl(summary)); err != nil {
			return nil
		}

		if err := a.createDeepWorkLog(d, entity.Toggl); err != nil {
			return err
		}
		togglSleep()
	}

	return nil
}

func workLogFromToggl(s *toggl.Summary) *entity.Day {
	return &entity.Day{
		DeepWorkMinutes: int(toggl.ToDuration(s.TotalGrand).Minutes()),
	}
}

func togglSleep() {
	// Toggle Api Only accepts 1 api cal per second
	time.Sleep(time.Second)
}

func (a *App) HabitUpsertFromDeepWorkLog(d *entity.Day, origin entity.DataSource) error {
	habitCategory, err := a.HabitCategoryGetByType(entity.HabitTypeDeepWork)
	if err != nil {
		return err
	}

	var success bool
	// If the deep work minutes are bigger than 5 hours
	dur := time.Duration(d.DeepWorkMinutes) * time.Minute
	if dur >= (time.Hour * 5) {
		success = true
	}

	habit := &entity.Habit{
		DayID:      d.ID,
		CategoryID: habitCategory.ID,
		Logs: []*entity.HabitLog{{
			Success:     success,
			IsAutomated: origin != entity.Manual,
			Origin:      origin,
			Note:        fmt.Sprintf("Deep work duration: %s", dur.String()),
		}},
	}

	_, err = a.HabitCreate(d, habit)
	return err
}

func (a *App) DayUpdate(day, data *entity.Day) (*entity.Day, error) {
	var err error
	day, err = a.dayUpdate(day, data)
	if err != nil {
		return nil, err
	}

	// create deep work resource (in the future) and upsert habit
	if err := a.createDeepWorkLog(day, entity.Manual); err != nil {
		return nil, err
	}

	database.ExecuteDaysPipeline([]*entity.Day{day}, database.DaysJoinHabits(a.db))
	return day, err
}

func (a *App) dayUpdate(day, data *entity.Day) (*entity.Day, error) {
	day.DeepWorkMinutes = data.DeepWorkMinutes
	return a.db.Days.Update(day, database.DaysPipeline(a.db)...)
}

func (a *App) createDeepWorkLog(day *entity.Day, origin entity.DataSource) error {
	if err := a.HabitUpsertFromDeepWorkLog(day, origin); err != nil {
		return err
	}

	a.Log.InfoP("updated deep work log", log.Prop{
		"origin": origin.Str(),
		"date":   day.Date.Format(entity.HumanDate),
	})
	return nil
}

func (a *App) daysGetByDateRange(start, end time.Time) ([]*entity.Day, error) {
	days := []*entity.Day{}

	for _, date := range datesRange(start, end) {
		d, err := a.DayGetByDate(date)
		if err != nil {
			return nil, err
		}

		days = append(days, d)
	}

	return days, nil
}

func datesRange(start, end time.Time) []time.Time {
	dates := []time.Time{}
	inter := start

	for inter.Before(end) {
		dates = append(dates, inter)
		inter = inter.AddDate(0, 0, 1)
	}
	dates = append(dates, end)

	return dates
}
