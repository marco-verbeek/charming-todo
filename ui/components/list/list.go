package list

import (
	"charming-todo/data"
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight, defaultWidth = 14, 20

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)

// Delegate
type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(data.TodoItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Description)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprintf(w, fn(str))
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
	l.Title = ""

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
