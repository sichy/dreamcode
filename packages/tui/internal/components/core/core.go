package core

import (
	"strings"

	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/sst/opencode/internal/styles"
	"github.com/sst/opencode/internal/theme"
)

// Title creates a title with a gradient slash pattern
func Title(title string, width int) string {
	t := theme.CurrentTheme()
	char := "╱"
	length := lipgloss.Width(title) + 1
	remainingWidth := width - length
	titleStyle := styles.NewStyle().Foreground(t.Primary())
	if remainingWidth > 0 {
		lines := strings.Repeat(char, remainingWidth)
		lines = ApplyForegroundGrad(lines)
		title = titleStyle.Render(title) + " " + lines
	}
	return title
}

// ApplyForegroundGrad applies a gradient effect to the text
func ApplyForegroundGrad(text string) string {
	if text == "" {
		return text
	}
	
	// Simple implementation - create a gradient effect
	t := theme.CurrentTheme()
	primaryStyle := styles.NewStyle().Foreground(t.Primary())
	secondaryStyle := styles.NewStyle().Foreground(t.Secondary())
	mutedStyle := styles.NewStyle().Foreground(t.TextMuted())
	
	length := len(text)
	if length <= 3 {
		return primaryStyle.Render(text)
	}
	
	// Split text into three parts for a simple gradient effect
	third := length / 3
	part1 := text[:third]
	part2 := text[third : third*2]
	part3 := text[third*2:]
	
	// Apply gradient by using different colors
	return primaryStyle.Render(part1) + 
		   mutedStyle.Render(part2) + 
		   secondaryStyle.Render(part3)
}

// Section creates a section divider with a line
func Section(text string, width int) string {
	t := theme.CurrentTheme()
	char := "─"
	length := lipgloss.Width(text) + 1
	remainingWidth := width - length
	lineStyle := styles.NewStyle().Foreground(t.Border())
	if remainingWidth > 0 {
		text = text + " " + lineStyle.Render(strings.Repeat(char, remainingWidth))
	}
	return text
}

// Status creates a status line with icon, title, and description
type StatusOpts struct {
	Icon             string // if empty no icon will be shown
	Title            string
	Description      string
	ExtraContent     string // additional content to append after the description
}

func Status(opts StatusOpts, width int) string {
	t := theme.CurrentTheme()
	icon := opts.Icon
	title := opts.Title
	description := opts.Description
	
	// Use theme colors for styling
	title = styles.NewStyle().Foreground(t.Text()).Render(title)
	
	if description != "" {
		extraContentWidth := lipgloss.Width(opts.ExtraContent)
		if extraContentWidth > 0 {
			extraContentWidth += 1
		}
		description = ansi.Truncate(description, width-lipgloss.Width(icon)-lipgloss.Width(title)-2-extraContentWidth, "…")
	}
	description = styles.NewStyle().Foreground(t.TextMuted()).Render(description)

	content := []string{}
	if icon != "" {
		content = append(content, icon)
	}
	content = append(content, title, description)
	if opts.ExtraContent != "" {
		content = append(content, opts.ExtraContent)
	}

	return strings.Join(content, " ")
}
