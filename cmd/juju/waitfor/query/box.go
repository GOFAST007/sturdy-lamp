// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package query

import (
	"reflect"

	"github.com/juju/collections/set"
)

// Ord represents an ordered type
type Ord interface {
	// Less checks if a Box is less than another Box
	Less(Ord) bool

	// Equal checks if an Box is equal to another Box.
	Equal(Ord) bool
}

// Box represents a boxed datatype.
type Box interface {
	Ord

	// IsZero returns if the underlying value is zero.
	IsZero() bool

	// Value defines the shadow type value of the Box.
	Value() any
}

// BoxScope defines an ordered integer.
type BoxScope struct {
	value Scope
}

// NewScope creates a new Box value
func NewScope(value Scope) *BoxScope {
	return &BoxScope{value: value}
}

// Less checks if a BoxScope is less than another BoxScope.
func (o *BoxScope) Less(other Ord) bool {
	return false
}

// Equal checks if an BoxScope is equal to another BoxScope.
func (o *BoxScope) Equal(other Ord) bool {
	if i, ok := other.(*BoxScope); ok {
		return o.value == i.value
	}
	return false
}

// IsZero returns if the underlying value is zero.
func (o *BoxScope) IsZero() bool {
	return o.value == nil
}

// Value defines the shadow type value of the Box.
func (o *BoxScope) Value() any {
	return o.value
}

// ForEach iterates over each value in the box.
func (o *BoxScope) ForEach(fn func(any) bool) {
	fn(o.value)
}

// BoxInteger defines an ordered integer.
type BoxInteger struct {
	value int64
}

// NewInteger creates a new Box value
func NewInteger(value int64) *BoxInteger {
	return &BoxInteger{value: value}
}

// Less checks if a BoxInteger is less than another BoxInteger.
func (o *BoxInteger) Less(other Ord) bool {
	if i, ok := other.(*BoxInteger); ok {
		return o.value < i.value
	}
	return false
}

// Equal checks if an BoxInteger is equal to another BoxInteger.
func (o *BoxInteger) Equal(other Ord) bool {
	if i, ok := other.(*BoxInteger); ok {
		return o.value == i.value
	}
	return false
}

// IsZero returns if the underlying value is zero.
func (o *BoxInteger) IsZero() bool {
	return o.value < 1
}

// Value defines the shadow type value of the Box.
func (o *BoxInteger) Value() any {
	return o.value
}

// ForEach iterates over each value in the box.
func (o *BoxInteger) ForEach(fn func(any) bool) {
	fn(o.value)
}

// BoxFloat defines an ordered float.
type BoxFloat struct {
	value float64
}

// NewFloat creates a new Box value
func NewFloat(value float64) *BoxFloat {
	return &BoxFloat{value: value}
}

// Less checks if a BoxFloat is less than another BoxFloat.
func (o *BoxFloat) Less(other Ord) bool {
	if i, ok := other.(*BoxFloat); ok {
		return o.value < i.value
	}
	return false
}

// Equal checks if an BoxFloat is equal to another BoxFloat.
func (o *BoxFloat) Equal(other Ord) bool {
	if i, ok := other.(*BoxFloat); ok {
		return o.value == i.value
	}
	return false
}

// IsZero returns if the underlying value is zero.
func (o *BoxFloat) IsZero() bool {
	return o.value < 1
}

// Value defines the shadow type value of the Box.
func (o *BoxFloat) Value() any {
	return o.value
}

// ForEach iterates over each value in the box.
func (o *BoxFloat) ForEach(fn func(any) bool) {
	fn(o.value)
}

// BoxString defines an ordered string.
type BoxString struct {
	value string
}

// NewString creates a new Box value
func NewString(value string) *BoxString {
	return &BoxString{value: value}
}

// Less checks if a BoxString is less than another BoxString.
func (o *BoxString) Less(other Ord) bool {
	if i, ok := other.(*BoxString); ok {
		return o.value < i.value
	}
	return false
}

