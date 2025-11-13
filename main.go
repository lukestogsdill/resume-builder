package main

import (
	"fmt"
	"log"
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

func main() {
	start := time.Now()
	fmt.Println("Starting resume generation...")
	
	m := getMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("resume.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}
	
	duration := time.Since(start)
	fmt.Printf("Resume generated: resume.pdf (took %v)\n", duration)
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
			Style: fontstyle.BoldItalic,
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
			Style: fontstyle.BoldItalic,
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
			Style: fontstyle.BoldItalic,
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

	mrt.AddRows(text.NewRow(15, "Full Stack Developer with 2+ years of experience delivering production websites for live clients. Co-founded web design agency serving real businesses including medical practitioners, achieving 98-100/100 PageSpeed scores on live client sites. Proven expertise in React, Next.js, Svelte, Node.js, and cutting-edge technologies like Convex and TanStack Query.", props.Text{
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
		text.NewCol(8, "Co-founder & Lead Developer - Lobby Media", props.Text{
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

	mrt.AddRows(text.NewRow(6, "• Led development team of 3 engineers delivering high-performance websites for medical practices", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Managed full-stack architecture serving 1000+ monthly visitors with 99.9% uptime", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Optimized client sites achieving 98-100/100 PageSpeed scores and <2s load times", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "Tech: Svelte5, Next.js, TypeScript, SveltiaCMS, Cloudflare", props.Text{
		Size:  9,
		Align: align.Left,
		Style: fontstyle.Italic,
	}))

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

	mrt.AddRows(text.NewRow(6, "• Architected scalable IoT monitoring system handling 50+ concurrent device connections", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Built real-time motor monitoring system with <500ms alert response times", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Deployed containerized microservices with automated CI/CD pipeline", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "Tech: Next.js, Golang, MongoDB, Docker, Hetzner Cloud", props.Text{
		Size:  9,
		Align: align.Left,
		Style: fontstyle.Italic,
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
		text.NewCol(12, "Sibyl - Daily Morality Quiz", props.Text{
			Size:  11,
			Style: fontstyle.BoldItalic,
			Align: align.Left,
			Color: &props.Color{Red: 0, Green: 0, Blue: 255},
			Hyperlink: &[]string{"https://sibyl.it.com"}[0],
		}),
	)


	mrt.AddRows(text.NewRow(6, "• Developed analytics platform processing 10K+ user responses with real-time insights", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Implemented serverless architecture with reactive data subscriptions and live updates", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Built AI analysis engine generating personalized moral compass insights from aggregate data", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "Tech: React, TypeScript, Convex, TanStack Query, ORPC", props.Text{
		Size:  9,
		Align: align.Left,
		Style: fontstyle.Italic,
	}))

	// Project 2 
	mrt.AddRows(text.NewRow(6, "Professional Website Portfolio", props.Text{
			Size:  11,
			Style: fontstyle.BoldItalic,
			Align: align.Left,
		}))
	mrt.AddRow(6,
		text.NewCol(1, "", props.Text{}),
		text.NewCol(3, "Vascular & Wound Care", props.Text{
			Size:  9,
			Align: align.Left,
			Color: &props.Color{Red: 0, Green: 0, Blue: 255},
			Style: fontstyle.BoldItalic,
			Hyperlink: &[]string{"https://vascularandwound-com.pages.dev"}[0],
		}),
		text.NewCol(3, "Landscape Services", props.Text{
			Size:  9,
			Align: align.Left,
			Color: &props.Color{Red: 0, Green: 0, Blue: 255},
			Style: fontstyle.BoldItalic,
			Hyperlink: &[]string{"https://landscape.lobby.media"}[0],
		}),
		text.NewCol(3, "Legal Practice", props.Text{
			Size:  9,
			Align: align.Left,
			Color: &props.Color{Red: 0, Green: 0, Blue: 255},
			Style: fontstyle.BoldItalic,
			Hyperlink: &[]string{"https://law.lobby.media"}[0],
		}),
		text.NewCol(2, "", props.Text{}),
	)

	// Project 2 - AI Vtuber
	mrt.AddRow(8,
		text.NewCol(12, "AI Vtuber Content Creator", props.Text{
			Size:  11,
			Style: fontstyle.Bold,
			Align: align.Left,
		}),
	)

	mrt.AddRows(text.NewRow(6, "• Created AI-powered Vtuber that watches videos and generates authentic real-time reactions", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Built offline rendering pipeline processing 4K video with synchronized Live2D animations", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "• Integrated advanced speech synthesis and emotion recognition for dynamic character responses", props.Text{
		Size:  9,
		Align: align.Left,
	}))

	mrt.AddRows(text.NewRow(6, "Tech: React, Python, C++, Live2D Cubism SDK, WhisperX, OpenAI API", props.Text{
		Size:  9,
		Align: align.Left,
		Style: fontstyle.Italic,
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

