package sr_Latn_RS

import (
	"math"

	"github.com/go-playground/locales"
)

type sr_Latn_RS struct {
	locale     string
	plurals    []locales.PluralRule
	decimal    []byte
	group      []byte
	minus      []byte
	percent    []byte
	perMille   []byte
	currencies [][]byte // idx = enum of currency code
}

// New returns a new instance of translator for the 'sr_Latn_RS' locale
func New() locales.Translator {
	return &sr_Latn_RS{
		locale:     "sr_Latn_RS",
		plurals:    []locales.PluralRule{2, 4, 6},
		decimal:    []byte{0x2c},
		group:      []byte{0x2e},
		minus:      []byte{0x2d},
		percent:    []byte{0x25},
		perMille:   []byte{0xe2, 0x80, 0xb0},
		currencies: [][]uint8{[]uint8{0x41, 0x44, 0x50}, []uint8{0x41, 0x45, 0x44}, []uint8{0x41, 0x46, 0x41}, []uint8{0x41, 0x46, 0x4e}, []uint8{0x41, 0x4c, 0x4b}, []uint8{0x41, 0x4c, 0x4c}, []uint8{0x41, 0x4d, 0x44}, []uint8{0x41, 0x4e, 0x47}, []uint8{0x41, 0x4f, 0x41}, []uint8{0x41, 0x4f, 0x4b}, []uint8{0x41, 0x4f, 0x4e}, []uint8{0x41, 0x4f, 0x52}, []uint8{0x41, 0x52, 0x41}, []uint8{0x41, 0x52, 0x4c}, []uint8{0x41, 0x52, 0x4d}, []uint8{0x41, 0x52, 0x50}, []uint8{0x41, 0x52, 0x53}, []uint8{0x41, 0x54, 0x53}, []uint8{0x41, 0x55, 0x44}, []uint8{0x41, 0x57, 0x47}, []uint8{0x41, 0x5a, 0x4d}, []uint8{0x41, 0x5a, 0x4e}, []uint8{0x42, 0x41, 0x44}, []uint8{0x42, 0x41, 0x4d}, []uint8{0x42, 0x41, 0x4e}, []uint8{0x42, 0x42, 0x44}, []uint8{0x42, 0x44, 0x54}, []uint8{0x42, 0x45, 0x43}, []uint8{0x42, 0x45, 0x46}, []uint8{0x42, 0x45, 0x4c}, []uint8{0x42, 0x47, 0x4c}, []uint8{0x42, 0x47, 0x4d}, []uint8{0x42, 0x47, 0x4e}, []uint8{0x42, 0x47, 0x4f}, []uint8{0x42, 0x48, 0x44}, []uint8{0x42, 0x49, 0x46}, []uint8{0x42, 0x4d, 0x44}, []uint8{0x42, 0x4e, 0x44}, []uint8{0x42, 0x4f, 0x42}, []uint8{0x42, 0x4f, 0x4c}, []uint8{0x42, 0x4f, 0x50}, []uint8{0x42, 0x4f, 0x56}, []uint8{0x42, 0x52, 0x42}, []uint8{0x42, 0x52, 0x43}, []uint8{0x42, 0x52, 0x45}, []uint8{0x42, 0x52, 0x4c}, []uint8{0x42, 0x52, 0x4e}, []uint8{0x42, 0x52, 0x52}, []uint8{0x42, 0x52, 0x5a}, []uint8{0x42, 0x53, 0x44}, []uint8{0x42, 0x54, 0x4e}, []uint8{0x42, 0x55, 0x4b}, []uint8{0x42, 0x57, 0x50}, []uint8{0x42, 0x59, 0x42}, []uint8{0x42, 0x59, 0x52}, []uint8{0x42, 0x5a, 0x44}, []uint8{0x43, 0x41, 0x44}, []uint8{0x43, 0x44, 0x46}, []uint8{0x43, 0x48, 0x45}, []uint8{0x43, 0x48, 0x46}, []uint8{0x43, 0x48, 0x57}, []uint8{0x43, 0x4c, 0x45}, []uint8{0x43, 0x4c, 0x46}, []uint8{0x43, 0x4c, 0x50}, []uint8{0x43, 0x4e, 0x58}, []uint8{0x43, 0x4e, 0x59}, []uint8{0x43, 0x4f, 0x50}, []uint8{0x43, 0x4f, 0x55}, []uint8{0x43, 0x52, 0x43}, []uint8{0x43, 0x53, 0x44}, []uint8{0x43, 0x53, 0x4b}, []uint8{0x43, 0x55, 0x43}, []uint8{0x43, 0x55, 0x50}, []uint8{0x43, 0x56, 0x45}, []uint8{0x43, 0x59, 0x50}, []uint8{0x43, 0x5a, 0x4b}, []uint8{0x44, 0x44, 0x4d}, []uint8{0x44, 0x45, 0x4d}, []uint8{0x44, 0x4a, 0x46}, []uint8{0x44, 0x4b, 0x4b}, []uint8{0x44, 0x4f, 0x50}, []uint8{0x44, 0x5a, 0x44}, []uint8{0x45, 0x43, 0x53}, []uint8{0x45, 0x43, 0x56}, []uint8{0x45, 0x45, 0x4b}, []uint8{0x45, 0x47, 0x50}, []uint8{0x45, 0x52, 0x4e}, []uint8{0x45, 0x53, 0x41}, []uint8{0x45, 0x53, 0x42}, []uint8{0x45, 0x53, 0x50}, []uint8{0x45, 0x54, 0x42}, []uint8{0x45, 0x55, 0x52}, []uint8{0x46, 0x49, 0x4d}, []uint8{0x46, 0x4a, 0x44}, []uint8{0x46, 0x4b, 0x50}, []uint8{0x46, 0x52, 0x46}, []uint8{0x47, 0x42, 0x50}, []uint8{0x47, 0x45, 0x4b}, []uint8{0x47, 0x45, 0x4c}, []uint8{0x47, 0x48, 0x43}, []uint8{0x47, 0x48, 0x53}, []uint8{0x47, 0x49, 0x50}, []uint8{0x47, 0x4d, 0x44}, []uint8{0x47, 0x4e, 0x46}, []uint8{0x47, 0x4e, 0x53}, []uint8{0x47, 0x51, 0x45}, []uint8{0x47, 0x52, 0x44}, []uint8{0x47, 0x54, 0x51}, []uint8{0x47, 0x57, 0x45}, []uint8{0x47, 0x57, 0x50}, []uint8{0x47, 0x59, 0x44}, []uint8{0x48, 0x4b, 0x44}, []uint8{0x48, 0x4e, 0x4c}, []uint8{0x48, 0x52, 0x44}, []uint8{0x48, 0x52, 0x4b}, []uint8{0x48, 0x54, 0x47}, []uint8{0x48, 0x55, 0x46}, []uint8{0x49, 0x44, 0x52}, []uint8{0x49, 0x45, 0x50}, []uint8{0x49, 0x4c, 0x50}, []uint8{0x49, 0x4c, 0x52}, []uint8{0x49, 0x4c, 0x53}, []uint8{0x49, 0x4e, 0x52}, []uint8{0x49, 0x51, 0x44}, []uint8{0x49, 0x52, 0x52}, []uint8{0x49, 0x53, 0x4a}, []uint8{0x49, 0x53, 0x4b}, []uint8{0x49, 0x54, 0x4c}, []uint8{0x4a, 0x4d, 0x44}, []uint8{0x4a, 0x4f, 0x44}, []uint8{0x4a, 0x50, 0x59}, []uint8{0x4b, 0x45, 0x53}, []uint8{0x4b, 0x47, 0x53}, []uint8{0x4b, 0x48, 0x52}, []uint8{0x4b, 0x4d, 0x46}, []uint8{0x4b, 0x50, 0x57}, []uint8{0x4b, 0x52, 0x48}, []uint8{0x4b, 0x52, 0x4f}, []uint8{0x4b, 0x52, 0x57}, []uint8{0x4b, 0x57, 0x44}, []uint8{0x4b, 0x59, 0x44}, []uint8{0x4b, 0x5a, 0x54}, []uint8{0x4c, 0x41, 0x4b}, []uint8{0x4c, 0x42, 0x50}, []uint8{0x4c, 0x4b, 0x52}, []uint8{0x4c, 0x52, 0x44}, []uint8{0x4c, 0x53, 0x4c}, []uint8{0x4c, 0x54, 0x4c}, []uint8{0x4c, 0x54, 0x54}, []uint8{0x4c, 0x55, 0x43}, []uint8{0x4c, 0x55, 0x46}, []uint8{0x4c, 0x55, 0x4c}, []uint8{0x4c, 0x56, 0x4c}, []uint8{0x4c, 0x56, 0x52}, []uint8{0x4c, 0x59, 0x44}, []uint8{0x4d, 0x41, 0x44}, []uint8{0x4d, 0x41, 0x46}, []uint8{0x4d, 0x43, 0x46}, []uint8{0x4d, 0x44, 0x43}, []uint8{0x4d, 0x44, 0x4c}, []uint8{0x4d, 0x47, 0x41}, []uint8{0x4d, 0x47, 0x46}, []uint8{0x4d, 0x4b, 0x44}, []uint8{0x4d, 0x4b, 0x4e}, []uint8{0x4d, 0x4c, 0x46}, []uint8{0x4d, 0x4d, 0x4b}, []uint8{0x4d, 0x4e, 0x54}, []uint8{0x4d, 0x4f, 0x50}, []uint8{0x4d, 0x52, 0x4f}, []uint8{0x4d, 0x54, 0x4c}, []uint8{0x4d, 0x54, 0x50}, []uint8{0x4d, 0x55, 0x52}, []uint8{0x4d, 0x56, 0x50}, []uint8{0x4d, 0x56, 0x52}, []uint8{0x4d, 0x57, 0x4b}, []uint8{0x4d, 0x58, 0x4e}, []uint8{0x4d, 0x58, 0x50}, []uint8{0x4d, 0x58, 0x56}, []uint8{0x4d, 0x59, 0x52}, []uint8{0x4d, 0x5a, 0x45}, []uint8{0x4d, 0x5a, 0x4d}, []uint8{0x4d, 0x5a, 0x4e}, []uint8{0x4e, 0x41, 0x44}, []uint8{0x4e, 0x47, 0x4e}, []uint8{0x4e, 0x49, 0x43}, []uint8{0x4e, 0x49, 0x4f}, []uint8{0x4e, 0x4c, 0x47}, []uint8{0x4e, 0x4f, 0x4b}, []uint8{0x4e, 0x50, 0x52}, []uint8{0x4e, 0x5a, 0x44}, []uint8{0x4f, 0x4d, 0x52}, []uint8{0x50, 0x41, 0x42}, []uint8{0x50, 0x45, 0x49}, []uint8{0x50, 0x45, 0x4e}, []uint8{0x50, 0x45, 0x53}, []uint8{0x50, 0x47, 0x4b}, []uint8{0x50, 0x48, 0x50}, []uint8{0x50, 0x4b, 0x52}, []uint8{0x50, 0x4c, 0x4e}, []uint8{0x50, 0x4c, 0x5a}, []uint8{0x50, 0x54, 0x45}, []uint8{0x50, 0x59, 0x47}, []uint8{0x51, 0x41, 0x52}, []uint8{0x52, 0x48, 0x44}, []uint8{0x52, 0x4f, 0x4c}, []uint8{0x52, 0x4f, 0x4e}, []uint8{0x52, 0x53, 0x44}, []uint8{0x52, 0x55, 0x42}, []uint8{0x52, 0x55, 0x52}, []uint8{0x52, 0x57, 0x46}, []uint8{0x53, 0x41, 0x52}, []uint8{0x53, 0x42, 0x44}, []uint8{0x53, 0x43, 0x52}, []uint8{0x53, 0x44, 0x44}, []uint8{0x53, 0x44, 0x47}, []uint8{0x53, 0x44, 0x50}, []uint8{0x53, 0x45, 0x4b}, []uint8{0x53, 0x47, 0x44}, []uint8{0x53, 0x48, 0x50}, []uint8{0x53, 0x49, 0x54}, []uint8{0x53, 0x4b, 0x4b}, []uint8{0x53, 0x4c, 0x4c}, []uint8{0x53, 0x4f, 0x53}, []uint8{0x53, 0x52, 0x44}, []uint8{0x53, 0x52, 0x47}, []uint8{0x53, 0x53, 0x50}, []uint8{0x53, 0x54, 0x44}, []uint8{0x53, 0x55, 0x52}, []uint8{0x53, 0x56, 0x43}, []uint8{0x53, 0x59, 0x50}, []uint8{0x53, 0x5a, 0x4c}, []uint8{0x54, 0x48, 0x42}, []uint8{0x54, 0x4a, 0x52}, []uint8{0x54, 0x4a, 0x53}, []uint8{0x54, 0x4d, 0x4d}, []uint8{0x54, 0x4d, 0x54}, []uint8{0x54, 0x4e, 0x44}, []uint8{0x54, 0x4f, 0x50}, []uint8{0x54, 0x50, 0x45}, []uint8{0x54, 0x52, 0x4c}, []uint8{0x54, 0x52, 0x59}, []uint8{0x54, 0x54, 0x44}, []uint8{0x54, 0x57, 0x44}, []uint8{0x54, 0x5a, 0x53}, []uint8{0x55, 0x41, 0x48}, []uint8{0x55, 0x41, 0x4b}, []uint8{0x55, 0x47, 0x53}, []uint8{0x55, 0x47, 0x58}, []uint8{0x55, 0x53, 0x44}, []uint8{0x55, 0x53, 0x4e}, []uint8{0x55, 0x53, 0x53}, []uint8{0x55, 0x59, 0x49}, []uint8{0x55, 0x59, 0x50}, []uint8{0x55, 0x59, 0x55}, []uint8{0x55, 0x5a, 0x53}, []uint8{0x56, 0x45, 0x42}, []uint8{0x56, 0x45, 0x46}, []uint8{0x56, 0x4e, 0x44}, []uint8{0x56, 0x4e, 0x4e}, []uint8{0x56, 0x55, 0x56}, []uint8{0x57, 0x53, 0x54}, []uint8{0x58, 0x41, 0x46}, []uint8{0x58, 0x41, 0x47}, []uint8{0x58, 0x41, 0x55}, []uint8{0x58, 0x42, 0x41}, []uint8{0x58, 0x42, 0x42}, []uint8{0x58, 0x42, 0x43}, []uint8{0x58, 0x42, 0x44}, []uint8{0x58, 0x43, 0x44}, []uint8{0x58, 0x44, 0x52}, []uint8{0x58, 0x45, 0x55}, []uint8{0x58, 0x46, 0x4f}, []uint8{0x58, 0x46, 0x55}, []uint8{0x58, 0x4f, 0x46}, []uint8{0x58, 0x50, 0x44}, []uint8{0x58, 0x50, 0x46}, []uint8{0x58, 0x50, 0x54}, []uint8{0x58, 0x52, 0x45}, []uint8{0x58, 0x53, 0x55}, []uint8{0x58, 0x54, 0x53}, []uint8{0x58, 0x55, 0x41}, []uint8{0x58, 0x58, 0x58}, []uint8{0x59, 0x44, 0x44}, []uint8{0x59, 0x45, 0x52}, []uint8{0x59, 0x55, 0x44}, []uint8{0x59, 0x55, 0x4d}, []uint8{0x59, 0x55, 0x4e}, []uint8{0x59, 0x55, 0x52}, []uint8{0x5a, 0x41, 0x4c}, []uint8{0x5a, 0x41, 0x52}, []uint8{0x5a, 0x4d, 0x4b}, []uint8{0x5a, 0x4d, 0x57}, []uint8{0x5a, 0x52, 0x4e}, []uint8{0x5a, 0x52, 0x5a}, []uint8{0x5a, 0x57, 0x44}, []uint8{0x5a, 0x57, 0x4c}, []uint8{0x5a, 0x57, 0x52}},
	}
}

// Locale returns the current translators string locale
func (t *sr_Latn_RS) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sr_Latn_RS'
func (t *sr_Latn_RS) Plurals() []locales.PluralRule {
	return t.plurals
}

// cardinalPluralRule returns the PluralRule given 'num' and digits/precision of 'v' for 'sr_Latn_RS'
func (t *sr_Latn_RS) cardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)

	if (v == 0 && i%10 == 1 && i%100 != 11) || (f%10 == 1 && f%100 != 11) {
		return locales.PluralRuleOne
	} else if (v == 0 && i%10 >= 2 && i%10 <= 4 && i%100 < 12 && i%100 > 14) || (f%10 >= 2 && f%10 <= 4 && f%100 < 12 && f%100 > 14) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}
