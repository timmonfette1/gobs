package gobs

import "strings"

/*
*****************
*  Structs and Types
*****************
*/
type sortOrder int
type orderByClauses []orderBy

type orderBy struct {
  columnName  string
  order       sortOrder
}

/*
*****************
*  Constants
*****************
*/
const (
  Asc sortOrder = iota
  Desc
)

/*
*****************
*  OrderBy functions
*****************
*/
func (sb *SqlBuilder) OrderBy(columnName string, order sortOrder) {
  sb.orderByClause = append(sb.orderByClause, orderBy{columnName, order})
}

/*
*****************
*  ToString functions
*****************
*/
func (order sortOrder) toString() string {
  return [...]string{"ASC", "DESC"}[order]
}

func (orderBy orderByClauses) toString() string {
  if len := len(orderBy); len != 0 {
    var orderByString strings.Builder
    orderByString.WriteString("ORDER BY ")

    len = len - 1
    for index, o := range orderBy {
      orderByString.WriteString(o.columnName)
      orderByString.WriteString(" ")
      orderByString.WriteString(o.order.toString())

      // Don't write a comma on the last column
      if index != len {
        orderByString.WriteString(",")
      }
    }

    return orderByString.String()
  } else {
    return ""
  }
}
