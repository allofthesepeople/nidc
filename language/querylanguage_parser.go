// Code generated from QueryLanguage.g4 by ANTLR 4.7.1. DO NOT EDIT.

package language // QueryLanguage
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 278, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2, 
	3, 2, 3, 3, 3, 3, 6, 3, 29, 10, 3, 13, 3, 14, 3, 30, 3, 3, 3, 3, 7, 3, 
	35, 10, 3, 12, 3, 14, 3, 38, 11, 3, 3, 3, 3, 3, 7, 3, 42, 10, 3, 12, 3, 
	14, 3, 45, 11, 3, 3, 3, 3, 3, 6, 3, 49, 10, 3, 13, 3, 14, 3, 50, 5, 3, 
	53, 10, 3, 3, 3, 6, 3, 56, 10, 3, 13, 3, 14, 3, 57, 3, 3, 3, 3, 6, 3, 62, 
	10, 3, 13, 3, 14, 3, 63, 3, 3, 3, 3, 7, 3, 68, 10, 3, 12, 3, 14, 3, 71, 
	11, 3, 3, 3, 3, 3, 7, 3, 75, 10, 3, 12, 3, 14, 3, 78, 11, 3, 3, 3, 5, 3, 
	81, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 88, 10, 4, 12, 4, 14, 4, 
	91, 11, 4, 3, 4, 3, 4, 7, 4, 95, 10, 4, 12, 4, 14, 4, 98, 11, 4, 3, 4, 
	3, 4, 7, 4, 102, 10, 4, 12, 4, 14, 4, 105, 11, 4, 3, 4, 3, 4, 7, 4, 109, 
	10, 4, 12, 4, 14, 4, 112, 11, 4, 3, 4, 5, 4, 115, 10, 4, 3, 4, 7, 4, 118, 
	10, 4, 12, 4, 14, 4, 121, 11, 4, 3, 4, 3, 4, 5, 4, 125, 10, 4, 3, 4, 3, 
	4, 3, 5, 3, 5, 7, 5, 131, 10, 5, 12, 5, 14, 5, 134, 11, 5, 3, 5, 3, 5, 
	7, 5, 138, 10, 5, 12, 5, 14, 5, 141, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 
	6, 7, 6, 148, 10, 6, 12, 6, 14, 6, 151, 11, 6, 3, 6, 3, 6, 7, 6, 155, 10, 
	6, 12, 6, 14, 6, 158, 11, 6, 3, 6, 3, 6, 7, 6, 162, 10, 6, 12, 6, 14, 6, 
	165, 11, 6, 3, 6, 3, 6, 7, 6, 169, 10, 6, 12, 6, 14, 6, 172, 11, 6, 3, 
	6, 5, 6, 175, 10, 6, 3, 6, 7, 6, 178, 10, 6, 12, 6, 14, 6, 181, 11, 6, 
	3, 6, 3, 6, 5, 6, 185, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 192, 
	10, 6, 12, 6, 14, 6, 195, 11, 6, 3, 6, 3, 6, 7, 6, 199, 10, 6, 12, 6, 14, 
	6, 202, 11, 6, 3, 6, 3, 6, 7, 6, 206, 10, 6, 12, 6, 14, 6, 209, 11, 6, 
	3, 6, 3, 6, 7, 6, 213, 10, 6, 12, 6, 14, 6, 216, 11, 6, 3, 6, 5, 6, 219, 
	10, 6, 3, 6, 7, 6, 222, 10, 6, 12, 6, 14, 6, 225, 11, 6, 3, 6, 3, 6, 5, 
	6, 229, 10, 6, 3, 6, 3, 6, 5, 6, 233, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 
	7, 5, 7, 240, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 248, 10, 
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 7, 11, 256, 10, 11, 12, 11, 
	14, 11, 259, 11, 11, 3, 11, 3, 11, 6, 11, 263, 10, 11, 13, 11, 14, 11, 
	264, 3, 11, 6, 11, 268, 10, 11, 13, 11, 14, 11, 269, 3, 11, 3, 11, 5, 11, 
	274, 10, 11, 3, 12, 3, 12, 3, 12, 2, 2, 13, 2, 4, 6, 8, 10, 12, 14, 16, 
	18, 20, 22, 2, 2, 2, 306, 2, 24, 3, 2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 84, 
	3, 2, 2, 2, 8, 128, 3, 2, 2, 2, 10, 232, 3, 2, 2, 2, 12, 239, 3, 2, 2, 
	2, 14, 247, 3, 2, 2, 2, 16, 249, 3, 2, 2, 2, 18, 251, 3, 2, 2, 2, 20, 273, 
	3, 2, 2, 2, 22, 275, 3, 2, 2, 2, 24, 25, 5, 4, 3, 2, 25, 3, 3, 2, 2, 2, 
	26, 28, 7, 15, 2, 2, 27, 29, 7, 17, 2, 2, 28, 27, 3, 2, 2, 2, 29, 30, 3, 
	2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 
	36, 5, 6, 4, 2, 33, 35, 7, 17, 2, 2, 34, 33, 3, 2, 2, 2, 35, 38, 3, 2, 
	2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 52, 3, 2, 2, 2, 38, 36, 
	3, 2, 2, 2, 39, 43, 5, 10, 6, 2, 40, 42, 7, 17, 2, 2, 41, 40, 3, 2, 2, 
	2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 46, 
	3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 47, 5, 6, 4, 2, 47, 49, 3, 2, 2, 2, 
	48, 39, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 
	2, 2, 2, 51, 53, 3, 2, 2, 2, 52, 48, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 
	55, 3, 2, 2, 2, 54, 56, 7, 17, 2, 2, 55, 54, 3, 2, 2, 2, 56, 57, 3, 2, 
	2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61, 
	7, 16, 2, 2, 60, 62, 7, 17, 2, 2, 61, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 
	2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 80, 
	5, 12, 7, 2, 66, 68, 7, 17, 2, 2, 67, 66, 3, 2, 2, 2, 68, 71, 3, 2, 2, 
	2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 72, 3, 2, 2, 2, 71, 69, 
	3, 2, 2, 2, 72, 76, 7, 3, 2, 2, 73, 75, 7, 17, 2, 2, 74, 73, 3, 2, 2, 2, 
	75, 78, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 79, 3, 
	2, 2, 2, 78, 76, 3, 2, 2, 2, 79, 81, 5, 12, 7, 2, 80, 69, 3, 2, 2, 2, 80, 
	81, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 7, 2, 2, 3, 83, 5, 3, 2, 2, 
	2, 84, 85, 7, 4, 2, 2, 85, 124, 5, 14, 8, 2, 86, 88, 7, 17, 2, 2, 87, 86, 
	3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 
	90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 96, 7, 5, 2, 2, 93, 95, 7, 
	17, 2, 2, 94, 93, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 
	97, 3, 2, 2, 2, 97, 99, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99, 114, 5, 8, 
	5, 2, 100, 102, 7, 17, 2, 2, 101, 100, 3, 2, 2, 2, 102, 105, 3, 2, 2, 2, 
	103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 106, 3, 2, 2, 2, 105, 
	103, 3, 2, 2, 2, 106, 110, 7, 3, 2, 2, 107, 109, 7, 17, 2, 2, 108, 107, 
	3, 2, 2, 2, 109, 112, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110, 111, 3, 2, 
	2, 2, 111, 113, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 113, 115, 5, 8, 5, 2, 
	114, 103, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 119, 3, 2, 2, 2, 116, 
	118, 7, 17, 2, 2, 117, 116, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 
	3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 122, 3, 2, 2, 2, 121, 119, 3, 2, 
	2, 2, 122, 123, 7, 6, 2, 2, 123, 125, 3, 2, 2, 2, 124, 89, 3, 2, 2, 2, 
	124, 125, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 7, 7, 2, 2, 127, 
	7, 3, 2, 2, 2, 128, 132, 5, 18, 10, 2, 129, 131, 7, 17, 2, 2, 130, 129, 
	3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 
	2, 2, 133, 135, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 135, 139, 7, 8, 2, 2, 
	136, 138, 7, 17, 2, 2, 137, 136, 3, 2, 2, 2, 138, 141, 3, 2, 2, 2, 139, 
	137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 142, 3, 2, 2, 2, 141, 139, 
	3, 2, 2, 2, 142, 143, 5, 20, 11, 2, 143, 9, 3, 2, 2, 2, 144, 145, 7, 9, 
	2, 2, 145, 184, 5, 14, 8, 2, 146, 148, 7, 17, 2, 2, 147, 146, 3, 2, 2, 
	2, 148, 151, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 
	152, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 152, 156, 7, 5, 2, 2, 153, 155, 
	7, 17, 2, 2, 154, 153, 3, 2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2, 
	2, 2, 156, 157, 3, 2, 2, 2, 157, 159, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 
	159, 174, 5, 8, 5, 2, 160, 162, 7, 17, 2, 2, 161, 160, 3, 2, 2, 2, 162, 
	165, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 166, 
	3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 166, 170, 7, 3, 2, 2, 167, 169, 7, 17, 
	2, 2, 168, 167, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 
	170, 171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 173, 
	175, 5, 8, 5, 2, 174, 163, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 179, 
	3, 2, 2, 2, 176, 178, 7, 17, 2, 2, 177, 176, 3, 2, 2, 2, 178, 181, 3, 2, 
	2, 2, 179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 182, 3, 2, 2, 2, 
	181, 179, 3, 2, 2, 2, 182, 183, 7, 6, 2, 2, 183, 185, 3, 2, 2, 2, 184, 
	149, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 187, 
	7, 10, 2, 2, 187, 233, 3, 2, 2, 2, 188, 189, 7, 11, 2, 2, 189, 228, 5, 
	14, 8, 2, 190, 192, 7, 17, 2, 2, 191, 190, 3, 2, 2, 2, 192, 195, 3, 2, 
	2, 2, 193, 191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 196, 3, 2, 2, 2, 
	195, 193, 3, 2, 2, 2, 196, 200, 7, 5, 2, 2, 197, 199, 7, 17, 2, 2, 198, 
	197, 3, 2, 2, 2, 199, 202, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 
	3, 2, 2, 2, 201, 203, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 203, 218, 5, 8, 
	5, 2, 204, 206, 7, 17, 2, 2, 205, 204, 3, 2, 2, 2, 206, 209, 3, 2, 2, 2, 
	207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 210, 3, 2, 2, 2, 209, 
	207, 3, 2, 2, 2, 210, 214, 7, 3, 2, 2, 211, 213, 7, 17, 2, 2, 212, 211, 
	3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 214, 215, 3, 2, 
	2, 2, 215, 217, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 217, 219, 5, 8, 5, 2, 
	218, 207, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 223, 3, 2, 2, 2, 220, 
	222, 7, 17, 2, 2, 221, 220, 3, 2, 2, 2, 222, 225, 3, 2, 2, 2, 223, 221, 
	3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 226, 3, 2, 2, 2, 225, 223, 3, 2, 
	2, 2, 226, 227, 7, 6, 2, 2, 227, 229, 3, 2, 2, 2, 228, 193, 3, 2, 2, 2, 
	228, 229, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 231, 7, 12, 2, 2, 231, 
	233, 3, 2, 2, 2, 232, 144, 3, 2, 2, 2, 232, 188, 3, 2, 2, 2, 233, 11, 3, 
	2, 2, 2, 234, 240, 5, 16, 9, 2, 235, 236, 5, 16, 9, 2, 236, 237, 7, 13, 
	2, 2, 237, 238, 5, 18, 10, 2, 238, 240, 3, 2, 2, 2, 239, 234, 3, 2, 2, 
	2, 239, 235, 3, 2, 2, 2, 240, 13, 3, 2, 2, 2, 241, 242, 5, 16, 9, 2, 242, 
	243, 7, 8, 2, 2, 243, 244, 5, 22, 12, 2, 244, 248, 3, 2, 2, 2, 245, 246, 
	7, 8, 2, 2, 246, 248, 5, 22, 12, 2, 247, 241, 3, 2, 2, 2, 247, 245, 3, 
	2, 2, 2, 248, 15, 3, 2, 2, 2, 249, 250, 7, 19, 2, 2, 250, 17, 3, 2, 2, 
	2, 251, 252, 7, 19, 2, 2, 252, 19, 3, 2, 2, 2, 253, 257, 7, 14, 2, 2, 254, 
	256, 7, 17, 2, 2, 255, 254, 3, 2, 2, 2, 256, 259, 3, 2, 2, 2, 257, 255, 
	3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 260, 3, 2, 2, 2, 259, 257, 3, 2, 
	2, 2, 260, 267, 7, 19, 2, 2, 261, 263, 7, 17, 2, 2, 262, 261, 3, 2, 2, 
	2, 263, 264, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 
	266, 3, 2, 2, 2, 266, 268, 7, 19, 2, 2, 267, 262, 3, 2, 2, 2, 268, 269, 
	3, 2, 2, 2, 269, 267, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 271, 3, 2, 
	2, 2, 271, 274, 7, 14, 2, 2, 272, 274, 7, 18, 2, 2, 273, 253, 3, 2, 2, 
	2, 273, 272, 3, 2, 2, 2, 274, 21, 3, 2, 2, 2, 275, 276, 7, 19, 2, 2, 276, 
	23, 3, 2, 2, 2, 42, 30, 36, 43, 50, 52, 57, 63, 69, 76, 80, 89, 96, 103, 
	110, 114, 119, 124, 132, 139, 149, 156, 163, 170, 174, 179, 184, 193, 200, 
	207, 214, 218, 223, 228, 232, 239, 247, 257, 264, 269, 273,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "'('", "'{'", "'}'", "')'", "':'", "'<-['", "']-'", "'-['", 
	"']->'", "'.'", "'\"'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "Match", "Return", 
	"WS", "NumberText", "AnyText",
}

