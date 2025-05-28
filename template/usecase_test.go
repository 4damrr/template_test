package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateCVUsecase(t *testing.T) {
	var (
		startDate, _ = time.Parse(time.DateOnly, "2025-01-01")
		endDate, _   = time.Parse(time.DateOnly, "2026-01-01")
	)

	tests := []struct {
		skipTest   bool
		testName   string
		reqPayload UserData
		assertFunc func(a *assert.Assertions, result string, err error)
	}{
		{
			skipTest: false,
			testName: "Success",
			reqPayload: UserData{
				Name:     "Adam",
				Email:    "adam@gmail.com",
				Phone:    "10219079",
				UserName: "4dam",
				Summary:  "Basically, Adam is just a chill guy.",
				Skills: []string{
					"Coding",
					"Sleeping",
					"Breathing",
				},
				Experiences: []Experience{
					{
						Name: "Badan Riset dan Inovasi Nasional",
						Descriptions: []string{
							"Conduct a research",
							"Doing some fun",
						},
						StartDate: startDate,
						EndDate:   endDate,
					},
					{
						Name: "The National Aeronautics and Space Administration",
						Descriptions: []string{
							"Conduct a research",
							"Building A Rocket",
						},
						StartDate: startDate,
						EndDate:   endDate,
					},
				},
				//Education: nil,
			},
			assertFunc: func(a *assert.Assertions, result string, err error) {
				t.Logf("This is the error: %v", err)
				t.Logf("This is the result: %v", result)
			},
		},
		{
			skipTest: false,
			testName: "Success",
			reqPayload: UserData{
				Name:     "Adam",
				Email:    "adam@gmail.com",
				Phone:    "10219079",
				UserName: "4dam",
				Summary:  "Basically, Adam is just a chill guy.",
				Skills: []string{
					"Coding",
					"Sleeping",
					"Breathing",
				},
				Experiences: []Experience{},
				//Education: nil,
			},
			assertFunc: func(a *assert.Assertions, result string, err error) {
				t.Logf("This is the error: %v", err)
				t.Logf("This is the result: %v", result)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			if test.skipTest {
				t.Logf("Skipping test '%v'", test.testName)
				t.Skip(test.testName, "Test skipped")
			}
			a := assert.New(t)

			result, err := GenerateCVUsecase(test.reqPayload)

			test.assertFunc(a, result, err)
		})
	}
}

func TestExprFormulaUsecase(t *testing.T) {
	tests := []struct {
		skipTest   bool
		testName   string
		reqPayload struct {
			formula string
			params  ExprEnv
		}
		assertFunc func(a *assert.Assertions, result any, err error)
	}{
		{
			skipTest: false,
			testName: "Success",
			reqPayload: struct {
				formula string
				params  ExprEnv
			}{
				formula: `price - (price * (discount / 100))`,
				params: ExprEnv{
					Price:    100,
					Discount: 20,
					Date:     time.Time{},
				},
			},
			assertFunc: func(a *assert.Assertions, result any, err error) {
				t.Logf("This is the error: %v", err)
				t.Logf("This is the result: %v", result)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			if test.skipTest {
				t.Logf("Skipping test '%v'", test.testName)
				t.Skip(test.testName, "Test skipped")
			}
			a := assert.New(t)

			result, err := ExprFormulaUsecase(test.reqPayload.formula, test.reqPayload.params)

			test.assertFunc(a, result, err)
		})
	}
}

func TestGenerateHTMLWithFormulaUsecase(t *testing.T) {
	var (
		endDate, _ = time.Parse(time.DateOnly, "2026-01-01")
	)

	tests := []struct {
		skipTest   bool
		testName   string
		reqPayload struct {
			data     ExprEnv
			formulas map[string]string
		}
		assertFunc func(a *assert.Assertions, result any, err error)
	}{
		{
			skipTest: false,
			testName: "Success",
			reqPayload: struct {
				data     ExprEnv
				formulas map[string]string
			}{
				data: ExprEnv{
					Price:    100,
					Discount: 25,
					Date:     endDate,
				},
				formulas: map[string]string{
					"discountedPrice": `price - (price * (discount / 100))`,
				},
			},
			assertFunc: func(a *assert.Assertions, result any, err error) {
				t.Logf("This is the error: %v", err)
				t.Logf("This is the result: %v", result)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			if test.skipTest {
				t.Logf("Skipping test '%v'", test.testName)
				t.Skip(test.testName, "Test skipped")
			}
			a := assert.New(t)

			result, err := GenerateHTMLWithFormulaUsecase(test.reqPayload.data, test.reqPayload.formulas)

			test.assertFunc(a, result, err)
		})
	}
}
