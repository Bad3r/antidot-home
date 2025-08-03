package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// confirmModel represents the state of our confirmation prompt
type confirmModel struct {
	message   string
	confirmed bool
	decided   bool
}

func (m confirmModel) Init() tea.Cmd {
	// No initial command needed
	return nil
}

func (m confirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch strings.ToLower(msg.String()) {
		case "y", "yes":
			m.confirmed = true
			m.decided = true
			return m, tea.Quit
		case "n", "no":
			m.confirmed = false
			m.decided = true
			return m, tea.Quit
		case "enter":
			// Default to "no" when Enter is pressed
			m.confirmed = false
			m.decided = true
			return m, tea.Quit
		case "q", "ctrl+c":
			// User cancelled - treat as "no"
			m.confirmed = false
			m.decided = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m confirmModel) View() string {
	// Format: "message (y/N) "
	return fmt.Sprintf("%s (y/N) ", m.message)
}

// Confirm displays a yes/no prompt and returns the user's choice
func Confirm(text string) bool {
	model := confirmModel{
		message:   text,
		confirmed: false,
		decided:   false,
	}

	// Create program without alternate screen or mouse tracking
	p := tea.NewProgram(model)

	// Run the program
	finalModel, err := p.Run()
	if err != nil {
		// On error, default to false
		return false
	}

	// Extract the final answer from the model
	if m, ok := finalModel.(confirmModel); ok {
		return m.confirmed
	}

	// Default to false if something went wrong
	return false
}
