package gobs

import "strings"

/*
*****************
*  Structs and Types
*****************
*/
type groupByClauses []string

/*
*****************
*  GroupBy functions
*****************
*/
func (sb *SqlBuilder) GroupBy(column string) {
  sb.groupByClause = append(sb.groupByClause, column)
}

func (sb *SqlBuilder) GroupByMultiple(clauses []string) {
  sb.groupByClause = clauses
}

/*
*****************
*  ToString functions
*****************
*/
func (groupBy groupByClauses) toString() string {
  if len := len(groupBy); len != 0 {
    var groupByString strings.Builder
    groupByString.WriteString("GROUP BY ")

    len = len - 1
    for index, group := range groupBy {
      groupByString.WriteString(group)

      // Don't write a comma on the last column
      if index != len {
        groupByString.WriteString(",")
      }
    }

    return groupByString.String()
  } else {
    return ""
  }
}
