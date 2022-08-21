// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"context"
	"errors"
	"fmt"
	"image/color"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/MichaelMure/git-bug/api/graphql/models"
	"github.com/MichaelMure/git-bug/entities/bug"
	"github.com/MichaelMure/git-bug/repository"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type ColorResolver interface {
	R(ctx context.Context, obj *color.RGBA) (int, error)
	G(ctx context.Context, obj *color.RGBA) (int, error)
	B(ctx context.Context, obj *color.RGBA) (int, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Color_R(ctx context.Context, field graphql.CollectedField, obj *color.RGBA) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Color_R(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Color().R(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Color_R(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Color",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Color_G(ctx context.Context, field graphql.CollectedField, obj *color.RGBA) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Color_G(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Color().G(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Color_G(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Color",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Color_B(ctx context.Context, field graphql.CollectedField, obj *color.RGBA) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Color_B(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Color().B(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Color_B(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Color",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _PageInfo_hasNextPage(ctx context.Context, field graphql.CollectedField, obj *models.PageInfo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PageInfo_hasNextPage(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.HasNextPage, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PageInfo_hasNextPage(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PageInfo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _PageInfo_hasPreviousPage(ctx context.Context, field graphql.CollectedField, obj *models.PageInfo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PageInfo_hasPreviousPage(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.HasPreviousPage, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bool)
	fc.Result = res
	return ec.marshalNBoolean2bool(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PageInfo_hasPreviousPage(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PageInfo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Boolean does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _PageInfo_startCursor(ctx context.Context, field graphql.CollectedField, obj *models.PageInfo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PageInfo_startCursor(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.StartCursor, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PageInfo_startCursor(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PageInfo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _PageInfo_endCursor(ctx context.Context, field graphql.CollectedField, obj *models.PageInfo) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_PageInfo_endCursor(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.EndCursor, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_PageInfo_endCursor(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "PageInfo",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

func (ec *executionContext) _Authored(ctx context.Context, sel ast.SelectionSet, obj models.Authored) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case bug.Comment:
		return ec._Comment(ctx, sel, &obj)
	case *bug.Comment:
		if obj == nil {
			return graphql.Null
		}
		return ec._Comment(ctx, sel, obj)
	case models.BugWrapper:
		if obj == nil {
			return graphql.Null
		}
		return ec._Bug(ctx, sel, obj)
	case *bug.CreateOperation:
		if obj == nil {
			return graphql.Null
		}
		return ec._CreateOperation(ctx, sel, obj)
	case *bug.SetTitleOperation:
		if obj == nil {
			return graphql.Null
		}
		return ec._SetTitleOperation(ctx, sel, obj)
	case *bug.AddCommentOperation:
		if obj == nil {
			return graphql.Null
		}
		return ec._AddCommentOperation(ctx, sel, obj)
	case *bug.EditCommentOperation:
		if obj == nil {
			return graphql.Null
		}
		return ec._EditCommentOperation(ctx, sel, obj)
	case *bug.SetStatusOperation:
		if obj == nil {
			return graphql.Null
		}
		return ec._SetStatusOperation(ctx, sel, obj)
	case *bug.LabelChangeOperation:
		if obj == nil {
			return graphql.Null
		}
		return ec._LabelChangeOperation(ctx, sel, obj)
	case *bug.CreateTimelineItem:
		if obj == nil {
			return graphql.Null
		}
		return ec._CreateTimelineItem(ctx, sel, obj)
	case *bug.AddCommentTimelineItem:
		if obj == nil {
			return graphql.Null
		}
		return ec._AddCommentTimelineItem(ctx, sel, obj)
	case bug.LabelChangeTimelineItem:
		return ec._LabelChangeTimelineItem(ctx, sel, &obj)
	case *bug.LabelChangeTimelineItem:
		if obj == nil {
			return graphql.Null
		}
		return ec._LabelChangeTimelineItem(ctx, sel, obj)
	case bug.SetStatusTimelineItem:
		return ec._SetStatusTimelineItem(ctx, sel, &obj)
	case *bug.SetStatusTimelineItem:
		if obj == nil {
			return graphql.Null
		}
		return ec._SetStatusTimelineItem(ctx, sel, obj)
	case bug.SetTitleTimelineItem:
		return ec._SetTitleTimelineItem(ctx, sel, &obj)
	case *bug.SetTitleTimelineItem:
		if obj == nil {
			return graphql.Null
		}
		return ec._SetTitleTimelineItem(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var colorImplementors = []string{"Color"}

func (ec *executionContext) _Color(ctx context.Context, sel ast.SelectionSet, obj *color.RGBA) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, colorImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Color")
		case "R":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Color_R(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
		case "G":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Color_G(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
		case "B":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Color_B(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var pageInfoImplementors = []string{"PageInfo"}

func (ec *executionContext) _PageInfo(ctx context.Context, sel ast.SelectionSet, obj *models.PageInfo) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, pageInfoImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("PageInfo")
		case "hasNextPage":

			out.Values[i] = ec._PageInfo_hasNextPage(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "hasPreviousPage":

			out.Values[i] = ec._PageInfo_hasPreviousPage(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "startCursor":

			out.Values[i] = ec._PageInfo_startCursor(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "endCursor":

			out.Values[i] = ec._PageInfo_endCursor(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNColor2imageᚋcolorᚐRGBA(ctx context.Context, sel ast.SelectionSet, v color.RGBA) graphql.Marshaler {
	return ec._Color(ctx, sel, &v)
}

func (ec *executionContext) marshalNColor2ᚖimageᚋcolorᚐRGBA(ctx context.Context, sel ast.SelectionSet, v *color.RGBA) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Color(ctx, sel, v)
}

func (ec *executionContext) unmarshalNHash2githubᚗcomᚋMichaelMureᚋgitᚑbugᚋrepositoryᚐHash(ctx context.Context, v interface{}) (repository.Hash, error) {
	var res repository.Hash
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNHash2githubᚗcomᚋMichaelMureᚋgitᚑbugᚋrepositoryᚐHash(ctx context.Context, sel ast.SelectionSet, v repository.Hash) graphql.Marshaler {
	return v
}

func (ec *executionContext) unmarshalNHash2ᚕgithubᚗcomᚋMichaelMureᚋgitᚑbugᚋrepositoryᚐHashᚄ(ctx context.Context, v interface{}) ([]repository.Hash, error) {
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]repository.Hash, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNHash2githubᚗcomᚋMichaelMureᚋgitᚑbugᚋrepositoryᚐHash(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalNHash2ᚕgithubᚗcomᚋMichaelMureᚋgitᚑbugᚋrepositoryᚐHashᚄ(ctx context.Context, sel ast.SelectionSet, v []repository.Hash) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalNHash2githubᚗcomᚋMichaelMureᚋgitᚑbugᚋrepositoryᚐHash(ctx, sel, v[i])
	}

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNPageInfo2ᚖgithubᚗcomᚋMichaelMureᚋgitᚑbugᚋapiᚋgraphqlᚋmodelsᚐPageInfo(ctx context.Context, sel ast.SelectionSet, v *models.PageInfo) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._PageInfo(ctx, sel, v)
}

func (ec *executionContext) unmarshalNTime2timeᚐTime(ctx context.Context, v interface{}) (time.Time, error) {
	res, err := graphql.UnmarshalTime(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNTime2timeᚐTime(ctx context.Context, sel ast.SelectionSet, v time.Time) graphql.Marshaler {
	res := graphql.MarshalTime(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalNTime2ᚖtimeᚐTime(ctx context.Context, v interface{}) (*time.Time, error) {
	res, err := graphql.UnmarshalTime(v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNTime2ᚖtimeᚐTime(ctx context.Context, sel ast.SelectionSet, v *time.Time) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	res := graphql.MarshalTime(*v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) unmarshalOHash2ᚕgithubᚗcomᚋMichaelMureᚋgitᚑbugᚋrepositoryᚐHashᚄ(ctx context.Context, v interface{}) ([]repository.Hash, error) {
	if v == nil {
		return nil, nil
	}
	var vSlice []interface{}
	if v != nil {
		vSlice = graphql.CoerceList(v)
	}
	var err error
	res := make([]repository.Hash, len(vSlice))
	for i := range vSlice {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithIndex(i))
		res[i], err = ec.unmarshalNHash2githubᚗcomᚋMichaelMureᚋgitᚑbugᚋrepositoryᚐHash(ctx, vSlice[i])
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (ec *executionContext) marshalOHash2ᚕgithubᚗcomᚋMichaelMureᚋgitᚑbugᚋrepositoryᚐHashᚄ(ctx context.Context, sel ast.SelectionSet, v []repository.Hash) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	for i := range v {
		ret[i] = ec.marshalNHash2githubᚗcomᚋMichaelMureᚋgitᚑbugᚋrepositoryᚐHash(ctx, sel, v[i])
	}

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

// endregion ***************************** type.gotpl *****************************