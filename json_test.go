/*
 * JSON IO for GOPL
 */
package json

import (
	"fmt"
	"os"
	"testing"
)

func TestIndex1(t *testing.T){
	var filename string = "tst/index1.json"
	var fo *os.File
	var er error

	fo, er = os.Open(filename)
	if nil != er {
		t.Fatalf("[TestIndex1] Opening '%s': %v",filename,er)
	} else {
		defer fo.Close()
		var count int = 0

		var reader Reader = ReadFile(fo)
		if reader.IsNotEmpty() {
			var array Reader = reader.HeadArray()
			if array.IsNotEmpty() {
				var object Reader = array.HeadObject()
				for object.IsNotEmpty() {

					fmt.Println(object)

					count += 1

					object = object.TailObject()
				}
			}
		}

		if 1 != count {
			t.Fatalf("[TestIndex1] Count '%d'",count)
		} else {
			fmt.Printf("[TestIndex1] Count '%d'\n",count)
		}
	}
}

func TestIndex12(t *testing.T){
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

					fmt.Println(object)

					count += 1

					object = object.TailObject()
				}
			}
		}
		
		if 12 != count {
			t.Fatalf("[TestIndex12] Count '%d'",count)
		} else {
			fmt.Printf("[TestIndex12] Count '%d'\n",count)
		}
	}
}
