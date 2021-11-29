package main

import (
	"log"
	"math/rand"
	"os"
	"text/template"
	"time"
)

const (
	testTmpl = `// CODE GENERATED AUTOMATICALLY
// THIS FILE SHOULD NOT BE EDITED BY HAND
package main

import (
	"testing"
)

type TestStruct struct {
	m        map[string]interface{}
	expected MyStruct
}

func TestSetStructValuesSuccess(t *testing.T) {
	tests := []TestStruct{
		{{range .Structs}}
                TestStruct{
					m: map[string]interface{}{
						"I": {{.M.I}},
						"Y": {{.M.Y}},
					},
					expected: MyStruct{
						{{.Expected.I}},
						{{.Expected.Y}},
					},
				},
            {{end}}
	}

	ts := NewMyStruct()
	for _, v := range tests {
		err := SetStructValues(ts, v.m)
		if err != nil {
			t.Errorf("Ошибка «%v»\n", err)
		}

		if ts.I != v.expected.I  {
			t.Errorf("Ожидалось: «%v». Пришло: «%v»\n", v.expected.I, ts.I)
		}
		if ts.Y != v.expected.Y  {
			t.Errorf("Ожидалось: «%v». Пришло: «%v»\n", v.expected.Y, ts.Y)
		}
	}
}

func TestSetStructValuesIntFailed(t *testing.T) {
	tests := []TestStruct{
		{{range .FailedIntStructs}}
                TestStruct{
					m: map[string]interface{}{
						"I": {{.M.I}},
						"Y": {{.M.Y}},
					},
					expected: MyStruct{
						{{.Expected.I}},
						{{.Expected.Y}},
					},
				},
            {{end}}
	}

	ts := NewMyStruct()
	for _, v := range tests {
		err := SetStructValues(ts, v.m)
		if err != nil {
			t.Errorf("Ошибка «%v»\n", err)
		}
		if ts.I == v.expected.I  {
			t.Errorf("Ожидалось: «%v». Пришло: «%v»", v.expected.I, ts.I)
		}
	}
}

func TestSetStructValuesBoolFailed(t *testing.T) {
	tests := []TestStruct{
		{{range .FailedBoolStructs}}
                TestStruct{
					m: map[string]interface{}{
						"I": {{.M.I}},
						"Y": {{.M.Y}},
					},
					expected: MyStruct{
						{{.Expected.I}},
						{{.Expected.Y}},
					},
				},
            {{end}}
	}

	ts := NewMyStruct()
	for _, v := range tests {
		err := SetStructValues(ts, v.m)
		if err != nil {
			t.Errorf("Ошибка «%v»\n", err)
		}
		if ts.Y == v.expected.Y  {
			t.Errorf("Ожидалось: «%v». Пришло: «%v»", v.expected.Y, ts.Y)
		}
	}
}
`
	genFileName = "../task1_gen_test.go"
)

type MyTestStruct struct {
	I int  `yaml:"i" json:"i"`
	Y bool `yaml:"y" json:"y"`
}

type GenTestStruct struct {
	M        map[string]interface{}
	Expected MyTestStruct
}

var booList = [2]bool{true, false}

func main() {
	genFile, err := os.Create(genFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer genFile.Close()

	templateData := struct {
		Structs           []GenTestStruct
		FailedIntStructs  []GenTestStruct
		FailedBoolStructs []GenTestStruct
	}{}

	var y GenTestStruct

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		randI := rand.Int()

		rand.Seed(time.Now().UnixNano())
		randB := booList[rand.Intn(2)]
		x := GenTestStruct{
			M: map[string]interface{}{
				"I": randI,
				"Y": randB,
			},
			Expected: MyTestStruct{
				randI,
				randB,
			},
		}

		if randB {
			y = GenTestStruct{
				M: map[string]interface{}{
					"I": randI,
					"Y": !randB,
				},
				Expected: MyTestStruct{
					randI,
					randB,
				},
			}
			templateData.FailedBoolStructs = append(templateData.FailedBoolStructs, y)
		} else {
			y = GenTestStruct{
				M: map[string]interface{}{
					"I": randI + 1,
					"Y": randB,
				},
				Expected: MyTestStruct{
					randI,
					randB,
				},
			}

			templateData.FailedIntStructs = append(templateData.FailedIntStructs, y)

		}

		templateData.Structs = append(templateData.Structs, x)
	}

	t := template.Must(template.New("const-list").Parse(testTmpl))
	if err := t.Execute(genFile, templateData); err != nil {
		log.Fatal(err)
	}

}
