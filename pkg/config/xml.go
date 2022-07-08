/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-07-08 13:31:15
 * @LastEditTime    : 2022-07-08 13:38:51
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /pkg/config/xml.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package config

import (
	"github.com/gogf/gf/v2/encoding/gxml"
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
