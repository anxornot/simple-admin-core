// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/suyuan32/simple-admin-core/pkg/ent/api"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionarydetail"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menu"
	"github.com/suyuan32/simple-admin-core/pkg/ent/menuparam"
	"github.com/suyuan32/simple-admin-core/pkg/ent/oauthprovider"
	"github.com/suyuan32/simple-admin-core/pkg/ent/role"
	"github.com/suyuan32/simple-admin-core/pkg/ent/token"
	"github.com/suyuan32/simple-admin-core/pkg/ent/user"
)

const errInvalidPage = "INVALID_PAGE"

const (
	listField     = "list"
	pageNumField  = "pageNum"
	pageSizeField = "pageSize"
)

type PageDetails struct {
	Page  uint64 `json:"page"`
	Limit uint64 `json:"limit"`
	Total uint64 `json:"total"`
}

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

// Cursor of an edge type.
type Cursor struct {
	ID    uint64
	Value Value
}

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

type apiPager struct {
	order  *APIOrder
	filter func(*APIQuery) (*APIQuery, error)
}

// APIPaginateOption enables pagination customization.
type APIPaginateOption func(*apiPager) error

// APIOrder defines the ordering of API.
type APIOrder struct {
	Direction OrderDirection `json:"direction"`
	Field     *APIOrderField `json:"field"`
}

// APIOrderField defines the ordering field of API.
type APIOrderField struct {
	field    string
	toCursor func(*API) Cursor
}

// DefaultAPIOrder is the default ordering of API.
var DefaultAPIOrder = &APIOrder{
	Direction: OrderDirectionAsc,
	Field: &APIOrderField{
		field: api.FieldID,
		toCursor: func(a *API) Cursor {
			return Cursor{ID: a.ID}
		},
	},
}

func newAPIPager(opts []APIPaginateOption) (*apiPager, error) {
	pager := &apiPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultAPIOrder
	}
	return pager, nil
}

