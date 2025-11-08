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
	mrt.AddRows(text.NewRow(15, "Luke Stogsdill", props.Text{
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
		text.NewCol(3, "Houston, TX 77064", props.Text{
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
		text.NewCol(3, "(832) 392-2613", props.Text{
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
		text.NewCol(3, "lukestogsdill@gmail.com", props.Text{
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
			Style: fontstyle.Italic,
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
			Style: fontstyle.Italic,
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
			Style: fontstyle.Italic,
		}),
	)

	// Separator line
	mrt.AddRow(.25, col.New(1)).WithStyle(&props.Cell{BackgroundColor: getPrimaryColor()})

	// Full Stack Developer title
	mrt.AddRows(text.NewRow(10, "Full Stack Developer", props.Text{
		Top:    3,
		Style:  fontstyle.Bold,
		Size:   14,
		Align:  align.Left,
		Family: "dejavu",
	}))

	mrt.AddRows(text.NewRow(15, "Full Stack Developer with 2+ years of experience building scalable web applications and modern software solutions. Proven expertise in React, Next.js, Svelte, Node.js, and cutting-edge technologies like Convex and TanStack Query. Co-founded a successful web design agency, delivering high-performance websites with 98-100/100 PageSpeed scores.", props.Text{
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

	mrt.AddRow(5,
		text.NewCol(3, "Frontend:", props.Text{
			Size:  10,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
		text.NewCol(9, "React, Next.js, Svelte5, TanStack Query, TypeScript", props.Text{
			Size:  10,
			Align: align.Left,
		}),
	)

	mrt.AddRow(5,
		text.NewCol(3, "Backend:", props.Text{
			Size:  10,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
		text.NewCol(9, "Node.js, Golang, Python, C++, ORPC, Convex", props.Text{
			Size:  10,
			Align: align.Left,
		}),
	)

	mrt.AddRow(5,
		text.NewCol(3, "Databases:", props.Text{
			Size:  10,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
		text.NewCol(9, "MongoDB, PostgreSQL, MySQL, SveltiaCMS", props.Text{
			Size:  10,
			Align: align.Left,
		}),
	)

	mrt.AddRow(5,
		text.NewCol(3, "DevOps:", props.Text{
			Size:  10,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
		text.NewCol(9, "Docker, CI/CD, Cloudflare, Hetzner, Git", props.Text{
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
		text.NewCol(8, "Co-founder - Lobby Media Web Design", props.Text{
			Size:  11,
			Style: fontstyle.BoldItalic,
			Align: align.Left,
			Color: &props.Color{Red: 0, Green: 0, Blue: 255},
			Hyperlink: &[]string{"https://lobby.media"}[0],
		}),
		text.NewCol(4, "June 2024 - Current", props.Text{
			Size:  10,
			Align: align.Right,
		}),
	)

	mrt.AddRows(text.NewRow(6, "• Co-founded web design agency delivering high-performance websites for small businesses", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Led full-cycle development using Svelte5, Next.js, and headless CMS (SveltiaCMS)", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Achieved 98-100/100 Google PageSpeed scores through optimized development practices", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	// mrt.AddRows(text.NewRow(6, "• Client portfolio: Vascular and Wound, Landscape, Law", props.Text{
	// 	Size:  9,
	// 	Align: align.Left,
	// 	Color: &props.Color{Red: 0, Green: 0, Blue: 255},
	// 	Style: fontstyle.Italic,
	// 	Hyperlink: &[]string{"https://vascularandwound-com.pages.dev"}[0],
	// }))

	// mrt.AddRow(6,
	// 	text.NewCol(2, "", props.Text{}),
	// 	text.NewCol(3, "Vascular and Wound", props.Text{
	// 		Size:  9,
	// 		Align: align.Left,
	// 		Color: &props.Color{Red: 0, Green: 0, Blue: 255},
	// 		Style: fontstyle.Italic,
	// 		Hyperlink: &[]string{"https://vascularandwound-com.pages.dev"}[0],
	// 	}),
	// 	text.NewCol(2, "Landscape", props.Text{
	// 		Size:  9,
	// 		Align: align.Left,
	// 		Color: &props.Color{Red: 0, Green: 0, Blue: 255},
	// 		Style: fontstyle.Italic,
	// 		Hyperlink: &[]string{"https://landscape.lobby.media"}[0],
	// 	}),
	// 	text.NewCol(2, "Law", props.Text{
	// 		Size:  9,
	// 		Align: align.Left,
	// 		Color: &props.Color{Red: 0, Green: 0, Blue: 255},
	// 		Style: fontstyle.Italic,
	// 		Hyperlink: &[]string{"https://legal.lobby.media"}[0],
	// 	}),
	// 	text.NewCol(3, "", props.Text{}),
	// )

	// Job 2
	mrt.AddRow(8,
		text.NewCol(8, "Freelance Software Consultant", props.Text{
			Size:  11,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
		text.NewCol(4, "July 2023 - June 2024", props.Text{
			Size:  10,
			Align: align.Right,
		}),
	)

	mrt.AddRows(text.NewRow(6, "IoT Alert Management System (TruVolt)", props.Text{
		Size:  10,
		Style: fontstyle.Italic,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Designed full-stack IoT solution with Next.js frontend and Golang microservices", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Implemented MongoDB database with custom CI/CD pipeline using Docker on Hetzner", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Built scalable architecture for real-time IoT device monitoring and alert distribution", props.Text{
		Size:  9,
		Align: align.Left,
	}))

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

	// Project 1 - Sibyl
	mrt.AddRow(8,
		text.NewCol(8, "Sibyl - Daily Morality Quiz", props.Text{
			Size:  11,
			Style: fontstyle.BoldItalic,
			Align: align.Left,
			Color: &props.Color{Red: 0, Green: 0, Blue: 255},
			Hyperlink: &[]string{"https://sibyl.it.com"}[0],
		}),
		text.NewCol(4, "React, Convex, ORPC", props.Text{
			Size:  10,
			Align: align.Right,
			Style: fontstyle.Italic,
		}),
	)


	mrt.AddRows(text.NewRow(6, "• Built full-stack web app with React and TanStack Query for real-time data sync", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Leveraged Convex for serverless backend with automatic API generation", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Implemented AI-powered analysis engine that processes aggregate user data for personalized insights", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	// Project 2 - AI Content Creation
	mrt.AddRow(8,
		text.NewCol(8, "AI Content Creation Suite", props.Text{
			Size:  11,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
		text.NewCol(4, "React, Python, C++", props.Text{
			Size:  10,
			Align: align.Right,
			Style: fontstyle.Italic,
		}),
	)

	mrt.AddRows(text.NewRow(6, "• Engineered React web client for AI-assisted text generation and video processing", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Built Python backend API for advanced audio/video processing and transcription", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Developed C++ renderer with Live2D Cubism SDK for high-quality animations", props.Text{
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
		text.NewCol(8, "Coding Temple - Software Engineering", props.Text{
			Size:  11,
			Style: fontstyle.BoldItalic,
			Align: align.Left,
			Color: &props.Color{Red: 0, Green: 0, Blue: 255},
			Hyperlink: &[]string{"https://www.credly.com/badges/2cb16330-a5d6-41b4-8cda-a1e177287f35/public_url"}[0],
		}),
		text.NewCol(4, "April 2023", props.Text{
			Size:  10,
			Align: align.Right,
		}),
	)

	return mrt
}

func getPrimaryColor() *props.Color {
	return &props.Color{
		Red:   70,
		Green: 130,
		Blue:  180,
	}
}

