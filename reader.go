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
		default:
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

func (this Reader) String() (empty string) {

	if this.IsNotEmpty() {

		var substring []byte = this.source[this.begin:this.end]

		return string(substring)
	} else {
		return empty
	}
}

func (this Reader) StringUnquote() (empty string) {

	if this.IsNotEmpty() {

		var first, last int = int(this.begin), int(this.end)-1

		if first < last && '"' == this.source[first] && '"' == this.source[last] {

			var substring []byte = this.source[(first+1):last]

			return string(substring)
		} else {

			var substring []byte = this.source[this.begin:this.end]

			return string(substring)
		}
	} else {
		return empty
	}
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

func ReadFile(fo *os.File) (empty Reader) {

	var fi os.FileInfo
	var er error

	fi, er = fo.Stat()
	if nil == er {
		var length uint32 = uint32(fi.Size())
		if 0 != length {
			var content []byte = make([]byte,length)

			_, er = fo.Read(content)
			if nil == er {
				fo.Close()

				return NewReader(fi.Name(),content)
			}
		}
	}
	return empty
}

func NewReader(location string, content []byte) (empty Reader) {
	var length uint32 = uint32(len(content))
	if 0 != length {
		var reader Reader = Reader{location,content,length,0,length}

		return reader
	} else {
		return empty
	}
}

func (this Reader) ReadArray(begin uint32) (empty Reader) {

	if begin < this.length {
		if '[' == this.source[begin] {
			var first, last uint32 = begin, uint32(span.Forward(this.source,int(begin),int(this.length),'[',']'))
			if first < last {
				var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

				return reader
			}
		} else {
			var ws int = span.Class(this.source,int(begin),int(this.length),span.WS)
			if 0 < ws {
				begin = uint32(ws+1)
				if '[' == this.source[begin] {
					var first, last uint32 = begin, uint32(span.Forward(this.source,int(begin),int(this.length),'[',']'))
					if first < last {
						var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

						return reader
					}
				}
			} else {
				begin += 1
				if begin < this.length {
					if '[' == this.source[begin] {
						var first, last uint32 = begin, uint32(span.Forward(this.source,int(begin),int(this.length),'[',']'))
						if first < last {
							var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

							return reader
						}
					} else {
						var ws int = span.Class(this.source,int(begin),int(this.length),span.WS)
						if 0 < ws {
							begin = uint32(ws+1)
							if '[' == this.source[begin] {
								var first, last uint32 = begin, uint32(span.Forward(this.source,int(begin),int(this.length),'[',']'))
								if first < last {
									var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

									return reader
								}
							}
						}
					}
				}
			}
		}
	}
	return empty
}

func (this Reader) HeadArray() Reader {
	return this.ReadArray(this.begin)
}

func (this Reader) TailArray() Reader {
	return this.ReadArray(this.end)
}

func (this Reader) ReadObject(begin uint32) (empty Reader) {

	if begin < this.length {
		if '{' == this.source[begin] {
			var first, last uint32 = begin, uint32(span.Forward(this.source,int(begin),int(this.length),'{','}'))
			if first < last {
				var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

				return reader
			}
		} else {
			var ws int = span.Class(this.source,int(begin),int(this.length),span.WS)
			if 0 < ws {
				begin = uint32(ws+1)
				if '{' == this.source[begin] {
					var first, last uint32 = begin, uint32(span.Forward(this.source,int(begin),int(this.length),'{','}'))
					if first < last {
						var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

						return reader
					}
				}
			} else {
				begin += 1
				if begin < this.length {
					if '{' == this.source[begin] {
						var first, last uint32 = begin, uint32(span.Forward(this.source,int(begin),int(this.length),'{','}'))
						if first < last {
							var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

							return reader
						}
					} else {
						var ws int = span.Class(this.source,int(begin),int(this.length),span.WS)
						if 0 < ws {
							begin = uint32(ws+1)
							if '{' == this.source[begin] {
								var first, last uint32 = begin, uint32(span.Forward(this.source,int(begin),int(this.length),'{','}'))
								if first < last {
									var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

									return reader
								}
							}
						}
					}
				}
			}
		}
	}
	return empty
}

func (this Reader) HeadObject() Reader {
	return this.ReadObject(this.begin)
}

func (this Reader) TailObject() Reader {
	return this.ReadObject(this.end)
}

func (this Reader) ReadField(begin uint32) (empty Reader) {
	var name Reader = this.ReadString(begin)
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

func (this Reader) CondField(field_name string, begin uint32) (empty Reader) {
	var field Reader = this.ReadField(begin)
	if field.IsNotEmpty() {

		var name Reader = field.HeadString()
		if ':' == name.Tail(0) && field_name == name.StringUnquote() {

			return field
		}
	}
	return empty
}

func (this Reader) HeadField() Reader {
	return this.ReadField(this.begin)
}

func (this Reader) TailField() Reader {
	return this.ReadField(this.end)
}

func (this Reader) CondHeadField(field_name string) Reader {
	return this.CondField(field_name,this.begin)
}

func (this Reader) CondTailField(field_name string) Reader {
	return this.CondField(field_name,this.end)
}

func (this Reader) ReadString(begin uint32) (empty Reader) {

	if begin < this.length {
		if '"' == this.source[begin] {

			var first, last uint32 = begin, uint32(span.First(this.source,int(begin+1),int(this.length),'"'))
			if first < last {
				var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

				return reader
			}
		} else {
			var ws int = span.Class(this.source,int(begin),int(this.length),span.WS)
			if 0 < ws {
				begin = uint32(ws+1)

				if '"' == this.source[begin] {

					var first, last uint32 = begin, uint32(span.First(this.source,int(begin+1),int(this.length),'"'))
					if first < last {
						var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

						return reader
					}
				}
			} else {
				var first, last int = int(begin), span.Class(this.source,int(begin),int(this.length),span.GI)
				if first < last {
					var reader Reader = Reader{this.location,this.source,this.length,uint32(first),uint32(last+1)}

					return reader
				} else {
					begin += 1
					if begin < this.length {
						if '"' == this.source[begin] {

							var first, last uint32 = begin, uint32(span.First(this.source,int(begin+1),int(this.length),'"'))
							if first < last {
								var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

								return reader
							}
						} else {
							var ws int = span.Class(this.source,int(begin),int(this.length),span.WS)
							if 0 < ws {
								begin = uint32(ws+1)

								if '"' == this.source[begin] {

									var first, last uint32 = begin, uint32(span.First(this.source,int(begin+1),int(this.length),'"'))
									if first < last {
										var reader Reader = Reader{this.location,this.source,this.length,first,(last+1)}

										return reader
									}
								}
							} else {
								var first, last int = int(begin), span.Class(this.source,int(begin),int(this.length),span.GI)
								if first < last {
									var reader Reader = Reader{this.location,this.source,this.length,uint32(first),uint32(last+1)}

									return reader
								}
							}
						}
					}
				}
			}
		}
	}
	return empty
}

func (this Reader) HeadString() Reader {
	return this.ReadString(this.begin)
}

func (this Reader) TailString() Reader {
	return this.ReadString(this.end)
}
