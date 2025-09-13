package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/lipgloss/v2/compat"
	"github.com/sst/opencode/internal/styles"
	"github.com/sst/opencode/internal/theme"
)

// Icon constants for better visual appeal
const (
	// Status icons
	IconSuccess   = "✓"
	IconError     = "✗"
	IconWarning   = "⚠"
	IconInfo      = "ℹ"
	IconPending   = "⋯"
	IconRunning   = "◉"
	IconCompleted = "◆"
	
	// Action icons
	IconSend      = "➤"
	IconReceive   = "◁"
	IconEdit      = "✎"
	IconRead      = "👁"
	IconWrite     = "✍"
	IconExecute   = "⚡"
	IconFetch     = "⬇"
	IconUpload    = "⬆"
	IconCopy      = "⎘"
	IconPaste     = "📋"
	
	// Navigation icons
	IconArrowUp    = "↑"
	IconArrowDown  = "↓"
	IconArrowLeft  = "←"
	IconArrowRight = "→"
	IconChevronUp   = "⌃"
	IconChevronDown = "⌄"
	IconChevronLeft = "‹"
	IconChevronRight = "›"
	
	// UI elements
	IconDot       = "•"
	IconCircle    = "○"
	IconSquare    = "□"
	IconDiamond   = "◇"
	IconStar      = "★"
	IconHeart     = "♥"
	IconBullet    = "▸"
	IconTri       = "▶"
	
	// Separators
	SeparatorDot   = " • "
	SeparatorPipe  = " | "
	SeparatorArrow = " → "
	SeparatorDash  = " — "
	
	// Box drawing characters for better borders
	BoxTopLeft     = "╭"
	BoxTopRight    = "╮"
	BoxBottomLeft  = "╰"
	BoxBottomRight = "╯"
	BoxHorizontal  = "─"
	BoxVertical    = "│"
	BoxCross       = "┼"
	BoxTeeUp       = "┴"
	BoxTeeDown     = "┬"
	BoxTeeLeft     = "┤"
	BoxTeeRight    = "├"
	
	// Double box drawing
	BoxDoubleTopLeft     = "╔"
	BoxDoubleTopRight    = "╗"
	BoxDoubleBottomLeft  = "╚"
	BoxDoubleBottomRight = "╝"
	BoxDoubleHorizontal  = "═"
	BoxDoubleVertical    = "║"
)

// GetToolIcon returns an appropriate icon for a tool name
func GetToolIcon(toolName string) string {
	switch toolName {
	case "bash", "shell":
		return IconExecute
	case "read":
		return IconRead
	case "write":
		return IconWrite
	case "edit":
		return IconEdit
	case "webfetch", "fetch":
		return IconFetch
	case "task", "todo", "todowrite":
		return IconSquare
	default:
		return IconDiamond
	}
}

// GetStatusIcon returns an appropriate icon for a status
func GetStatusIcon(status string) string {
	switch status {
	case "success", "completed", "done":
		return IconSuccess
	case "error", "failed":
		return IconError
	case "warning", "warn":
		return IconWarning
	case "info":
		return IconInfo
	case "pending", "waiting":
		return IconPending
	case "running", "in_progress":
		return IconRunning
	default:
		return IconCircle
	}
}

// RenderProgressBar creates a visual progress bar
func RenderProgressBar(current, total int, width int) string {
	if width < 10 || total == 0 {
		return ""
	}
	
	percentage := float64(current) / float64(total)
	filled := int(percentage * float64(width-2))
	
	t := theme.CurrentTheme()
	
	bar := "["
	for i := 0; i < width-2; i++ {
		if i < filled {
			bar += styles.NewStyle().Foreground(t.Primary()).Render("█")
		} else {
			bar += styles.NewStyle().Foreground(t.BorderSubtle()).Render("░")
		}
	}
	bar += "]"
	
	return bar
}

// RenderSpinner creates an animated spinner character based on time
func RenderSpinner() string {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	index := (time.Now().UnixMilli() / 100) % int64(len(frames))
	return frames[index]
}

// RenderDivider creates a stylized divider line
func RenderDivider(width int, style string) string {
	if width <= 0 {
		return ""
	}
	
	t := theme.CurrentTheme()
	dividerStyle := styles.NewStyle().Foreground(t.BorderSubtle())
	
	switch style {
	case "double":
		return dividerStyle.Render(strings.Repeat("═", width))
	case "dotted":
		return dividerStyle.Render(strings.Repeat("·", width))
	case "dashed":
		pattern := "─ "
		count := width / len(pattern)
		return dividerStyle.Render(strings.Repeat(pattern, count))
	case "wave":
		pattern := "～"
		return dividerStyle.Render(strings.Repeat(pattern, width))
	default:
		return dividerStyle.Render(strings.Repeat("─", width))
	}
}

// RenderBadge creates a colored badge with text
func RenderBadge(text string, color compat.AdaptiveColor) string {
	return styles.NewStyle().
		Background(color).
		Foreground(compat.AdaptiveColor{Light: lipgloss.Color("#000000"), Dark: lipgloss.Color("#000000")}).
		Padding(0, 1).
		Bold(true).
		Render(text)
}

