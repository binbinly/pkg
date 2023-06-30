package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegexMatch(t *testing.T) {
	assert.Equal(t, true, RegexMatch("abc", AlphaMatcher))
	assert.Equal(t, false, RegexMatch("111", AlphaMatcher))
	assert.Equal(t, true, RegexMatch("abc", LetterRegexMatcher))
	assert.Equal(t, true, RegexMatch("123", NumberRegexMatcher))
	assert.Equal(t, true, RegexMatch("abc.com", DnsMatcher))
	assert.Equal(t, false, RegexMatch("a.b.com", DnsMatcher))
	assert.Equal(t, true, RegexMatch("abc@xyz.com", EmailMatcher))
	assert.Equal(t, false, RegexMatch("a.b@@com", EmailMatcher))
	assert.Equal(t, true, RegexMatch("13263527980", ChineseMobileMatcher))
	assert.Equal(t, false, RegexMatch("434324324", ChineseMobileMatcher))
	assert.Equal(t, true, RegexMatch("210911192105130715", ChineseIdMatcher))
	assert.Equal(t, false, RegexMatch("123456", ChineseIdMatcher))
	assert.Equal(t, true, RegexMatch("010-32116675", ChinesePhoneMatcher))
	assert.Equal(t, true, RegexMatch("0731-82251545", ChinesePhoneMatcher))
	assert.Equal(t, false, RegexMatch("123-87562", ChinesePhoneMatcher))
	assert.Equal(t, true, RegexMatch("你好", ChineseMatcher))
	assert.Equal(t, false, RegexMatch("hello", ChineseMatcher))
	assert.Equal(t, true, RegexMatch("4111111111111111", CreditCardMatcher))
	assert.Equal(t, false, RegexMatch("123456", CreditCardMatcher))
	assert.Equal(t, true, RegexMatch("aGVsbG8=", Base64Matcher))
	assert.Equal(t, false, RegexMatch("123456", Base64Matcher))
}

func TestIsAllUpper(t *testing.T) {
	assert.Equal(t, true, IsAllUpper("ABC"))
	assert.Equal(t, false, IsAllUpper(""))
	assert.Equal(t, false, IsAllUpper("abc"))
	assert.Equal(t, false, IsAllUpper("aBC"))
	assert.Equal(t, false, IsAllUpper("1BC"))
	assert.Equal(t, false, IsAllUpper("1bc"))
	assert.Equal(t, false, IsAllUpper("123"))
	assert.Equal(t, false, IsAllUpper("你好"))
	assert.Equal(t, false, IsAllUpper("A&"))
	assert.Equal(t, false, IsAllUpper("&@#$%^&*"))
}

func TestIsAllLower(t *testing.T) {
	assert.Equal(t, true, IsAllLower("abc"))
	assert.Equal(t, false, IsAllLower("ABC"))
	assert.Equal(t, false, IsAllLower(""))
	assert.Equal(t, false, IsAllLower("aBC"))
	assert.Equal(t, false, IsAllLower("1BC"))
	assert.Equal(t, false, IsAllLower("1bc"))
	assert.Equal(t, false, IsAllLower("123"))
	assert.Equal(t, false, IsAllLower("你好"))
	assert.Equal(t, false, IsAllLower("A&"))
	assert.Equal(t, false, IsAllLower("&@#$%^&*"))
}

func TestContainLower(t *testing.T) {
	assert.Equal(t, true, ContainLower("abc"))
	assert.Equal(t, true, ContainLower("aBC"))
	assert.Equal(t, true, ContainLower("1bc"))
	assert.Equal(t, true, ContainLower("a&"))

	assert.Equal(t, false, ContainLower("ABC"))
	assert.Equal(t, false, ContainLower(""))
	assert.Equal(t, false, ContainLower("1BC"))
	assert.Equal(t, false, ContainLower("123"))
	assert.Equal(t, false, ContainLower("你好"))
	assert.Equal(t, false, ContainLower("&@#$%^&*"))
}

func TestContainUpper(t *testing.T) {
	assert.Equal(t, true, ContainUpper("ABC"))
	assert.Equal(t, true, ContainUpper("aBC"))
	assert.Equal(t, true, ContainUpper("1BC"))
	assert.Equal(t, true, ContainUpper("A&"))

	assert.Equal(t, false, ContainUpper("abc"))
	assert.Equal(t, false, ContainUpper(""))
	assert.Equal(t, false, ContainUpper("1bc"))
	assert.Equal(t, false, ContainUpper("123"))
	assert.Equal(t, false, ContainUpper("你好"))
	assert.Equal(t, false, ContainUpper("&@#$%^&*"))
}

func TestIsJSON(t *testing.T) {
	assert.Equal(t, true, IsJSON("{}"))
	assert.Equal(t, true, IsJSON("{\"name\": \"test\"}"))
	assert.Equal(t, true, IsJSON("[]"))
	assert.Equal(t, true, IsJSON("123"))

	assert.Equal(t, false, IsJSON(""))
	assert.Equal(t, false, IsJSON("abc"))
	assert.Equal(t, false, IsJSON("你好"))
	assert.Equal(t, false, IsJSON("&@#$%^&*"))
}

func TestIsNumber(t *testing.T) {
	assert.Equal(t, false, IsNumber(""))
	assert.Equal(t, false, IsNumber("3"))
	assert.Equal(t, true, IsNumber(0))
	assert.Equal(t, true, IsNumber(0.1))
}

