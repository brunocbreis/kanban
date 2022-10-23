package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	focused  status
	lists    []list.Model
	loaded   bool
	quitting bool
}

const divisor = 4

func InitialModel() *Model {
	return &Model{}
}

func (m *Model) Next() {
	f := m.focused + 1
	m.focused = f % status(len(m.lists))
}

func (m *Model) Previous() {
	f := m.focused + status(len(m.lists)-1)
	m.focused = f % status(len(m.lists))
}

func (m *Model) loadLists(w, h int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), w/divisor, h/2)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	m.lists[todo].Title = "To Do"
	m.lists[inProgress].Title = "In Progress"
	m.lists[done].Title = "Done"

	// Loading tasks from JSON
	tasks := FromJSONFile("kanban.json")

	for _, t := range tasks {
		m.lists[t.Status].SetItems([]list.Item{t})
	}

}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// on startup, we get this message with the window size. this is what we'll use to initialize the list
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.loadLists(msg.Width, msg.Height)
			m.loaded = true
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "tab", "right":
			m.Next()
		case "shift+tab", "left":
			m.Previous()
		}
	}

	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)

	return m, cmd
}

func (m Model) View() string {

	if m.quitting {
		return ""
	}

	if !m.loaded {
		return "Loading..."
	}

	todoView := m.lists[todo].View()
	inProgView := m.lists[inProgress].View()
	doneView := m.lists[done].View()

	switch m.focused {
	case todo:
		todoView = focusedStyle.Render(todoView)
		inProgView = normalStyle.Render(inProgView)
		doneView = normalStyle.Render(doneView)

	case inProgress:
		todoView = normalStyle.Render(todoView)
		inProgView = focusedStyle.Render(inProgView)
		doneView = normalStyle.Render(doneView)

	case done:
		todoView = normalStyle.Render(todoView)
		inProgView = normalStyle.Render(inProgView)
		doneView = focusedStyle.Render(doneView)
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		todoView,
		inProgView,
		doneView,
	)

}
