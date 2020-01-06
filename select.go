package gobs

import "strings"

/*
*****************
*  Structs and Types
*****************
*/
type selectColumns map[string]string

/*
*****************
*  SELECT functions
*****************
*/
func (sb *SqlBuilder) Select(column string, alias string) {
  if len(sb.selectClause) == 0 {
    sb.selectClause = make(map[string]string)
  }

  sb.selectClause[column] = alias
}

func (sb *SqlBuilder) SelectMultiple(clauses map[string]string) {
  sb.selectClause = clauses
}

/*
*****************
*  ToString function
*****************
*/
func (sel selectColumns) toString() string {
  var selectString strings.Builder
  selectString.WriteString("SELECT ")

  len := len(sel) - 1
  index := 0
  for col, alias := range sel {
    selectString.WriteString(col)

    // Check for a desired alias for the column
    if alias != "" {
      selectString.WriteString(" AS [")
      selectString.WriteString(alias)
      selectString.WriteString("]")
    }

    // Don't write a comma on the last column
    if index != len {
      selectString.WriteString(",")
    }

    index++
  }

  return selectString.String()
}
