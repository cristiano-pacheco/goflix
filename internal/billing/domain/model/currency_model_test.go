package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
)

func TestCreateCurrencyModel(t *testing.T) {
	t.Run("valid currency code creates model successfully", func(t *testing.T) {
		// Arrange
		validCode := "USD"

		// Act
		result, err := model.CreateCurrencyModel(validCode)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, "AMERICAN SAMOA", result.Country())
		assert.Equal(t, "US Dollar", result.Currency())
		assert.Equal(t, "USD", result.Code())
		assert.Equal(t, "840", result.Number())
	})

	t.Run("valid currency code with whitespace creates model successfully", func(t *testing.T) {
		// Arrange
		codeWithWhitespace := "  EUR  "

		// Act
		result, err := model.CreateCurrencyModel(codeWithWhitespace)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, "ÅLAND ISLANDS", result.Country())
		assert.Equal(t, "Euro", result.Currency())
		assert.Equal(t, "EUR", result.Code())
		assert.Equal(t, "978", result.Number())
	})

	t.Run("valid currency code BRL creates model successfully", func(t *testing.T) {
		// Arrange
		validCode := "BRL"

		// Act
		result, err := model.CreateCurrencyModel(validCode)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, "BRAZIL", result.Country())
		assert.Equal(t, "Brazilian Real", result.Currency())
		assert.Equal(t, "BRL", result.Code())
		assert.Equal(t, "986", result.Number())
	})

	t.Run("valid currency code GBP creates model successfully", func(t *testing.T) {
		// Arrange
		validCode := "GBP"

		// Act
		result, err := model.CreateCurrencyModel(validCode)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, "UNITED KINGDOM", result.Country())
		assert.Equal(t, "Pound Sterling", result.Currency())
		assert.Equal(t, "GBP", result.Code())
		assert.Equal(t, "826", result.Number())
	})

	t.Run("empty currency code returns error", func(t *testing.T) {
		// Arrange
		emptyCode := ""

		// Act
		result, err := model.CreateCurrencyModel(emptyCode)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "currency code cannot be empty", err.Error())
		assert.Equal(t, model.CurrencyModel{}, result)
	})

	t.Run("whitespace only currency code returns error", func(t *testing.T) {
		// Arrange
		whitespaceCode := "   "

		// Act
		result, err := model.CreateCurrencyModel(whitespaceCode)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "currency code cannot be empty", err.Error())
		assert.Equal(t, model.CurrencyModel{}, result)
	})

	t.Run("currency code with less than 3 characters returns error", func(t *testing.T) {
		// Arrange
		shortCode := "US"

		// Act
		result, err := model.CreateCurrencyModel(shortCode)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "currency code must be exactly 3 characters", err.Error())
		assert.Equal(t, model.CurrencyModel{}, result)
	})

	t.Run("currency code with more than 3 characters returns error", func(t *testing.T) {
		// Arrange
		longCode := "USDD"

		// Act
		result, err := model.CreateCurrencyModel(longCode)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "currency code must be exactly 3 characters", err.Error())
		assert.Equal(t, model.CurrencyModel{}, result)
	})

	t.Run("invalid currency code returns error", func(t *testing.T) {
		// Arrange
		invalidCode := "XYZ"

		// Act
		result, err := model.CreateCurrencyModel(invalidCode)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "invalid currency code: XYZ", err.Error())
		assert.Equal(t, model.CurrencyModel{}, result)
	})

	t.Run("lowercase currency code returns error", func(t *testing.T) {
		// Arrange
		lowercaseCode := "usd"

		// Act
		result, err := model.CreateCurrencyModel(lowercaseCode)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "invalid currency code: usd", err.Error())
		assert.Equal(t, model.CurrencyModel{}, result)
	})
}

func TestCurrencyModel_Getters(t *testing.T) {
	t.Run("getters return correct values for JPY", func(t *testing.T) {
		// Arrange
		currencyModel, err := model.CreateCurrencyModel("JPY")
		require.NoError(t, err)

		// Act & Assert
		assert.Equal(t, "JAPAN", currencyModel.Country())
		assert.Equal(t, "Yen", currencyModel.Currency())
		assert.Equal(t, "JPY", currencyModel.Code())
		assert.Equal(t, "392", currencyModel.Number())
	})

	t.Run("getters return correct values for CHF", func(t *testing.T) {
		// Arrange
		currencyModel, err := model.CreateCurrencyModel("CHF")
		require.NoError(t, err)

		// Act & Assert
		assert.Equal(t, "SWITZERLAND", currencyModel.Country())
		assert.Equal(t, "Swiss Franc", currencyModel.Currency())
		assert.Equal(t, "CHF", currencyModel.Code())
		assert.Equal(t, "756", currencyModel.Number())
	})
}

func TestCreateCurrencyModel_AllSupportedCurrencies(t *testing.T) {
	// Arrange
	allSupportedCurrencies := []string{
		"AFN", "EUR", "ALL", "DZD", "USD", "AOA", "XCD", "ARS", "AMD", "AWG", "AUD", "AZN",
		"BSD", "BHD", "BDT", "BBD", "BYN", "BZD", "XOF", "BMD", "BTN", "INR", "BOB", "BOV",
		"BAM", "BWP", "NOK", "BRL", "BND", "BGN", "BIF", "CVE", "KHR", "XAF", "CAD", "KYD",
		"CLF", "CLP", "CNY", "COP", "COU", "KMF", "CDF", "NZD", "CRC", "HRK", "CUC", "CUP",
		"ANG", "CZK", "DKK", "DJF", "DOP", "EGP", "SVC", "ERN", "SZL", "ETB", "FKP", "FJD",
		"XPF", "GMD", "GEL", "GHS", "GIP", "GTQ", "GBP", "GNF", "GYD", "HTG", "HNL", "HKD",
		"HUF", "ISK", "IDR", "XDR", "IRR", "IQD", "ILS", "JMD", "JPY", "JOD", "KZT", "KES",
		"KPW", "KRW", "KWD", "KGS", "LAK", "LBP", "LSL", "ZAR", "LRD", "LYD", "CHF", "MOP",
		"MKD", "MGA", "MWK", "MYR", "MVR", "MRU", "MUR", "XUA", "MXN", "MXV", "MDL", "MNT",
		"MAD", "MZN", "MMK", "NAD", "NPR", "NIO", "NGN", "OMR", "PKR", "PAB", "PGK", "PYG",
		"PEN", "PHP", "PLN", "QAR", "RON", "RUB", "RWF", "SHP", "WST", "STN", "SAR", "RSD",
		"SCR", "SLE", "SGD", "XSU", "SBD", "SOS", "SSP", "LKR", "SDG", "SRD", "SEK", "CHE",
		"CHW", "SYP", "TWD", "TJS", "TZS", "THB", "TOP", "TTD", "TND", "TRY", "TMT", "UGX",
		"UAH", "AED", "USN", "UYI", "UYU", "UZS", "VUV", "VED", "VEF", "VND", "YER", "ZMW",
		"ZWL",
	}

	for _, currencyCode := range allSupportedCurrencies {
		t.Run("currency code "+currencyCode+" creates model successfully", func(t *testing.T) {
			// Act
			result, err := model.CreateCurrencyModel(currencyCode)

			// Assert
			require.NoError(t, err)
			assert.Equal(t, currencyCode, result.Code())
			assert.NotEmpty(t, result.Country())
			assert.NotEmpty(t, result.Currency())
			assert.NotEmpty(t, result.Number())
		})
	}
}
