package gobs

import "strings"

/*
*****************
*  Structs and Types
*****************
*/
type whereClauses []whereClause

type whereClause struct {
  columnName  string
  operator    operator
  conditional conditional
  groupStart  bool
  groupEnd    bool
}

/*
*****************
*  WHERE functions
*****************
*/
func (sb *SqlBuilder) Where(columnName string, operator operator, conditional conditional) {
  sb.whereClause = append(sb.whereClause, whereClause{columnName, operator, conditional, false, false})
}

func (sb *SqlBuilder) WhereGroup(columnName string, operator operator, conditional conditional) {
  sb.whereClause = append(sb.whereClause, whereClause{columnName, operator, conditional, true, false})
}

func (sb *SqlBuilder) WhereGroupEnd(columnName string, operator operator, conditional conditional) {
  sb.whereClause = append(sb.whereClause, whereClause{columnName, operator, conditional, false, true})
}

/*
*****************
*  ToString functions
*****************
*/
func (where whereClauses) toString() string {
  if len := len(where); len != 0 {
    var whereString strings.Builder
    whereString.WriteString("WHERE ")

    len = len - 1
    for index, w := range where {
      if w.groupStart {
        whereString.WriteString("(")
      }

      whereString.WriteString(w.columnName)
      whereString.WriteString(GetOperator(w.operator))

      if index != len && w.conditional != 0 {
        if w.groupEnd {
          whereString.WriteString(")")
        }

        whereString.WriteString(" ")
        whereString.WriteString(w.conditional.toString())
        whereString.WriteString(" ")
      }
    }

    return whereString.String()
  } else {
    return ""
  }
}
