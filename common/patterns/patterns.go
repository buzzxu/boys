package patterns

import (
	"regexp"
)

var (
	phoneRegex          = regexp.MustCompile(`0\d{2,3}-\d{7,8}|\(?0\d{2,3}[)-]?\d{7,8}|\(?0\d{2,3}[)-]*\d{7,8}`)
	mobileRegex         = regexp.MustCompile(`^(13[0-9]|14[579]|15[0-3,5-9]|16[2-7]|17[01235678]|18[0-9]|19[0-9])\d{8}$`)
	emailRegex          = regexp.MustCompile(`^([a-z0-9A-Z]+[-|\.]?)+[a-z0-9A-Z]@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\.)+[a-zA-Z]{2,}$`)
	usernameRegex       = regexp.MustCompile(`^([a-zA-Z]|[0-9])\w{4,17}$`)
	usernameCnRegex     = regexp.MustCompile(`^[(a-zA-Z0-9\u4e00-\u9fa5){1}_#]{2,20}$`)
	passwordRegex       = regexp.MustCompile(`^[a-zA-Z0-9]{6,16}$`)
	passwordStrongRegex = regexp.MustCompile(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?!.*\s).{6,16}$`)
	chineseRegex        = regexp.MustCompile(`^(?:[\u3400-\u4DB5\u4E00-\u9FEA\uFA0E\uFA0F\uFA11\uFA13\uFA14\uFA1F\uFA21\uFA23\uFA24\uFA27-\uFA29]|[\uD840-\uD868\uD86A-\uD86C\uD86F-\uD872\uD874-\uD879][\uDC00-\uDFFF]|\uD869[\uDC00-\uDED6\uDF00-\uDFFF]|\uD86D[\uDC00-\uDF34\uDF40-\uDFFF]|\uD86E[\uDC00-\uDC1D\uDC20-\uDFFF]|\uD873[\uDC00-\uDEA1\uDEB0-\uDFFF]|\uD87A[\uDC00-\uDFE0])+$`)
	nameChineseRegex    = regexp.MustCompile(`^(?:[\u4e00-\u9fa5·]{2,16})$`)
	nameEnglishRegex    = regexp.MustCompile(`(^[a-zA-Z]{1}[a-zA-Z\s]{0,20}[a-zA-Z]{1}$)`)
	idCardRegex         = regexp.MustCompile(`^(^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$)|(^[1-9]\d{5}[1-9]\d{3}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])((\d{4})|\d{3}[Xx])$)$`)
	idCard15Regex       = regexp.MustCompile(`^[1-9]\d{7}((0\d)|(1[0-2]))(([0|1|2]\d)|3[0-1])\d{3}$`)
	idCard18Regex       = regexp.MustCompile(`^[1-9]\d{5}(?:18|19|20)\d{2}(?:0[1-9]|10|11|12)(?:0[1-9]|[1-2]\d|30|31)\d{3}[\dXx]$`)
	passportRegex       = regexp.MustCompile(`(^[EeKkGgDdSsPpHh]\d{8}$)|(^(([Ee][a-fA-F])|([DdSsPp][Ee])|([Kk][Jj])|([Mm][Aa])|(1[45]))\d{7}$)`)
	urlRegex            = regexp.MustCompile(`^((ht|f)tps?:\/\/)?[\w-]+(\.[\w-]+)+:\d{1,5}\/?$`)
	ipRegex             = regexp.MustCompile(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`)
	numericRegex        = regexp.MustCompile(`^[-]?\d+[.]?\d*$`)
	numberSignRegex     = regexp.MustCompile(`^[+-]?[0-9]+$`)
	decimalRegex        = regexp.MustCompile(`^[+-]?[0-9]+$`)
	decimalSignRegex    = regexp.MustCompile(`^[+-]?\d+(\.\d+)?$`)
	letterRegex         = regexp.MustCompile(`^[A-Za-z]+$`)
	alphanumericRegex   = regexp.MustCompile(`^[0-9a-zA-Z]+$`)
	longitudeRegex      = regexp.MustCompile(`^([-+])?(((\d|[1-9]\d|1[0-7]\d|0{1,3})\.\d{0,6})|(\d|[1-9]\d|1[0-7]\d|0{1,3})|180\.0{0,6}|180)$`)
	latitudeRegex       = regexp.MustCompile(`^([-+])?([0-8]?\d{1}\.\d{0,6}|90\.0{0,6}|[0-8]?\d{1}|90)$`)
	carRegex            = regexp.MustCompile(`^(?:[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领 A-Z]{1}[A-HJ-NP-Z]{1}(?:(?:[0-9]{5}[DF])|(?:[DF](?:[A-HJ-NP-Z0-9])[0-9]{4})))|(?:[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领 A-Z]{1}[A-Z]{1}[A-HJ-NP-Z0-9]{4}[A-HJ-NP-Z0-9 挂学警港澳]{1})$`)
	orgRegex            = regexp.MustCompile(`^[0-9A-HJ-NPQRTUWXY]{2}\d{6}[0-9A-HJ-NPQRTUWXY]{10}$`)
	nameRegex           = regexp.MustCompile(`(^[\u4e00-\u9fa5]{1}[\u4e00-\u9fa5\.·。]{0,8}[\u4e00-\u9fa5]{1}$)|(^[a-zA-Z]{1}[a-zA-Z\s]{0,8}[a-zA-Z]{1}$)`)
	douyinVideoRegex    = regexp.MustCompile(`(?<=video\/).*(?=\/)`)
	specialRegex        = regexp.MustCompile(`\pP|\pS\pZ\pC|\s+`)
)

