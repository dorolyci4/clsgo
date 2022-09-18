package config_test

import (
	"bytes"
	"fmt"
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
	utils.DeletePath("logs")
	utils.DeleteFiles(utils.Cwd(), "/*.yaml$")
	utils.DeleteFiles(utils.Cwd(), "/*.xml$")
}
func TestConfig(t *testing.T) {
	clean()
	log.Green("Running config test cases")
	gtest.C(t, func(t *gtest.T) {
		t.Run("config", func(ot *testing.T) {
			// Test case use default config.yaml
			t.AssertNE(config.ClsConfig("notexist", "notexist", true), nil)
			t.AssertNE(config.ClsConfig("notexist", "notexist", false), nil)
			t.AssertNE(config.Cfg, nil)
			// First time load, default value returned
			config.GetIntWithDefault("logger.rotatebackuplimit", 0)
			config.GetIntWithDefault("logger.rotatebackuplimit", 1)
			config.GetStringWithDefault("logger.rotateSize", "2MB")
			config.GetDurationWithDefault("logger.rotateCheckInterval", time.Minute)
			config.GetBoolWithDefault("logger.writerColorEnable", true)
			config.GetDurationWithDefault("server.tcpTimeout", time.Duration(1))
			config.GetStringWithDefault("server.openapiPath", "/api.json")
			config.GetFloat32WithDefault("logger.stStatus", 0.1)
			config.GetFloat64WithDefault("logger.stStatus", 0.1)
			config.GetInt64WithDefault("logger.stStatus", 0)
			config.GetStringSliceWithDefault("logger.stStatus", []string{})
			config.GetIntSliceWithDefault("logger.stStatus", []int{})
			t.Assert(config.Get("logger"), nil)
		})
		t.Run("json", func(ot *testing.T) {
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
			t.Assert(config.JsonIsValidDataType(".xml"), true)
			t.Assert(config.JsonValid(testDecode), true)

			_, err = config.JsonMarshal(testDecode)
			t.Assert(err, nil)
			_, err = config.JsonMarshalIndent(testDecode, "", "    ")
			t.Assert(err, nil)

			t.Assert(config.JsonUnmarshal([]byte(testInput), &mapDecode), nil)
		})
		t.Run("xml", func(ot *testing.T) {
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
			t.AssertNE(x.CreateElement("name", "resources", "lovelacelee", "test"), nil)
			t.AssertNE(x.CreateElement("name", "resources", "lovelacelee", "test", config.XMLAttr{K: "h", V: "178cm"}), nil)

			x.Dump("XML1:", "\n")
			x.Save()
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

func ExampleClsConfig() {
	cfg := config.ClsConfig("server", "http", true)
	fmt.Print(cfg.Get("project.name"))
	clean()
	// Output:
	// http
}

// Use global unique config instance
func Example() {
	// import "github.com/lovelacelee/clsgo/config"
	config.Init("clsgo")
	fmt.Print(config.Cfg.Get("project.name"))
	clean()
	// Output:
	// clsgo
}
