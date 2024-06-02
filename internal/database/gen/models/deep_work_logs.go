// Code generated by BobGen sqlite v0.25.0. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/aarondl/opt/null"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/clause"
	"github.com/stephenafamo/bob/dialect/sqlite"
	"github.com/stephenafamo/bob/dialect/sqlite/dialect"
	"github.com/stephenafamo/bob/dialect/sqlite/im"
	"github.com/stephenafamo/bob/dialect/sqlite/sm"
	"github.com/stephenafamo/bob/dialect/sqlite/um"
	"github.com/stephenafamo/bob/expr"
	"github.com/stephenafamo/bob/mods"
	"github.com/stephenafamo/bob/orm"
)

// DeepWorkLog is an object representing the database table.
type DeepWorkLog struct {
	ID          int32          `db:"id,pk" `
	DayID       int32          `db:"day_id" `
	Date        time.Time      `db:"date" `
	Seconds     int32          `db:"seconds" `
	IsAutomated null.Val[bool] `db:"is_automated" `
	Origin      string         `db:"origin" `

	R deepWorkLogR `db:"-" `
}

// DeepWorkLogSlice is an alias for a slice of pointers to DeepWorkLog.
// This should almost always be used instead of []*DeepWorkLog.
type DeepWorkLogSlice []*DeepWorkLog

// DeepWorkLogs contains methods to work with the deep_work_logs table
var DeepWorkLogs = sqlite.NewTablex[*DeepWorkLog, DeepWorkLogSlice, *DeepWorkLogSetter]("", "deep_work_logs")

// DeepWorkLogsQuery is a query on the deep_work_logs table
type DeepWorkLogsQuery = *sqlite.ViewQuery[*DeepWorkLog, DeepWorkLogSlice]

// DeepWorkLogsStmt is a prepared statment on deep_work_logs
type DeepWorkLogsStmt = bob.QueryStmt[*DeepWorkLog, DeepWorkLogSlice]

// deepWorkLogR is where relationships are stored.
type deepWorkLogR struct {
	Day *Day // fk_deep_work_logs_0
}

// DeepWorkLogSetter is used for insert/upsert/update operations
// All values are optional, and do not have to be set
// Generated columns are not included
type DeepWorkLogSetter struct {
	ID          omit.Val[int32]     `db:"id,pk"`
	DayID       omit.Val[int32]     `db:"day_id"`
	Date        omit.Val[time.Time] `db:"date"`
	Seconds     omit.Val[int32]     `db:"seconds"`
	IsAutomated omitnull.Val[bool]  `db:"is_automated"`
	Origin      omit.Val[string]    `db:"origin"`
}

func (s DeepWorkLogSetter) SetColumns() []string {
	vals := make([]string, 0, 6)
	if !s.ID.IsUnset() {
		vals = append(vals, "id")
	}

	if !s.DayID.IsUnset() {
		vals = append(vals, "day_id")
	}

	if !s.Date.IsUnset() {
		vals = append(vals, "date")
	}

	if !s.Seconds.IsUnset() {
		vals = append(vals, "seconds")
	}

	if !s.IsAutomated.IsUnset() {
		vals = append(vals, "is_automated")
	}

	if !s.Origin.IsUnset() {
		vals = append(vals, "origin")
	}

	return vals
}

func (s DeepWorkLogSetter) Overwrite(t *DeepWorkLog) {
	if !s.ID.IsUnset() {
		t.ID, _ = s.ID.Get()
	}
	if !s.DayID.IsUnset() {
		t.DayID, _ = s.DayID.Get()
	}
	if !s.Date.IsUnset() {
		t.Date, _ = s.Date.Get()
	}
	if !s.Seconds.IsUnset() {
		t.Seconds, _ = s.Seconds.Get()
	}
	if !s.IsAutomated.IsUnset() {
		t.IsAutomated, _ = s.IsAutomated.GetNull()
	}
	if !s.Origin.IsUnset() {
		t.Origin, _ = s.Origin.Get()
	}
}

func (s DeepWorkLogSetter) InsertMod() bob.Mod[*dialect.InsertQuery] {
	vals := make([]bob.Expression, 0, 6)
	if !s.ID.IsUnset() {
		vals = append(vals, sqlite.Arg(s.ID))
	}

	if !s.DayID.IsUnset() {
		vals = append(vals, sqlite.Arg(s.DayID))
	}

	if !s.Date.IsUnset() {
		vals = append(vals, sqlite.Arg(s.Date))
	}

	if !s.Seconds.IsUnset() {
		vals = append(vals, sqlite.Arg(s.Seconds))
	}

	if !s.IsAutomated.IsUnset() {
		vals = append(vals, sqlite.Arg(s.IsAutomated))
	}

	if !s.Origin.IsUnset() {
		vals = append(vals, sqlite.Arg(s.Origin))
	}

	return im.Values(vals...)
}

