package validation

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"go-runtimevalidation/args"
	"go-runtimevalidation/rules"
	"go-runtimevalidation/tags"
)

type ValidationRule struct {
	Tag             string                            // rule tag
	Text            string                            // rule text (args and expressions)
	Error           *ParsingError                     // error message, gets filled if parsing fails
	ValidationGroup int                               // AND group number for the rule, used for ORing multiple rules within a group
	Validate        func(field any, object any) error // validation function, executes upon validation
}

type ValidationRules map[int][]ValidationRule

func BadValidationRule(tag, text string, group int, err error) *ValidationRule {
	return &ValidationRule{
		Tag:             tag,
		Text:            text,
		ValidationGroup: group,
		Error:           NewParsingError(text, err),
		Validate:        func(field any, object any) error { return err },
	}
}

func NewValidationRule(tag, text string, group int, validationFunc func(field any, object any) error) *ValidationRule {
	return &ValidationRule{
		Tag:             tag,
		Text:            text,
		ValidationGroup: group,
		Validate:        validationFunc,
		Error:           nil,
	}
}

func (rules *ValidationRules) Error() error {
	var errbuff bytes.Buffer
	for _, group := range *rules {
		for _, rule := range group {
			if rule.Error != nil {
				errbuff.WriteString(rule.Error.String())
				errbuff.WriteString("\n")
			}
		}
	}
	if errbuff.Len() == 0 {
		return nil
	}

	return errors.New(errbuff.String())
}

// Validate runs the validation rules on the input.
// For each validation group, if any of the rules succeed, the group succeeds.
// If all groups succeed, the validation succeeds.
// Errors are returned as a list of failed rules.
func (rules ValidationRules) Validate(input any, parent any) ValidationErrors {
	if len(rules) == 0 {
		return nil
	}

	errs := make(ValidationErrors, 0)

	// Process each validation group sequentially
	for _, group := range rules {
		groupPassed := false
		groupErrs := make(ValidationErrors, 0)

		// Check if any rule in the group succeeds
		for _, rule := range group {

			// If there is a single parsing error anywhere, we can't validate.
			// Instead we report the list of all parsing errors for all rules as a single validation error.
			if rule.Error != nil {
				parsingErr := rules.Error()
				if parsingErr != nil {
					return ValidationErrors{
						*NewValidationError("one or more validation rules cannot be parsed", parsingErr),
					}
				}
			}

			// Run the validation function
			if err := rule.Validate(input, parent); err != nil {
				// Collect the group error if validation fails
				groupErrs = append(groupErrs, *NewValidationError(rule.Text, err))
			} else {
				groupPassed = true // If any rule in the group passes, mark the group as passed
				break              // Stop checking other rules in the group
			}
		}

		// If the group has no successful rule, append the group's errors
		// Else, we will discard the group's errors because the group passed
		if !groupPassed {
			errs = append(errs, groupErrs...)
		}
	}

	// If no errors, return nil (validation passed)
	if len(errs) == 0 {
		return nil
	}

	return errs // Return collected validation errors
}

// Parse parses a string of validation rules into a map of grouped rules.
func Parse(rulestext string) (ValidationRules, error) {
	groupedRules := make(ValidationRules)

	groups := strings.Split(rulestext, "&&") // Split the rules by AND operator
	for i, group := range groups {           // Parse each group of rules

		parsed := make([]ValidationRule, 0)

		rules := strings.Split(group, "||")
		for _, rule := range rules {
			parsedrule := parseRule(rule, i)
			parsed = append(parsed, *parsedrule)
		}

		// Append the rules to the grouped rules
		for _, rule := range parsed {
			groupedRules[rule.ValidationGroup] = append(groupedRules[rule.ValidationGroup], rule)
		}
	}

	// If no rules found, return an error
	// Note we also return the grouped rules, which may contain parsing errors
	// This helps the caller to know which rules failed to parse
	if len(groupedRules) == 0 {
		return groupedRules, fmt.Errorf("no rules found")
	}

	// If there are parsing errors, return an consolidated error
	// Note that the grouped rules will contain the parsing errors
	// This helps the caller to know which rules failed to parse
	if err := groupedRules.Error(); err != nil {
		return groupedRules, err
	}

	// Return the grouped rules
	return groupedRules, nil
}

