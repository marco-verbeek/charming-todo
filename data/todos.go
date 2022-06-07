package data

type TodoList struct {
	Title     string
	Displayed bool
	Items     []TodoItem
}

type TodoItem struct {
	Description string
	Indentation int
	Checked     bool
}

const MAX_INDENTATION = 6

func (i TodoItem) FilterValue() string { return i.Description }

func FetchTodoList() TodoList {
	return TodoList{
		Title:     "My Todo",
		Displayed: true,

		Items: []TodoItem{
			{Description: "Learn Go!", Indentation: 0, Checked: false},
			{Description: "Read documentation", Indentation: 1, Checked: true},
			{Description: "Write Charming Todo", Indentation: 1, Checked: false},
			{Description: "Core logic", Indentation: 2, Checked: true},
			{Description: "Advanced logic", Indentation: 2, Checked: false},
			{Description: "Have fun", Indentation: 0, Checked: true},
		},
	}
}

func TemplateTodoList() TodoList {
	return TodoList{
		Title:     "• Unsaved List",
		Displayed: true,

		Items: []TodoItem{
			{Description: "Item", Indentation: 0, Checked: false},
		},
	}
}
