package gobs

/*
*****************
*  Structs and Types
*****************
*/
type conditional int

/*
*****************
*  Constants
*****************
*/
const (
  None conditional = iota
  And
  Or
)

/*
*****************
*  ToString functions
*****************
*/
func (condition conditional) toString() string {
  return [...]string{"", "AND", "OR"}[condition]
}