// RenderTag creates a tag-like element
func RenderTag(text string) string {
	t := theme.CurrentTheme()
	return styles.NewStyle().
		Foreground(t.Primary()).
		Background(t.BackgroundElement()).
		Padding(0, 1).
		Render(text)
}

// RenderHighlight wraps text with highlighting
func RenderHighlight(text string) string {
	t := theme.CurrentTheme()
	return styles.NewStyle().
		Background(t.BackgroundElement()).
		Foreground(t.Accent()).
		Bold(true).
		Render(text)
}

// RenderGradientText simulates gradient text using color transitions
func RenderGradientText(text string, startColor, endColor string) string {
	// This is a simplified version - true gradients aren't possible in terminals
	// but we can simulate with alternating colors
	runes := []rune(text)
	result := ""
	
	for i, r := range runes {
		if i%2 == 0 {
			result += styles.NewStyle().Foreground(compat.AdaptiveColor{Light: lipgloss.Color(startColor), Dark: lipgloss.Color(startColor)}).Render(string(r))
		} else {
			result += styles.NewStyle().Foreground(compat.AdaptiveColor{Light: lipgloss.Color(endColor), Dark: lipgloss.Color(endColor)}).Render(string(r))
		}
	}
	
	return result
}

// RenderBoxTitle creates a box with a title
func RenderBoxTitle(title, content string, width int) string {
	t := theme.CurrentTheme()
	
	titleStyle := styles.NewStyle().
		Foreground(t.Primary()).
		Bold(true)
	
	borderStyle := styles.NewStyle().
		Foreground(t.BorderActive())
	
	// Create top border with title
	titleLen := len(title)
	if titleLen > width-4 {
		title = title[:width-4]
		titleLen = width - 4
	}
	
	topBorder := borderStyle.Render(BoxTopLeft + BoxHorizontal)
	topBorder += titleStyle.Render(" " + title + " ")
	remaining := width - titleLen - 4
	if remaining > 0 {
		topBorder += borderStyle.Render(strings.Repeat(BoxHorizontal, remaining))
	}
	topBorder += borderStyle.Render(BoxTopRight)
	
	// Process content lines
	lines := strings.Split(content, "\n")
	var boxedLines []string
	boxedLines = append(boxedLines, topBorder)
	
	contentStyle := styles.NewStyle().
		Foreground(t.Text())
	
	for _, line := range lines {
		if len(line) > width-4 {
			line = line[:width-4]
		}
		padding := width - len(line) - 4
		boxedLine := borderStyle.Render(BoxVertical) + " "
		boxedLine += contentStyle.Render(line)
		if padding > 0 {
			boxedLine += strings.Repeat(" ", padding)
		}
		boxedLine += " " + borderStyle.Render(BoxVertical)
		boxedLines = append(boxedLines, boxedLine)
	}
	
	// Bottom border
	bottomBorder := borderStyle.Render(BoxBottomLeft)
	bottomBorder += borderStyle.Render(strings.Repeat(BoxHorizontal, width-2))
	bottomBorder += borderStyle.Render(BoxBottomRight)
	boxedLines = append(boxedLines, bottomBorder)
	
	return strings.Join(boxedLines, "\n")
}

// RenderPulse creates a pulsing effect for text (simulated)
func RenderPulse(text string, intensity float64) string {
	t := theme.CurrentTheme()
	// Simulate pulse by alternating between normal and bright
	if time.Now().UnixMilli()%1000 < 500 {
		return styles.NewStyle().
			Foreground(t.Primary()).
			Bold(true).
			Render(text)
	}
	return styles.NewStyle().
		Foreground(t.Primary()).
		Render(text)
}

// FormatTimestamp formats a timestamp in a visually appealing way
func FormatTimestamp(timestamp time.Time) string {
	t := theme.CurrentTheme()
	now := time.Now()
	diff := now.Sub(timestamp)
	
	var timeStr string
	switch {
	case diff < time.Minute:
		timeStr = "just now"
	case diff < time.Hour:
		mins := int(diff.Minutes())
		if mins == 1 {
			timeStr = "1 minute ago"
		} else {
			timeStr = fmt.Sprintf("%d minutes ago", mins)
		}
	case diff < 24*time.Hour:
		hours := int(diff.Hours())
		if hours == 1 {
			timeStr = "1 hour ago"
		} else {
			timeStr = fmt.Sprintf("%d hours ago", hours)
		}
	default:
		timeStr = timestamp.Format("Jan 2, 15:04")
	}
	
	return styles.NewStyle().
		Foreground(t.TextMuted()).
		Italic(true).
		Render(timeStr)
}

// RenderNotification creates a notification-style message
func RenderNotification(title, message string, notificationType string) string {
	t := theme.CurrentTheme()
	
	var icon string
	var color compat.AdaptiveColor
	
	switch notificationType {
	case "success":
		icon = IconSuccess
		color = t.Success()
	case "error":
		icon = IconError
		color = t.Error()
	case "warning":
		icon = IconWarning
		color = t.Warning()
	default:
		icon = IconInfo
		color = t.Info()
	}
	
	header := styles.NewStyle().
		Foreground(color).
		Bold(true).
		Render(icon + " " + title)
	
	body := styles.NewStyle().
		Foreground(t.Text()).
		PaddingLeft(2).
		Render(message)
	
	return header + "\n" + body
}
