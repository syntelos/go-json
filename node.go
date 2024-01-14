/*
 * JSON IO for GOPL
 */
package json

type Node interface {

	Location() string

	Source() []byte

	Length() uint32

	First() uint32

	Last() uint32

	Begin() uint32

	End() uint32

	Count() uint32

	IsNotEmpty() bool

	String() string

	Head(uint32) byte

	Tail(uint32) byte
}
