package gobs

import "strings"

/*
*****************
*  Structs and Types
*****************
*/
type onClauses []OnClause

type OnClause struct {
  columnName  string
  operator    operator
  conditional conditional
}

/*
*****************
*  Creation functions
*****************
*/
func NewOnClause(columnName string, operator operator, conditional conditional) OnClause {
  return OnClause{columnName, operator, conditional}
}

/*
*****************
*  ToString functions
*****************
*/
func (on onClauses) toString() string {
  if len := len(on); len != 0 {
    var onString strings.Builder
    onString.WriteString(" ON ")

    for index, o := range on {
      onString.WriteString(o.columnName)
      onString.WriteString(GetOperator(o.operator))

      if index != len && o.conditional != 0 {
        onString.WriteString(" ")
        onString.WriteString(o.conditional.toString())
        onString.WriteString(" ")
      }
    }

    return onString.String()
  } else {
    return ""
  }
}