var ruleNames = []string{
	"query", "matchClause", "node", "filter", "relationship", "returnValue", 
	"aliasIdentity", "alias", "field", "fieldValue", "nodeName",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type QueryLanguageParser struct {
	*antlr.BaseParser
}

func NewQueryLanguageParser(input antlr.TokenStream) *QueryLanguageParser {
	this := new(QueryLanguageParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "QueryLanguage.g4"

	return this
}

// QueryLanguageParser tokens.
const (
	QueryLanguageParserEOF = antlr.TokenEOF
	QueryLanguageParserT__0 = 1
	QueryLanguageParserT__1 = 2
	QueryLanguageParserT__2 = 3
	QueryLanguageParserT__3 = 4
	QueryLanguageParserT__4 = 5
	QueryLanguageParserT__5 = 6
	QueryLanguageParserT__6 = 7
	QueryLanguageParserT__7 = 8
	QueryLanguageParserT__8 = 9
	QueryLanguageParserT__9 = 10
	QueryLanguageParserT__10 = 11
	QueryLanguageParserT__11 = 12
	QueryLanguageParserMatch = 13
	QueryLanguageParserReturn = 14
	QueryLanguageParserWS = 15
	QueryLanguageParserNumberText = 16
	QueryLanguageParserAnyText = 17
)

// QueryLanguageParser rules.
const (
	QueryLanguageParserRULE_query = 0
	QueryLanguageParserRULE_matchClause = 1
	QueryLanguageParserRULE_node = 2
	QueryLanguageParserRULE_filter = 3
	QueryLanguageParserRULE_relationship = 4
	QueryLanguageParserRULE_returnValue = 5
	QueryLanguageParserRULE_aliasIdentity = 6
	QueryLanguageParserRULE_alias = 7
	QueryLanguageParserRULE_field = 8
	QueryLanguageParserRULE_fieldValue = 9
	QueryLanguageParserRULE_nodeName = 10
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) MatchClause() IMatchClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatchClauseContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *QueryLanguageParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QueryLanguageParserRULE_query)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.MatchClause()
	}



	return localctx
}


