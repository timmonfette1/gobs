package gobs

import "strings"

/*
*****************
*  Structs and Types
*****************
*/
type join int
type joinClauses []joinClause

type joinClause struct {
  joinType      join
  databaseName  string
  schemaName    string
  tableName     string
  tableAlias    string
  onClauses     onClauses
}

/*
*****************
*  Constants
*****************
*/
const (
  Inner join = iota
  Left
  Right
  Full
)

/*
*****************
*  JOIN functions
*****************
*/
func (sb *SqlBuilder) Join(joinType join, databaseName string, schemaName string, tableName string, tableAlias string, onClauses []OnClause) {
  sb.joinClause = append(sb.joinClause, joinClause{
    joinType,
    databaseName,
    schemaName,
    tableName,
    tableAlias,
    onClauses,
  })
}

/*
*****************
*  ToString functions
*****************
*/
func (join join) toString() string {
  return [...]string{"INNER JOIN", "LEFT JOIN", "RIGHT JOIN", "FULL JOIN"}[join]
}

func (join joinClauses) toString() string {
  if lenJoin := len(join); lenJoin != 0 {
    var joinString strings.Builder
    lenJoin = lenJoin - 1

    for jIndex, j := range join {
      joinString.WriteString(j.joinType.toString())
      joinString.WriteString(" ")

      // Check for Database name
      if j.databaseName != "" {
        joinString.WriteString("[")
        joinString.WriteString(j.databaseName)
        joinString.WriteString("].")
      }

      // Check for schema name
      if j.schemaName != "" {
        joinString.WriteString("[")
        joinString.WriteString(j.schemaName)
        joinString.WriteString("].")
      }

      // A Database with no schema infers .. notation
      if j.databaseName != "" && j.schemaName == "" {
        joinString.WriteString(".")
      }

      // Table Name
      joinString.WriteString("[")
      joinString.WriteString(j.tableName)
      joinString.WriteString("]")

      // Check for a desired alias for the table
      if j.tableAlias != "" {
        joinString.WriteString(" ")
        joinString.WriteString(j.tableAlias)
      }

      // ON clauses
      joinString.WriteString(j.onClauses.toString())

      // Newline to prepare for next JOIN
      if jIndex != lenJoin {
        joinString.WriteString("\n")
      }
    }

    return joinString.String()
  } else {
    return ""
  }
}
