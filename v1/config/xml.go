package config

import (
	"fmt"
	"os"

	"github.com/beevik/etree"
	"github.com/gogf/gf/v2/encoding/gxml"
	"github.com/lovelacelee/clsgo/v1/utils"
)

func XmlDecode(content []byte) (map[string]interface{}, error) {
	return gxml.Decode(content)
}

func XmlDecodeWithoutRoot(content []byte) (map[string]interface{}, error) {
	return gxml.DecodeWithoutRoot(content)
}

func XmlEncode(m map[string]interface{}, rootTag ...string) ([]byte, error) {
	return gxml.Encode(m, rootTag...)
}

func XmlEncodeWithIndent(m map[string]interface{}, rootTag ...string) ([]byte, error) {
	return gxml.EncodeWithIndent(m, rootTag...)
}

func XmlToJson(content []byte) ([]byte, error) {
	return gxml.ToJson(content)
}

type XMLElement = etree.Element

// XML implement base on "github.com/beevik/etree"
type XML struct {
	Doc  *etree.Document
	File string //filepath
}

type XMLAttr struct {
	K string
	V string
}

// Title head is like "xml" or "xml-stylesheet",
// attrs is like `version="1.0" encoding="UTF-8"`, `type="text/xsl" href="style.xsl"`
func (xml *XML) AddTitle(node string, attrs ...string) *XML {
	xmlhead := utils.Param(attrs, `version="1.0" encoding="UTF-8"`)
	xml.Doc.CreateProcInst(node, xmlhead)
	return xml
}

// Element Chain operation is supported,
// If path cannot be found, then tag will be created under root tag
func (xml *XML) CreateElement(tag string, path string, value string, comment string, attrs ...XMLAttr) *XMLElement {
	if utils.IsEmpty(tag) {
		return nil
	}
	node := xml.Find(path)
	if node == nil {
		node = &xml.Doc.Element
	}
	if !utils.IsEmpty(comment) {
		node.CreateComment(comment)
	}
	t := NewElement(tag, value, attrs...)
	node.AddChild(t)
	return t
}

func (xml *XML) Dump(prefix, suffix string) {
	xml.Doc.WriteSettings.UseCRLF = true
	xml.Doc.WriteSettings.CanonicalAttrVal = true
	xml.Doc.WriteSettings.CanonicalEndTags = true
	// CannonicalText generate &#xD; suffix for every tag
	// xml.Doc.WriteSettings.CanonicalText = true
	fmt.Println(prefix)
	xml.Doc.WriteTo(os.Stdout)
	fmt.Println(suffix)
}

func (xml *XML) String() string {
	xml.Doc.Indent(2)
	s, err := xml.Doc.WriteToString()
	utils.IfErrorWithoutHeader(err, utils.ErrorWithoutHeader)
	return s
}

// Save to file --force
func (xml *XML) Save(filenamepath ...string) error {
	xml.Doc.Indent(2)
	file := utils.Param(filenamepath, xml.File)

	if utils.FileIsExisted(file) {
		os.Remove(file)
	}
	return xml.Doc.WriteToFile(file)
}

// Select the first one match
func (xml *XML) Get(tag string) *XMLElement {
	return xml.Doc.SelectElement(tag)
}

func (xml *XML) GetAll(tag string) []*XMLElement {
	return xml.Doc.SelectElements(tag)
}

// Finds the first element in the XPATH path
// p is like "./bookstore/book[p:price='49.99']/title"
func (xml *XML) Find(fmt string) *XMLElement {
	return xml.Doc.FindElement(fmt)
}

// Finds all the elements in the XPATH path
// fmt is like "//book[@category='WEB']/title"
// or "./bookstore/book[1]/*"
func (xml *XML) FindAll(fmt string) []*XMLElement {
	return xml.Doc.FindElements(fmt)
}

func NewElement(tag string, value string, attrs ...XMLAttr) *XMLElement {
	t := etree.NewElement(tag)
	if !utils.IsEmpty(attrs) {
		for _, attr := range attrs {
			t.CreateAttr(attr.K, attr.V)
		}
	}
	t.SetText(value)
	return t
}

func NewXML() *XML {
	xml := XML{
		Doc:  etree.NewDocument(),
		File: "output.xml",
	}
	// xml.Doc.ReadSettings.Permissive = true
	return &xml
}

func XMLFile(file ...string) *XML {
	xml := NewXML()
	xml.File = utils.Param(file, "output.xml")
	if utils.FileIsExisted(xml.File) {
		// error ignored, if any error occurred, the document will empty
		xml.Doc.ReadFromFile(xml.File)
	}
	return xml
}

func XMLString(s ...string) *XML {
	xml := NewXML()
	if !utils.IsEmpty(s) {
		xml.Doc.ReadFromString(utils.Param(s, ""))
	}
	return xml
}
