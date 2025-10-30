package main

import (
	"fmt"
	"log"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	marotoimg "github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/repository"
)

func main() {
	m := getMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("resume.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}
	
	fmt.Println("Resume generated: resume.pdf")
}

func getMaroto() core.Maroto {
	// Load custom fonts
	customFonts, err := repository.New().
		AddUTF8Font("dejavu", fontstyle.Normal, "fonts/DejaVuSans.ttf").
		AddUTF8Font("dejavu", fontstyle.Bold, "fonts/DejaVuSans-Bold.ttf").
		AddUTF8Font("dejavu", fontstyle.Italic, "fonts/DejaVuSans-Oblique.ttf").
		AddUTF8Font("dejavu", fontstyle.BoldItalic, "fonts/DejaVuSans-BoldOblique.ttf").
		AddUTF8Font("dejavu-light", fontstyle.Normal, "fonts/DejaVuSans-ExtraLight.ttf").
		Load()
	if err != nil {
		log.Printf("Warning: Failed to load custom fonts: %v", err)
		// Fallback to default config
		cfg := config.NewBuilder().
			WithPageNumber().
			WithLeftMargin(10).
			WithTopMargin(15).
			WithRightMargin(10).
			Build()
		return maroto.New(cfg)
	}

	cfg := config.NewBuilder().
		WithCustomFonts(customFonts).
		WithDefaultFont(&props.Font{Family: "dejavu"}).
		WithPageNumber().
		WithLeftMargin(10).
		WithTopMargin(15).
		WithRightMargin(10).
		Build()

	mrt := maroto.New(cfg)

	// Header with name and title
	mrt.AddRows(text.NewRow(15, "John Doe", props.Text{
		Top:   3,
		Style: fontstyle.Bold,
		Size:  20,
		Align: align.Left,
	}))


	// Contact info in 2x3 grid with icons
	mrt.AddRow(8,
		marotoimg.NewFromFileCol(1, "icons-png/map-pin.png", props.Rect{
			Center:  false,
			Percent: 60,
			Top: 1,
			Left: 10,
		}),
		text.NewCol(3, "123 Main Street, City, ST 12345", props.Text{
			Size:  10,
			Align: align.Left,
			Top:   1,
		}),
		marotoimg.NewFromFileCol(1, "icons-png/phone.png", props.Rect{
			Center:  false,
			Percent: 60,
			Top: 1,
			Left: 10,
		}),
		text.NewCol(3, "+1 (555) 123-4567", props.Text{
			Size:  10,
			Align: align.Left,
			Top:   1,
		}),
		marotoimg.NewFromFileCol(1, "icons-png/mail.png", props.Rect{
			Center:  false,
			Percent: 60,
			Top: 1,
			Left: 10,
		}),
		text.NewCol(3, "john.doe@email.com", props.Text{
			Size:  10,
			Align: align.Left,
			Top:   1,
		}),
	)
	mrt.AddRow(8,
		marotoimg.NewFromFileCol(1, "icons-png/link.png", props.Rect{
			Center:  false,
			Percent: 60,
			Top: 1,
			Left: 10,
		}),
		text.NewCol(3, "lustogs.com", props.Text{
			Size:  10,
			Align: align.Left,
			Top:   1,
			Hyperlink: &[]string{"https://lustogs.com"}[0],
			Color: &props.Color{Red: 0, Green: 0, Blue: 255},
		}),
		marotoimg.NewFromFileCol(1, "icons-png/linkedin.png", props.Rect{
			Center:  false,
			Percent: 60,
			Top: 1,
			Left: 10,
		}),
		text.NewCol(3, "luke-stogsdill", props.Text{
			Size:  10,
			Align: align.Left,
			Top:   1,
			Hyperlink: &[]string{"https://linkedin.com/in/luke-stogsdill"}[0],
			Color: &props.Color{Red: 0, Green: 0, Blue: 255},
		}),
		marotoimg.NewFromFileCol(1, "icons-png/github.png", props.Rect{
			Center:  false,
			Percent: 60,
			Top: 1,
			Left: 10,
		}),
		text.NewCol(3, "lukestogsdill", props.Text{
			Size:  10,
			Align: align.Left,
			Top:   1,
			Hyperlink: &[]string{"https://github.com/lukestogsdill"}[0],
			Color: &props.Color{Red: 0, Green: 0, Blue: 255},
		}),
	)

	// Separator line
	mrt.AddRow(.25, col.New(1)).WithStyle(&props.Cell{BackgroundColor: getPrimaryColor()})

	// Software Engineer title
	mrt.AddRows(text.NewRow(10, "Software Engineer", props.Text{
		Top:    3,
		Style:  fontstyle.Bold,
		Size:   14,
		Align:  align.Left,
		Family: "dejavu",
	}))

	mrt.AddRows(text.NewRow(15, "Experienced software engineer with 5+ years developing web applications and APIs. Proficient in Go, Python, and JavaScript with expertise in cloud architecture and DevOps practices.", props.Text{
		Top:    2,
		Size:   10,
		Align:  align.Left,
		Family: "dejavu-light",
	}))

	// Technical Skills
	mrt.AddRow(10,
		marotoimg.NewFromFileCol(0, "icons-png/box.png", props.Rect{
			Center:  false,
			Percent: 80,
			Left: 0,
		}),
		text.NewCol(12, "TECHNICAL SKILLS", props.Text{
			Top:   3,
			Style: fontstyle.Bold,
			Size:  12,
			Align: align.Left,
			Color: getPrimaryColor(),
			Left:  10,
		}),
	)

	// Technical Skills underline
	mrt.AddRow(.25, col.New(1)).WithStyle(&props.Cell{BackgroundColor: getPrimaryColor()})

	mrt.AddRow(8,
		text.NewCol(3, "Languages:", props.Text{
			Size:  10,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
		text.NewCol(9, "Go, Python, JavaScript, TypeScript, Java", props.Text{
			Size:  10,
			Align: align.Left,
		}),
	)

	mrt.AddRow(8,
		text.NewCol(3, "Frameworks:", props.Text{
			Size:  10,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
		text.NewCol(9, "React, Node.js, Express, Gin, Django", props.Text{
			Size:  10,
			Align: align.Left,
		}),
	)

	mrt.AddRow(8,
		text.NewCol(3, "Databases:", props.Text{
			Size:  10,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
		text.NewCol(9, "PostgreSQL, MongoDB, Redis", props.Text{
			Size:  10,
			Align: align.Left,
		}),
	)

	mrt.AddRow(8,
		text.NewCol(3, "Tools:", props.Text{
			Size:  10,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
		text.NewCol(9, "Docker, Kubernetes, AWS, Git, CI/CD", props.Text{
			Size:  10,
			Align: align.Left,
		}),
	)

	// Work Experience
	mrt.AddRow(10,
		marotoimg.NewFromFileCol(0, "icons-png/building-2.png", props.Rect{
			Center:  false,
			Percent: 80,
			Left: 0,
		}),
		text.NewCol(12, "WORK EXPERIENCE", props.Text{
			Top:   3,
			Style: fontstyle.Bold,
			Size:  12,
			Align: align.Left,
			Color: getPrimaryColor(),
			Left:  10,
		}),
	)

	// Work Experience underline
	mrt.AddRow(.25, col.New(1)).WithStyle(&props.Cell{BackgroundColor: getPrimaryColor()})

	// Job 1
	mrt.AddRow(8,
		text.NewCol(8, "Senior Software Engineer", props.Text{
			Size:  11,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
		text.NewCol(4, "2021 - Present", props.Text{
			Size:  10,
			Align: align.Right,
		}),
	)

	mrt.AddRows(text.NewRow(6, "TechCorp Inc. - San Francisco, CA", props.Text{
		Size:  10,
		Style: fontstyle.Italic,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Led development of microservices architecture serving 1M+ users", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Improved API performance by 40% through optimization and caching", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Mentored 3 junior developers and conducted code reviews", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	// Job 2
	mrt.AddRow(8,
		text.NewCol(8, "Software Engineer", props.Text{
			Size:  11,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
		text.NewCol(4, "2019 - 2021", props.Text{
			Size:  10,
			Align: align.Right,
		}),
	)

	mrt.AddRows(text.NewRow(6, "StartupXYZ - Austin, TX", props.Text{
		Size:  10,
		Style: fontstyle.Italic,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Built full-stack web applications using React and Node.js", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Implemented automated testing reducing bugs by 60%", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Collaborated with design team to create responsive user interfaces", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	// Education
	mrt.AddRow(10,
		marotoimg.NewFromFileCol(0, "icons-png/graduation-cap.png", props.Rect{
			Center:  false,
			Percent: 80,
			Left: 0,
		}),
		text.NewCol(12, "EDUCATION", props.Text{
			Top:   3,
			Style: fontstyle.Bold,
			Size:  12,
			Align: align.Left,
			Color: getPrimaryColor(),
			Left:  10,
		}),
	)

	// Education underline
	mrt.AddRow(.25, col.New(1)).WithStyle(&props.Cell{BackgroundColor: getPrimaryColor()})

	mrt.AddRow(8,
		text.NewCol(8, "Bachelor of Science in Computer Science", props.Text{
			Size:  11,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
		text.NewCol(4, "2015 - 2019", props.Text{
			Size:  10,
			Align: align.Right,
		}),
	)

	mrt.AddRows(text.NewRow(6, "University of California, Berkeley", props.Text{
		Size:  10,
		Style: fontstyle.Italic,
		Align: align.Left,
	}))

	return mrt
}

func getPrimaryColor() *props.Color {
	return &props.Color{
		Red:   70,
		Green: 130,
		Blue:  180,
	}
}

