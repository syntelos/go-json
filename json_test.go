/*
 * JSON IO for GOPL
 */
package json

import (
	"fmt"
	"os"
	"testing"
)

func TestIndex(t *testing.T){
	var filename string = "tst/index12.json"
	var fo *os.File
	var er error

	fo, er = os.Open(filename)
	if nil != er {
		t.Fatalf("Opening '%s': %v",filename,er)
	} else {
		defer fo.Close()
		var count int = 0

		var reader Reader = ReadFile(fo)
		if reader.IsNotEmpty() {

			var array Reader = reader.HeadArray()
			if array.IsNotEmpty() {

				var object Reader = array.HeadObject()
				for object.IsNotEmpty() {

					count += 1

					object = object.TailObject()
				}
			}
		}
		
		if 12 != count {
			t.Fatalf("[TestIndex] Count '%d'",count)
		} else {
			fmt.Printf("[TestIndex] Count '%d'\n",count)
		}
	}
}

func TestField(t *testing.T){
	var filename string = "tst/index1.json"
	var fo *os.File
	var er error

	fo, er = os.Open(filename)
	if nil != er {
		t.Fatalf("[TestField] Opening '%s': %v",filename,er)
	} else {
		defer fo.Close()
		var count int = 0

		var reader Reader = ReadFile(fo)
		if reader.IsNotEmpty() {

			var array Reader = reader.HeadArray()
			if array.IsNotEmpty() {

				var object Reader = array.HeadObject()
				for object.IsNotEmpty() {

					var field Reader = object.HeadField()
					for field.IsNotEmpty() {
						count += 1

						field = field.TailField()
					}

					object = object.TailObject()
				}
			}
		}

		if 6 != count {
			t.Fatalf("[TestField] Count '%d'",count)
		} else {
			fmt.Printf("[TestField] Count '%d'\n",count)
		}
	}
}
