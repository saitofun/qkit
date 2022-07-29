// This is a generated source file. DO NOT EDIT
// Version: 0.0.1
// Source: strfmt/strfmt_generated.go
// Date: Jul 28 23:27:02

package strfmt

import (
	"github.com/saitofun/qkit/kit/validator"
)

func init() {
	validator.DefaultFactory.Register(Base64UrlValidator)
}

var Base64UrlValidator = validator.NewRegexpStrfmtValidator(regexpStringBase64URL, "base-64-url", "base64Url")

func init() {
	validator.DefaultFactory.Register(Uuid3Rfc4122Validator)
}

var Uuid3Rfc4122Validator = validator.NewRegexpStrfmtValidator(regexpStringUUID3RFC4122, "uuid-3-rfc-4122", "uuid3Rfc4122")

func init() {
	validator.DefaultFactory.Register(Uuid4Rfc4122Validator)
}

var Uuid4Rfc4122Validator = validator.NewRegexpStrfmtValidator(regexpStringUUID4RFC4122, "uuid-4-rfc-4122", "uuid4Rfc4122")

func init() {
	validator.DefaultFactory.Register(EthAddressValidator)
}

var EthAddressValidator = validator.NewRegexpStrfmtValidator(regexpStringEthAddress, "eth-address", "ethAddress")

func init() {
	validator.DefaultFactory.Register(RgbaValidator)
}

var RgbaValidator = validator.NewRegexpStrfmtValidator(regexpStringRGBA, "rgba")

func init() {
	validator.DefaultFactory.Register(RgbValidator)
}

var RgbValidator = validator.NewRegexpStrfmtValidator(regexpStringRGB, "rgb")

func init() {
	validator.DefaultFactory.Register(HslaValidator)
}

var HslaValidator = validator.NewRegexpStrfmtValidator(regexpStringHSLA, "hsla")

func init() {
	validator.DefaultFactory.Register(E164Validator)
}

var E164Validator = validator.NewRegexpStrfmtValidator(regexpStringE164, "e-164", "e164")

func init() {
	validator.DefaultFactory.Register(Isbn10Validator)
}

var Isbn10Validator = validator.NewRegexpStrfmtValidator(regexpStringISBN10, "isbn-10", "isbn10")

func init() {
	validator.DefaultFactory.Register(HostnameRfc1123Validator)
}

var HostnameRfc1123Validator = validator.NewRegexpStrfmtValidator(regexpStringHostnameRFC1123, "hostname-rfc-1123", "hostnameRfc1123")

func init() {
	validator.DefaultFactory.Register(BtcAddressValidator)
}

var BtcAddressValidator = validator.NewRegexpStrfmtValidator(regexpStringBtcAddress, "btc-address", "btcAddress")

func init() {
	validator.DefaultFactory.Register(AlphaUnicodeNumericValidator)
}

var AlphaUnicodeNumericValidator = validator.NewRegexpStrfmtValidator(regexpStringAlphaUnicodeNumeric, "alpha-unicode-numeric", "alphaUnicodeNumeric")

func init() {
	validator.DefaultFactory.Register(Uuid5Validator)
}

var Uuid5Validator = validator.NewRegexpStrfmtValidator(regexpStringUUID5, "uuid-5", "uuid5")

func init() {
	validator.DefaultFactory.Register(LongitudeValidator)
}

var LongitudeValidator = validator.NewRegexpStrfmtValidator(regexpStringLongitude, "longitude")

func init() {
	validator.DefaultFactory.Register(BtcAddressUpperBech32Validator)
}

var BtcAddressUpperBech32Validator = validator.NewRegexpStrfmtValidator(regexpStringBtcAddressUpperBech32, "btc-address-upper-bech-32", "btcAddressUpperBech32")

func init() {
	validator.DefaultFactory.Register(EthAddressLowerValidator)
}

var EthAddressLowerValidator = validator.NewRegexpStrfmtValidator(regexpStringEthAddressLower, "eth-address-lower", "ethAddressLower")

func init() {
	validator.DefaultFactory.Register(HexadecimalValidator)
}

var HexadecimalValidator = validator.NewRegexpStrfmtValidator(regexpStringHexadecimal, "hexadecimal")

func init() {
	validator.DefaultFactory.Register(HslValidator)
}

var HslValidator = validator.NewRegexpStrfmtValidator(regexpStringHSL, "hsl")

func init() {
	validator.DefaultFactory.Register(Uuid4Validator)
}

var Uuid4Validator = validator.NewRegexpStrfmtValidator(regexpStringUUID4, "uuid-4", "uuid4")

func init() {
	validator.DefaultFactory.Register(HostnameRfc952Validator)
}

var HostnameRfc952Validator = validator.NewRegexpStrfmtValidator(regexpStringHostnameRFC952, "hostname-rfc-952", "hostnameRfc952")

func init() {
	validator.DefaultFactory.Register(UrlEncodedValidator)
}

