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
	selectedTextStyle    = lipgloss.NewStyle().Background(lipgloss.Color("#4c6a8e"))
)

type Model struct {
	todoList *data.TodoList
	uiMode   *data.Mode

	currItemId int
	textCursor [2]int

	editedDesc string
}

func NewModel(todoList *data.TodoList, uiMode *data.Mode) Model {
	return Model{
		todoList: todoList,
		uiMode:   uiMode,

		textCursor: [2]int{0, 1},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// TODO: refactor this entire logic.
// Cleanup includes making use of lipgloss' SetString()
// See: https://github.com/charmbracelet/lipgloss/blob/a86f21a0ae430173036c5b6158b0654af447e5a1/example/main.go#L134
func (m Model) View() string {
	// If the current list is hidden, it means that we couldn't find a displayed list to display.
	if !m.todoList.Displayed {
		return "\nUse 'ctrl+n' to create a new Todo List or 'ctrl+o' to open an existing one."
	}

	s := strings.Builder{}

	for idx, item := range m.todoList.Items {
		isSelected := idx == m.currItemId
		indentationStyle := lipgloss.NewStyle().PaddingLeft(4 * (item.Indentation + 1))

		// Add initial indentation
		strItem := indentationStyle.Render("")

		// Add checkmark and strike description if completed
		if item.Checked {
			strikedDesc := checkedItemStyle.Render(item.Description)
			strItem = strItem + greenForegroundStyle.Render("✓ ") + strikedDesc
		} else {
			// If text is being selected
			if isSelected && *m.uiMode == data.Edit && (m.textCursor[0] > 0 || m.textCursor[1] > 0) {
				minIdx := min(m.textCursor[0], m.textCursor[1])
				maxIdx := max(m.textCursor[0], m.textCursor[1])

				minIdx = max(0, minIdx)
				maxIdx = min(maxIdx, len(m.editedDesc))

				strItem = strItem + "☐ " + m.editedDesc[:minIdx] + selectedTextStyle.Render(m.editedDesc[minIdx:maxIdx]) + m.editedDesc[maxIdx:]
			} else {
				strItem = strItem + "☐ " + item.Description
			}
		}

		// Add prefixes to selected item
		if isSelected {
			prefixes := "›› "

			// Add 'E' prefix to item if in edit mode
			if *m.uiMode == data.Edit {
				prefixes += "E "
			} else {
				prefixes += "  "
			}

			strItem = selectedColorStyle.Render(prefixes + strItem)
		} else {
			strItem = strings.Repeat(" ", 5) + strItem
		}

		s.WriteString("\n")
		s.WriteString(strItem)
	}

	if *m.uiMode == data.Edit {
		s.WriteString("\n\nCurrently in Edit Mode. Use 'enter' to save edits, 'esc' to cancel edits.")
	}

	return s.String()
}

func (m *Model) resetSelectedText() {
	m.textCursor = [2]int{0, 1}
}

func (m *Model) SetTodoList(todoList *data.TodoList) {
	m.todoList = todoList
	m.currItemId = 0

	m.resetSelectedText()
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

// TODO: implement text selection bool
// TODO: allow to go one further in order to be able to remove the last char ^^
func (m *Model) MoveCursor(amount int, selectText bool) {
	m.textCursor[0] += amount
	m.textCursor[1] += amount

	m.textCursor[0] = min(max(0, m.textCursor[0]), len(m.editedDesc))
	m.textCursor[1] = max(0, min(m.textCursor[1], len(m.editedDesc)))

	if m.textCursor[0] == m.textCursor[1] {
		m.textCursor[1] = m.textCursor[0] + 1
	}
}

func (m *Model) ListenChanges() {
	currentItem := &(*m.todoList).Items[m.currItemId]
	m.editedDesc = currentItem.Description
}

func (m *Model) ApplyChanges() {
	currentItem := &(*m.todoList).Items[m.currItemId]

	if m.editedDesc != currentItem.Description {
		m.MarkDirty(true)
	}

	currentItem.Description = m.editedDesc
}

func (m *Model) DeleteSelectedText() {
	fromIdx := max(0, m.textCursor[0]-1)
	toIdx := min(len(m.editedDesc)-1, m.textCursor[1])

	m.editedDesc = m.editedDesc[:fromIdx] + m.editedDesc[toIdx:]
}

func (m *Model) Write(text string) {
	cursorIdx := min(m.textCursor[0], m.textCursor[1])

	m.editedDesc = m.editedDesc[:cursorIdx+1] + text + m.editedDesc[cursorIdx+1:]
	m.MoveCursor(1, false)
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