// IMatchClauseContext is an interface to support dynamic dispatch.
type IMatchClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMatchClauseContext differentiates from other interfaces.
	IsMatchClauseContext()
}

type MatchClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchClauseContext() *MatchClauseContext {
	var p = new(MatchClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_matchClause
	return p
}

func (*MatchClauseContext) IsMatchClauseContext() {}

func NewMatchClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchClauseContext {
	var p = new(MatchClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_matchClause

	return p
}

func (s *MatchClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchClauseContext) Match() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserMatch, 0)
}

func (s *MatchClauseContext) AllNode() []INodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INodeContext)(nil)).Elem())
	var tst = make([]INodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INodeContext)
		}
	}

	return tst
}

func (s *MatchClauseContext) Node(i int) INodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INodeContext)
}

func (s *MatchClauseContext) Return() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserReturn, 0)
}

func (s *MatchClauseContext) AllReturnValue() []IReturnValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReturnValueContext)(nil)).Elem())
	var tst = make([]IReturnValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReturnValueContext)
		}
	}

	return tst
}

func (s *MatchClauseContext) ReturnValue(i int) IReturnValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReturnValueContext)
}

func (s *MatchClauseContext) EOF() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserEOF, 0)
}

func (s *MatchClauseContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(QueryLanguageParserWS)
}