// Equal checks if an BoxString is equal to another BoxString.
func (o *BoxString) Equal(other Ord) bool {
	if i, ok := other.(*BoxString); ok {
		return o.value == i.value
	}
	return false
}

// IsZero returns if the underlying value is zero.
func (o *BoxString) IsZero() bool {
	return o.value == ""
}

// Value defines the shadow type value of the Box.
func (o *BoxString) Value() any {
	return o.value
}

// ForEach iterates over each value in the box.
func (o *BoxString) ForEach(fn func(any) bool) {
	fn(o.value)
}

// BoxBool defines an ordered float.
type BoxBool struct {
	value bool
}

// NewBool creates a new Box value
func NewBool(value bool) *BoxBool {
	return &BoxBool{value: value}
}

// Less checks if a BoxBool is less than another BoxBool.
func (o *BoxBool) Less(other Ord) bool {
	return false
}

// Equal checks if an BoxBool is equal to another BoxBool.
func (o *BoxBool) Equal(other Ord) bool {
	if i, ok := other.(*BoxBool); ok {
		return o.value == i.value
	}
	return false
}

// IsZero returns if the underlying value is zero.
func (o *BoxBool) IsZero() bool {
	return !o.value
}

// Value defines the shadow type value of the Box.
func (o *BoxBool) Value() any {
	return o.value
}

// ForEach iterates over each value in the box.
func (o *BoxBool) ForEach(fn func(any) bool) {
	fn(o.value)
}

// BoxMapStringInterface defines an ordered map[string]any.
type BoxMapStringInterface struct {
	value map[string]any
}

// NewMapStringInterface creates a new Box value
func NewMapStringInterface(value map[string]any) *BoxMapStringInterface {
	return &BoxMapStringInterface{value: value}
}

// Less checks if a BoxMapStringInterface is less than another BoxMapStringInterface.
func (o *BoxMapStringInterface) Less(other Ord) bool {
	return false
}

// Equal checks if an BoxMapStringInterface is equal to another BoxMapStringInterface.
func (o *BoxMapStringInterface) Equal(other Ord) bool {
	if i, ok := other.(*BoxMapStringInterface); ok {
		return reflect.DeepEqual(o.value, i.value)
	}
	return false
}

// IsZero returns if the underlying value is zero.
func (o *BoxMapStringInterface) IsZero() bool {
	return len(o.value) == 0
}

// Value defines the shadow type value of the Box.
func (o *BoxMapStringInterface) Value() any {
	return o.value
}

// ForEach iterates over each value in the box.
func (o *BoxMapStringInterface) ForEach(fn func(any) bool) {
	names := set.NewStrings()
	for k := range o.value {
		names.Add(k)
	}
	for _, v := range names.SortedValues() {
		if !fn(o.value[v]) {
			return
		}
	}
}

// BoxMapInterfaceInterface defines an ordered map[any]any.
type BoxMapInterfaceInterface struct {
	value map[any]any
}

// NewMapInterfaceInterface creates a new Box value
func NewMapInterfaceInterface(value map[any]any) *BoxMapInterfaceInterface {
	return &BoxMapInterfaceInterface{value: value}
}

// Less checks if a BoxMapInterfaceInterface is less than another BoxMapInterfaceInterface.
func (o *BoxMapInterfaceInterface) Less(other Ord) bool {
	return false
}

// Equal checks if an BoxMapInterfaceInterface is equal to another BoxMapInterfaceInterface.
func (o *BoxMapInterfaceInterface) Equal(other Ord) bool {
	if i, ok := other.(*BoxMapInterfaceInterface); ok {
		return reflect.DeepEqual(o.value, i.value)
	}
	return false
}

// IsZero returns if the underlying value is zero.
func (o *BoxMapInterfaceInterface) IsZero() bool {
	return len(o.value) == 0
}

