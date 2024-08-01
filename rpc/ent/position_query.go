// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/position"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/ent/user"
)

// PositionQuery is the builder for querying Position entities.
type PositionQuery struct {
	config
	ctx        *QueryContext
	order      []position.OrderOption
	inters     []Interceptor
	predicates []predicate.Position
	withUsers  *UserQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PositionQuery builder.
func (pq *PositionQuery) Where(ps ...predicate.Position) *PositionQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PositionQuery) Limit(limit int) *PositionQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PositionQuery) Offset(offset int) *PositionQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PositionQuery) Unique(unique bool) *PositionQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PositionQuery) Order(o ...position.OrderOption) *PositionQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryUsers chains the current query on the "users" edge.
func (pq *PositionQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(position.Table, position.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, position.UsersTable, position.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Position entity from the query.
// Returns a *NotFoundError when no Position was found.
func (pq *PositionQuery) First(ctx context.Context) (*Position, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{position.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PositionQuery) FirstX(ctx context.Context) *Position {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Position ID from the query.
// Returns a *NotFoundError when no Position ID was found.
func (pq *PositionQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{position.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PositionQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Position entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Position entity is found.
// Returns a *NotFoundError when no Position entities are found.
func (pq *PositionQuery) Only(ctx context.Context) (*Position, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{position.Label}
	default:
		return nil, &NotSingularError{position.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PositionQuery) OnlyX(ctx context.Context) *Position {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Position ID in the query.
// Returns a *NotSingularError when more than one Position ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PositionQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{position.Label}
	default:
		err = &NotSingularError{position.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PositionQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Positions.
func (pq *PositionQuery) All(ctx context.Context) ([]*Position, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Position, *PositionQuery]()
	return withInterceptors[[]*Position](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PositionQuery) AllX(ctx context.Context) []*Position {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Position IDs.
func (pq *PositionQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryIDs)
	if err = pq.Select(position.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PositionQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PositionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PositionQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PositionQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PositionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PositionQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PositionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PositionQuery) Clone() *PositionQuery {
	if pq == nil {
		return nil
	}
	return &PositionQuery{
		config:     pq.config,
		ctx:        pq.ctx.Clone(),
		order:      append([]position.OrderOption{}, pq.order...),
		inters:     append([]Interceptor{}, pq.inters...),
		predicates: append([]predicate.Position{}, pq.predicates...),
		withUsers:  pq.withUsers.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PositionQuery) WithUsers(opts ...func(*UserQuery)) *PositionQuery {
	query := (&UserClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withUsers = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Position.Query().
//		GroupBy(position.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PositionQuery) GroupBy(field string, fields ...string) *PositionGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PositionGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = position.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Position.Query().
//		Select(position.FieldCreatedAt).
//		Scan(ctx, &v)
func (pq *PositionQuery) Select(fields ...string) *PositionSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PositionSelect{PositionQuery: pq}
	sbuild.label = position.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PositionSelect configured with the given aggregations.
func (pq *PositionQuery) Aggregate(fns ...AggregateFunc) *PositionSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PositionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !position.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PositionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Position, error) {
	var (
		nodes       = []*Position{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Position).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Position{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withUsers; query != nil {
		if err := pq.loadUsers(ctx, query, nodes,
			func(n *Position) { n.Edges.Users = []*User{} },
			func(n *Position, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PositionQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Position, init func(*Position), assign func(*Position, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uint64]*Position)
	nids := make(map[uuid.UUID]map[*Position]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(position.UsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(position.UsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(position.UsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(position.UsersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := uint64(values[0].(*sql.NullInt64).Int64)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Position]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (pq *PositionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PositionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(position.Table, position.Columns, sqlgraph.NewFieldSpec(position.FieldID, field.TypeUint64))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, position.FieldID)
		for i := range fields {
			if fields[i] != position.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PositionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(position.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = position.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range pq.modifiers {
		m(selector)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pq *PositionQuery) Modify(modifiers ...func(s *sql.Selector)) *PositionSelect {
	pq.modifiers = append(pq.modifiers, modifiers...)
	return pq.Select()
}

// PositionGroupBy is the group-by builder for Position entities.
type PositionGroupBy struct {
	selector
	build *PositionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PositionGroupBy) Aggregate(fns ...AggregateFunc) *PositionGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PositionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PositionQuery, *PositionGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PositionGroupBy) sqlScan(ctx context.Context, root *PositionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PositionSelect is the builder for selecting fields of Position entities.
type PositionSelect struct {
	*PositionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PositionSelect) Aggregate(fns ...AggregateFunc) *PositionSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PositionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PositionQuery, *PositionSelect](ctx, ps.PositionQuery, ps, ps.inters, v)
}

func (ps *PositionSelect) sqlScan(ctx context.Context, root *PositionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ps *PositionSelect) Modify(modifiers ...func(s *sql.Selector)) *PositionSelect {
	ps.modifiers = append(ps.modifiers, modifiers...)
	return ps
}
