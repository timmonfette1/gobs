package gobs

import "strings"

/*
*****************
*  Struct
*****************
*/
type SqlBuilder struct {
  selectClause  selectColumns
  fromClause    fromClause
  whereClause   whereClauses
  joinClause    joinClauses
  groupByClause groupByClauses
  orderByClause orderByClauses
}

/*
*****************
*  Functions
*****************
*/
func (sb *SqlBuilder) BuildSql() string {
  var sqlString strings.Builder

  // Build Output
  sqlString.WriteString(sb.selectClause.toString())
  sqlString.WriteString("\n")
  sqlString.WriteString(sb.fromClause.toString())

  // Only Write JOIN clauses if they exist
  if sb.joinClause != nil {
    sqlString.WriteString("\n")
    sqlString.WriteString(sb.joinClause.toString())
  }

  // Only Write WHERE clauses if they exist
  if sb.whereClause != nil {
    sqlString.WriteString("\n")
    sqlString.WriteString(sb.whereClause.toString())
  }

  // Only Write GROUP BY clauses if they exist
  if sb.groupByClause != nil {
    sqlString.WriteString("\n")
    sqlString.WriteString(sb.groupByClause.toString())
  }

  // Only Write ORDER BY clauses if they exist
  if sb.orderByClause != nil {
    sqlString.WriteString("\n")
    sqlString.WriteString(sb.orderByClause.toString())
  }

  return sqlString.String()
}

// ToString the SQL Builder, but enumerate the result set
func (sb *SqlBuilder) BuildAndEnumerateSql() string {
  sbEnum := SqlBuilder{}

  var rowNumString strings.Builder
  rowNumString.WriteString("ROW_NUMBER() OVER(")
  rowNumString.WriteString(sb.orderByClause.toString())
  rowNumString.WriteString(")")

  sb.Select("COUNT(*) OVER()", "totalRecords")
  sb.Select(rowNumString.String(), "rowNum")

  sbEnum.Select("t.*", "")
  sbEnum.FromDerived("", "", sb.BuildSql(), "t")
  sbEnum.Where("t.rowNum",
    NewBetween(NewSqlField("@startIndex"), NewSqlField("@endIndex")),
    None,
  )

  return sbEnum.BuildSql()
}
