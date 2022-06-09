package data

type TodoList struct {
	Title     string
	Displayed bool
	Dirty     bool
	Items     []TodoItem
}

type TodoItem struct {
	Description string
	Indentation int
	Checked     bool
}

const MAX_INDENTATION = 6
const DIRTY_PREFIX = "• "

func (i TodoItem) FilterValue() string { return i.Description }

func FetchTodoList() TodoList {
	return TodoList{
		Title:     "My Todo",
		Displayed: true,
		Dirty:     false,

		Items: []TodoItem{
			{Description: "~ Welcome to Charming Todo!", Indentation: 0, Checked: false},
			{Description: "Introduction", Indentation: 0, Checked: false},
			{Description: "Use 'ctrl+n' to create a new list", Indentation: 1, Checked: false},
			{Description: "Then, create new items by pressing 'n'", Indentation: 1, Checked: false},
			{Description: "Pressing 'enter' will toggle the item's checkmark", Indentation: 1, Checked: false},
			{Description: "Finally, hit ctrl+s to save it!", Indentation: 2, Checked: false},
			{Description: "You can indent the list with 'tab' and 'shift+tab'.", Indentation: 0, Checked: false},
			{Description: "Have fun!", Indentation: 0, Checked: false},
		},
	}
}

func TemplateTodoList() TodoList {
	return TodoList{
		Title:     "Unsaved List",
		Displayed: true,
		Dirty:     true,

		Items: []TodoItem{
			{Description: "Item", Indentation: 0, Checked: false},
		},
	}
}

func TemplateTodoItem() TodoItem {
	return TodoItem{
		Description: "New item",
		Indentation: 0,
		Checked:     false,
	}
}