func (s *MatchClauseContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserWS, i)
}

func (s *MatchClauseContext) AllRelationship() []IRelationshipContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelationshipContext)(nil)).Elem())
	var tst = make([]IRelationshipContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelationshipContext)
		}
	}

	return tst
}

func (s *MatchClauseContext) Relationship(i int) IRelationshipContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationshipContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelationshipContext)
}

func (s *MatchClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MatchClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterMatchClause(s)
	}
}

func (s *MatchClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitMatchClause(s)
	}
}

func (s *MatchClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageVisitor:
		return t.VisitMatchClause(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *QueryLanguageParser) MatchClause() (localctx IMatchClauseContext) {
	localctx = NewMatchClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QueryLanguageParserRULE_matchClause)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(QueryLanguageParserMatch)
	}
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == QueryLanguageParserWS {
		{
			p.SetState(25)
			p.Match(QueryLanguageParserWS)
		}


		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(30)
		p.Node()
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(31)
				p.Match(QueryLanguageParserWS)
			}


		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == QueryLanguageParserT__6 || _la == QueryLanguageParserT__8 {
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == QueryLanguageParserT__6 || _la == QueryLanguageParserT__8 {
			{
				p.SetState(37)
				p.Relationship()
			}
			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for _la == QueryLanguageParserWS {
				{
					p.SetState(38)
					p.Match(QueryLanguageParserWS)
				}


				p.SetState(43)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(44)
				p.Node()
			}


			p.SetState(48)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == QueryLanguageParserWS {
		{
			p.SetState(52)
			p.Match(QueryLanguageParserWS)
		}


		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
		p.Match(QueryLanguageParserReturn)
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == QueryLanguageParserWS {
		{
			p.SetState(58)
			p.Match(QueryLanguageParserWS)
		}


		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(63)
		p.ReturnValue()
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == QueryLanguageParserT__0 || _la == QueryLanguageParserWS {
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == QueryLanguageParserWS {
			{
				p.SetState(64)
				p.Match(QueryLanguageParserWS)
			}


			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(70)
			p.Match(QueryLanguageParserT__0)
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == QueryLanguageParserWS {
			{
				p.SetState(71)
				p.Match(QueryLanguageParserWS)
			}


			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(77)
			p.ReturnValue()
		}

	}
	{
		p.SetState(80)
		p.Match(QueryLanguageParserEOF)
	}



	return localctx
}


// INodeContext is an interface to support dynamic dispatch.
type INodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNodeContext differentiates from other interfaces.
	IsNodeContext()
}

type NodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeContext() *NodeContext {
	var p = new(NodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_node
	return p
}

func (*NodeContext) IsNodeContext() {}

func NewNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeContext {
	var p = new(NodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_node

	return p
}

func (s *NodeContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeContext) AliasIdentity() IAliasIdentityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasIdentityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasIdentityContext)
}

func (s *NodeContext) AllFilter() []IFilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilterContext)(nil)).Elem())
	var tst = make([]IFilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilterContext)
		}
	}

	return tst
}