// Value defines the shadow type value of the Box.
func (o *BoxMapInterfaceInterface) Value() any {
	return o.value
}

// ForEach iterates over each value in the box.
func (o *BoxMapInterfaceInterface) ForEach(fn func(any) bool) {
	for _, v := range o.value {
		if !fn(v) {
			return
		}
	}
}

// BoxSliceString defines an ordered []string.
type BoxSliceString struct {
	value []string
}

// NewSliceString creates a new Box value
func NewSliceString(value []string) *BoxSliceString {
	return &BoxSliceString{value: value}
}

// Less checks if a BoxSliceString is less than another BoxSliceString.
func (o *BoxSliceString) Less(other Ord) bool {
	return false
}

// Equal checks if an BoxSliceString is equal to another BoxSliceString.
func (o *BoxSliceString) Equal(other Ord) bool {
	if i, ok := other.(*BoxSliceString); ok {
		return reflect.DeepEqual(o.value, i.value)
	}
	return false
}

// IsZero returns if the underlying value is zero.
func (o *BoxSliceString) IsZero() bool {
	return len(o.value) == 0
}

// Value defines the shadow type value of the Box.
func (o *BoxSliceString) Value() any {
	return o.value
}

// ForEach iterates over each value in the box.
func (o *BoxSliceString) ForEach(fn func(any) bool) {
	for _, v := range o.value {
		if !fn(v) {
			return
		}
	}
}

// BoxLambda defines an ordered integer.
type BoxLambda struct {
	arg *Identifier
	fn  func(Scope) ([]Box, error)
}

// NewLambda creates a new Box value
func NewLambda(arg *Identifier, fn func(Scope) ([]Box, error)) *BoxLambda {
	return &BoxLambda{
		arg: arg,
		fn:  fn,
	}
}

// Less checks if a BoxLambda is less than another BoxLambda.
func (o *BoxLambda) Less(other Ord) bool {
	return false
}

// Equal checks if an BoxLambda is equal to another BoxLambda.
func (o *BoxLambda) Equal(other Ord) bool {
	return false
}

// IsZero returns if the underlying value is zero.
func (o *BoxLambda) IsZero() bool {
	return false
}

// Value defines the shadow type value of the Box.
func (o *BoxLambda) Value() any {
	return o
}

// ArgName is the name expected to be seen in the lambda.
func (o *BoxLambda) ArgName() string {
	return o.arg.Token.Literal
}

// Call the underlying lambda
func (o *BoxLambda) Call(scope Scope) ([]Box, error) {
	return o.fn(scope)
}

func expectStringIndex(i any) (*BoxString, error) {
	box, ok := i.(Box)
	if !ok {
		return nil, RuntimeErrorf("expected string, but got %T", i)
	}

	idx, ok := i.(*BoxString)
	if !ok {
		return nil, RuntimeErrorf("expected string, but got %v", shadowType(box))
	}

	return idx, nil
}

func expectIntegerIndex(i any) (*BoxInteger, error) {
	box, ok := i.(Box)
	if !ok {
		return nil, RuntimeErrorf("expected int, but got %T", i)
	}

	idx, ok := i.(*BoxInteger)
	if !ok {
		return nil, RuntimeErrorf("expected int, but got %v", shadowType(box))
	}

	return idx, nil
}

func expectBoxIndex(i any) (Box, error) {
	box, ok := i.(Box)
	if !ok {
		return nil, RuntimeErrorf("expected box, but got %T", i)
	}

	return box, nil
}

func shadowType(box Box) string {
	switch box.(type) {
	case *BoxBool:
		return "bool"
	case *BoxInteger:
		return "int64"
	case *BoxFloat:
		return "float64"
	case *BoxString:
		return "string"
	case *BoxMapInterfaceInterface:
		return "map[any]any"
	case *BoxMapStringInterface:
		return "map[string]any"
	case *BoxSliceString:
		return "[]string"
	}
	return "<unknown>"
}