func parseRule(text string, group int) *ValidationRule {
	if text == "" {
		return BadValidationRule(string(tags.Unknown), text, group, fmt.Errorf("empty rule"))
	}

	parts := strings.Split(text, ":")
	rulename := strings.ToLower(strings.TrimSpace(parts[0]))
	// args string is the rest of the rule text after the rule name and first colon
	// e.g., for "oneof:choice1,choice2,choice3", argsStr is "choice1,choice2,choice3"
	argsStr := strings.TrimSpace(strings.Join(parts[1:], ":"))

	// Later, it might be better to lazy-load rules into a map for faster lookup
	if len(argsStr) == 0 { // Rule has no arguments, we do this for faster lookup of the correct rule, eliminating the need to switch on rules which we know have arguments
		switch tags.Tag(rulename) {
		case tags.Required:
			return NewValidationRule(text, string(tags.Required), group, func(field any, object any) error {
				return rules.Required(field)
			})
		case tags.Alpha:
			return NewValidationRule(text, string(tags.Alpha), group, func(field any, object any) error {
				return rules.Alpha(field)
			})
		case tags.AlphaNumeric:
			return NewValidationRule(text, string(tags.AlphaNumeric), group, func(field any, object any) error {
				return rules.AlphaNumeric(field)
			})
		case tags.AlphaUnicode:
			return NewValidationRule(text, string(tags.AlphaUnicode), group, func(field any, object any) error {
				return rules.AlphaUnicode(field)
			})
		case tags.AlphaNumericUnicode:
			return NewValidationRule(text, string(tags.AlphaNumericUnicode), group, func(field any, object any) error {
				return rules.AlphaNumericUnicode(field)
			})
		case tags.Numeric:
			return NewValidationRule(text, string(tags.Numeric), group, func(field any, object any) error {
				return rules.Numeric(field)
			})
		case tags.NumericUnsigned:
			return NewValidationRule(text, string(tags.NumericUnsigned), group, func(field any, object any) error {
				return rules.NumericUnsigned(field)
			})
		case tags.Hexadecimal:
			return NewValidationRule(text, string(tags.Hexadecimal), group, func(field any, object any) error {
				return rules.Hexadecimal(field)
			})
		case tags.HexColor:
			return NewValidationRule(text, string(tags.HexColor), group, func(field any, object any) error {
				return rules.HexColor(field)
			})
		case tags.RGB:
			return NewValidationRule(text, string(tags.RGB), group, func(field any, object any) error {
				return rules.RGB(field)
			})
		case tags.RGBA:
			return NewValidationRule(text, string(tags.RGBA), group, func(field any, object any) error {
				return rules.RGBA(field)
			})
		case tags.HSL:
			return NewValidationRule(text, string(tags.HSL), group, func(field any, object any) error {
				return rules.HSL(field)
			})
		case tags.HSLA:
			return NewValidationRule(text, string(tags.HSLA), group, func(field any, object any) error {
				return rules.HSLA(field)
			})
		case tags.Email:
			return NewValidationRule(string(tags.Email), text, group, func(field any, object any) error {
				return rules.Email(field)
			})
		case tags.ISSN:
			return NewValidationRule(string(tags.ISSN), text, group, func(field any, object any) error {
				return rules.ISSN(field)
			})
		case tags.E164:
			return NewValidationRule(string(tags.E164), text, group, func(field any, object any) error {
				return rules.E164(field)
			})
		case tags.Base32:
			return NewValidationRule(string(tags.Base32), text, group, func(field any, object any) error {
				return rules.Base32(field)
			})
		case tags.Base32Hex:
			return NewValidationRule(string(tags.Base32Hex), text, group, func(field any, object any) error {
				return rules.Base32Hex(field)
			})
		case tags.Base64:
			return NewValidationRule(string(tags.Base64), text, group, func(field any, object any) error {
				return rules.Base64(field)
			})
		case tags.Base64Raw:
			return NewValidationRule(string(tags.Base64Raw), text, group, func(field any, object any) error {
				return rules.Base64Raw(field)
			})
		case tags.Base64URL:
			return NewValidationRule(string(tags.Base64URL), text, group, func(field any, object any) error {
				return rules.Base64Url(field)
			})
		case tags.Base64RawURL:
			return NewValidationRule(string(tags.Base64RawURL), text, group, func(field any, object any) error {
				return rules.Base64RawUrl(field)
			})
		case tags.Isbn10:
			return NewValidationRule(string(tags.Isbn10), text, group, func(field any, object any) error {
				return rules.Isbn10(field)
			})
		case tags.Isbn13:
			return NewValidationRule(string(tags.Isbn13), text, group, func(field any, object any) error {
				return rules.Isbn13(field)
			})
		case tags.SSN:
			return NewValidationRule(string(tags.SSN), text, group, func(field any, object any) error {
				return rules.SSN(field)
			})
		case tags.UUID:
			return NewValidationRule(string(tags.UUID), text, group, func(field any, object any) error {
				return rules.UUID(field)
			})
		case tags.UUID3:
			return NewValidationRule(string(tags.UUID3), text, group, func(field any, object any) error {
				return rules.UUID3(field)
			})
		case tags.UUID4:
			return NewValidationRule(string(tags.UUID4), text, group, func(field any, object any) error {
				return rules.UUID4(field)
			})
		case tags.UUID5:
			return NewValidationRule(string(tags.UUID5), text, group, func(field any, object any) error {
				return rules.UUID5(field)
			})
		case tags.ULID:
			return NewValidationRule(string(tags.ULID), text, group, func(field any, object any) error {
				return rules.ULID(field)
			})
		case tags.MD4:
			return NewValidationRule(string(tags.MD4), text, group, func(field any, object any) error {
				return rules.MD4(field)
			})
		case tags.MD5:
			return NewValidationRule(string(tags.MD5), text, group, func(field any, object any) error {
				return rules.MD5(field)
			})
		case tags.SHA:
			return NewValidationRule(string(tags.SHA), text, group, func(field any, object any) error {
				return rules.SHA(field)
			})
		case tags.SHA0:
			return NewValidationRule(string(tags.SHA0), text, group, func(field any, object any) error {
				return rules.SHA160(field)
			})
		case tags.SHA1:
			return NewValidationRule(string(tags.SHA1), text, group, func(field any, object any) error {
				return rules.SHA160(field)
			})
		case tags.SHA2:
			return NewValidationRule(string(tags.SHA2), text, group, func(field any, object any) error {
				return rules.SHA3(field)
			})
		case tags.SHA3:
			return NewValidationRule(string(tags.SHA3), text, group, func(field any, object any) error {
				return rules.SHA3(field)
			})
		case tags.SHA224:
			return NewValidationRule(string(tags.SHA224), text, group, func(field any, object any) error {
				return rules.SHA224(field)
			})
		case tags.SHA256:
			return NewValidationRule(string(tags.SHA256), text, group, func(field any, object any) error {
				return rules.SHA256(field)
			})
		case tags.SHA384:
			return NewValidationRule(string(tags.SHA384), text, group, func(field any, object any) error {
				return rules.SHA384(field)
			})
		case tags.SHA512:
			return NewValidationRule(string(tags.SHA512), text, group, func(field any, object any) error {
				return rules.SHA512(field)
			})
		case tags.ASCII:
			return NewValidationRule(string(tags.ASCII), text, group, func(field any, object any) error {
				return rules.Ascii(field)
			})
		case tags.PrintableASCII:
			return NewValidationRule(string(tags.PrintableASCII), text, group, func(field any, object any) error {
				return rules.AsciiPrint(field)
			})
		case tags.MultiByte:
			return NewValidationRule(string(tags.MultiByte), text, group, func(field any, object any) error {
				return rules.MultiByte(field)
			})
		case tags.Uppercase:
			return NewValidationRule(string(tags.Uppercase), text, group, func(field any, object any) error {
				return rules.Uppercase(field)
			})
		case tags.Lowercase:
			return NewValidationRule(string(tags.Lowercase), text, group, func(field any, object any) error {
				return rules.Lowercase(field)
			})
		case tags.DataURI:
			return NewValidationRule(string(tags.DataURI), text, group, func(field any, object any) error {
				return rules.DataUri(field)
			})
		case tags.Latitude:
			return NewValidationRule(string(tags.Latitude), text, group, func(field any, object any) error {
				return rules.Latitude(field)
			})
		case tags.Longitude:
			return NewValidationRule(string(tags.Longitude), text, group, func(field any, object any) error {
				return rules.Longitude(field)
			})
		case tags.Hostname:
			return NewValidationRule(string(tags.Hostname), text, group, func(field any, object any) error {
				return rules.Hostname(field)
			})
		case tags.Fqdn:
			return NewValidationRule(string(tags.Fqdn), text, group, func(field any, object any) error {
				return rules.FQDN(field)
			})
		case tags.UrlEncoded:
			return NewValidationRule(string(tags.UrlEncoded), text, group, func(field any, object any) error {
				return rules.UrlEncoded(field)
			})
		case tags.HTML:
			return NewValidationRule(string(tags.HTML), text, group, func(field any, object any) error {
				return rules.HTML(field)
			})
		case tags.HTMLEncoded:
			return NewValidationRule(string(tags.HTMLEncoded), text, group, func(field any, object any) error {
				return rules.HTMLEncoded(field)
			})
		case tags.JWT:
			return NewValidationRule(string(tags.JWT), text, group, func(field any, object any) error {
				return rules.JWT(field)
			})
		case tags.BIC:
			return NewValidationRule(string(tags.BIC), text, group, func(field any, object any) error {
				return rules.BIC(field)
			})
		case tags.SemVer:
			return NewValidationRule(string(tags.SemVer), text, group, func(field any, object any) error {
				return rules.SemVer(field)
			})
		case tags.DNS:
			return NewValidationRule(string(tags.DNS), text, group, func(field any, object any) error {
				return rules.DNS(field)
			})
		case tags.CVE:
			return NewValidationRule(string(tags.CVE), text, group, func(field any, object any) error {
				return rules.CVE(field)
			})
		case tags.Cron:
			return NewValidationRule(string(tags.Cron), text, group, func(field any, object any) error {
				return rules.Cron(field)
			})
		case tags.Regex, tags.RequiredIf, tags.Between, tags.XBetween, tags.BetweenF, tags.XBetweenF, tags.OneOf, tags.Min, tags.Max, tags.Length, tags.StartsWith, tags.StartsNotWith, tags.EndsWith, tags.EndsNotWith, tags.Contains, tags.ContainsNot:
			return BadValidationRule(rulename, text, group, fmt.Errorf("missing arguments for rule: %s", text))
		default:
			return BadValidationRule(string(tags.Unknown), text, group, fmt.Errorf("unknown rule: %s", text))
		}
	} else { // Rule has arguments, we do this for faster lookup of the correct rule, eliminating the need to switch on rules which we know have no arguments
		ruleargs, err := args.ParseArgs(argsStr)
		switch tags.Tag(rulename) {
		case tags.Regex:
			if err != nil {
				return BadValidationRule(string(tags.Regex), text, group, err)
			}
			return NewValidationRule(string(tags.Regex), text, group, func(field any, object any) error {
				return rules.Regex(field, object, ruleargs)
			})
		case tags.RequiredIf:
			if err != nil {
				return BadValidationRule(string(tags.RequiredIf), text, group, err)
			}
			return NewValidationRule(string(tags.RequiredIf), text, group, func(field any, object any) error {
				return rules.RequiredIf(field, object, ruleargs)
			})
		case tags.Between:
			if err != nil {
				return BadValidationRule(string(tags.Between), text, group, err)
			}
			return NewValidationRule(string(tags.Between), text, group, func(field any, object any) error {
				return rules.Between(field, object, ruleargs)
			})
		case tags.XBetween:
			if err != nil {
				return BadValidationRule(string(tags.XBetween), text, group, err)
			}
			return NewValidationRule(string(tags.XBetween), text, group, func(field any, object any) error {
				return rules.XBetween(field, object, ruleargs)
			})
		case tags.BetweenF:
			if err != nil {
				return BadValidationRule(string(tags.BetweenF), text, group, err)
			}
			return NewValidationRule(string(tags.BetweenF), text, group, func(field any, object any) error {
				return rules.BetweenF(field, object, ruleargs)
			})
		case tags.XBetweenF:
			if err != nil {
				return BadValidationRule(string(tags.XBetweenF), text, group, err)
			}
			return NewValidationRule(string(tags.XBetweenF), text, group, func(field any, object any) error {
				return rules.XBetweenF(field, object, ruleargs)
			})
		case tags.OneOf:
			if err != nil {
				return BadValidationRule(string(tags.OneOf), text, group, err)
			}
			return NewValidationRule(string(tags.OneOf), text, group, func(field any, object any) error {
				return rules.OneOf(field, object, ruleargs)
			})
		case tags.Min:
			if err != nil {
				return BadValidationRule(string(tags.Min), text, group, err)
			}
			return NewValidationRule(string(tags.Min), text, group, func(field any, object any) error {
				return rules.Min(field, object, ruleargs)
			})
		case tags.Max:
			if err != nil {
				return BadValidationRule(string(tags.Max), text, group, err)
			}
			return NewValidationRule(string(tags.Max), text, group, func(field any, object any) error {
				return rules.Max(field, object, ruleargs)
			})
		case tags.StartsWith:
			if err != nil {
				return BadValidationRule(string(tags.StartsWith), text, group, err)
			}
			return NewValidationRule(string(tags.StartsWith), text, group, func(field any, object any) error {
				return rules.StartsWith(field, object, ruleargs)
			})
		case tags.StartsNotWith:
			if err != nil {
				return BadValidationRule(string(tags.StartsNotWith), text, group, err)
			}
			return NewValidationRule(string(tags.StartsNotWith), text, group, func(field any, object any) error {
				return rules.StartsNotWith(field, object, ruleargs)
			})
		case tags.EndsWith:
			if err != nil {
				return BadValidationRule(string(tags.EndsWith), text, group, err)
			}
			return NewValidationRule(string(tags.EndsWith), text, group, func(field any, object any) error {
				return rules.EndsWith(field, object, ruleargs)
			})
		case tags.EndsNotWith:
			if err != nil {
				return BadValidationRule(string(tags.EndsNotWith), text, group, err)
			}
			return NewValidationRule(string(tags.EndsNotWith), text, group, func(field any, object any) error {
				return rules.EndsNotWith(field, object, ruleargs)
			})
		case tags.Contains:
			if err != nil {
				return BadValidationRule(string(tags.Contains), text, group, err)
			}
			return NewValidationRule(string(tags.Contains), text, group, func(field any, object any) error {
				return rules.Contains(field, object, ruleargs)
			})
		case tags.ContainsNot:
			if err != nil {
				return BadValidationRule(string(tags.ContainsNot), text, group, err)
			}
			return NewValidationRule(string(tags.ContainsNot), text, group, func(field any, object any) error {
				return rules.ContainsNot(field, object, ruleargs)
			})
		case tags.Length:
			if err != nil {
				return BadValidationRule(string(tags.Length), text, group, err)
			}
			return NewValidationRule(string(tags.Length), text, group, func(field any, object any) error {
				return rules.Length(field, object, ruleargs)
			})
		case tags.Required, tags.Alpha, tags.AlphaNumeric, tags.AlphaUnicode, tags.AlphaNumericUnicode, tags.Numeric, tags.NumericUnsigned, tags.Hexadecimal, tags.HexColor, tags.RGB, tags.RGBA, tags.HSL, tags.HSLA, tags.Email, tags.ISSN, tags.E164, tags.Base32, tags.Base32Hex, tags.Base64, tags.Base64Raw, tags.Base64URL, tags.Base64RawURL, tags.Isbn10, tags.Isbn13, tags.SSN, tags.UUID, tags.UUID3, tags.UUID4, tags.UUID5, tags.ULID, tags.MD4, tags.MD5, tags.SHA, tags.SHA0, tags.SHA1, tags.SHA2, tags.SHA3, tags.SHA224, tags.SHA256, tags.SHA384, tags.SHA512, tags.ASCII, tags.PrintableASCII, tags.MultiByte, tags.Uppercase, tags.Lowercase, tags.DataURI, tags.Latitude, tags.Longitude, tags.Hostname, tags.Fqdn, tags.UrlEncoded, tags.HTML, tags.HTMLEncoded, tags.JWT, tags.BIC, tags.SemVer, tags.DNS, tags.CVE, tags.Cron:
			return BadValidationRule(rulename, text, group, fmt.Errorf("rule: %s accepts no arguments", text))
		default:
			return BadValidationRule(string(tags.Unknown), text, group, fmt.Errorf("unknown rule: %s", text))
		}
	}
}
