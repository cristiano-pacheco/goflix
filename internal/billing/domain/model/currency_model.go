package model

import (
	"errors"
	"strings"
)

const (
	currencyCodeLength = 3
)

// CurrencyModel represents an ISO 4217 currency with complete information
type CurrencyModel struct {
	country  string
	currency string
	code     string
	number   string
}

func CreateCurrencyModel(code string) (CurrencyModel, error) {
	code = strings.TrimSpace(code)
	if err := validateCurrencyCode(code); err != nil {
		return CurrencyModel{}, err
	}

	currencyInfo := getCurrencyInfo(code)
	return CurrencyModel(currencyInfo), nil
}

func (c *CurrencyModel) Country() string {
	return c.country
}

func (c *CurrencyModel) Currency() string {
	return c.currency
}

func (c *CurrencyModel) Code() string {
	return c.code
}

func (c *CurrencyModel) Number() string {
	return c.number
}

type currencyInfo struct {
	country  string
	currency string
	code     string
	number   string
}

// getCurrencyInfo returns complete currency information for a given code
func getCurrencyInfo(code string) currencyInfo { //nolint:funlen
	currencyData := map[string]currencyInfo{
		"AFN": {"AFGHANISTAN", "Afghani", "AFN", "971"},
		"EUR": {"ÅLAND ISLANDS", "Euro", "EUR", "978"},
		"ALL": {"ALBANIA", "Lek", "ALL", "008"},
		"DZD": {"ALGERIA", "Algerian Dinar", "DZD", "012"},
		"USD": {"AMERICAN SAMOA", "US Dollar", "USD", "840"},
		"AOA": {"ANGOLA", "Kwanza", "AOA", "973"},
		"XCD": {"ANGUILLA", "East Caribbean Dollar", "XCD", "951"},
		"ARS": {"ARGENTINA", "Argentine Peso", "ARS", "032"},
		"AMD": {"ARMENIA", "Armenian Dram", "AMD", "051"},
		"AWG": {"ARUBA", "Aruban Florin", "AWG", "533"},
		"AUD": {"AUSTRALIA", "Australian Dollar", "AUD", "036"},
		"AZN": {"AZERBAIJAN", "Azerbaijanian Manat", "AZN", "944"},
		"BSD": {"BAHAMAS (THE)", "Bahamian Dollar", "BSD", "044"},
		"BHD": {"BAHRAIN", "Bahraini Dinar", "BHD", "048"},
		"BDT": {"BANGLADESH", "Taka", "BDT", "050"},
		"BBD": {"BARBADOS", "Barbados Dollar", "BBD", "052"},
		"BYN": {"BELARUS", "Belarussian Ruble", "BYN", "933"},
		"BZD": {"BELIZE", "Belize Dollar", "BZD", "084"},
		"XOF": {"BENIN", "CFA Franc BCEAO", "XOF", "952"},
		"BMD": {"BERMUDA", "Bermudian Dollar", "BMD", "060"},
		"BTN": {"BHUTAN", "Ngultrum", "BTN", "064"},
		"INR": {"BHUTAN", "Indian Rupee", "INR", "356"},
		"BOB": {"BOLIVIA (PLURINATIONAL STATE OF)", "Boliviano", "BOB", "068"},
		"BOV": {"BOLIVIA (PLURINATIONAL STATE OF)", "Mvdol", "BOV", "984"},
		"BAM": {"BOSNIA AND HERZEGOVINA", "Convertible Mark", "BAM", "977"},
		"BWP": {"BOTSWANA", "Pula", "BWP", "072"},
		"NOK": {"BOUVET ISLAND", "Norwegian Krone", "NOK", "578"},
		"BRL": {"BRAZIL", "Brazilian Real", "BRL", "986"},
		"BND": {"BRUNEI DARUSSALAM", "Brunei Dollar", "BND", "096"},
		"BGN": {"BULGARIA", "Bulgarian Lev", "BGN", "975"},
		"BIF": {"BURUNDI", "Burundi Franc", "BIF", "108"},
		"CVE": {"CABO VERDE", "Cabo Verde Escudo", "CVE", "132"},
		"KHR": {"CAMBODIA", "Riel", "KHR", "116"},
		"XAF": {"CAMEROON", "CFA Franc BEAC", "XAF", "950"},
		"CAD": {"CANADA", "Canadian Dollar", "CAD", "124"},
		"KYD": {"CAYMAN ISLANDS (THE)", "Cayman Islands Dollar", "KYD", "136"},
		"CLF": {"CHILE", "Unidad de Fomento", "CLF", "990"},
		"CLP": {"CHILE", "Chilean Peso", "CLP", "152"},
		"CNY": {"CHINA", "Yuan Renminbi", "CNY", "156"},
		"COP": {"COLOMBIA", "Colombian Peso", "COP", "170"},
		"COU": {"COLOMBIA", "Unidad de Valor Real", "COU", "970"},
		"KMF": {"COMOROS (THE)", "Comoro Franc", "KMF", "174"},
		"CDF": {"CONGO (THE DEMOCRATIC REPUBLIC OF THE)", "Congolese Franc", "CDF", "976"},
		"NZD": {"COOK ISLANDS (THE)", "New Zealand Dollar", "NZD", "554"},
		"CRC": {"COSTA RICA", "Costa Rican Colon", "CRC", "188"},
		"HRK": {"CROATIA", "Kuna", "HRK", "191"},
		"CUC": {"CUBA", "Peso Convertible", "CUC", "931"},
		"CUP": {"CUBA", "Cuban Peso", "CUP", "192"},
		"ANG": {"CURAÇAO", "Netherlands Antillean Guilder", "ANG", "532"},
		"CZK": {"CZECH REPUBLIC (THE)", "Czech Koruna", "CZK", "203"},
		"DKK": {"DENMARK", "Danish Krone", "DKK", "208"},
		"DJF": {"DJIBOUTI", "Djibouti Franc", "DJF", "262"},
		"DOP": {"DOMINICAN REPUBLIC (THE)", "Dominican Peso", "DOP", "214"},
		"EGP": {"EGYPT", "Egyptian Pound", "EGP", "818"},
		"SVC": {"EL SALVADOR", "El Salvador Colon", "SVC", "222"},
		"ERN": {"ERITREA", "Nakfa", "ERN", "232"},
		"SZL": {"ESWATINI", "Lilangeni", "SZL", "748"},
		"ETB": {"ETHIOPIA", "Ethiopian Birr", "ETB", "230"},
		"FKP": {"FALKLAND ISLANDS (THE)", "Falkland Islands Pound", "FKP", "238"},
		"FJD": {"FIJI", "Fiji Dollar", "FJD", "242"},
		"XPF": {"FRENCH POLYNESIA", "CFP Franc", "XPF", "953"},
		"GMD": {"GAMBIA (THE)", "Dalasi", "GMD", "270"},
		"GEL": {"GEORGIA", "Lari", "GEL", "981"},
		"GHS": {"GHANA", "Ghana Cedi", "GHS", "936"},
		"GIP": {"GIBRALTAR", "Gibraltar Pound", "GIP", "292"},
		"GTQ": {"GUATEMALA", "Quetzal", "GTQ", "320"},
		"GBP": {"UNITED KINGDOM", "Pound Sterling", "GBP", "826"},
		"GNF": {"GUINEA", "Guinean Franc", "GNF", "324"},
		"GYD": {"GUYANA", "Guyana Dollar", "GYD", "328"},
		"HTG": {"HAITI", "Gourde", "HTG", "332"},
		"HNL": {"HONDURAS", "Lempira", "HNL", "340"},
		"HKD": {"HONG KONG", "Hong Kong Dollar", "HKD", "344"},
		"HUF": {"HUNGARY", "Forint", "HUF", "348"},
		"ISK": {"ICELAND", "Iceland Krona", "ISK", "352"},
		"IDR": {"INDONESIA", "Rupiah", "IDR", "360"},
		"XDR": {"INTERNATIONAL MONETARY FUND (IMF)", "SDR (Special Drawing Right)", "XDR", "960"},
		"IRR": {"IRAN", "Iranian Rial", "IRR", "364"},
		"IQD": {"IRAQ", "Iraqi Dinar", "IQD", "368"},
		"ILS": {"ISRAEL", "New Israeli Sheqel", "ILS", "376"},
		"JMD": {"JAMAICA", "Jamaican Dollar", "JMD", "388"},
		"JPY": {"JAPAN", "Yen", "JPY", "392"},
		"JOD": {"JORDAN", "Jordanian Dinar", "JOD", "400"},
		"KZT": {"KAZAKHSTAN", "Tenge", "KZT", "398"},
		"KES": {"KENYA", "Kenyan Shilling", "KES", "404"},
		"KPW": {"KOREA (THE DEMOCRATIC PEOPLE'S REPUBLIC OF)", "North Korean Won", "KPW", "408"},
		"KRW": {"KOREA (THE REPUBLIC OF)", "Won", "KRW", "410"},
		"KWD": {"KUWAIT", "Kuwaiti Dinar", "KWD", "414"},
		"KGS": {"KYRGYZSTAN", "Som", "KGS", "417"},
		"LAK": {"LAO PEOPLE'S DEMOCRATIC REPUBLIC (THE)", "Lao Kip", "LAK", "418"},
		"LBP": {"LEBANON", "Lebanese Pound", "LBP", "422"},
		"LSL": {"LESOTHO", "Loti", "LSL", "426"},
		"ZAR": {"SOUTH AFRICA", "Rand", "ZAR", "710"},
		"LRD": {"LIBERIA", "Liberian Dollar", "LRD", "430"},
		"LYD": {"LIBYA", "Libyan Dinar", "LYD", "434"},
		"CHF": {"SWITZERLAND", "Swiss Franc", "CHF", "756"},
		"MOP": {"MACAO", "Pataca", "MOP", "446"},
		"MKD": {"NORTH MACEDONIA", "Denar", "MKD", "807"},
		"MGA": {"MADAGASCAR", "Malagasy Ariary", "MGA", "969"},
		"MWK": {"MALAWI", "Malawi Kwacha", "MWK", "454"},
		"MYR": {"MALAYSIA", "Malaysian Ringgit", "MYR", "458"},
		"MVR": {"MALDIVES", "Rufiyaa", "MVR", "462"},
		"MRU": {"MAURITANIA", "Ouguiya", "MRU", "929"},
		"MUR": {"MAURITIUS", "Mauritius Rupee", "MUR", "480"},
		"XUA": {"AFRICAN DEVELOPMENT BANK", "ADB Unit of Account", "XUA", "965"},
		"MXN": {"MEXICO", "Mexican Peso", "MXN", "484"},
		"MXV": {"MEXICO", "Mexican Unidad de Inversion (UDI)", "MXV", "979"},
		"MDL": {"REPUBLIC OF MOLDOVA", "Moldovan Leu", "MDL", "498"},
		"MNT": {"MONGOLIA", "Tugrik", "MNT", "496"},
		"MAD": {"MOROCCO", "Moroccan Dirham", "MAD", "504"},
		"MZN": {"MOZAMBIQUE", "Mozambique Metical", "MZN", "943"},
		"MMK": {"MYANMAR", "Kyat", "MMK", "104"},
		"NAD": {"NAMIBIA", "Namibia Dollar", "NAD", "516"},
		"NPR": {"NEPAL", "Nepalese Rupee", "NPR", "524"},
		"NIO": {"NICARAGUA", "Cordoba Oro", "NIO", "558"},
		"NGN": {"NIGERIA", "Naira", "NGN", "566"},
		"OMR": {"OMAN", "Rial Omani", "OMR", "512"},
		"PKR": {"PAKISTAN", "Pakistan Rupee", "PKR", "586"},
		"PAB": {"PANAMA", "Balboa", "PAB", "590"},
		"PGK": {"PAPUA NEW GUINEA", "Kina", "PGK", "598"},
		"PYG": {"PARAGUAY", "Guarani", "PYG", "600"},
		"PEN": {"PERU", "Sol", "PEN", "604"},
		"PHP": {"PHILIPPINES (THE)", "Philippine Peso", "PHP", "608"},
		"PLN": {"POLAND", "Zloty", "PLN", "985"},
		"QAR": {"QATAR", "Qatari Rial", "QAR", "634"},
		"RON": {"ROMANIA", "Romanian Leu", "RON", "946"},
		"RUB": {"RUSSIAN FEDERATION (THE)", "Russian Ruble", "RUB", "643"},
		"RWF": {"RWANDA", "Rwanda Franc", "RWF", "646"},
		"SHP": {"SAINT HELENA", "Saint Helena Pound", "SHP", "654"},
		"WST": {"SAMOA", "Tala", "WST", "882"},
		"STN": {"SAO TOME AND PRINCIPE", "Dobra", "STN", "930"},
		"SAR": {"SAUDI ARABIA", "Saudi Riyal", "SAR", "682"},
		"RSD": {"SERBIA", "Serbian Dinar", "RSD", "941"},
		"SCR": {"SEYCHELLES", "Seychelles Rupee", "SCR", "690"},
		"SLE": {"SIERRA LEONE", "Leone", "SLE", "925"},
		"SGD": {"SINGAPORE", "Singapore Dollar", "SGD", "702"},
		"XSU": {"SISTEMA UNITARIO DE COMPENSACION REGIONAL DE PAGOS", "Sucre", "XSU", "994"},
		"SBD": {"SOLOMON ISLANDS", "Solomon Islands Dollar", "SBD", "090"},
		"SOS": {"SOMALIA", "Somali Shilling", "SOS", "706"},
		"SSP": {"SOUTH SUDAN", "South Sudanese Pound", "SSP", "728"},
		"LKR": {"SRI LANKA", "Sri Lanka Rupee", "LKR", "144"},
		"SDG": {"SUDAN (THE)", "Sudanese Pound", "SDG", "938"},
		"SRD": {"SURINAME", "Surinam Dollar", "SRD", "968"},
		"SEK": {"SWEDEN", "Swedish Krona", "SEK", "752"},
		"CHE": {"SWITZERLAND", "WIR Euro", "CHE", "947"},
		"CHW": {"SWITZERLAND", "WIR Franc", "CHW", "948"},
		"SYP": {"SYRIAN ARAB REPUBLIC", "Syrian Pound", "SYP", "760"},
		"TWD": {"TAIWAN", "New Taiwan Dollar", "TWD", "901"},
		"TJS": {"TAJIKISTAN", "Somoni", "TJS", "972"},
		"TZS": {"TANZANIA", "Tanzanian Shilling", "TZS", "834"},
		"THB": {"THAILAND", "Baht", "THB", "764"},
		"TOP": {"TONGA", "Pa'anga", "TOP", "776"},
		"TTD": {"TRINIDAD AND TOBAGO", "Trinidad and Tobago Dollar", "TTD", "780"},
		"TND": {"TUNISIA", "Tunisian Dinar", "TND", "788"},
		"TRY": {"TURKEY", "Turkish Lira", "TRY", "949"},
		"TMT": {"TURKMENISTAN", "Turkmenistan New Manat", "TMT", "934"},
		"UGX": {"UGANDA", "Uganda Shilling", "UGX", "800"},
		"UAH": {"UKRAINE", "Hryvnia", "UAH", "980"},
		"AED": {"UNITED ARAB EMIRATES (THE)", "UAE Dirham", "AED", "784"},
		"USN": {"UNITED STATES OF AMERICA (THE)", "US Dollar (Next day)", "USN", "997"},
		"UYI": {"URUGUAY", "Uruguay Peso en Unidades Indexadas", "UYI", "940"},
		"UYU": {"URUGUAY", "Peso Uruguayo", "UYU", "858"},
		"UZS": {"UZBEKISTAN", "Uzbekistan Sum", "UZS", "860"},
		"VUV": {"VANUATU", "Vatu", "VUV", "548"},
		"VED": {"VENEZUELA", "Bolívar Soberano", "VED", "926"},
		"VEF": {"VENEZUELA", "Bolívar", "VEF", "937"},
		"VND": {"VIET NAM", "Dong", "VND", "704"},
		"YER": {"YEMEN", "Yemeni Rial", "YER", "886"},
		"ZMW": {"ZAMBIA", "Zambian Kwacha", "ZMW", "967"},
		"ZWL": {"ZIMBABWE", "Zimbabwe Dollar", "ZWL", "932"},
	}

	if info, exists := currencyData[code]; exists {
		return info
	}

	// Fallback for unknown codes
	return currencyInfo{
		country:  "UNKNOWN",
		currency: "Unknown Currency",
		code:     code,
		number:   "000",
	}
}