func (p *apiPager) applyFilter(query *APIQuery) (*APIQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// APIPageList is API PageList result.
type APIPageList struct {
	List        []*API       `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (a *APIQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...APIPaginateOption,
) (*APIPageList, error) {

	pager, err := newAPIPager(opts)
	if err != nil {
		return nil, err
	}

	if a, err = pager.applyFilter(a); err != nil {
		return nil, err
	}

	ret := &APIPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := a.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	direction := pager.order.Direction
	a = a.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultAPIOrder.Field {
		a = a.Order(direction.orderFunc(DefaultAPIOrder.Field.field))
	}

	a = a.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := a.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type dictionaryPager struct {
	order  *DictionaryOrder
	filter func(*DictionaryQuery) (*DictionaryQuery, error)
}

// DictionaryPaginateOption enables pagination customization.
type DictionaryPaginateOption func(*dictionaryPager) error

// DictionaryOrder defines the ordering of Dictionary.
type DictionaryOrder struct {
	Direction OrderDirection        `json:"direction"`
	Field     *DictionaryOrderField `json:"field"`
}

// DictionaryOrderField defines the ordering field of Dictionary.
type DictionaryOrderField struct {
	field    string
	toCursor func(*Dictionary) Cursor
}

// DefaultDictionaryOrder is the default ordering of Dictionary.
var DefaultDictionaryOrder = &DictionaryOrder{
	Direction: OrderDirectionAsc,
	Field: &DictionaryOrderField{
		field: dictionary.FieldID,
		toCursor: func(d *Dictionary) Cursor {
			return Cursor{ID: d.ID}
		},
	},
}

func newDictionaryPager(opts []DictionaryPaginateOption) (*dictionaryPager, error) {
	pager := &dictionaryPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultDictionaryOrder
	}
	return pager, nil
}

func (p *dictionaryPager) applyFilter(query *DictionaryQuery) (*DictionaryQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// DictionaryPageList is Dictionary PageList result.
type DictionaryPageList struct {
	List        []*Dictionary `json:"list"`
	PageDetails *PageDetails  `json:"pageDetails"`
}

func (d *DictionaryQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...DictionaryPaginateOption,
) (*DictionaryPageList, error) {

	pager, err := newDictionaryPager(opts)
	if err != nil {
		return nil, err
	}

	if d, err = pager.applyFilter(d); err != nil {
		return nil, err
	}

	ret := &DictionaryPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := d.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	direction := pager.order.Direction
	d = d.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultDictionaryOrder.Field {
		d = d.Order(direction.orderFunc(DefaultDictionaryOrder.Field.field))
	}

	d = d.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := d.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type dictionarydetailPager struct {
	order  *DictionaryDetailOrder
	filter func(*DictionaryDetailQuery) (*DictionaryDetailQuery, error)
}

// DictionaryDetailPaginateOption enables pagination customization.
type DictionaryDetailPaginateOption func(*dictionarydetailPager) error

// DictionaryDetailOrder defines the ordering of DictionaryDetail.
type DictionaryDetailOrder struct {
	Direction OrderDirection              `json:"direction"`
	Field     *DictionaryDetailOrderField `json:"field"`
}

// DictionaryDetailOrderField defines the ordering field of DictionaryDetail.
type DictionaryDetailOrderField struct {
	field    string
	toCursor func(*DictionaryDetail) Cursor
}

// DefaultDictionaryDetailOrder is the default ordering of DictionaryDetail.
var DefaultDictionaryDetailOrder = &DictionaryDetailOrder{
	Direction: OrderDirectionAsc,
	Field: &DictionaryDetailOrderField{
		field: dictionarydetail.FieldID,
		toCursor: func(dd *DictionaryDetail) Cursor {
			return Cursor{ID: dd.ID}
		},
	},
}

func newDictionaryDetailPager(opts []DictionaryDetailPaginateOption) (*dictionarydetailPager, error) {
	pager := &dictionarydetailPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultDictionaryDetailOrder
	}
	return pager, nil
}

func (p *dictionarydetailPager) applyFilter(query *DictionaryDetailQuery) (*DictionaryDetailQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// DictionaryDetailPageList is DictionaryDetail PageList result.
type DictionaryDetailPageList struct {
	List        []*DictionaryDetail `json:"list"`
	PageDetails *PageDetails        `json:"pageDetails"`
}

func (dd *DictionaryDetailQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...DictionaryDetailPaginateOption,
) (*DictionaryDetailPageList, error) {

	pager, err := newDictionaryDetailPager(opts)
	if err != nil {
		return nil, err
	}

	if dd, err = pager.applyFilter(dd); err != nil {
		return nil, err
	}

	ret := &DictionaryDetailPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := dd.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	direction := pager.order.Direction
	dd = dd.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultDictionaryDetailOrder.Field {
		dd = dd.Order(direction.orderFunc(DefaultDictionaryDetailOrder.Field.field))
	}

	dd = dd.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := dd.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type menuPager struct {
	order  *MenuOrder
	filter func(*MenuQuery) (*MenuQuery, error)
}

// MenuPaginateOption enables pagination customization.
type MenuPaginateOption func(*menuPager) error

// MenuOrder defines the ordering of Menu.
type MenuOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *MenuOrderField `json:"field"`
}

// MenuOrderField defines the ordering field of Menu.
type MenuOrderField struct {
	field    string
	toCursor func(*Menu) Cursor
}

// DefaultMenuOrder is the default ordering of Menu.
var DefaultMenuOrder = &MenuOrder{
	Direction: OrderDirectionAsc,
	Field: &MenuOrderField{
		field: menu.FieldID,
		toCursor: func(m *Menu) Cursor {
			return Cursor{ID: m.ID}
		},
	},
}

func newMenuPager(opts []MenuPaginateOption) (*menuPager, error) {
	pager := &menuPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultMenuOrder
	}
	return pager, nil
}

func (p *menuPager) applyFilter(query *MenuQuery) (*MenuQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// MenuPageList is Menu PageList result.
type MenuPageList struct {
	List        []*Menu      `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (m *MenuQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...MenuPaginateOption,
) (*MenuPageList, error) {

	pager, err := newMenuPager(opts)
	if err != nil {
		return nil, err
	}

	if m, err = pager.applyFilter(m); err != nil {
		return nil, err
	}

	ret := &MenuPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := m.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	direction := pager.order.Direction
	m = m.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultMenuOrder.Field {
		m = m.Order(direction.orderFunc(DefaultMenuOrder.Field.field))
	}

	m = m.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := m.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type menuparamPager struct {
	order  *MenuParamOrder
	filter func(*MenuParamQuery) (*MenuParamQuery, error)
}

// MenuParamPaginateOption enables pagination customization.
type MenuParamPaginateOption func(*menuparamPager) error

// MenuParamOrder defines the ordering of MenuParam.
type MenuParamOrder struct {
	Direction OrderDirection       `json:"direction"`
	Field     *MenuParamOrderField `json:"field"`
}

// MenuParamOrderField defines the ordering field of MenuParam.
type MenuParamOrderField struct {
	field    string
	toCursor func(*MenuParam) Cursor
}

// DefaultMenuParamOrder is the default ordering of MenuParam.
var DefaultMenuParamOrder = &MenuParamOrder{
	Direction: OrderDirectionAsc,
	Field: &MenuParamOrderField{
		field: menuparam.FieldID,
		toCursor: func(mp *MenuParam) Cursor {
			return Cursor{ID: mp.ID}
		},
	},
}

func newMenuParamPager(opts []MenuParamPaginateOption) (*menuparamPager, error) {
	pager := &menuparamPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultMenuParamOrder
	}
	return pager, nil
}

func (p *menuparamPager) applyFilter(query *MenuParamQuery) (*MenuParamQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// MenuParamPageList is MenuParam PageList result.
type MenuParamPageList struct {
	List        []*MenuParam `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (mp *MenuParamQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...MenuParamPaginateOption,
) (*MenuParamPageList, error) {

	pager, err := newMenuParamPager(opts)
	if err != nil {
		return nil, err
	}

	if mp, err = pager.applyFilter(mp); err != nil {
		return nil, err
	}

	ret := &MenuParamPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := mp.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	direction := pager.order.Direction
	mp = mp.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultMenuParamOrder.Field {
		mp = mp.Order(direction.orderFunc(DefaultMenuParamOrder.Field.field))
	}

	mp = mp.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := mp.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type oauthproviderPager struct {
	order  *OauthProviderOrder
	filter func(*OauthProviderQuery) (*OauthProviderQuery, error)
}

// OauthProviderPaginateOption enables pagination customization.
type OauthProviderPaginateOption func(*oauthproviderPager) error

// OauthProviderOrder defines the ordering of OauthProvider.
type OauthProviderOrder struct {
	Direction OrderDirection           `json:"direction"`
	Field     *OauthProviderOrderField `json:"field"`
}

// OauthProviderOrderField defines the ordering field of OauthProvider.
type OauthProviderOrderField struct {
	field    string
	toCursor func(*OauthProvider) Cursor
}

// DefaultOauthProviderOrder is the default ordering of OauthProvider.
var DefaultOauthProviderOrder = &OauthProviderOrder{
	Direction: OrderDirectionAsc,
	Field: &OauthProviderOrderField{
		field: oauthprovider.FieldID,
		toCursor: func(op *OauthProvider) Cursor {
			return Cursor{ID: op.ID}
		},
	},
}

func newOauthProviderPager(opts []OauthProviderPaginateOption) (*oauthproviderPager, error) {
	pager := &oauthproviderPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultOauthProviderOrder
	}
	return pager, nil
}

func (p *oauthproviderPager) applyFilter(query *OauthProviderQuery) (*OauthProviderQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// OauthProviderPageList is OauthProvider PageList result.
type OauthProviderPageList struct {
	List        []*OauthProvider `json:"list"`
	PageDetails *PageDetails     `json:"pageDetails"`
}

func (op *OauthProviderQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...OauthProviderPaginateOption,
) (*OauthProviderPageList, error) {

	pager, err := newOauthProviderPager(opts)
	if err != nil {
		return nil, err
	}

	if op, err = pager.applyFilter(op); err != nil {
		return nil, err
	}

	ret := &OauthProviderPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := op.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	direction := pager.order.Direction
	op = op.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultOauthProviderOrder.Field {
		op = op.Order(direction.orderFunc(DefaultOauthProviderOrder.Field.field))
	}

	op = op.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := op.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type rolePager struct {
	order  *RoleOrder
	filter func(*RoleQuery) (*RoleQuery, error)
}

// RolePaginateOption enables pagination customization.
type RolePaginateOption func(*rolePager) error

// RoleOrder defines the ordering of Role.
type RoleOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *RoleOrderField `json:"field"`
}

// RoleOrderField defines the ordering field of Role.
type RoleOrderField struct {
	field    string
	toCursor func(*Role) Cursor
}

// DefaultRoleOrder is the default ordering of Role.
var DefaultRoleOrder = &RoleOrder{
	Direction: OrderDirectionAsc,
	Field: &RoleOrderField{
		field: role.FieldID,
		toCursor: func(r *Role) Cursor {
			return Cursor{ID: r.ID}
		},
	},
}

func newRolePager(opts []RolePaginateOption) (*rolePager, error) {
	pager := &rolePager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultRoleOrder
	}
	return pager, nil
}

func (p *rolePager) applyFilter(query *RoleQuery) (*RoleQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// RolePageList is Role PageList result.
type RolePageList struct {
	List        []*Role      `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (r *RoleQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...RolePaginateOption,
) (*RolePageList, error) {

	pager, err := newRolePager(opts)
	if err != nil {
		return nil, err
	}

	if r, err = pager.applyFilter(r); err != nil {
		return nil, err
	}

	ret := &RolePageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := r.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	direction := pager.order.Direction
	r = r.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultRoleOrder.Field {
		r = r.Order(direction.orderFunc(DefaultRoleOrder.Field.field))
	}

	r = r.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := r.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type tokenPager struct {
	order  *TokenOrder
	filter func(*TokenQuery) (*TokenQuery, error)
}

// TokenPaginateOption enables pagination customization.
type TokenPaginateOption func(*tokenPager) error

// TokenOrder defines the ordering of Token.
type TokenOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *TokenOrderField `json:"field"`
}

// TokenOrderField defines the ordering field of Token.
type TokenOrderField struct {
	field    string
	toCursor func(*Token) Cursor
}

// DefaultTokenOrder is the default ordering of Token.
var DefaultTokenOrder = &TokenOrder{
	Direction: OrderDirectionAsc,
	Field: &TokenOrderField{
		field: token.FieldID,
		toCursor: func(t *Token) Cursor {
			return Cursor{ID: t.ID}
		},
	},
}

func newTokenPager(opts []TokenPaginateOption) (*tokenPager, error) {
	pager := &tokenPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultTokenOrder
	}
	return pager, nil
}

func (p *tokenPager) applyFilter(query *TokenQuery) (*TokenQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// TokenPageList is Token PageList result.
type TokenPageList struct {
	List        []*Token     `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (t *TokenQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...TokenPaginateOption,
) (*TokenPageList, error) {

	pager, err := newTokenPager(opts)
	if err != nil {
		return nil, err
	}

	if t, err = pager.applyFilter(t); err != nil {
		return nil, err
	}

	ret := &TokenPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := t.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	direction := pager.order.Direction
	t = t.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultTokenOrder.Field {
		t = t.Order(direction.orderFunc(DefaultTokenOrder.Field.field))
	}

	t = t.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := t.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type userPager struct {
	order  *UserOrder
	filter func(*UserQuery) (*UserQuery, error)
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	field    string
	toCursor func(*User) Cursor
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: OrderDirectionAsc,
	Field: &UserOrderField{
		field: user.FieldID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

func newUserPager(opts []UserPaginateOption) (*userPager, error) {
	pager := &userPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

// UserPageList is User PageList result.
type UserPageList struct {
	List        []*User      `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (u *UserQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...UserPaginateOption,
) (*UserPageList, error) {

	pager, err := newUserPager(opts)
	if err != nil {
		return nil, err
	}

	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}

	ret := &UserPageList{}

	ret.PageDetails = &PageDetails{
		Page:  pageNum,
		Limit: pageSize,
	}

	count, err := u.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	direction := pager.order.Direction
	u = u.Order(direction.orderFunc(pager.order.Field.field))
	if pager.order.Field != DefaultUserOrder.Field {
		u = u.Order(direction.orderFunc(DefaultUserOrder.Field.field))
	}

	u = u.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}