// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// WSSToken is an object representing the database table.
type WSSToken struct {
	Token      string    `boil:"token" json:"token" toml:"token" yaml:"token"`
	ValidUntil time.Time `boil:"valid_until" json:"valid_until" toml:"valid_until" yaml:"valid_until"`
	UserID     string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *wssTokenR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L wssTokenL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WSSTokenColumns = struct {
	Token      string
	ValidUntil string
	UserID     string
	CreatedAt  string
	UpdatedAt  string
}{
	Token:      "token",
	ValidUntil: "valid_until",
	UserID:     "user_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

var WSSTokenTableColumns = struct {
	Token      string
	ValidUntil string
	UserID     string
	CreatedAt  string
	UpdatedAt  string
}{
	Token:      "wss_tokens.token",
	ValidUntil: "wss_tokens.valid_until",
	UserID:     "wss_tokens.user_id",
	CreatedAt:  "wss_tokens.created_at",
	UpdatedAt:  "wss_tokens.updated_at",
}

// Generated where

var WSSTokenWhere = struct {
	Token      whereHelperstring
	ValidUntil whereHelpertime_Time
	UserID     whereHelperstring
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
}{
	Token:      whereHelperstring{field: "\"wss_tokens\".\"token\""},
	ValidUntil: whereHelpertime_Time{field: "\"wss_tokens\".\"valid_until\""},
	UserID:     whereHelperstring{field: "\"wss_tokens\".\"user_id\""},
	CreatedAt:  whereHelpertime_Time{field: "\"wss_tokens\".\"created_at\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"wss_tokens\".\"updated_at\""},
}

// WSSTokenRels is where relationship names are stored.
var WSSTokenRels = struct {
	User string
}{
	User: "User",
}

// wssTokenR is where relationships are stored.
type wssTokenR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*wssTokenR) NewStruct() *wssTokenR {
	return &wssTokenR{}
}

// wssTokenL is where Load methods for each relationship are stored.
type wssTokenL struct{}

var (
	wssTokenAllColumns            = []string{"token", "valid_until", "user_id", "created_at", "updated_at"}
	wssTokenColumnsWithoutDefault = []string{"valid_until", "user_id", "created_at", "updated_at"}
	wssTokenColumnsWithDefault    = []string{"token"}
	wssTokenPrimaryKeyColumns     = []string{"token"}
)

type (
	// WSSTokenSlice is an alias for a slice of pointers to WSSToken.
	// This should almost always be used instead of []WSSToken.
	WSSTokenSlice []*WSSToken

	wssTokenQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	wssTokenType                 = reflect.TypeOf(&WSSToken{})
	wssTokenMapping              = queries.MakeStructMapping(wssTokenType)
	wssTokenPrimaryKeyMapping, _ = queries.BindMapping(wssTokenType, wssTokenMapping, wssTokenPrimaryKeyColumns)
	wssTokenInsertCacheMut       sync.RWMutex
	wssTokenInsertCache          = make(map[string]insertCache)
	wssTokenUpdateCacheMut       sync.RWMutex
	wssTokenUpdateCache          = make(map[string]updateCache)
	wssTokenUpsertCacheMut       sync.RWMutex
	wssTokenUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single wssToken record from the query.
func (q wssTokenQuery) One(ctx context.Context, exec boil.ContextExecutor) (*WSSToken, error) {
	o := &WSSToken{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for wss_tokens")
	}

	return o, nil
}

// All returns all WSSToken records from the query.
func (q wssTokenQuery) All(ctx context.Context, exec boil.ContextExecutor) (WSSTokenSlice, error) {
	var o []*WSSToken

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to WSSToken slice")
	}

	return o, nil
}

// Count returns the count of all WSSToken records in the query.
func (q wssTokenQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count wss_tokens rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q wssTokenQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if wss_tokens exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *WSSToken) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (wssTokenL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeWSSToken interface{}, mods queries.Applicator) error {
	var slice []*WSSToken
	var object *WSSToken

	if singular {
		object = maybeWSSToken.(*WSSToken)
	} else {
		slice = *maybeWSSToken.(*[]*WSSToken)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &wssTokenR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &wssTokenR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.WSSTokens = append(foreign.R.WSSTokens, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.WSSTokens = append(foreign.R.WSSTokens, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the wssToken to the related item.
// Sets o.R.User to related.
// Adds o to related.R.WSSTokens.
func (o *WSSToken) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"wss_tokens\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, wssTokenPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.Token}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &wssTokenR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			WSSTokens: WSSTokenSlice{o},
		}
	} else {
		related.R.WSSTokens = append(related.R.WSSTokens, o)
	}

	return nil
}

// WSSTokens retrieves all the records using an executor.
func WSSTokens(mods ...qm.QueryMod) wssTokenQuery {
	mods = append(mods, qm.From("\"wss_tokens\""))
	return wssTokenQuery{NewQuery(mods...)}
}

// FindWSSToken retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWSSToken(ctx context.Context, exec boil.ContextExecutor, token string, selectCols ...string) (*WSSToken, error) {
	wssTokenObj := &WSSToken{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"wss_tokens\" where \"token\"=$1", sel,
	)

	q := queries.Raw(query, token)

	err := q.Bind(ctx, exec, wssTokenObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from wss_tokens")
	}

	return wssTokenObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *WSSToken) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no wss_tokens provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(wssTokenColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	wssTokenInsertCacheMut.RLock()
	cache, cached := wssTokenInsertCache[key]
	wssTokenInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			wssTokenAllColumns,
			wssTokenColumnsWithDefault,
			wssTokenColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(wssTokenType, wssTokenMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(wssTokenType, wssTokenMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"wss_tokens\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"wss_tokens\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into wss_tokens")
	}

	if !cached {
		wssTokenInsertCacheMut.Lock()
		wssTokenInsertCache[key] = cache
		wssTokenInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the WSSToken.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *WSSToken) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	wssTokenUpdateCacheMut.RLock()
	cache, cached := wssTokenUpdateCache[key]
	wssTokenUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			wssTokenAllColumns,
			wssTokenPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update wss_tokens, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"wss_tokens\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, wssTokenPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(wssTokenType, wssTokenMapping, append(wl, wssTokenPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update wss_tokens row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for wss_tokens")
	}

	if !cached {
		wssTokenUpdateCacheMut.Lock()
		wssTokenUpdateCache[key] = cache
		wssTokenUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q wssTokenQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for wss_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for wss_tokens")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WSSTokenSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), wssTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"wss_tokens\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, wssTokenPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in wssToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all wssToken")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *WSSToken) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no wss_tokens provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(wssTokenColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	wssTokenUpsertCacheMut.RLock()
	cache, cached := wssTokenUpsertCache[key]
	wssTokenUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			wssTokenAllColumns,
			wssTokenColumnsWithDefault,
			wssTokenColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			wssTokenAllColumns,
			wssTokenPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert wss_tokens, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(wssTokenPrimaryKeyColumns))
			copy(conflict, wssTokenPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"wss_tokens\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(wssTokenType, wssTokenMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(wssTokenType, wssTokenMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert wss_tokens")
	}

	if !cached {
		wssTokenUpsertCacheMut.Lock()
		wssTokenUpsertCache[key] = cache
		wssTokenUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single WSSToken record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *WSSToken) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no WSSToken provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), wssTokenPrimaryKeyMapping)
	sql := "DELETE FROM \"wss_tokens\" WHERE \"token\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from wss_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for wss_tokens")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q wssTokenQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no wssTokenQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from wss_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for wss_tokens")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WSSTokenSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), wssTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"wss_tokens\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, wssTokenPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from wssToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for wss_tokens")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *WSSToken) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindWSSToken(ctx, exec, o.Token)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WSSTokenSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WSSTokenSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), wssTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"wss_tokens\".* FROM \"wss_tokens\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, wssTokenPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in WSSTokenSlice")
	}

	*o = slice

	return nil
}

// WSSTokenExists checks if the WSSToken row exists.
func WSSTokenExists(ctx context.Context, exec boil.ContextExecutor, token string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"wss_tokens\" where \"token\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, token)
	}
	row := exec.QueryRowContext(ctx, sql, token)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if wss_tokens exists")
	}

	return exists, nil
}
