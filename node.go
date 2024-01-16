/*
 * JSON IO for GOPL
 */
package json

type NodeType uint8
const (
	NodeTypeUnrecognized NodeType = 0x00
	NodeTypeArray        NodeType = 0x10
	NodeTypeObject       NodeType = 0x20
	NodeTypeField        NodeType = 0x40
	NodeTypeString       NodeType = 0x80
)

type Node interface {

	Location() string

	Source() []byte

	Length() uint32

	Type() NodeType

	First() uint32

	Last() uint32

	Begin() uint32

	End() uint32

	Count() uint32

	IsNotEmpty() bool

	String() string

	Head(int) byte

	Tail(int) byte

	Contains(Node) bool
}
