package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

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

type Link struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

type Achievement struct {
	Text     string `json:"text"`
	Overflow bool   `json:"overflow,omitempty"`
}

type Contact struct {
	Location string `json:"location"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Website  Link   `json:"website"`
	LinkedIn Link   `json:"linkedin"`
	GitHub   Link   `json:"github"`
}

type Skill struct {
	Category string `json:"category"`
	Items    string `json:"items"`
}

type Experience struct {
	Company      string        `json:"company"`
	Title        string        `json:"title"`
	URL          string        `json:"url"`
	Dates        string        `json:"dates"`
	Achievements []Achievement `json:"achievements"`
	Tech         string        `json:"tech"`
}

type Project struct {
	Name         string        `json:"name"`
	URL          string        `json:"url"`
	Achievements []Achievement `json:"achievements"`
	Tech         string        `json:"tech"`
}

type Education struct {
	School string `json:"school"`
	Degree string `json:"degree"`
	URL    string `json:"url"`
	Date   string `json:"date"`
}

type Resume struct {
	Name       string       `json:"name"`
	Title      string       `json:"title"`
	Contact    Contact      `json:"contact"`
	Summary    string       `json:"summary"`
	Skills     []Skill      `json:"skills"`
	Experience []Experience `json:"experience"`
	Projects   []Project    `json:"projects"`
	Education  []Education  `json:"education"`
}

func loadResume(filename string) (*Resume, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read resume file: %w", err)
	}

	var resume Resume
	if err := json.Unmarshal(data, &resume); err != nil {
		return nil, fmt.Errorf("failed to parse resume JSON: %w", err)
	}

	return &resume, nil
}

func main() {
	start := time.Now()
	fmt.Println("Starting resume generation...")

	resume, err := loadResume("resume.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	m := getMaroto(resume)
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("resume.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	duration := time.Since(start)
	fmt.Printf("Resume generated successfully in %v\n", duration)
}

func getMaroto(resume *Resume) core.Maroto {
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
			WithTopMargin(10).
			WithRightMargin(10).
			Build()
		return maroto.New(cfg)
	}

	cfg := config.NewBuilder().
		WithCustomFonts(customFonts).
		WithDefaultFont(&props.Font{Family: "dejavu"}).
		WithPageNumber().
		WithLeftMargin(10).
		WithTopMargin(10).
		WithRightMargin(10).
		Build()

	mrt := maroto.New(cfg)

	// Header with name and title
	mrt.AddRows(text.NewRow(15, resume.Name, props.Text{
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
		text.NewCol(3, resume.Contact.Location, props.Text{
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
		text.NewCol(3, resume.Contact.Phone, props.Text{
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
		text.NewCol(3, resume.Contact.Email, props.Text{
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
		text.NewCol(3, resume.Contact.Website.Text, props.Text{
			Size:  10,
			Align: align.Left,
			Top:   1,
			Hyperlink: &[]string{resume.Contact.Website.URL}[0],
			Color: &props.Color{Red: 0, Green: 0, Blue: 255},
			Style: fontstyle.BoldItalic,
		}),
		marotoimg.NewFromFileCol(1, "icons-png/linkedin.png", props.Rect{
			Center:  false,
			Percent: 60,
			Top: 1,
			Left: 10,
		}),
		text.NewCol(3, resume.Contact.LinkedIn.Text, props.Text{
			Size:  10,
			Align: align.Left,
			Top:   1,
			Hyperlink: &[]string{resume.Contact.LinkedIn.URL}[0],
			Color: &props.Color{Red: 0, Green: 0, Blue: 255},
			Style: fontstyle.BoldItalic,
		}),
		marotoimg.NewFromFileCol(1, "icons-png/github.png", props.Rect{
			Center:  false,
			Percent: 60,
			Top: 1,
			Left: 10,
		}),
		text.NewCol(3, resume.Contact.GitHub.Text, props.Text{
			Size:  10,
			Align: align.Left,
			Top:   1,
			Hyperlink: &[]string{resume.Contact.GitHub.URL}[0],
			Color: &props.Color{Red: 0, Green: 0, Blue: 255},
			Style: fontstyle.BoldItalic,
		}),
	)

	// Separator line
	mrt.AddRow(.25, col.New(1)).WithStyle(&props.Cell{BackgroundColor: getPrimaryColor()})

	// Full Stack Developer title
	mrt.AddRows(text.NewRow(10, resume.Title, props.Text{
		Top:    3,
		Style:  fontstyle.Bold,
		Size:   14,
		Align:  align.Left,
		Family: "dejavu",
	}))

	mrt.AddRows(text.NewRow(15, resume.Summary, props.Text{
		Top:    2,
		Size:   10,
		Align:  align.Left,
		Family: "dejavu-light",
	}))

	// Add spacing after summary
	mrt.AddRow(5, col.New(1))

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

	// Add each skill category dynamically
	for _, skill := range resume.Skills {
		mrt.AddRow(5,
			text.NewCol(3, skill.Category + ":", props.Text{
				Size:  10,
				Style: fontstyle.Bold,
				Align: align.Left,
			}),
			text.NewCol(9, skill.Items, props.Text{
				Size:  10,
				Align: align.Left,
			}),
		)
	}

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

	// underline
	mrt.AddRow(.25, col.New(1)).WithStyle(&props.Cell{BackgroundColor: getPrimaryColor()})

	for _, exp := range resume.Experience {
		jobTitle := exp.Title + " - " + exp.Company
		if exp.URL != "" {
			mrt.AddRow(8,
				text.NewCol(8, jobTitle, props.Text{
					Size:      11,
					Style:     fontstyle.BoldItalic,
					Align:     align.Left,
					Color:     &props.Color{Red: 0, Green: 0, Blue: 255},
					Hyperlink: &[]string{exp.URL}[0],
				}),
				text.NewCol(4, exp.Dates, props.Text{
					Size:  10,
					Align: align.Right,
				}),
			)
		} else {
			mrt.AddRow(8,
				text.NewCol(8, jobTitle, props.Text{
					Size:  11,
					Style: fontstyle.Bold,
					Align: align.Left,
				}),
				text.NewCol(4, exp.Dates, props.Text{
					Size:  10,
					Align: align.Right,
				}),
			)
		}

		// Achievements
		for _, achievement := range exp.Achievements {
			rowHeight := 5
			if achievement.Overflow {
				rowHeight = 7
			}
			mrt.AddRows(text.NewRow(float64(rowHeight), "• "+achievement.Text, props.Text{
				Size:  9,
				Align: align.Left,
			}))
		}

		// Tech stack
		mrt.AddRows(text.NewRow(5, "Tech: "+exp.Tech, props.Text{
			Size:  9,
			Align: align.Left,
			Style: fontstyle.Italic,
		}))
	}

	// Projects
	mrt.AddRow(10,
		marotoimg.NewFromFileCol(0, "icons-png/layers.png", props.Rect{
			Center:  false,
			Percent: 80,
			Left: 0,
		}),
		text.NewCol(12, "PROJECTS", props.Text{
			Top:   3,
			Style: fontstyle.Bold,
			Size:  12,
			Align: align.Left,
			Color: getPrimaryColor(),
			Left:  10,
		}),
	)

	// Projects underline
	mrt.AddRow(.25, col.New(1)).WithStyle(&props.Cell{BackgroundColor: getPrimaryColor()})

	// Add each project dynamically
	for _, proj := range resume.Projects {
		// Project name with link
		if proj.URL != "" {
			mrt.AddRow(8,
				text.NewCol(12, proj.Name, props.Text{
					Size:      11,
					Style:     fontstyle.BoldItalic,
					Align:     align.Left,
					Color:     &props.Color{Red: 0, Green: 0, Blue: 255},
					Hyperlink: &[]string{proj.URL}[0],
				}),
			)
		} else {
			mrt.AddRow(8,
				text.NewCol(12, proj.Name, props.Text{
					Size:  11,
					Style: fontstyle.Bold,
					Align: align.Left,
				}),
			)
		}

		// Achievements
		for _, achievement := range proj.Achievements {
			rowHeight := 6
			if achievement.Overflow {
				rowHeight = 8
			}
			mrt.AddRows(text.NewRow(float64(rowHeight), "• "+achievement.Text, props.Text{
				Size:  9,
				Align: align.Left,
			}))
		}

		// Tech stack
		mrt.AddRows(text.NewRow(5, "Tech: "+proj.Tech, props.Text{
			Size:  9,
			Align: align.Left,
			Style: fontstyle.Italic,
		}))
	}

	

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

	// Add each education entry dynamically
	for _, edu := range resume.Education {
		if edu.URL != "" {
			mrt.AddRow(8,
				text.NewCol(8, edu.School, props.Text{
					Size:      11,
					Style:     fontstyle.BoldItalic,
					Align:     align.Left,
					Color:     &props.Color{Red: 0, Green: 0, Blue: 255},
					Hyperlink: &[]string{edu.URL}[0],
				}),
				text.NewCol(4, edu.Date, props.Text{
					Size:  10,
					Align: align.Right,
				}),
			)
		} else {
			mrt.AddRow(8,
				text.NewCol(8, edu.School, props.Text{
					Size:  11,
					Style: fontstyle.Bold,
					Align: align.Left,
				}),
				text.NewCol(4, edu.Date, props.Text{
					Size:  10,
					Align: align.Right,
				}),
			)
		}
	}

	return mrt
}

func getPrimaryColor() *props.Color {
	return &props.Color{
		Red:   70,
		Green: 130,
		Blue:  180,
	}
}

