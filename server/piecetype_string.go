// Code generated by "stringer -type=PieceType"; DO NOT EDIT.

package server

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PieceTypeUnknown - -1]
	_ = x[PieceTypeI-0]
	_ = x[PieceTypeO-1]
	_ = x[PieceTypeT-2]
	_ = x[PieceTypeL-3]
	_ = x[PieceTypeP-4]
	_ = x[PieceTypeS-5]
	_ = x[PieceTypeZ-6]
}

const _PieceType_name = "PieceTypeUnknownPieceTypeIPieceTypeOPieceTypeTPieceTypeLPieceTypePPieceTypeSPieceTypeZ"

var _PieceType_index = [...]uint8{0, 16, 26, 36, 46, 56, 66, 76, 86}

func (i PieceType) String() string {
	i -= -1
	if i < 0 || i >= PieceType(len(_PieceType_index)-1) {
		return "PieceType(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _PieceType_name[_PieceType_index[i]:_PieceType_index[i+1]]
}