func (s *NodeContext) Filter(i int) IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *NodeContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(QueryLanguageParserWS)
}

func (s *NodeContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserWS, i)
}

func (s *NodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterNode(s)
	}
}

func (s *NodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitNode(s)
	}
}

func (s *NodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageVisitor:
		return t.VisitNode(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *QueryLanguageParser) Node() (localctx INodeContext) {
	localctx = NewNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QueryLanguageParserRULE_node)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(QueryLanguageParserT__1)
	}
	{
		p.SetState(83)
		p.AliasIdentity()
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == QueryLanguageParserT__2 || _la == QueryLanguageParserWS {
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == QueryLanguageParserWS {
			{
				p.SetState(84)
				p.Match(QueryLanguageParserWS)
			}


			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(90)
			p.Match(QueryLanguageParserT__2)
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == QueryLanguageParserWS {
			{
				p.SetState(91)
				p.Match(QueryLanguageParserWS)
			}


			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(97)
			p.Filter()
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for _la == QueryLanguageParserWS {
				{
					p.SetState(98)
					p.Match(QueryLanguageParserWS)
				}


				p.SetState(103)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(104)
				p.Match(QueryLanguageParserT__0)
			}
			p.SetState(108)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for _la == QueryLanguageParserWS {
				{
					p.SetState(105)
					p.Match(QueryLanguageParserWS)
				}


				p.SetState(110)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(111)
				p.Filter()
			}


		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == QueryLanguageParserWS {
			{
				p.SetState(114)
				p.Match(QueryLanguageParserWS)
			}


			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(120)
			p.Match(QueryLanguageParserT__3)
		}

	}
	{
		p.SetState(124)
		p.Match(QueryLanguageParserT__4)
	}



	return localctx
}


// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FilterContext) FieldValue() IFieldValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldValueContext)
}

func (s *FilterContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(QueryLanguageParserWS)
}

func (s *FilterContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserWS, i)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (s *FilterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageVisitor:
		return t.VisitFilter(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *QueryLanguageParser) Filter() (localctx IFilterContext) {
	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QueryLanguageParserRULE_filter)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Field()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == QueryLanguageParserWS {
		{
			p.SetState(127)
			p.Match(QueryLanguageParserWS)
		}


		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(133)
		p.Match(QueryLanguageParserT__5)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == QueryLanguageParserWS {
		{
			p.SetState(134)
			p.Match(QueryLanguageParserWS)
		}


		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(140)
		p.FieldValue()
	}



	return localctx
}


// IRelationshipContext is an interface to support dynamic dispatch.
type IRelationshipContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationshipContext differentiates from other interfaces.
	IsRelationshipContext()
}

type RelationshipContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationshipContext() *RelationshipContext {
	var p = new(RelationshipContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_relationship
	return p
}

func (*RelationshipContext) IsRelationshipContext() {}

func NewRelationshipContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationshipContext {
	var p = new(RelationshipContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_relationship

	return p
}

func (s *RelationshipContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationshipContext) AliasIdentity() IAliasIdentityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasIdentityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasIdentityContext)
}

func (s *RelationshipContext) AllFilter() []IFilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilterContext)(nil)).Elem())
	var tst = make([]IFilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilterContext)
		}
	}

	return tst
}

