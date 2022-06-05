package list

import (
	"charming-todo/data"
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight, defaultWidth = 14, 20

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#a3cded"))
	checkedItemStyle  = lipgloss.NewStyle().Strikethrough(true).StrikethroughSpaces(false).Faint(true)
	checkedColorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff99"))
)

// Delegate
type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

// called on each list item to render displayed text
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(data.TodoItem)
	if !ok {
		return
	}

	indentation := strings.Repeat("   ", i.Indentation)

	var str string
	if i.Checked {
		striked := checkedItemStyle.Render(i.Description)
		checkmark := checkedColorStyle.Render("✓")

		str = fmt.Sprintf("%s %s %s", indentation, checkmark, striked)
	} else {
		str = fmt.Sprintf("%s ☐ %s", indentation, i.Description)
	}

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("›› " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

type Model struct {
	todoList *data.TodoList
	list     list.Model
}

func NewModel(todoList *data.TodoList) Model {
	items := []list.Item{}
	for _, v := range todoList.Items {
		items = append(items, v)
	}

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)

	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	return Model{
		todoList: todoList,
		list:     l,
	}
}

func (m *Model) SetTodoList(todoList *data.TodoList) {
	m.todoList = todoList

	items := []list.Item{}
	for _, v := range todoList.Items {
		items = append(items, v)
	}

	m.list.SetItems(items)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return m.list.View()
}