// Static methods to match the patterns
func Mobile(val string) bool {
	return val != "" && mobileRegex.MatchString(val)
}

func Phone(val string) bool {
	return val != "" && phoneRegex.MatchString(val)
}

func Email(val string) bool {
	return val != "" && emailRegex.MatchString(val)
}

func UserName(val string) bool {
	return val != "" && usernameRegex.MatchString(val)
}

func UserNameCN(val string) bool {
	return val != "" && usernameCnRegex.MatchString(val)
}

func Password(val string) bool {
	return val != "" && passwordRegex.MatchString(val)
}

func PasswordStrong(val string) bool {
	return passwordStrongRegex.MatchString(val)
}

func Chinese(val string) bool {
	return chineseRegex.MatchString(val)
}

func NameCN(val string) bool {
	return val != "" && nameChineseRegex.MatchString(val)
}

func NameEN(val string) bool {
	return val != "" && nameEnglishRegex.MatchString(val)
}

func IDCard(val string) bool {
	return val != "" && idCardRegex.MatchString(val)
}

func IDCard15(val string) bool {
	return val != "" && idCard15Regex.MatchString(val)
}

func IDCard18(val string) bool {
	return val != "" && idCard18Regex.MatchString(val)
}

func Passport(val string) bool {
	return val != "" && passportRegex.MatchString(val)
}

func URL(val string) bool {
	return val != "" && urlRegex.MatchString(val)
}

func IP(val string) bool {
	return val != "" && ipRegex.MatchString(val)
}

func Numeric(val string) bool {
	return val != "" && numericRegex.MatchString(val)
}

func NumberSign(val string) bool {
	return val != "" && numberSignRegex.MatchString(val)
}

func Decimal(val string) bool {
	return val != "" && decimalRegex.MatchString(val)
}

func DecimalSign(val string) bool {
	return val != "" && decimalSignRegex.MatchString(val)
}

func Letter(val string) bool {
	return val != "" && letterRegex.MatchString(val)
}

func Alphanumeric(val string) bool {
	return val != "" && alphanumericRegex.MatchString(val)
}

func Longitude(val string) bool {
	return val != "" && longitudeRegex.MatchString(val)
}

func Latitude(val string) bool {
	return val != "" && latitudeRegex.MatchString(val)
}

func Car(val string) bool {
	return val != "" && carRegex.MatchString(val)
}

func Organization(val string) bool {
	return val != "" && orgRegex.MatchString(val)
}

func Name(val string) bool {
	return val != "" && nameRegex.MatchString(val)
}

func DouyinVideo(val string) bool {
	return val != "" && douyinVideoRegex.MatchString(val)
}

func Special(val string) bool {
	return val != "" && specialRegex.MatchString(val)
}
