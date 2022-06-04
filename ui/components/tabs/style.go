package tabs

import "github.com/charmbracelet/lipgloss"

var (
	tabsBorderHeight  = 1
	tabsContentHeight = 2
	TabsHeight        = tabsBorderHeight + tabsContentHeight

	indigo       = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#383B5B"}
	subtleIndigo = lipgloss.AdaptiveColor{Light: "#5A57B5", Dark: "#242347"}

	tab = lipgloss.NewStyle().
		Faint(true).
		Padding(0, 2)

	activeTab = tab.
			Copy().
			Faint(false).
			Bold(true).
			Background(lipgloss.AdaptiveColor{Light: subtleIndigo.Light, Dark: "#39386b"}).
			Foreground(lipgloss.AdaptiveColor{Light: "#242347", Dark: "#E2E1ED"})

	tabsRow = lipgloss.NewStyle().
		Height(tabsContentHeight).
		PaddingTop(1).
		PaddingBottom(0).
		BorderBottom(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderBottomForeground(lipgloss.AdaptiveColor{Light: indigo.Light, Dark: indigo.Dark})
)