func TestIsFloat(t *testing.T) {
	assert.Equal(t, false, IsFloat(""))
	assert.Equal(t, false, IsFloat("3"))
	assert.Equal(t, false, IsFloat(0))
	assert.Equal(t, true, IsFloat(0.1))
}

func TestIsInt(t *testing.T) {
	assert.Equal(t, false, IsInt(""))
	assert.Equal(t, false, IsInt("3"))
	assert.Equal(t, false, IsInt(0.1))
	assert.Equal(t, true, IsInt(0))
	assert.Equal(t, true, IsInt(-1))
}

func TestIsNumberStr(t *testing.T) {
	assert.Equal(t, true, IsNumberStr("3."))
	assert.Equal(t, true, IsNumberStr("+3."))
	assert.Equal(t, true, IsNumberStr("-3."))
	assert.Equal(t, true, IsNumberStr("+3e2"))
	assert.Equal(t, false, IsNumberStr("abc"))
}

func TestIsFloatStr(t *testing.T) {
	assert.Equal(t, true, IsFloatStr("3."))
	assert.Equal(t, true, IsFloatStr("+3."))
	assert.Equal(t, true, IsFloatStr("-3."))
	assert.Equal(t, true, IsFloatStr("12"))
	assert.Equal(t, false, IsFloatStr("abc"))
}

func TestIsIntStr(t *testing.T) {
	assert.Equal(t, true, IsIntStr("+3"))
	assert.Equal(t, true, IsIntStr("-3"))
	assert.Equal(t, false, IsIntStr("3."))
	assert.Equal(t, false, IsIntStr("abc"))
}

func TestIsPort(t *testing.T) {
	assert.Equal(t, true, IsPort("1"))
	assert.Equal(t, true, IsPort("65535"))
	assert.Equal(t, false, IsPort("abc"))
	assert.Equal(t, false, IsPort("123abc"))
	assert.Equal(t, false, IsPort(""))
	assert.Equal(t, false, IsPort("-1"))
	assert.Equal(t, false, IsPort("65536"))
}

func TestIsIp(t *testing.T) {
	assert.Equal(t, true, IsIp("127.0.0.1"))
	assert.Equal(t, true, IsIp("::0:0:0:0:0:0:1"))
	assert.Equal(t, false, IsIp("127.0.0"))
	assert.Equal(t, false, IsIp("127"))
}

func TestIsIpV4(t *testing.T) {
	assert.Equal(t, true, IsIpV4("127.0.0.1"))
	assert.Equal(t, false, IsIpV4("::0:0:0:0:0:0:1"))
}

func TestIsIpV6(t *testing.T) {
	assert.Equal(t, false, IsIpV6("127.0.0.1"))
	assert.Equal(t, true, IsIpV6("::0:0:0:0:0:0:1"))
}

func TestIsUrl(t *testing.T) {
	assert.Equal(t, true, IsUrl("http://abc.com"))
	assert.Equal(t, true, IsUrl("abc.com"))
	assert.Equal(t, true, IsUrl("a.b.com"))
	assert.Equal(t, false, IsUrl("abc"))
}

func TestIsEmptyString(t *testing.T) {
	assert.Equal(t, true, IsEmptyString(""))
	assert.Equal(t, false, IsEmptyString("111"))
	assert.Equal(t, false, IsEmptyString(" "))
	assert.Equal(t, false, IsEmptyString("\t"))
}

func TestIsRegexMatch(t *testing.T) {
	assert.Equal(t, true, IsRegexMatch("abc", `^[a-zA-Z]+$`))
	assert.Equal(t, false, IsRegexMatch("1ab", `^[a-zA-Z]+$`))
	assert.Equal(t, false, IsRegexMatch("", `^[a-zA-Z]+$`))
}

func TestIsStrongPassword(t *testing.T) {
	assert.Equal(t, false, IsStrongPassword("abc", 3))
	assert.Equal(t, false, IsStrongPassword("abc123", 6))
	assert.Equal(t, false, IsStrongPassword("abcABC", 6))
	assert.Equal(t, false, IsStrongPassword("abc123@#$", 9))
	assert.Equal(t, false, IsStrongPassword("abcABC123@#$", 16))
	assert.Equal(t, true, IsStrongPassword("abcABC123@#$", 12))
	assert.Equal(t, true, IsStrongPassword("abcABC123@#$", 10))
}

func TestIsWeakPassword(t *testing.T) {
	assert.Equal(t, true, IsWeakPassword("abc"))
	assert.Equal(t, true, IsWeakPassword("123"))
	assert.Equal(t, true, IsWeakPassword("abc123"))
	assert.Equal(t, true, IsWeakPassword("abcABC123"))
	assert.Equal(t, false, IsWeakPassword("abc123@#$"))
}

func TestIsASCII(t *testing.T) {
	assert.Equal(t, true, IsASCII("ABC"))
	assert.Equal(t, true, IsASCII("123"))
	assert.Equal(t, true, IsASCII(""))
	assert.Equal(t, false, IsASCII("😄"))
	assert.Equal(t, false, IsASCII("你好"))
}

func TestIsPrintable(t *testing.T) {
	assert.Equal(t, true, IsPrintable("ABC"))
	assert.Equal(t, true, IsPrintable("{id: 123}"))
	assert.Equal(t, true, IsPrintable(""))
	assert.Equal(t, true, IsPrintable("😄"))
	assert.Equal(t, false, IsPrintable("\u0000"))
}