func (s DeepWorkLogSetter) Apply(q *dialect.UpdateQuery) {
	um.Set(s.Expressions()...).Apply(q)
}

func (s DeepWorkLogSetter) Expressions(prefix ...string) []bob.Expression {
	exprs := make([]bob.Expression, 0, 6)

	if !s.ID.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			sqlite.Quote(append(prefix, "id")...),
			sqlite.Arg(s.ID),
		}})
	}

	if !s.DayID.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			sqlite.Quote(append(prefix, "day_id")...),
			sqlite.Arg(s.DayID),
		}})
	}

	if !s.Date.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			sqlite.Quote(append(prefix, "date")...),
			sqlite.Arg(s.Date),
		}})
	}

	if !s.Seconds.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			sqlite.Quote(append(prefix, "seconds")...),
			sqlite.Arg(s.Seconds),
		}})
	}

	if !s.IsAutomated.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			sqlite.Quote(append(prefix, "is_automated")...),
			sqlite.Arg(s.IsAutomated),
		}})
	}

	if !s.Origin.IsUnset() {
		exprs = append(exprs, expr.Join{Sep: " = ", Exprs: []bob.Expression{
			sqlite.Quote(append(prefix, "origin")...),
			sqlite.Arg(s.Origin),
		}})
	}

	return exprs
}

type deepWorkLogColumnNames struct {
	ID          string
	DayID       string
	Date        string
	Seconds     string
	IsAutomated string
	Origin      string
}

type deepWorkLogRelationshipJoins[Q dialect.Joinable] struct {
	Day bob.Mod[Q]
}

func buildDeepWorkLogRelationshipJoins[Q dialect.Joinable](ctx context.Context, typ string) deepWorkLogRelationshipJoins[Q] {
	return deepWorkLogRelationshipJoins[Q]{
		Day: deepWorkLogsJoinDay[Q](ctx, typ),
	}
}

func deepWorkLogsJoin[Q dialect.Joinable](ctx context.Context) joinSet[deepWorkLogRelationshipJoins[Q]] {
	return joinSet[deepWorkLogRelationshipJoins[Q]]{
		InnerJoin: buildDeepWorkLogRelationshipJoins[Q](ctx, clause.InnerJoin),
		LeftJoin:  buildDeepWorkLogRelationshipJoins[Q](ctx, clause.LeftJoin),
		RightJoin: buildDeepWorkLogRelationshipJoins[Q](ctx, clause.RightJoin),
	}
}

var DeepWorkLogColumns = struct {
	ID          sqlite.Expression
	DayID       sqlite.Expression
	Date        sqlite.Expression
	Seconds     sqlite.Expression
	IsAutomated sqlite.Expression
	Origin      sqlite.Expression
}{
	ID:          sqlite.Quote("deep_work_logs", "id"),
	DayID:       sqlite.Quote("deep_work_logs", "day_id"),
	Date:        sqlite.Quote("deep_work_logs", "date"),
	Seconds:     sqlite.Quote("deep_work_logs", "seconds"),
	IsAutomated: sqlite.Quote("deep_work_logs", "is_automated"),
	Origin:      sqlite.Quote("deep_work_logs", "origin"),
}

type deepWorkLogWhere[Q sqlite.Filterable] struct {
	ID          sqlite.WhereMod[Q, int32]
	DayID       sqlite.WhereMod[Q, int32]
	Date        sqlite.WhereMod[Q, time.Time]
	Seconds     sqlite.WhereMod[Q, int32]
	IsAutomated sqlite.WhereNullMod[Q, bool]
	Origin      sqlite.WhereMod[Q, string]
}

func DeepWorkLogWhere[Q sqlite.Filterable]() deepWorkLogWhere[Q] {
	return deepWorkLogWhere[Q]{
		ID:          sqlite.Where[Q, int32](DeepWorkLogColumns.ID),
		DayID:       sqlite.Where[Q, int32](DeepWorkLogColumns.DayID),
		Date:        sqlite.Where[Q, time.Time](DeepWorkLogColumns.Date),
		Seconds:     sqlite.Where[Q, int32](DeepWorkLogColumns.Seconds),
		IsAutomated: sqlite.WhereNull[Q, bool](DeepWorkLogColumns.IsAutomated),
		Origin:      sqlite.Where[Q, string](DeepWorkLogColumns.Origin),
	}
}

