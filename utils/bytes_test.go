package utils

import (
	"bytes"
	"testing"
)

type testCase struct {
	idHex   string
	vtype   VType
	length  uint8
	content interface{}
}

type testRockBLOCKCase struct {
	port    uint8
	idHex   string
	vtype   VType
	length  uint8
	content interface{}
}

func TestConvertBytes(t *testing.T) {
	testCases := map[testCase][]byte{
		{
			idHex:   "0x01",
			vtype:   "uint32",
			length:  4,
			content: float64(3600),
		}: {0x1, 0x4, 0x10, 0xe, 0x0, 0x0},
		{
			idHex:   "01",
			vtype:   "uint32",
			length:  4,
			content: float64(3600),
		}: {0x1, 0x4, 0x10, 0xe, 0x0, 0x0},
		{
			idHex:   "0x09",
			vtype:   "bool",
			length:  1,
			content: "false",
		}: {0x9, 0x1, 0x0},
		{
			idHex:   "0x0E",
			vtype:   "uint8",
			length:  1,
			content: float64(254),
		}: {0xe, 0x1, 0xfe},
		{
			idHex:   "0x10",
			vtype:   "byte_array",
			length:  16,
			content: "010203040506070910111213141516",
		}: {0x10, 0x10, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x9, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16},
		{
			idHex:   "0x21",
			vtype:   "string",
			length:  8,
			content: "abcdefhi",
		}: {0x21, 0x8, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x68, 0x69},
	}
	for tc, shouldBS := range testCases {
		bs, err := ConvertBytes(tc.idHex, tc.vtype, tc.length, tc.content)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(shouldBS, bs) {
			t.Errorf("%#v should be : %#v ,got %#v \n", tc, shouldBS, bs)
		}
	}

}

func TestConvertRockBLOCKBytes(t *testing.T) {
	testCases := map[testRockBLOCKCase][]byte{
		{
			port:    3,
			idHex:   "0x01",
			vtype:   "uint32",
			length:  4,
			content: float64(3600),
		}: {0x03, 0x1, 0x4, 0x10, 0xe, 0x0, 0x0},
		{
			port:    3,
			idHex:   "01",
			vtype:   "uint32",
			length:  4,
			content: float64(3600),
		}: {0x03, 0x1, 0x4, 0x10, 0xe, 0x0, 0x0},
		{
			port:    3,
			idHex:   "0x09",
			vtype:   "bool",
			length:  1,
			content: "false",
		}: {0x03, 0x9, 0x1, 0x0},
		{
			port:    3,
			idHex:   "0x0E",
			vtype:   "uint8",
			length:  1,
			content: float64(254),
		}: {0x03, 0xe, 0x1, 0xfe},
		{
			port:    3,
			idHex:   "0x10",
			vtype:   "byte_array",
			length:  16,
			content: "010203040506070910111213141516",
		}: {0x03, 0x10, 0x10, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x9, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16},
		{
			port:    32,
			idHex:   "0x21",
			vtype:   "string",
			length:  8,
			content: "abcdefhi",
		}: {0x20, 0x21, 0x8, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x68, 0x69},
	}
	for tc, shouldBS := range testCases {
		bs, err := ConvertRockBLOCKBytes(tc.port, tc.idHex, tc.vtype, tc.length, tc.content)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(shouldBS, bs) {
			t.Errorf("%#v should be : %#v ,got %#v \n", tc, shouldBS, bs)
		}
	}

}
