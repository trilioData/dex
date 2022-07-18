// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dexidp/dex/v2/storage/ent/db/devicerequest"
	"github.com/dexidp/dex/v2/storage/ent/db/predicate"
)

// DeviceRequestQuery is the builder for querying DeviceRequest entities.
type DeviceRequestQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DeviceRequest
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeviceRequestQuery builder.
func (drq *DeviceRequestQuery) Where(ps ...predicate.DeviceRequest) *DeviceRequestQuery {
	drq.predicates = append(drq.predicates, ps...)
	return drq
}

// Limit adds a limit step to the query.
func (drq *DeviceRequestQuery) Limit(limit int) *DeviceRequestQuery {
	drq.limit = &limit
	return drq
}

// Offset adds an offset step to the query.
func (drq *DeviceRequestQuery) Offset(offset int) *DeviceRequestQuery {
	drq.offset = &offset
	return drq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (drq *DeviceRequestQuery) Unique(unique bool) *DeviceRequestQuery {
	drq.unique = &unique
	return drq
}

// Order adds an order step to the query.
func (drq *DeviceRequestQuery) Order(o ...OrderFunc) *DeviceRequestQuery {
	drq.order = append(drq.order, o...)
	return drq
}

// First returns the first DeviceRequest entity from the query.
// Returns a *NotFoundError when no DeviceRequest was found.
func (drq *DeviceRequestQuery) First(ctx context.Context) (*DeviceRequest, error) {
	nodes, err := drq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{devicerequest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (drq *DeviceRequestQuery) FirstX(ctx context.Context) *DeviceRequest {
	node, err := drq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DeviceRequest ID from the query.
// Returns a *NotFoundError when no DeviceRequest ID was found.
func (drq *DeviceRequestQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = drq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{devicerequest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (drq *DeviceRequestQuery) FirstIDX(ctx context.Context) int {
	id, err := drq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DeviceRequest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DeviceRequest entity is found.
// Returns a *NotFoundError when no DeviceRequest entities are found.
func (drq *DeviceRequestQuery) Only(ctx context.Context) (*DeviceRequest, error) {
	nodes, err := drq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{devicerequest.Label}
	default:
		return nil, &NotSingularError{devicerequest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (drq *DeviceRequestQuery) OnlyX(ctx context.Context) *DeviceRequest {
	node, err := drq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DeviceRequest ID in the query.
// Returns a *NotSingularError when more than one DeviceRequest ID is found.
// Returns a *NotFoundError when no entities are found.
func (drq *DeviceRequestQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = drq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{devicerequest.Label}
	default:
		err = &NotSingularError{devicerequest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (drq *DeviceRequestQuery) OnlyIDX(ctx context.Context) int {
	id, err := drq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeviceRequests.
func (drq *DeviceRequestQuery) All(ctx context.Context) ([]*DeviceRequest, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return drq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (drq *DeviceRequestQuery) AllX(ctx context.Context) []*DeviceRequest {
	nodes, err := drq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DeviceRequest IDs.
func (drq *DeviceRequestQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := drq.Select(devicerequest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (drq *DeviceRequestQuery) IDsX(ctx context.Context) []int {
	ids, err := drq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (drq *DeviceRequestQuery) Count(ctx context.Context) (int, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return drq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (drq *DeviceRequestQuery) CountX(ctx context.Context) int {
	count, err := drq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (drq *DeviceRequestQuery) Exist(ctx context.Context) (bool, error) {
	if err := drq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return drq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (drq *DeviceRequestQuery) ExistX(ctx context.Context) bool {
	exist, err := drq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeviceRequestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (drq *DeviceRequestQuery) Clone() *DeviceRequestQuery {
	if drq == nil {
		return nil
	}
	return &DeviceRequestQuery{
		config:     drq.config,
		limit:      drq.limit,
		offset:     drq.offset,
		order:      append([]OrderFunc{}, drq.order...),
		predicates: append([]predicate.DeviceRequest{}, drq.predicates...),
		// clone intermediate query.
		sql:    drq.sql.Clone(),
		path:   drq.path,
		unique: drq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserCode string `json:"user_code,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DeviceRequest.Query().
//		GroupBy(devicerequest.FieldUserCode).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
//
func (drq *DeviceRequestQuery) GroupBy(field string, fields ...string) *DeviceRequestGroupBy {
	group := &DeviceRequestGroupBy{config: drq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := drq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return drq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserCode string `json:"user_code,omitempty"`
//	}
//
//	client.DeviceRequest.Query().
//		Select(devicerequest.FieldUserCode).
//		Scan(ctx, &v)
//
func (drq *DeviceRequestQuery) Select(fields ...string) *DeviceRequestSelect {
	drq.fields = append(drq.fields, fields...)
	return &DeviceRequestSelect{DeviceRequestQuery: drq}
}

func (drq *DeviceRequestQuery) prepareQuery(ctx context.Context) error {
	for _, f := range drq.fields {
		if !devicerequest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if drq.path != nil {
		prev, err := drq.path(ctx)
		if err != nil {
			return err
		}
		drq.sql = prev
	}
	return nil
}

func (drq *DeviceRequestQuery) sqlAll(ctx context.Context) ([]*DeviceRequest, error) {
	var (
		nodes = []*DeviceRequest{}
		_spec = drq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DeviceRequest{config: drq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("db: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, drq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (drq *DeviceRequestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := drq.querySpec()
	_spec.Node.Columns = drq.fields
	if len(drq.fields) > 0 {
		_spec.Unique = drq.unique != nil && *drq.unique
	}
	return sqlgraph.CountNodes(ctx, drq.driver, _spec)
}

func (drq *DeviceRequestQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := drq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("db: check existence: %w", err)
	}
	return n > 0, nil
}

func (drq *DeviceRequestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   devicerequest.Table,
			Columns: devicerequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: devicerequest.FieldID,
			},
		},
		From:   drq.sql,
		Unique: true,
	}
	if unique := drq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := drq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, devicerequest.FieldID)
		for i := range fields {
			if fields[i] != devicerequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := drq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := drq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := drq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := drq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (drq *DeviceRequestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(drq.driver.Dialect())
	t1 := builder.Table(devicerequest.Table)
	columns := drq.fields
	if len(columns) == 0 {
		columns = devicerequest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if drq.sql != nil {
		selector = drq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if drq.unique != nil && *drq.unique {
		selector.Distinct()
	}
	for _, p := range drq.predicates {
		p(selector)
	}
	for _, p := range drq.order {
		p(selector)
	}
	if offset := drq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := drq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeviceRequestGroupBy is the group-by builder for DeviceRequest entities.
type DeviceRequestGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (drgb *DeviceRequestGroupBy) Aggregate(fns ...AggregateFunc) *DeviceRequestGroupBy {
	drgb.fns = append(drgb.fns, fns...)
	return drgb
}

// Scan applies the group-by query and scans the result into the given value.
func (drgb *DeviceRequestGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := drgb.path(ctx)
	if err != nil {
		return err
	}
	drgb.sql = query
	return drgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (drgb *DeviceRequestGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := drgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (drgb *DeviceRequestGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("db: DeviceRequestGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (drgb *DeviceRequestGroupBy) StringsX(ctx context.Context) []string {
	v, err := drgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (drgb *DeviceRequestGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = drgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicerequest.Label}
	default:
		err = fmt.Errorf("db: DeviceRequestGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (drgb *DeviceRequestGroupBy) StringX(ctx context.Context) string {
	v, err := drgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (drgb *DeviceRequestGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("db: DeviceRequestGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (drgb *DeviceRequestGroupBy) IntsX(ctx context.Context) []int {
	v, err := drgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (drgb *DeviceRequestGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = drgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicerequest.Label}
	default:
		err = fmt.Errorf("db: DeviceRequestGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (drgb *DeviceRequestGroupBy) IntX(ctx context.Context) int {
	v, err := drgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (drgb *DeviceRequestGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("db: DeviceRequestGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (drgb *DeviceRequestGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := drgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (drgb *DeviceRequestGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = drgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicerequest.Label}
	default:
		err = fmt.Errorf("db: DeviceRequestGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (drgb *DeviceRequestGroupBy) Float64X(ctx context.Context) float64 {
	v, err := drgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (drgb *DeviceRequestGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(drgb.fields) > 1 {
		return nil, errors.New("db: DeviceRequestGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := drgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (drgb *DeviceRequestGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := drgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (drgb *DeviceRequestGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = drgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicerequest.Label}
	default:
		err = fmt.Errorf("db: DeviceRequestGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (drgb *DeviceRequestGroupBy) BoolX(ctx context.Context) bool {
	v, err := drgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (drgb *DeviceRequestGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range drgb.fields {
		if !devicerequest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := drgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := drgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (drgb *DeviceRequestGroupBy) sqlQuery() *sql.Selector {
	selector := drgb.sql.Select()
	aggregation := make([]string, 0, len(drgb.fns))
	for _, fn := range drgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(drgb.fields)+len(drgb.fns))
		for _, f := range drgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(drgb.fields...)...)
}

// DeviceRequestSelect is the builder for selecting fields of DeviceRequest entities.
type DeviceRequestSelect struct {
	*DeviceRequestQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (drs *DeviceRequestSelect) Scan(ctx context.Context, v interface{}) error {
	if err := drs.prepareQuery(ctx); err != nil {
		return err
	}
	drs.sql = drs.DeviceRequestQuery.sqlQuery(ctx)
	return drs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (drs *DeviceRequestSelect) ScanX(ctx context.Context, v interface{}) {
	if err := drs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (drs *DeviceRequestSelect) Strings(ctx context.Context) ([]string, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("db: DeviceRequestSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (drs *DeviceRequestSelect) StringsX(ctx context.Context) []string {
	v, err := drs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (drs *DeviceRequestSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = drs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicerequest.Label}
	default:
		err = fmt.Errorf("db: DeviceRequestSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (drs *DeviceRequestSelect) StringX(ctx context.Context) string {
	v, err := drs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (drs *DeviceRequestSelect) Ints(ctx context.Context) ([]int, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("db: DeviceRequestSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (drs *DeviceRequestSelect) IntsX(ctx context.Context) []int {
	v, err := drs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (drs *DeviceRequestSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = drs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicerequest.Label}
	default:
		err = fmt.Errorf("db: DeviceRequestSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (drs *DeviceRequestSelect) IntX(ctx context.Context) int {
	v, err := drs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (drs *DeviceRequestSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("db: DeviceRequestSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (drs *DeviceRequestSelect) Float64sX(ctx context.Context) []float64 {
	v, err := drs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (drs *DeviceRequestSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = drs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicerequest.Label}
	default:
		err = fmt.Errorf("db: DeviceRequestSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (drs *DeviceRequestSelect) Float64X(ctx context.Context) float64 {
	v, err := drs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (drs *DeviceRequestSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(drs.fields) > 1 {
		return nil, errors.New("db: DeviceRequestSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := drs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (drs *DeviceRequestSelect) BoolsX(ctx context.Context) []bool {
	v, err := drs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (drs *DeviceRequestSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = drs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{devicerequest.Label}
	default:
		err = fmt.Errorf("db: DeviceRequestSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (drs *DeviceRequestSelect) BoolX(ctx context.Context) bool {
	v, err := drs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (drs *DeviceRequestSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := drs.sql.Query()
	if err := drs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
