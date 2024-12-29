package regex

import (
	"regexp"
	"sync"
)

const (
	alphaRegexString               = "^[a-zA-Z]+$"
	alphaNumericRegexString        = "^[a-zA-Z0-9]+$"
	alphaUnicodeRegexString        = `^[\p{L}]+$`
	alphaNumericUnicodeRegexString = `^[\p{L}\p{N}]+$`
	numericRegexString             = `^[-+]?[0-9]+(?:\.[0-9]+)?$`
	numericUnsignedRegexString     = `^[0-9]+(?:\.[0-9]+)?$`
	hexadecimalRegexString         = "^(0[xX])?[0-9a-fA-F]+$"
	hexColorRegexString            = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$"
	rgbRegexString                 = `^rgb\(\s*(?:(?:0|[1-9]\d?|1\d\d?|2[0-4]\d|25[0-5])\s*,\s*(?:0|[1-9]\d?|1\d\d?|2[0-4]\d|25[0-5])\s*,\s*(?:0|[1-9]\d?|1\d\d?|2[0-4]\d|25[0-5])\s*)\)$`
	rgbaRegexString                = `^rgba\(\s*(?:(?:0|[1-9]\d?|1\d\d?|2[0-4]\d|25[0-5])\s*,\s*(?:0|[1-9]\d?|1\d\d?|2[0-4]\d|25[0-5])\s*,\s*(?:0|[1-9]\d?|1\d\d?|2[0-4]\d|25[0-5])\s*,\s*(?:0|1(?:\.0)?|0?\.\d+)\s*)\)$`
	hslRegexString                 = `^hsl\(\s*(?:(360|[1-9]?[0-9]|1[0-9][0-9]|2[0-9][0-9])\s*,\s*(100|[1-9]?\d)%\s*,\s*(100|[1-9]?\d)%\s*)\)$`
	hslaRegexString                = `^hsla\(\s*(?:(360|[1-9]?[0-9]|1[0-9][0-9]|2[0-9][0-9])\s*,\s*(100|[1-9]?\d)%\s*,\s*(100|[1-9]?\d)%\s*,\s*(?:0|1(?:\.0)?|0?\.\d+)\s*)\)$`
	emailRegexString               = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	e164RegexString                = `^\+[1-9]\d{6,14}$`
	iSSNRegexString                = "^(?:[0-9]{4}-[0-9]{3}[0-9X])$"
	uUID3RFC4122RegexString        = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	uUID4RFC4122RegexString        = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uUID5RFC4122RegexString        = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uUIDRFC4122RegexString         = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	uLIDRegexString                = "^(?i)[A-HJKMNP-TV-Z0-9]{26}$"
	md4RegexString                 = "^[0-9a-fA-F]{32}$"
	md5RegexString                 = "^[0-9a-fA-F]{32}$"
	shaRegexString                 = "^[0-9a-fA-F]{40,128}$"
	sha160RegexString              = "^[0-9a-fA-F]{40}$"
	sha3RegexString                = "^[0-9a-fA-F]{56,128}$"
	sha224RegexString              = "^[0-9a-fA-F]{56}$"
	sha256RegexString              = "^[0-9a-fA-F]{64}$"
	sha384RegexString              = "^[0-9a-fA-F]{96}$"
	sha512RegexString              = "^[0-9a-fA-F]{128}$"
	aSCIIRegexString               = "^[\x00-\x7F]*$"
	printableASCIIRegexString      = "^[\x20-\x7E]*$"
	multibyteRegexString           = "^[^\x00-\x1F\x21-\x7F]*$"
	dataURIRegexString             = `^data:((?:\w+\/(?:([^;]|;[^;]).)+)?)`
	latitudeRegexString            = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	longitudeRegexString           = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	sSNRegexString                 = `^[0-9]{3}[ -]?(0[1-9]|[1-9][0-9])[ -]?([1-9][0-9]{3}|[0-9][1-9][0-9]{2}|[0-9]{2}[1-9][0-9]|[0-9]{3}[1-9])$`
	hostnameRegexString            = `^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,61}[a-zA-Z0-9]{1}){1}(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,61}[a-zA-Z0-9]{1})*?$`
	fqdnRegexString                = `^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,61}[a-zA-Z0-9]{1}){1}(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,61}[a-zA-Z0-9]{1})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,61}[a-zA-Z0-9]{1})\.?$`
	uRLEncodedRegexString          = `^(?:[a-zA-Z0-9\-_.~]|%[0-9A-Fa-f]{2})*$`
	hTMLEncodedRegexString         = `&#[xX]?[0-9a-fA-F]{1,5};|&[a-zA-Z0-9]+;`
	hTMLRegexString                = `(<[/]?([a-zA-Z]+).*?>|&[a-zA-Z]+;)`
	jWTRegexString                 = `^[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]*$`
	bicRegexString                 = `^[A-Z]{4}[A-Z]{2}[A-Z0-9]{2}([A-Z0-9]{3})?$`
	semverRegexString              = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$` // numbered capture groups https://semver.org/
	dnsRegexString                 = "^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$"
	cveRegexString                 = `^CVE-(1999|2\d{3})-(0[^0]\d{2}|0\d[^0]\d{1}|0\d{2}[^0]|[1-9]{1}\d{3,})$` // CVE Format Id https://cve.mitre.org/cve/identifiers/syntaxchange.html
	extractDigitsRegexString       = "[^0-9]"
)

var (
	AlphaRegex               = CompileOnce(alphaRegexString)
	AlphaNumericRegex        = CompileOnce(alphaNumericRegexString)
	AlphaUnicodeRegex        = CompileOnce(alphaUnicodeRegexString)
	AlphaNumericUnicodeRegex = CompileOnce(alphaNumericUnicodeRegexString)
	NumericRegex             = CompileOnce(numericRegexString)
	NumericUnsignedRegex     = CompileOnce(numericUnsignedRegexString)
	HexadecimalRegex         = CompileOnce(hexadecimalRegexString)
	HexColorRegex            = CompileOnce(hexColorRegexString)
	RgbRegex                 = CompileOnce(rgbRegexString)
	RgbaRegex                = CompileOnce(rgbaRegexString)
	HslRegex                 = CompileOnce(hslRegexString)
	HslaRegex                = CompileOnce(hslaRegexString)
	E164Regex                = CompileOnce(e164RegexString)
	EmailRegex               = CompileOnce(emailRegexString)
	ISSNRegex                = CompileOnce(iSSNRegexString)
	UUID3RFC4122Regex        = CompileOnce(uUID3RFC4122RegexString)
	UUID4RFC4122Regex        = CompileOnce(uUID4RFC4122RegexString)
	UUID5RFC4122Regex        = CompileOnce(uUID5RFC4122RegexString)
	UUIDRFC4122Regex         = CompileOnce(uUIDRFC4122RegexString)
	ULIDRegex                = CompileOnce(uLIDRegexString)
	Md4Regex                 = CompileOnce(md4RegexString)
	Md5Regex                 = CompileOnce(md5RegexString)
	ShaRegex                 = CompileOnce(shaRegexString)
	Sha160Regex              = CompileOnce(sha160RegexString)
	Sha3Regex                = CompileOnce(sha3RegexString)
	Sha224Regex              = CompileOnce(sha224RegexString)
	Sha256Regex              = CompileOnce(sha256RegexString)
	Sha384Regex              = CompileOnce(sha384RegexString)
	Sha512Regex              = CompileOnce(sha512RegexString)
	ASCIIRegex               = CompileOnce(aSCIIRegexString)
	PrintableASCIIRegex      = CompileOnce(printableASCIIRegexString)
	MultibyteRegex           = CompileOnce(multibyteRegexString)
	DataURIRegex             = CompileOnce(dataURIRegexString)
	LatitudeRegex            = CompileOnce(latitudeRegexString)
	LongitudeRegex           = CompileOnce(longitudeRegexString)
	SSNRegex                 = CompileOnce(sSNRegexString)
	HostnameRegex            = CompileOnce(hostnameRegexString)
	FqdnRegex                = CompileOnce(fqdnRegexString)
	URLEncodedRegex          = CompileOnce(uRLEncodedRegexString)
	HTMLEncodedRegex         = CompileOnce(hTMLEncodedRegexString)
	HTMLRegex                = CompileOnce(hTMLRegexString)
	JWTRegex                 = CompileOnce(jWTRegexString)
	BicRegex                 = CompileOnce(bicRegexString)
	SemverRegex              = CompileOnce(semverRegexString)
	DnsRegex                 = CompileOnce(dnsRegexString)
	CveRegex                 = CompileOnce(cveRegexString)
	ExtractDigitsRegex       = CompileOnce(extractDigitsRegexString)
)

func CompileOnce(str string) func() *regexp.Regexp {
	var regex *regexp.Regexp
	var once sync.Once
	return func() *regexp.Regexp {
		once.Do(func() {
			regex = regexp.MustCompile(str)
		})
		return regex
	}
}
