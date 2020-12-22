// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Format string

// Enum values for Format
const (
	FormatAuto       Format = "AUTO"
	FormatNumber     Format = "NUMBER"
	FormatCurrency   Format = "CURRENCY"
	FormatDate       Format = "DATE"
	FormatTime       Format = "TIME"
	FormatDatetime   Format = "DATE_TIME"
	FormatPercentage Format = "PERCENTAGE"
	FormatText       Format = "TEXT"
	FormatAccounting Format = "ACCOUNTING"
	FormatContact    Format = "CONTACT"
	FormatRowlink    Format = "ROWLINK"
)

// Values returns all known values for Format. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Format) Values() []Format {
	return []Format{
		"AUTO",
		"NUMBER",
		"CURRENCY",
		"DATE",
		"TIME",
		"DATE_TIME",
		"PERCENTAGE",
		"TEXT",
		"ACCOUNTING",
		"CONTACT",
		"ROWLINK",
	}
}