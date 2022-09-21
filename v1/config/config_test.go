package config_test

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/config"
	"github.com/lovelacelee/clsgo/v1/log"
	"github.com/lovelacelee/clsgo/v1/utils"
)

type JsonModel struct {
	Version string `json:"version"`
	Author  string `json:"author"`
	Github  string `json:"github"`
}

func clean() {
	utils.DeleteFiles(utils.Cwd(), "/*.yaml$")
	utils.DeleteFiles(utils.Cwd(), "/*.xml$")
	utils.DeleteFiles(utils.Cwd(), "/*.json$")
}

var data = `
project: clsgo
test:
  int: 1
  string: "test"
  bool: true
  duration: "1h"
  float32: 3.14
  float64: 3.1415926
  intslice:
    - 1
    - 2
    - 3
  stringslice:
    - "a"
    - "b"
    - "c"
  int64: 23492938579
`

func TestConfig(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Run("config", func(_ *testing.T) {
			t.AssertNE(config.Cfg, nil)
			config.CreateDefault("")
			// Test case use default config.yaml
			t.AssertNE(config.New("notexist", "test"), nil)
			// First time load, default value returned
			t.Assert(config.GetIntWithDefault("test.int", 0), 0)
			t.Assert(config.GetStringWithDefault("test.string", "nil"), "nil")
			t.Assert(config.GetBoolWithDefault("test.bool", false), false)
			t.Assert(config.GetDurationWithDefault("test.duration", time.Second), time.Second)
			t.Assert(config.GetFloat32WithDefault("test.float32", 0.8), 0.8)
			t.Assert(config.GetFloat64WithDefault("test.float64", 0.8888), 0.8888)
			t.Assert(config.GetIntSliceWithDefault("test.intslice", []int{3, 2, 1}), []int{3, 2, 1})
			t.Assert(config.GetStringSliceWithDefault("test.stringslice", []string{"c", "b", "a"}), []string{"c", "b", "a"})
			t.Assert(config.GetInt64WithDefault("test.int64", 8888), 8888)

			t.Assert(config.Cfg.GetInt("test.int"), 0)
			watched := config.New("config", "test")
			t.Assert(watched.GetString("test.string"), "")
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				// Do some change
				os.WriteFile("config.yaml", []byte(data), 0755)
				config.Cfg.SafeWriteConfig()
				time.Sleep(time.Microsecond * 50)
				wg.Done()
			}()
			wg.Wait() // Check whether fsnotify works well

			t.Assert(config.GetIntWithDefault("test.int", 0), 1)
			t.Assert(config.GetStringWithDefault("test.string", "nil"), "test")
			t.Assert(config.GetBoolWithDefault("test.bool", false), true)
			t.Assert(config.GetDurationWithDefault("test.duration", time.Second), time.Hour)
			t.Assert(config.GetFloat32WithDefault("test.float32", 0.8), 3.14)
			t.Assert(config.GetFloat64WithDefault("test.float64", 0.8888), 3.1415926)
			t.Assert(config.GetIntSliceWithDefault("test.intslice", []int{3, 2, 1}), []int{1, 2, 3})
			t.Assert(config.GetStringSliceWithDefault("test.stringslice", []string{"c", "b", "a"}), []string{"a", "b", "c"})
			t.Assert(config.GetInt64WithDefault("test.int64", 8888), 23492938579)

			// t.Assert(watched.GetString("test.string"), "test")
		})
		t.Run("json", func(_ *testing.T) {
			testInput := `{"author": "lovelacelee","github": "https://github.com/lovelacelee","version": "1.0.0"}`

			testDecode, err := config.JsonDecode(testInput)
			t.Assert(err, nil)
			t.Assert(testDecode.(map[string]any)["author"], "lovelacelee")

			testEncode, err := config.JsonEncode(testDecode)
			t.Assert(err, nil)
			mustEncode := config.JsonMustEncode(testDecode)
			t.Assert(bytes.Equal(testEncode, mustEncode), true)
			var mapDecode map[string]string
			t.Assert(config.JsonDecodeTo(testEncode, &mapDecode), nil)
			t.Assert(mapDecode["author"], testDecode.(map[string]any)["author"])

			encString1, err := config.JsonEncodeString(testDecode)
			t.Assert(err, nil)
			encString2 := config.JsonMustEncodeString(testDecode)
			t.Assert(encString1, encString2)

			t.Assert(config.JsonIsValidDataType(""), false)
			t.Assert(config.JsonIsValidDataType(".json"), true)
			t.Assert(config.JsonValid(testDecode), true)

			_, err = config.JsonMarshal(testDecode)
			t.Assert(err, nil)
			_, err = config.JsonMarshalIndent(testDecode, "", "    ")
			t.Assert(err, nil)

			var st JsonModel
			t.Assert(config.JsonUnmarshal([]byte(testInput), &st), nil)
			t.Assert(st.Author, "lovelacelee")

			f := config.NewJsonWith("test.json", testEncode)
			f.Save()
			nf := config.NewJson()
			nf.FromString(testInput)
			nf.OutputTo("nf.json")

			t.Assert(nf.Save("newnf.json"), nil)
			f.FromFile("newnf.json")
			var nfst JsonModel
			t.Assert(f.Unmarshal(&nfst), nil)
			t.Assert(nfst.Author, "lovelacelee")
			nfst.Author = "Lee"
			t.Assert(f.Marshal(&nfst), nil)
			t.AssertNE(f.FromFile("xx.json"), nil)
			log.Green("--> %s", f.String())
			t.AssertNE(f.Save("./test/test.json"), nil)
		})
		t.Run("xml", func(_ *testing.T) {
			xml := `<?xml version="1.0" encoding="utf-8"?>
			<resources xmlns:xliff="urn:oasis:names:tc:xliff:document:1.2">
				<integer name="config_mobile_mtu">1440</integer>
				<string-array translatable="false" name="config_operatorConsideredNonRoaming">
					<item>23410</item>
					<item>23426</item>
				</string-array>
			</resources>
			`
			json, err := config.XmlToJson([]byte(xml))
			t.Assert(err, nil)
			_, err = config.JsonDecode(json)
			t.Assert(err, nil)

			x := config.XMLString(xml)
			s := x.String()
			x = config.XMLString(s)
			// Get one
			t.AssertNE(x.Get("resources"), nil)
			t.AssertNE(x.Get("resources").SelectElement("integer"), nil)
			t.Assert(x.Get("resources").SelectElement("integer").Tag, "integer")
			t.Assert(x.Get("resources").SelectElement("integer").Text(), "1440")
			t.Assert(x.Get("integer"), nil)
			t.Assert(x.Find("resources/string-array/item[0]").Text(), "23410")
			// Get all
			t.Assert(len(x.GetAll("resources")), 1)
			t.Assert(len(x.FindAll("resources/string-array/item")), 2)
			// Modification
			// x.AddTitle("xml")
			t.AssertNE(x.CreateElement("nameinroot", "", "", ""), nil)
			t.Assert(x.CreateElement("", "", "", ""), nil)
			t.AssertNE(x.CreateElement("check", "tag", "", ""), nil)
			t.AssertNE(x.CreateElement("name", "resources", "lovelacelee", "test"), nil)
			t.AssertNE(x.CreateElement("name", "resources", "lovelacelee", "test", config.XMLAttr{K: "h", V: "178cm"}), nil)

			x.Dump("XML1:", "\n")
			x.Save()
			x.Save() // cover precondition: file exist
			x = config.XMLFile("output.xml")
			n := x.Get("nameinroot")
			t.AssertNE(x.Doc.RemoveChild(n), nil)

			x.Dump("XML2:", "\n")

			xn := config.NewXML()
			xn.AddTitle("xml")
			root := xn.CreateElement("root", "", "", "")
			root.CreateElement("Message").CreateElement("SN")
			root.SelectElement("Message").AddChild(
				config.NewElement("Device", "IPC", config.XMLAttr{K: "SN", V: "ABC"}))
			xn.Save("new.xml")

			ackMsg := make(map[string]interface{})
			ackMsg["UserName"] = "guest"
			ackMsg["UserType"] = "0"
			ackMsg["Ver"] = "2.0"
			ackMsg["Right"] = "65535"
			bytes, err := config.XmlEncode(ackMsg, "Message")
			t.Assert(err, nil)
			ackBytes, err := config.XmlEncodeWithIndent(ackMsg, "Message")
			t.Assert(err, nil)
			config.XmlDecodeWithoutRoot(ackBytes)
			config.XmlDecode(bytes)
			config.XmlToJson(bytes)
		})
	})

	clean()
}

func TestMain(m *testing.M) {

	clean()
	m.Run()
}

func ExampleNew() {
	cfg := config.New("server", "http")
	fmt.Print(cfg.Get("project.name"))
	clean()
	// Output:
	// http
}

// Use global unique config instance
func Example() {
	// import "github.com/lovelacelee/clsgo/config"
	config.CreateDefault("clsgo")
	fmt.Print(config.Default().Get("project.name"))
	clean()
	// Output:
	// clsgo
}
