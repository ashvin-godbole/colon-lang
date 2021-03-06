package colobj

import (
	"bytes"
	ast "colon/colast"
	"fmt"
	"strings"
)

// At runtime, every node from the ast will be wrapped into an equivalent "object" for the evaluator

// ObjectType : type of "object" that a particular node in the ast produces at runtime
type ObjectType string

// Datatypes in Colon
const (
	INTEGER  = "INTEGER"
	BOOLEAN  = "BOOLEAN"
	FLOATING = "FLOATING"
	STRING   = "STRING"
	LIST     = "LIST"
	EMPTY    = "EMPTY"
	RETVAL   = "RETURN_VALUE"
	FUNCTION = "FUNCTION"
	LOOP     = "LOOP"
	BUILTIN  = "BUILT_IN"
	INPUT    = "INPUT"
)

// Object : an interface that wraps values of all types which can be fed into the evaluator
type Object interface {
	ObType() ObjectType
	ObValue() string
}

// ----------------------------------------------------------------------------

// Integer : A wrapper for integer values
type Integer struct {
	Value int64
}

// ObValue : Integer
func (i *Integer) ObValue() string {
	return fmt.Sprintf("%v", i.Value)
}

// ObType : Integer
func (i *Integer) ObType() ObjectType {
	return INTEGER
}

// ----------------------------------------------------------------------------

// Boolean : A wrapper for boolean values
type Boolean struct {
	Value bool
}

// ObValue : Boolean
func (b *Boolean) ObValue() string {
	return fmt.Sprintf("%v", b.Value)
}

// ObType : Boolean
func (b *Boolean) ObType() ObjectType {
	return BOOLEAN
}

// ----------------------------------------------------------------------------

// Floating : A wrapper for floating-point values
type Floating struct {
	Value float64
}

// ObValue : Floating
func (f *Floating) ObValue() string {
	return fmt.Sprintf("%v", f.Value)
}

// ObType : Floating
func (f *Floating) ObType() ObjectType {
	return FLOATING
}

// ----------------------------------------------------------------------------

// String : A wrapper for string values
type String struct {
	Value string
}

// ObValue : String
func (s *String) ObValue() string {
	return fmt.Sprintf("%v", s.Value)
}

// ObType : String
func (s *String) ObType() ObjectType {
	return STRING
}

// ----------------------------------------------------------------------------

// Empty : Colon's version of Null/Nil.
type Empty struct{}

// ObValue : Empty
func (e *Empty) ObValue() string {
	return ""
}

//ObType : Empty
func (e *Empty) ObType() ObjectType {
	return EMPTY
}

// ----------------------------------------------------------------------------

// ReturnValue : structure that wraps the return value into an object so that
// it can be returned as an object instead of just a value
type ReturnValue struct {
	Value Object
}

// ObValue : ReturnValue
func (r *ReturnValue) ObValue() string {
	return r.Value.ObValue()
}

// ObType : ReturnValue
func (r *ReturnValue) ObType() ObjectType {
	return RETVAL
}

// ----------------------------------------------------------------------------

// Function : structure that wraps function definitions
type Function struct {
	Parameters []*ast.Identifier
	FuncBody   *ast.Block
	Env        *Env // functions have their own environment
}

// ObValue : Function
func (f *Function) ObValue() string {
	var str bytes.Buffer
	str.WriteString("FUNCTION {\n(")
	for _, v := range f.Parameters {
		str.WriteString(v.String())
		str.WriteString(", ")
	}
	str.WriteString(")\n")
	str.WriteString(f.FuncBody.String())
	str.WriteString("}")
	return str.String()
}

// ObType : Function
func (f *Function) ObType() ObjectType {
	return FUNCTION
}

// ----------------------------------------------------------------------------

// Loop : structure that wraps loop statements
type Loop struct {
	Condition ast.Expression
	LoopBody  *ast.Block
	Env       *Env // loops also have their own environment
}

// ObValue : Function
func (l *Loop) ObValue() string {
	var str bytes.Buffer
	str.WriteString("Loop {\n(")
	str.WriteString("I'm working on getting the condition to print here")
	str.WriteString(")\n")
	str.WriteString(l.LoopBody.String())
	str.WriteString("}")
	return str.String()
}

// ObType : Function
func (l *Loop) ObType() ObjectType {
	return LOOP
}

// ----------------------------------------------------------------------------

// BuiltInFunction : a type of function which is built
// into the colon interpreter
type BuiltInFunction func(args ...Object) Object

// ----------------------------------------------------------------------------

// BuiltIn : to provide a warpper around some
// functions that are built into the colon interpreter
type BuiltIn struct {
	Bfunct BuiltInFunction
}

// ObValue : BuiltIn
func (bin *BuiltIn) ObValue() string {
	return "Builtin Function"
}

// ObType : BuiltIn
func (bin *BuiltIn) ObType() ObjectType {
	return BUILTIN
}

// ----------------------------------------------------------------------------

// InputFunction : to provide a way to get user input from stdin
type InputFunction struct {
	InFunc func(env *Env, varname string, dtype DataType) Object
	ENV    *Env
}

// ObValue : InputFunction
func (bin *InputFunction) ObValue() string {
	return "Builtin Function"
}

// ObType : InputFunction
func (bin *InputFunction) ObType() ObjectType {
	return INPUT
}

// ----------------------------------------------------------------------------

//DataType : To encode Datatype information; mainly required by the input function
type DataType struct {
	Dtype string
}

// ObValue : DataType
func (dt *DataType) ObValue() string {
	return "DATA-TYPE"
}

// ObType : DataType
func (dt *DataType) ObType() ObjectType {
	return BUILTIN
}

// ----------------------------------------------------------------------------

// List : structure that wraps a list into an object
type List struct {
	Elements []Object
}

// ObValue : ReturnValue
func (l *List) ObValue() string {
	var str bytes.Buffer
	str.WriteString("[")
	elems := []string{}
	for _, v := range l.Elements {
		elems = append(elems, v.ObValue())
	}
	str.WriteString(strings.Join(elems, ", "))
	str.WriteString("]")
	return str.String()
}

// ObType : ReturnValue
func (l *List) ObType() ObjectType {
	return LIST
}

// ----------------------------------------------------------------------------
