package template

import (
	"github.com/aymerick/raymond"
	"github.com/expr-lang/expr"
	"time"
)

var (
	template = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{name}} - CV</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 40px;
            line-height: 1.6;
        }
        h1, h2 {
            color: #333;
        }
        .section {
            margin-bottom: 30px;
        }
        .contact, .skills {
            font-size: 0.95em;
        }
        ul {
            padding-left: 20px;
        }
    </style>
</head>
<body>

    <h1>{{name}}</h1>
    <p class="contact">
        Email: {{email}}<br>
        Phone: {{phone}}<br>
        LinkedIn: linkedin.com/in/{{userName}}
    </p>

    <div class="section">
        <h2>Summary</h2>
        <p>{{summary}}</p>
    </div>

	<div class="section">
        <h2>Skills</h2>
			<ul>
				{{#each skills}}
					<li>{{this}}</li>
  				{{/each}}
			</ul>
	</div>
	
	{{#if experiences}}
		<div class="section">
	        <h2>Experience</h2>
			{{#each experiences}}
	        	<h3>{{this.name}}</h3>
	        	<p><em>{{formatDate this.startDate "2006-01-02"}} – {{formatDate this.endDate "2006-01-02"}}</em></p>
	        	<ul>
					{{#each this.descriptions}}
	        	    	<li>{{this}}</li>
					{{/each}}	
	        	</ul>
			{{/each}}
	    </div>
	{{/if}}
</body>
</html>`

	templateFormula = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Discount Offer</title>
  <style>
    body {
      font-family: 'Segoe UI', sans-serif;
      background-color: #f4f4f4;
      padding: 40px;
    }

    .card {
      max-width: 400px;
      margin: auto;
      background: #ffffff;
      border-radius: 12px;
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
      text-align: center;
      padding: 24px;
    }

    .card h2 {
      color: #e91e63;
      margin-bottom: 10px;
    }

    .price-section {
      margin: 20px 0;
      font-size: 24px;
    }

    .old-price {
      text-decoration: line-through;
      color: #999;
      margin-right: 10px;
    }

    .new-price {
      color: #2e7d32;
      font-weight: bold;
    }

    .discount-code {
      background-color: #e91e63;
      color: #fff;
      font-weight: bold;
      font-size: 20px;
      letter-spacing: 1px;
      padding: 10px 16px;
      margin: 20px 0;
      border-radius: 6px;
      display: inline-block;
    }

    .card-footer {
      font-size: 14px;
      color: #888;
      margin-top: 16px;
    }
  </style>
</head>
<body>

  <div class="card">
    <h2>Special Offer</h2>
    <p>Save big with this limited-time discount!</p>

    <div class="price-section">
      <span class="old-price">${{oldPrice}}</span>
      <span class="new-price">${{discountedPrice}}</span>
    </div>

    <div class="discount-code">SAVE {{discount}}%</div>

    <p class="card-footer">Valid until: {{formatDate date "2006-01-02"}}</p>
  </div>

</body>
</html>`
)

func GenerateCVUsecase(data UserData) (string, error) {
	raymond.RegisterHelper("formatDate", func(date time.Time, format string) string {
		return date.Format(format)
	})

	ctx := map[string]interface{}{
		"name":        data.Name,
		"email":       data.Email,
		"phone":       data.Phone,
		"userName":    data.UserName,
		"summary":     data.Summary,
		"skills":      data.Skills,
		"experiences": data.Experiences,
		//"educations":  data.Educations,
	}

	result := raymond.MustRender(template, ctx)

	return result, nil
}

func ExprFormulaUsecase(formula string, params ExprEnv) (any, error) {
	program, err := expr.Compile(formula, expr.Env(ExprEnv{}))
	if err != nil {
		return "", err
	}

	result, err := expr.Run(program, params)
	if err != nil {
		return "", err
	}

	//res, _ := fmt.Println(result)

	return result, nil
}

func GenerateHTMLWithFormulaUsecase(data ExprEnv, formulas map[string]string) (string, error) {
	raymond.RegisterHelper("formatDate", func(date time.Time, format string) string {
		return date.Format(format)
	})

	formulaResult := make(map[string]any)

	for i, formula := range formulas {
		program, err := expr.Compile(formula, expr.Env(ExprEnv{}))
		if err != nil {
			return "", err
		}

		formulaResult[i], err = expr.Run(program, data)
		if err != nil {
			return "", err
		}
	}

	ctx := map[string]interface{}{
		"oldPrice":        data.Price,
		"discountedPrice": formulaResult["discountedPrice"],
		"discount":        data.Discount,
		"date":            data.Date,
	}

	result := raymond.MustRender(templateFormula, ctx)

	return result, nil
}
