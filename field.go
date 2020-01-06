package gobs

import "strings"

/*
*****************
*  Structs and Types
*****************
*/
type fieldList []Field

type Field struct {
  value     string
  paramName string
  subQuery  string
}

/*
*****************
*  Creation functions
*****************
*/
func NewValueField(value string) Field {
  return Field{value, "", ""}
}

func NewSqlField(param string) Field {
  return Field{"", param, ""}
}

func NewSubQuery(query string) Field {
  return Field{"", "", query}
}

/*
*****************
*  ToString functions
*****************
*/
func (field Field) toString() string {
  var fieldString strings.Builder

  if len(field.subQuery) != 0 {
    fieldString.WriteString("(")
    fieldString.WriteString(field.subQuery)
    fieldString.WriteString(")")
  } else if len(field.value) != 0 {
    fieldString.WriteString("'")
    fieldString.WriteString(field.value)
    fieldString.WriteString("'")
  } else {
    fieldString.WriteString(field.paramName)
  }

  return fieldString.String()
}

func (fieldList fieldList) toString() string {
  isSubQ := false
  if len(fieldList) == 1 && len(fieldList[0].subQuery) != 0 {
    isSubQ = true
  }

  if len := len(fieldList); len != 0 {
    var fieldString strings.Builder

    if !isSubQ {
      fieldString.WriteString("(")
    }

    len = len - 1
    for index, f := range fieldList {
      fieldString.WriteString(f.toString())
      if index != len {
        fieldString.WriteString(",")
      }
    }

    if !isSubQ {
      fieldString.WriteString(")")
    }
    
    return fieldString.String()
  } else {
    return ""
  }
}