// FindDeepWorkLog retrieves a single record by primary key
// If cols is empty Find will return all columns.
func FindDeepWorkLog(ctx context.Context, exec bob.Executor, IDPK int32, cols ...string) (*DeepWorkLog, error) {
	if len(cols) == 0 {
		return DeepWorkLogs.Query(
			ctx, exec,
			SelectWhere.DeepWorkLogs.ID.EQ(IDPK),
		).One()
	}

	return DeepWorkLogs.Query(
		ctx, exec,
		SelectWhere.DeepWorkLogs.ID.EQ(IDPK),
		sm.Columns(DeepWorkLogs.Columns().Only(cols...)),
	).One()
}

// DeepWorkLogExists checks the presence of a single record by primary key
func DeepWorkLogExists(ctx context.Context, exec bob.Executor, IDPK int32) (bool, error) {
	return DeepWorkLogs.Query(
		ctx, exec,
		SelectWhere.DeepWorkLogs.ID.EQ(IDPK),
	).Exists()
}

// PrimaryKeyVals returns the primary key values of the DeepWorkLog
func (o *DeepWorkLog) PrimaryKeyVals() bob.Expression {
	return sqlite.Arg(o.ID)
}

// Update uses an executor to update the DeepWorkLog
func (o *DeepWorkLog) Update(ctx context.Context, exec bob.Executor, s *DeepWorkLogSetter) error {
	return DeepWorkLogs.Update(ctx, exec, s, o)
}

// Delete deletes a single DeepWorkLog record with an executor
func (o *DeepWorkLog) Delete(ctx context.Context, exec bob.Executor) error {
	return DeepWorkLogs.Delete(ctx, exec, o)
}

// Reload refreshes the DeepWorkLog using the executor
func (o *DeepWorkLog) Reload(ctx context.Context, exec bob.Executor) error {
	o2, err := DeepWorkLogs.Query(
		ctx, exec,
		SelectWhere.DeepWorkLogs.ID.EQ(o.ID),
	).One()
	if err != nil {
		return err
	}
	o2.R = o.R
	*o = *o2

	return nil
}

func (o DeepWorkLogSlice) UpdateAll(ctx context.Context, exec bob.Executor, vals DeepWorkLogSetter) error {
	return DeepWorkLogs.Update(ctx, exec, &vals, o...)
}

func (o DeepWorkLogSlice) DeleteAll(ctx context.Context, exec bob.Executor) error {
	return DeepWorkLogs.Delete(ctx, exec, o...)
}

func (o DeepWorkLogSlice) ReloadAll(ctx context.Context, exec bob.Executor) error {
	var mods []bob.Mod[*dialect.SelectQuery]

	IDPK := make([]int32, len(o))

	for i, o := range o {
		IDPK[i] = o.ID
	}

	mods = append(mods,
		SelectWhere.DeepWorkLogs.ID.In(IDPK...),
	)

	o2, err := DeepWorkLogs.Query(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, old := range o {
		for _, new := range o2 {
			if new.ID != old.ID {
				continue
			}
			new.R = old.R
			*old = *new
			break
		}
	}

	return nil
}

func deepWorkLogsJoinDay[Q dialect.Joinable](ctx context.Context, typ string) bob.Mod[Q] {
	return mods.QueryMods[Q]{
		dialect.Join[Q](typ, Days.NameAs(ctx)).On(
			DayColumns.ID.EQ(DeepWorkLogColumns.DayID),
		),
	}
}

// Day starts a query for related objects on days
func (o *DeepWorkLog) Day(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) DaysQuery {
	return Days.Query(ctx, exec, append(mods,
		sm.Where(DayColumns.ID.EQ(sqlite.Arg(o.DayID))),
	)...)
}

func (os DeepWorkLogSlice) Day(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) DaysQuery {
	PKArgs := make([]bob.Expression, len(os))
	for i, o := range os {
		PKArgs[i] = sqlite.ArgGroup(o.DayID)
	}

	return Days.Query(ctx, exec, append(mods,
		sm.Where(sqlite.Group(DayColumns.ID).In(PKArgs...)),
	)...)
}

func (o *DeepWorkLog) Preload(name string, retrieved any) error {
	if o == nil {
		return nil
	}

	switch name {
	case "Date":
		rel, ok := retrieved.(*Day)
		if !ok {
			return fmt.Errorf("deepWorkLog cannot load %T as %q", retrieved, name)
		}

		o.R.Day = rel

		if rel != nil {
			rel.R.DeepWorkLogs = DeepWorkLogSlice{o}
		}
		return nil
	default:
		return fmt.Errorf("deepWorkLog has no relationship %q", name)
	}
}

func PreloadDeepWorkLogDay(opts ...sqlite.PreloadOption) sqlite.Preloader {
	return sqlite.Preload[*Day, DaySlice](orm.Relationship{
		Name: "Date",
		Sides: []orm.RelSide{
			{
				From: "deep_work_logs",
				To:   TableNames.Days,
				ToExpr: func(ctx context.Context) bob.Expression {
					return Days.Name(ctx)
				},
				FromColumns: []string{
					ColumnNames.DeepWorkLogs.DayID,
				},
				ToColumns: []string{
					ColumnNames.Days.ID,
				},
			},
		},
	}, Days.Columns().Names(), opts...)
}

func ThenLoadDeepWorkLogDay(queryMods ...bob.Mod[*dialect.SelectQuery]) sqlite.Loader {
	return sqlite.Loader(func(ctx context.Context, exec bob.Executor, retrieved any) error {
		loader, isLoader := retrieved.(interface {
			LoadDeepWorkLogDay(context.Context, bob.Executor, ...bob.Mod[*dialect.SelectQuery]) error
		})
		if !isLoader {
			return fmt.Errorf("object %T cannot load DeepWorkLogDay", retrieved)
		}

		err := loader.LoadDeepWorkLogDay(ctx, exec, queryMods...)

		// Don't cause an issue due to missing relationships
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}

		return err
	})
}

