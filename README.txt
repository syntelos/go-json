JSON IO for GOPL

  A superset of JSON allowing unquoted strings in lieu of
  JSON string conventions.


Stateless document object model

  The document object model is objectified and classified
  but not collected into an inclusive structure.

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


References

  [JSON] https://datatracker.ietf.org/doc/html/rfc8259
  [SPAN] https://github.com/syntelos/go-span

