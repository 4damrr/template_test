package template

import (
	"github.com/aymerick/raymond"
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
