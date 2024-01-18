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

func TestCondField(t *testing.T){
	var filename string = "tst/index1.json"
	var fo *os.File
	var er error

	fo, er = os.Open(filename)
	if nil != er {
		t.Fatalf("[TestCondField] Opening '%s': %v",filename,er)
	} else {
		defer fo.Close()
		var count int = 0

		var reader Reader = ReadFile(fo)
		if reader.IsNotEmpty() {

			var array Reader = reader.HeadArray()
			if array.IsNotEmpty() {

				var object Reader = array.HeadObject()
				for object.IsNotEmpty() {

					var field_id Reader = object.CondHeadField("id")
					if field_id.IsNotEmpty() && object.Contains(field_id) {

						count += 1

						var field_ic = field_id.CondTailField("icon")
						if field_ic.IsNotEmpty() && object.Contains(field_ic) {

							count += 1

							var field_pa = field_ic.CondTailField("path")
							if field_pa.IsNotEmpty() && object.Contains(field_pa) {

								count += 1

								var field_li = field_pa.CondTailField("link")
								if field_li.IsNotEmpty() && object.Contains(field_li) {

									count += 1

									var field_na = field_li.CondTailField("name")
									if field_na.IsNotEmpty() && object.Contains(field_na) {

										count += 1

										var field_em = field_na.CondTailField("embed")
										if field_em.IsNotEmpty() && object.Contains(field_em) {

											count += 1
										}
									}

								}

							}
						}
					}

					object = object.TailObject()
				}
			}
		}

		if 6 != count {
			t.Fatalf("[TestCondField] Count '%d'",count)
		} else {
			fmt.Printf("[TestCondField] Count '%d'\n",count)
		}
	}
}