// validateCurrencyCode validates against ISO 4217 currency codes
func validateCurrencyCode(code string) error {
	if code == "" {
		return errors.New("currency code cannot be empty")
	}

	if len(code) != currencyCodeLength {
		return errors.New("currency code must be exactly 3 characters")
	}

	validCurrencies := map[string]struct{}{
		"AFN": {}, "EUR": {}, "ALL": {}, "DZD": {}, "USD": {}, "AOA": {}, "XCD": {}, "ARS": {},
		"AMD": {}, "AWG": {}, "AUD": {}, "AZN": {}, "BSD": {}, "BHD": {}, "BDT": {}, "BBD": {},
		"BYN": {}, "BZD": {}, "XOF": {}, "BMD": {}, "BTN": {}, "INR": {}, "BOB": {}, "BOV": {},
		"BAM": {}, "BWP": {}, "NOK": {}, "BRL": {}, "BND": {}, "BGN": {}, "BIF": {}, "CVE": {},
		"KHR": {}, "XAF": {}, "CAD": {}, "KYD": {}, "CLF": {}, "CLP": {}, "CNY": {}, "COP": {},
		"COU": {}, "KMF": {}, "CDF": {}, "NZD": {}, "CRC": {}, "HRK": {}, "CUC": {}, "CUP": {},
		"ANG": {}, "CZK": {}, "DKK": {}, "DJF": {}, "DOP": {}, "EGP": {}, "SVC": {}, "ERN": {},
		"SZL": {}, "ETB": {}, "FKP": {}, "FJD": {}, "XPF": {}, "GMD": {}, "GEL": {}, "GHS": {},
		"GIP": {}, "GTQ": {}, "GBP": {}, "GNF": {}, "GYD": {}, "HTG": {}, "HNL": {}, "HKD": {},
		"HUF": {}, "ISK": {}, "IDR": {}, "XDR": {}, "IRR": {}, "IQD": {}, "ILS": {}, "JMD": {},
		"JPY": {}, "JOD": {}, "KZT": {}, "KES": {}, "KPW": {}, "KRW": {}, "KWD": {}, "KGS": {},
		"LAK": {}, "LBP": {}, "LSL": {}, "ZAR": {}, "LRD": {}, "LYD": {}, "CHF": {}, "MOP": {},
		"MKD": {}, "MGA": {}, "MWK": {}, "MYR": {}, "MVR": {}, "MRU": {}, "MUR": {}, "XUA": {},
		"MXN": {}, "MXV": {}, "MDL": {}, "MNT": {}, "MAD": {}, "MZN": {}, "MMK": {}, "NAD": {},
		"NPR": {}, "NIO": {}, "NGN": {}, "OMR": {}, "PKR": {}, "PAB": {}, "PGK": {}, "PYG": {},
		"PEN": {}, "PHP": {}, "PLN": {}, "QAR": {}, "RON": {}, "RUB": {}, "RWF": {}, "SHP": {},
		"WST": {}, "STN": {}, "SAR": {}, "RSD": {}, "SCR": {}, "SLE": {}, "SGD": {}, "XSU": {},
		"SBD": {}, "SOS": {}, "SSP": {}, "LKR": {}, "SDG": {}, "SRD": {}, "SEK": {}, "CHE": {},
		"CHW": {}, "SYP": {}, "TWD": {}, "TJS": {}, "TZS": {}, "THB": {}, "TOP": {}, "TTD": {},
		"TND": {}, "TRY": {}, "TMT": {}, "UGX": {}, "UAH": {}, "AED": {}, "USN": {}, "UYI": {},
		"UYU": {}, "UZS": {}, "VUV": {}, "VED": {}, "VEF": {}, "VND": {}, "YER": {}, "ZMW": {},
		"ZWL": {},
	}

	if _, ok := validCurrencies[code]; !ok {
		return errors.New("invalid currency code: " + code)
	}

	return nil
}