func (s *RelationshipContext) Filter(i int) IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *RelationshipContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(QueryLanguageParserWS)
}

func (s *RelationshipContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserWS, i)
}

func (s *RelationshipContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationshipContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RelationshipContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterRelationship(s)
	}
}

func (s *RelationshipContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitRelationship(s)
	}
}

func (s *RelationshipContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageVisitor:
		return t.VisitRelationship(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *QueryLanguageParser) Relationship() (localctx IRelationshipContext) {
	localctx = NewRelationshipContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QueryLanguageParserRULE_relationship)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(230)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QueryLanguageParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Match(QueryLanguageParserT__6)
		}
		{
			p.SetState(143)
			p.AliasIdentity()
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == QueryLanguageParserT__2 || _la == QueryLanguageParserWS {
			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for _la == QueryLanguageParserWS {
				{
					p.SetState(144)
					p.Match(QueryLanguageParserWS)
				}


				p.SetState(149)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(150)
				p.Match(QueryLanguageParserT__2)
			}
			p.SetState(154)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for _la == QueryLanguageParserWS {
				{
					p.SetState(151)
					p.Match(QueryLanguageParserWS)
				}


				p.SetState(156)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(157)
				p.Filter()
			}
			p.SetState(172)
			p.GetErrorHandler().Sync(p)


			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
				p.SetState(161)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)


				for _la == QueryLanguageParserWS {
					{
						p.SetState(158)
						p.Match(QueryLanguageParserWS)
					}


					p.SetState(163)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(164)
					p.Match(QueryLanguageParserT__0)
				}
				p.SetState(168)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)


				for _la == QueryLanguageParserWS {
					{
						p.SetState(165)
						p.Match(QueryLanguageParserWS)
					}


					p.SetState(170)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(171)
					p.Filter()
				}


			}
			p.SetState(177)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for _la == QueryLanguageParserWS {
				{
					p.SetState(174)
					p.Match(QueryLanguageParserWS)
				}


				p.SetState(179)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(180)
				p.Match(QueryLanguageParserT__3)
			}

		}
		{
			p.SetState(184)
			p.Match(QueryLanguageParserT__7)
		}


	case QueryLanguageParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.Match(QueryLanguageParserT__8)
		}
		{
			p.SetState(187)
			p.AliasIdentity()
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == QueryLanguageParserT__2 || _la == QueryLanguageParserWS {
			p.SetState(191)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for _la == QueryLanguageParserWS {
				{
					p.SetState(188)
					p.Match(QueryLanguageParserWS)
				}


				p.SetState(193)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(194)
				p.Match(QueryLanguageParserT__2)
			}
			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for _la == QueryLanguageParserWS {
				{
					p.SetState(195)
					p.Match(QueryLanguageParserWS)
				}


				p.SetState(200)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(201)
				p.Filter()
			}
			p.SetState(216)
			p.GetErrorHandler().Sync(p)


			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
				p.SetState(205)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)


				for _la == QueryLanguageParserWS {
					{
						p.SetState(202)
						p.Match(QueryLanguageParserWS)
					}


					p.SetState(207)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(208)
					p.Match(QueryLanguageParserT__0)
				}
				p.SetState(212)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)


				for _la == QueryLanguageParserWS {
					{
						p.SetState(209)
						p.Match(QueryLanguageParserWS)
					}


					p.SetState(214)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(215)
					p.Filter()
				}


			}
			p.SetState(221)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for _la == QueryLanguageParserWS {
				{
					p.SetState(218)
					p.Match(QueryLanguageParserWS)
				}


				p.SetState(223)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(224)
				p.Match(QueryLanguageParserT__3)
			}

		}
		{
			p.SetState(228)
			p.Match(QueryLanguageParserT__9)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IReturnValueContext is an interface to support dynamic dispatch.
type IReturnValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnValueContext differentiates from other interfaces.
	IsReturnValueContext()
}

type ReturnValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnValueContext() *ReturnValueContext {
	var p = new(ReturnValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_returnValue
	return p
}

func (*ReturnValueContext) IsReturnValueContext() {}

func NewReturnValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnValueContext {
	var p = new(ReturnValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_returnValue

	return p
}

func (s *ReturnValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnValueContext) Alias() IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *ReturnValueContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *ReturnValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ReturnValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterReturnValue(s)
	}
}

func (s *ReturnValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitReturnValue(s)
	}
}

func (s *ReturnValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageVisitor:
		return t.VisitReturnValue(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *QueryLanguageParser) ReturnValue() (localctx IReturnValueContext) {
	localctx = NewReturnValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, QueryLanguageParserRULE_returnValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Alias()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.Alias()
		}
		{
			p.SetState(234)
			p.Match(QueryLanguageParserT__10)
		}
		{
			p.SetState(235)
			p.Field()
		}

	}


	return localctx
}


// IAliasIdentityContext is an interface to support dynamic dispatch.
type IAliasIdentityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasIdentityContext differentiates from other interfaces.
	IsAliasIdentityContext()
}

type AliasIdentityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasIdentityContext() *AliasIdentityContext {
	var p = new(AliasIdentityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_aliasIdentity
	return p
}

func (*AliasIdentityContext) IsAliasIdentityContext() {}

func NewAliasIdentityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasIdentityContext {
	var p = new(AliasIdentityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_aliasIdentity

	return p
}

func (s *AliasIdentityContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasIdentityContext) Alias() IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *AliasIdentityContext) NodeName() INodeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INodeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INodeNameContext)
}

func (s *AliasIdentityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasIdentityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AliasIdentityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterAliasIdentity(s)
	}
}

func (s *AliasIdentityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitAliasIdentity(s)
	}
}

