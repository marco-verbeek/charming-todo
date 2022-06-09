package list

import (
	"charming-todo/data"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	checkedItemStyle     = lipgloss.NewStyle().Strikethrough(true).StrikethroughSpaces(false).Faint(true)
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
	if len(m.todoList.Items) == 0 {
		m.currItemId = 0
		return
	}

	if m.currItemId+1 >= len(m.todoList.Items) {
		return
	}

	m.currItemId++
}

func (m *Model) PrevItem() {
	if len(m.todoList.Items) == 0 || m.currItemId-1 < 0 {
		m.currItemId = 0
		return
	}

	m.currItemId--
}

func (m *Model) NavToTopItem() {
	m.currItemId = 0
}

func (m *Model) NavToBottomItem() {
	if len(m.todoList.Items) == 0 {
		m.currItemId = 0
	}

	m.currItemId = len(m.todoList.Items) - 1
}

func (m *Model) NewItem() {
	currentItemIndentation := 0

	if m.currItemId > 0 && m.currItemId <= len(m.todoList.Items) {
		currentItem := &(*m.todoList).Items[m.currItemId]
		currentItemIndentation = currentItem.Indentation
	}

	newItem := data.TemplateTodoItem()
	newItem.Indentation = currentItemIndentation

	if len(m.todoList.Items) == 0 {
		m.todoList.Items = append(m.todoList.Items, newItem)
		return
	}

	m.currItemId++
	m.todoList.Items = append(m.todoList.Items[:m.currItemId], append([]data.TodoItem{newItem}, m.todoList.Items[m.currItemId:]...)...)
}

func (m *Model) DeleteItem() {
	if len(m.todoList.Items) == 0 {
		return
	}

	m.todoList.Items = append(m.todoList.Items[:m.currItemId], m.todoList.Items[m.currItemId+1:]...)
	m.PrevItem()
}

func (m *Model) MarkDirty(value bool) {
	m.todoList.Dirty = value
}
