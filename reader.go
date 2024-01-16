/*
 * JSON IO for GOPL
 */
package json

import (
	span "github.com/syntelos/go-span"
	"os"
)
/*
 * The substring node performs reader functions.
 */
type Reader struct {
	location string
	source []byte
	length uint32
	begin, end uint32
}

func (this Reader) Location() string {

	return this.location
}

func (this Reader) Source() []byte {

	return this.source
}

func (this Reader) Length() uint32 {

	return this.length
}

func (this Reader) Type() NodeType {

	if this.IsNotEmpty() {

		switch this.Head(0) {

		case '[':
			if ']' == this.Tail(-1) {

				return NodeTypeArray
			}
		case '{':
			if '}' == this.Tail(-1) {

				return NodeTypeObject
			}
		case '"':
			var str Reader = this.HeadString()
			if str.IsNotEmpty() {

				if ':' == str.Tail(0) {

					return NodeTypeField
				} else {
					return NodeTypeString
				}
			}
		}
	}
	return NodeTypeUnrecognized
}

func (this Reader) Begin() uint32 {

	return this.begin
}

func (this Reader) First() uint32 {

	return this.begin
}

func (this Reader) Last() uint32 {

	if this.begin < this.end {
	
		return (this.end-1)
	} else {
		return this.begin
	}
}

func (this Reader) End() uint32 {

	return this.end
}

func (this Reader) Count() uint32 {

	return (this.end-this.begin)
}

func (this Reader) IsNotEmpty() bool {

	return (0 <= this.begin && this.begin < this.end)
}

func (this Reader) String() string {

	var substring []byte = this.source[this.begin:this.end]

	return string(substring)
}

func (this Reader) Head(offset int) (ch byte) {

	var ofs int = int(this.begin)+offset
	if -1 < ofs && uint32(ofs) < this.end {

		return this.source[ofs]
	} else {
		return 0
	}
}

func (this Reader) Tail(offset int) (ch byte) {

	var ofs int = int(this.end)+offset
	if -1 < ofs && uint32(ofs) < this.length {

		return this.source[ofs]
	} else {
		return 0
	}
}

func (this Reader) Contains(node Node) bool {

	if this.IsNotEmpty() {
		var node_head, node_tail uint32 = node.Begin(), node.End()

		return (node_head >= this.begin && node_tail <= this.end)
	} else {
		return false
	}
}

func ReadFile(fo *os.File) (this Reader) {

	var fi os.FileInfo
	var er error

	fi, er = fo.Stat()
	if nil == er {
		var length uint32 = uint32(fi.Size())
		if 0 != length {
			var content []byte = make([]byte,length)

			_, er = fo.Read(content)
			if nil == er {
				var location string = "file:"+fi.Name()

				var reader Reader = Reader{location,content,length,0,length}

				return reader
			}
		}
	}
	return this
}

func (this Reader) HeadArray() (empty Reader) {
	var begin uint32 = this.begin
	if begin < this.length {
		if '[' == this.source[begin] {
			var first, last uint32 = begin, uint32(span.Forward(this.source,int(begin),int(this.length),'[',']'))
			if first < last {
				var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

				return reader
			}
		} else {
			var first uint32 = uint32(span.First(this.source,int(begin),int(this.length),'['))
			if '[' == this.source[first] {

				var last uint32 = uint32(span.Forward(this.source,int(first),int(this.length),'[',']'))
				if first < last {
					var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

					return reader
				}
			}
		}
	}
	return empty
}

func (this Reader) HeadObject() (empty Reader) {
	var begin uint32 = this.begin
	if begin < this.length {
		if '{' == this.source[begin] {
			var first, last uint32 = begin, uint32(span.Forward(this.source,int(begin),int(this.length),'{','}'))
			if first < last {
				var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

				return reader
			}
		} else {
			var first uint32 = uint32(span.First(this.source,int(begin),int(this.length),'{'))
			if '{' == this.source[first] {

				var last uint32 = uint32(span.Forward(this.source,int(first),int(this.length),'{','}'))
				if first < last {
					var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

					return reader
				}
			}
		}
	}
	return empty
}

func (this Reader) HeadField() (empty Reader) {
	var name Reader = this.HeadString()
	if name.IsNotEmpty() {

		if ':' == name.Tail(0) {

			var value Reader = name.TailString()
			if value.IsNotEmpty() {
				var begin, end uint32 = name.begin, value.end

				var reader Reader = Reader{this.location,this.source,this.length,begin,end}

				return reader
			}
		}
	}
	return empty
}

func (this Reader) HeadString() (empty Reader) {
	var begin uint32 = this.begin
	if begin < this.length {
		if '"' == this.source[begin] {

			var first, last uint32 = begin, uint32(span.First(this.source,int(begin+1),int(this.length),'"'))
			if first < last {
				var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

				return reader
			}
		} else {
			var first uint32 = uint32(span.First(this.source,int(begin),int(this.length),'"'))
			if '"' == this.source[first] {

				var last uint32 = uint32(span.First(this.source,int(first+1),int(this.length),'"'))
				if first < last {
					var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

					return reader
				}
			}
		}
	}
	return empty
}

func (this Reader) TailArray() (empty Reader) {
	var begin uint32 = this.end
	if begin < this.length {
		if '[' == this.source[begin] {

			var first, last uint32 = begin, uint32(span.Forward(this.source,int(begin),int(this.length),'[',']'))
			if first < last {
				var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

				return reader
			}
		} else {

			var first uint32 = uint32(span.First(this.source,int(begin),int(this.length),'['))
			if '[' == this.source[first] {

				var last uint32 = uint32(span.Forward(this.source,int(first),int(this.length),'[',']'))
				if first < last {
					var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

					return reader
				}
			}
		}
	}
	return empty
}

func (this Reader) TailObject() (empty Reader) {
	var begin uint32 = this.end
	if begin < this.length {
		if '{' == this.source[begin] {
			var first, last uint32 = begin, uint32(span.Forward(this.source,int(begin),int(this.length),'{','}'))
			if first < last {
				var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

				return reader
			}
		} else {
			var first uint32 = uint32(span.First(this.source,int(begin),int(this.length),'{'))
			if '{' == this.source[first] {

				var last uint32 = uint32(span.Forward(this.source,int(first),int(this.length),'{','}'))
				if first < last {
					var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

					return reader
				}
			}
		}
	}
	return empty
}

func (this Reader) TailField() (empty Reader) {
	var name Reader = this.TailString()
	if name.IsNotEmpty() {

		if ':' == name.Tail(0) {

			var value Reader = name.TailString()
			if value.IsNotEmpty() {
				var begin, end uint32 = name.begin, value.end

				var reader Reader = Reader{this.location,this.source,this.length,begin,end}

				return reader
			}
		}
	}
	return empty
}

func (this Reader) TailString() (empty Reader) {
	var begin uint32 = this.end
	if begin < this.length {
		if '"' == this.source[begin] {

			var first, last uint32 = begin, uint32(span.First(this.source,int(begin+1),int(this.length),'"'))
			if first < last {
				var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

				return reader
			}
		} else {
			var first uint32 = uint32(span.First(this.source,int(begin),int(this.length),'"'))
			if '"' == this.source[first] {

				var last uint32 = uint32(span.First(this.source,int(first+1),int(this.length),'"'))
				if first < last {
					var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

					return reader
				}
			}
		}
	}
	return empty
}
