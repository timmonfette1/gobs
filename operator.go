package gobs

import "strings"

/*
*****************
*  Structs and Types
*****************
*/
type constOperator int
type nullOperator int

type operator interface {
  stringifyOperator() string
}

type basicOperator struct {
  cOper constOperator
  field Field
}

type between struct {
  field1  Field
  field2  Field
}

type in struct {
  fieldList fieldList
}

/*
*****************
*  Constants
*****************
*/
const (
  Equals constOperator = iota
  NotEquals
  GreaterThan
  LessThan
  GreaterThanEqual
  LessThanEqual
  Like
)

const (
  IsNull nullOperator = iota
  IsNotNull
)

/*
*****************
*  Creation functions
*****************
*/
func NewOperator(cOper constOperator, param Field) basicOperator {
  return basicOperator{cOper, param}
}

func NewBetween(param1 Field, param2 Field) between {
  return between{param1, param2}
}

func NewIn(fields ...Field) in {
  return in {fields}
}

/*
*****************
*  StringifyOperator functions
*****************
*/
func (nOperator nullOperator) stringifyOperator() string {
  return [...]string{" IS NULL", " IS NOT NULL"}[nOperator]
}

func (basic basicOperator) stringifyOperator() string {
  var basicString strings.Builder
  basicString.WriteString(basic.cOper.toString())
  basicString.WriteString(basic.field.toString())
  return basicString.String()
}

func (between between) stringifyOperator() string {
  var betweenString strings.Builder
  betweenString.WriteString(" BETWEEN ")
  betweenString.WriteString(between.field1.toString())
  betweenString.WriteString(" AND ")
  betweenString.WriteString(between.field2.toString())
  return betweenString.String()
}

func (in in) stringifyOperator() string {
  var inString strings.Builder
  inString.WriteString(" IN ")
  inString.WriteString(in.fieldList.toString())
  return inString.String()
}

/*
*****************
*  ToString functions
*****************
*/
func (cOperator constOperator) toString() string {
  return [...]string{"=", "<>", ">", "<", ">=", "<=", " LIKE "}[cOperator]
}

func GetOperator(operator operator) string {
  return operator.stringifyOperator()
}
