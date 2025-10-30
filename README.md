# Resume Builder

A Go-based PDF resume generator with custom fonts and clean, professional layouts.

## Features

- **Clean PDF Generation**: Creates professional resumes using the Maroto v2 library
- **Custom Fonts**: Supports DejaVu Sans with ExtraLight and Bold variants
- **Icon Integration**: Uses PNG icons for contact information and section headers
- **Hyperlinks**: Clickable links for website, LinkedIn, and GitHub
- **Responsive Layout**: Contact information in organized 2x3 grid
- **Professional Styling**: Blue accent colors and clean typography

## Requirements

- Go 1.21+
- DejaVu fonts (included in `/fonts` directory)

## Installation

```bash
git clone https://github.com/lukestogsdill/resume-builder.git
cd resume-builder
go mod tidy
```

## Usage

```bash
go run .
```

This generates `resume.pdf` in the current directory.

## Customization

### Personal Information
Edit the personal details directly in `main.go`:
- Name and contact information
- Professional summary
- Work experience entries
- Education details
- Technical skills

### Styling
- **Colors**: Modify `getPrimaryColor()` function for accent colors
- **Fonts**: DejaVu Sans variants are loaded automatically
- **Icons**: Replace PNG files in `/icons-png` directory
- **Layout**: Adjust spacing and alignment in the layout functions

### Icons
The contact section uses these icons:
- `map-pin.png` - Address
- `phone.png` - Phone number
- `mail.png` - Email
- `link.png` - Website
- `linkedin.png` - LinkedIn profile
- `github.png` - GitHub profile

## Project Structure

```
├── main.go           # Main application logic
├── fonts/            # DejaVu font files
├── icons-png/        # Contact and section icons
├── go.mod           # Go module definition
└── README.md        # This file
```

## Dependencies

- `github.com/johnfercher/maroto/v2` - PDF generation library

## License

MIT License - feel free to use and modify for your own resume needs.