var UrlEncodedValidator = validator.NewRegexpStrfmtValidator(regexpStringURLEncoded, "url-encoded", "urlEncoded")

func init() {
	validator.DefaultFactory.Register(NumberValidator)
}

var NumberValidator = validator.NewRegexpStrfmtValidator(regexpStringNumber, "number")

func init() {
	validator.DefaultFactory.Register(AlphaUnicodeValidator)
}

var AlphaUnicodeValidator = validator.NewRegexpStrfmtValidator(regexpStringAlphaUnicode, "alpha-unicode", "alphaUnicode")

func init() {
	validator.DefaultFactory.Register(Base64Validator)
}

var Base64Validator = validator.NewRegexpStrfmtValidator(regexpStringBase64, "base-64", "base64")

func init() {
	validator.DefaultFactory.Register(UuidValidator)
}

var UuidValidator = validator.NewRegexpStrfmtValidator(regexpStringUUID, "uuid")

func init() {
	validator.DefaultFactory.Register(Uuid5Rfc4122Validator)
}

var Uuid5Rfc4122Validator = validator.NewRegexpStrfmtValidator(regexpStringUUID5RFC4122, "uuid-5-rfc-4122", "uuid5Rfc4122")

func init() {
	validator.DefaultFactory.Register(Uuidrfc4122Validator)
}

var Uuidrfc4122Validator = validator.NewRegexpStrfmtValidator(regexpStringUUIDRFC4122, "uuidrfc-4122", "uuidrfc4122")

func init() {
	validator.DefaultFactory.Register(LatitudeValidator)
}

var LatitudeValidator = validator.NewRegexpStrfmtValidator(regexpStringLatitude, "latitude")

func init() {
	validator.DefaultFactory.Register(AlphaValidator)
}

var AlphaValidator = validator.NewRegexpStrfmtValidator(regexpStringAlpha, "alpha")

func init() {
	validator.DefaultFactory.Register(HexColorValidator)
}

var HexColorValidator = validator.NewRegexpStrfmtValidator(regexpStringHexColor, "hex-color", "hexColor")

func init() {
	validator.DefaultFactory.Register(EmailValidator)
}

var EmailValidator = validator.NewRegexpStrfmtValidator(regexpStringEmail, "email")

func init() {
	validator.DefaultFactory.Register(DataUriValidator)
}

var DataUriValidator = validator.NewRegexpStrfmtValidator(regexpStringDataURI, "data-uri", "dataUri")

func init() {
	validator.DefaultFactory.Register(SsnValidator)
}

var SsnValidator = validator.NewRegexpStrfmtValidator(regexpStringSSN, "ssn")

func init() {
	validator.DefaultFactory.Register(HtmlEncodedValidator)
}

var HtmlEncodedValidator = validator.NewRegexpStrfmtValidator(regexpStringHTMLEncoded, "html-encoded", "htmlEncoded")

func init() {
	validator.DefaultFactory.Register(HtmlValidator)
}

var HtmlValidator = validator.NewRegexpStrfmtValidator(regexpStringHTML, "html")

func init() {
	validator.DefaultFactory.Register(NumericValidator)
}

var NumericValidator = validator.NewRegexpStrfmtValidator(regexpStringNumeric, "numeric")

func init() {
	validator.DefaultFactory.Register(PrintableAsciiValidator)
}

var PrintableAsciiValidator = validator.NewRegexpStrfmtValidator(regexpStringPrintableASCII, "printable-ascii", "printableAscii")

func init() {
	validator.DefaultFactory.Register(AsciiValidator)
}

var AsciiValidator = validator.NewRegexpStrfmtValidator(regexpStringASCII, "ascii")

func init() {
	validator.DefaultFactory.Register(Isbn13Validator)
}

var Isbn13Validator = validator.NewRegexpStrfmtValidator(regexpStringISBN13, "isbn-13", "isbn13")

func init() {
	validator.DefaultFactory.Register(Uuid3Validator)
}

var Uuid3Validator = validator.NewRegexpStrfmtValidator(regexpStringUUID3, "uuid-3", "uuid3")

func init() {
	validator.DefaultFactory.Register(MultibyteValidator)
}

var MultibyteValidator = validator.NewRegexpStrfmtValidator(regexpStringMultibyte, "multibyte")

func init() {
	validator.DefaultFactory.Register(BtcAddressLowerBech32Validator)
}

var BtcAddressLowerBech32Validator = validator.NewRegexpStrfmtValidator(regexpStringBtcAddressLowerBech32, "btc-address-lower-bech-32", "btcAddressLowerBech32")

func init() {
	validator.DefaultFactory.Register(EthAddressUpperValidator)
}

var EthAddressUpperValidator = validator.NewRegexpStrfmtValidator(regexpStringEthAddressUpper, "eth-address-upper", "ethAddressUpper")

func init() {
	validator.DefaultFactory.Register(AlphaNumericValidator)
}

var AlphaNumericValidator = validator.NewRegexpStrfmtValidator(regexpStringAlphaNumeric, "alpha-numeric", "alphaNumeric")
