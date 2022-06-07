package list

import (
	"charming-todo/data"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	checkedItemStyle     = lipgloss.NewStyle().Strikethrough(true).Faint(true)
	greenForegroundStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff99"))
	selectedColorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#a3cded")).Bold(true)
)

type Model struct {
	todoList   *data.TodoList
	currItemId int
}

func NewModel(todoList *data.TodoList) Model {
	return Model{
		todoList: todoList,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	s := strings.Builder{}

	for idx, item := range m.todoList.Items {
		indentationStyle := lipgloss.NewStyle().PaddingLeft(4 * (item.Indentation + 1))

		// Add initial indentation
		strItem := indentationStyle.Render("")

		// Add checkmark and strike description if completed
		if item.Checked {
			strikedDesc := checkedItemStyle.Render(item.Description)
			strItem = strItem + greenForegroundStyle.Render("✓ ") + strikedDesc
		} else {
			strItem = strItem + "☐ " + item.Description
		}

		// Add prefix to selected item
		if idx == m.currItemId {
			strItem = selectedColorStyle.Render("›› " + strItem)
		} else {
			// Compensate for the above prefix
			strItem = "   " + strItem
		}

		s.WriteString("\n")
		s.WriteString(strItem)
	}

	return s.String()
}

func (m *Model) SetTodoList(todoList *data.TodoList) {
	m.todoList = todoList
	m.currItemId = 0
}

func (m *Model) ToggleCurrentItem() {
	if m.currItemId >= 0 && m.currItemId < len(m.todoList.Items) {
		currentItem := &(*m.todoList).Items[m.currItemId]
		currentItem.Checked = !currentItem.Checked
	}
}

func (m *Model) AddIndent() {
	currentItem := &(*m.todoList).Items[m.currItemId]

	if currentItem.Indentation >= data.MAX_INDENTATION {
		return
	}

	currentItem.Indentation++
}

func (m *Model) RemIndent() {
	currentItem := &(*m.todoList).Items[m.currItemId]

	if currentItem.Indentation <= 0 {
		return
	}

	currentItem.Indentation--
}

func (m *Model) NextItem() {
	nextIdx := m.currItemId + 1

	if nextIdx >= len(m.todoList.Items) {
		nextIdx = 0
	}

	m.currItemId = nextIdx
}

func (m *Model) PrevItem() {
	prevIdx := m.currItemId - 1

	if prevIdx < 0 {
		prevIdx = len(m.todoList.Items) - 1
	}

	m.currItemId = prevIdx
}