func (s *AliasIdentityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageVisitor:
		return t.VisitAliasIdentity(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *QueryLanguageParser) AliasIdentity() (localctx IAliasIdentityContext) {
	localctx = NewAliasIdentityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, QueryLanguageParserRULE_aliasIdentity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(245)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QueryLanguageParserAnyText:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(239)
			p.Alias()
		}
		{
			p.SetState(240)
			p.Match(QueryLanguageParserT__5)
		}
		{
			p.SetState(241)
			p.NodeName()
		}


	case QueryLanguageParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(243)
			p.Match(QueryLanguageParserT__5)
		}
		{
			p.SetState(244)
			p.NodeName()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) AnyText() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserAnyText, 0)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (s *AliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageVisitor:
		return t.VisitAlias(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *QueryLanguageParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, QueryLanguageParserRULE_alias)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(QueryLanguageParserAnyText)
	}



	return localctx
}


// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) AnyText() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserAnyText, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *QueryLanguageParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, QueryLanguageParserRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(QueryLanguageParserAnyText)
	}



	return localctx
}


// IFieldValueContext is an interface to support dynamic dispatch.
type IFieldValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldValueContext differentiates from other interfaces.
	IsFieldValueContext()
}

type FieldValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldValueContext() *FieldValueContext {
	var p = new(FieldValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_fieldValue
	return p
}

func (*FieldValueContext) IsFieldValueContext() {}

func NewFieldValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldValueContext {
	var p = new(FieldValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_fieldValue

	return p
}

func (s *FieldValueContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldValueContext) AllAnyText() []antlr.TerminalNode {
	return s.GetTokens(QueryLanguageParserAnyText)
}

func (s *FieldValueContext) AnyText(i int) antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserAnyText, i)
}

func (s *FieldValueContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(QueryLanguageParserWS)
}

func (s *FieldValueContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserWS, i)
}

func (s *FieldValueContext) NumberText() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserNumberText, 0)
}

func (s *FieldValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FieldValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterFieldValue(s)
	}
}

func (s *FieldValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitFieldValue(s)
	}
}

func (s *FieldValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageVisitor:
		return t.VisitFieldValue(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *QueryLanguageParser) FieldValue() (localctx IFieldValueContext) {
	localctx = NewFieldValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, QueryLanguageParserRULE_fieldValue)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(271)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QueryLanguageParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(251)
			p.Match(QueryLanguageParserT__11)
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == QueryLanguageParserWS {
			{
				p.SetState(252)
				p.Match(QueryLanguageParserWS)
			}


			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(258)
			p.Match(QueryLanguageParserAnyText)
		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == QueryLanguageParserWS {
			p.SetState(260)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			for ok := true; ok; ok = _la == QueryLanguageParserWS {
				{
					p.SetState(259)
					p.Match(QueryLanguageParserWS)
				}


				p.SetState(262)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(264)
				p.Match(QueryLanguageParserAnyText)
			}


			p.SetState(267)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(269)
			p.Match(QueryLanguageParserT__11)
		}


	case QueryLanguageParserNumberText:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)
			p.Match(QueryLanguageParserNumberText)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// INodeNameContext is an interface to support dynamic dispatch.
type INodeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNodeNameContext differentiates from other interfaces.
	IsNodeNameContext()
}

type NodeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeNameContext() *NodeNameContext {
	var p = new(NodeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QueryLanguageParserRULE_nodeName
	return p
}

func (*NodeNameContext) IsNodeNameContext() {}

func NewNodeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeNameContext {
	var p = new(NodeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QueryLanguageParserRULE_nodeName

	return p
}

func (s *NodeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeNameContext) AnyText() antlr.TerminalNode {
	return s.GetToken(QueryLanguageParserAnyText, 0)
}

func (s *NodeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NodeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.EnterNodeName(s)
	}
}

func (s *NodeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QueryLanguageListener); ok {
		listenerT.ExitNodeName(s)
	}
}

func (s *NodeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case QueryLanguageVisitor:
		return t.VisitNodeName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *QueryLanguageParser) NodeName() (localctx INodeNameContext) {
	localctx = NewNodeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, QueryLanguageParserRULE_nodeName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(QueryLanguageParserAnyText)
	}



	return localctx
}


