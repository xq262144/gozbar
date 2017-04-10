// ZBar constants
package zbar

// #cgo LDFLAGS: -lzbar
// #cgo linux LDFLAGS: -lrt
// #include <zbar.h>
import "C"

// Symbol types (Scanner.SetConfig() argument symbology)
const (
	NONE    = C.ZBAR_NONE
	PARTIAL = C.ZBAR_PARTIAL
	EAN8    = C.ZBAR_EAN8
	UPCE    = C.ZBAR_UPCE
	ISBN10  = C.ZBAR_ISBN10
	UPCA    = C.ZBAR_UPCA
	EAN13   = C.ZBAR_EAN13
	ISBN13  = C.ZBAR_ISBN13
	I25     = C.ZBAR_I25
	CODE39  = C.ZBAR_CODE39
	PDF417  = C.ZBAR_PDF417
	QRCODE  = C.ZBAR_QRCODE
	CODE128 = C.ZBAR_CODE128
	SYMBOL  = C.ZBAR_SYMBOL
	ADDON   = C.ZBAR_ADDON
	ADDON2  = C.ZBAR_ADDON2
	ADDON5  = C.ZBAR_ADDON5
)

// Configuration keys
const (
	CFG_ENABLE     = C.ZBAR_CFG_ENABLE
	CFG_ADD_CHECK  = C.ZBAR_CFG_ADD_CHECK
	CFG_EMIT_CHECK = C.ZBAR_CFG_EMIT_CHECK
	CFG_ASCII      = C.ZBAR_CFG_ASCII
	CFG_NUM        = C.ZBAR_CFG_NUM
	CFG_MIN_LEN    = C.ZBAR_CFG_MIN_LEN
	CFG_MAX_LEN    = C.ZBAR_CFG_MAX_LEN
	CFG_POSITION   = C.ZBAR_CFG_POSITION
	CFG_X_DENSITY  = C.ZBAR_CFG_X_DENSITY
	CFG_Y_DENSITY  = C.ZBAR_CFG_Y_DENSITY
)
