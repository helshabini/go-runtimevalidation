package tags


type Tag string

const (
	Unknown             Tag = "unknown"
	Required            Tag = "required"
	Alpha               Tag = "alpha"
	AlphaNumeric        Tag = "alphanum"
	AlphaUnicode        Tag = "alphaunicode"
	AlphaNumericUnicode Tag = "alphanumunicode"
	Numeric             Tag = "num"
	NumericUnsigned     Tag = "unum"
	Hexadecimal         Tag = "hex"
	HexColor            Tag = "hexcolor"
	RGB                 Tag = "rgb"
	RGBA                Tag = "rgba"
	HSL                 Tag = "hsl"
	HSLA                Tag = "hsla"
	Email               Tag = "email"
	ISSN                Tag = "issn"
	E164                Tag = "e164"
	Base32              Tag = "base32"
	Base32Hex           Tag = "base32hex"
	Base64              Tag = "base64"
	Base64Raw           Tag = "base64raw"
	Base64URL           Tag = "base64url"
	Base64RawURL        Tag = "base64rawurl"
	Isbn10              Tag = "isbn10"
	Isbn13              Tag = "isbn13"
	SSN                 Tag = "ssn"
	UUID                Tag = "uuid"
	UUID3               Tag = "uuid3"
	UUID4               Tag = "uuid4"
	UUID5               Tag = "uuid5"
	ULID                Tag = "ulid"
	MD4                 Tag = "md4"
	MD5                 Tag = "md5"
	SHA                 Tag = "sha"
	SHA0                Tag = "sha0"
	SHA1                Tag = "sha1"
	SHA2                Tag = "sha2"
	SHA3                Tag = "sha3"
	SHA224              Tag = "sha224"
	SHA256              Tag = "sha256"
	SHA384              Tag = "sha384"
	SHA512              Tag = "sha512"
	ASCII               Tag = "ascii"
	PrintableASCII      Tag = "asciiprint"
	MultiByte           Tag = "multibyte"
	DataURI             Tag = "datauri"
	Latitude            Tag = "lat"
	Longitude           Tag = "long"
	Hostname            Tag = "hostname"
	Fqdn                Tag = "fqdn"
	UrlEncoded          Tag = "urlencoded"
	HTML                Tag = "html"
	HTMLEncoded         Tag = "htmlencoded"
	JWT                 Tag = "jwt"
	BIC                 Tag = "bic"
	SemVer              Tag = "semver"
	DNS                 Tag = "dns"
	CVE                 Tag = "cve"
	Cron                Tag = "cron"
	Regex               Tag = "regex"
	RequiredIf          Tag = "requiredif"
	Between             Tag = "between"
	XBetween            Tag = "xbetween"
	BetweenF            Tag = "betweenf"
	XBetweenF           Tag = "xbetweenf"
	Min                 Tag = "min"
	Max                 Tag = "max"
	Length              Tag = "length"
	OneOf               Tag = "oneof"
	StartsWith          Tag = "startswith"
	StartsNotWith       Tag = "startsnotwith"
	EndsWith            Tag = "endswith"
	Contains            Tag = "contains"
	ContainsNot         Tag = "containsnot"
	Uppercase           Tag = "upper"
	Lowercase           Tag = "lower"
	EndsNotWith         Tag = "endsnotwith"
)