// LoadDeepWorkLogDay loads the deepWorkLog's Day into the .R struct
func (o *DeepWorkLog) LoadDeepWorkLogDay(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if o == nil {
		return nil
	}

	// Reset the relationship
	o.R.Day = nil

	related, err := o.Day(ctx, exec, mods...).One()
	if err != nil {
		return err
	}

	related.R.DeepWorkLogs = DeepWorkLogSlice{o}

	o.R.Day = related
	return nil
}

// LoadDeepWorkLogDay loads the deepWorkLog's Day into the .R struct
func (os DeepWorkLogSlice) LoadDeepWorkLogDay(ctx context.Context, exec bob.Executor, mods ...bob.Mod[*dialect.SelectQuery]) error {
	if len(os) == 0 {
		return nil
	}

	days, err := os.Day(ctx, exec, mods...).All()
	if err != nil {
		return err
	}

	for _, o := range os {
		for _, rel := range days {
			if o.DayID != rel.ID {
				continue
			}

			rel.R.DeepWorkLogs = append(rel.R.DeepWorkLogs, o)

			o.R.Day = rel
			break
		}
	}

	return nil
}

func attachDeepWorkLogDay0(ctx context.Context, exec bob.Executor, count int, deepWorkLog0 *DeepWorkLog, day1 *Day) (*DeepWorkLog, error) {
	setter := &DeepWorkLogSetter{
		DayID: omit.From(day1.ID),
	}

	err := DeepWorkLogs.Update(ctx, exec, setter, deepWorkLog0)
	if err != nil {
		return nil, fmt.Errorf("attachDeepWorkLogDay0: %w", err)
	}

	return deepWorkLog0, nil
}

func (deepWorkLog0 *DeepWorkLog) InsertDay(ctx context.Context, exec bob.Executor, related *DaySetter) error {
	day1, err := Days.Insert(ctx, exec, related)
	if err != nil {
		return fmt.Errorf("inserting related objects: %w", err)
	}

	_, err = attachDeepWorkLogDay0(ctx, exec, 1, deepWorkLog0, day1)
	if err != nil {
		return err
	}

	deepWorkLog0.R.Day = day1

	day1.R.DeepWorkLogs = append(day1.R.DeepWorkLogs, deepWorkLog0)

	return nil
}

func (deepWorkLog0 *DeepWorkLog) AttachDay(ctx context.Context, exec bob.Executor, day1 *Day) error {
	var err error

	_, err = attachDeepWorkLogDay0(ctx, exec, 1, deepWorkLog0, day1)
	if err != nil {
		return err
	}

	deepWorkLog0.R.Day = day1

	day1.R.DeepWorkLogs = append(day1.R.DeepWorkLogs, deepWorkLog0)

	return nil
}