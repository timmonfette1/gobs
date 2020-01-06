package gobs

import "strings"

/*
*****************
*  Clause Struct
*****************
*/
type fromClause struct {
  databaseName  string
  derived       bool
  schemaName    string
  tableName     string
  tableAlias    string
}

/*
*****************
*  FROM functions
*****************
*/
func (sb *SqlBuilder) From(database string, schema string, table string, alias string) {
  sb.fromClause = fromClause{database, false, schema, table, alias}
}

func (sb *SqlBuilder) FromDerived(database string, schema string, table string, alias string) {
  sb.fromClause = fromClause{database, true, schema, table, alias}
}


/*
*****************
*  ToString function
*****************
*/
func (from fromClause) toString() string {
  var fromString strings.Builder

  if from.derived {
    fromString.WriteString("FROM (\n")
    fromString.WriteString(from.tableName)
    fromString.WriteString(") ")
    fromString.WriteString(from.tableAlias)
  } else {
    fromString.WriteString("FROM ")

    // Check for Database name
    if from.databaseName != "" {
      fromString.WriteString("[")
      fromString.WriteString(from.databaseName)
      fromString.WriteString("].")
    }

    // Check for schema name
    if from.schemaName != "" {
      fromString.WriteString("[")
      fromString.WriteString(from.schemaName)
      fromString.WriteString("].")
    }

    // A Database with no schema infers .. notation
    if from.databaseName != "" && from.schemaName == "" {
      fromString.WriteString(".")
    }

    // Table Name
    fromString.WriteString("[")
    fromString.WriteString(from.tableName)
    fromString.WriteString("]")

    // Check for a desired alias for the table
    if from.tableAlias != "" {
      fromString.WriteString(" ")
      fromString.WriteString(from.tableAlias)
    }
  }

  return fromString.String()
}
