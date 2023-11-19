package cmd

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gophersumit/goscaffold/pkg/templates"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

type item struct {
	groupName    string
	groupDesc    string
	templateName string
	templateDesc string
}

func (i item) Title() string       { return i.templateName }
func (i item) Description() string { return i.templateDesc }
func (i item) FilterValue() string { return i.templateName }

type model struct {
	list         list.Model
	delegateKeys *delegateKeyMap
}

func newModel() model {
	var (
		delegateKeys = newDelegateKeyMap()
	)

	items := make([]list.Item, 0)
	allgroups, err := templates.GetGroups()
	if err != nil {
		fmt.Printf("Error getting project template groups: %v\n", err)

	}

	for _, group := range allgroups.Groups {
		for _, template := range group.Templates {
			items = append(items, item{
				groupName:    group.Name,
				groupDesc:    group.Description,
				templateName: template.Name,
				templateDesc: template.Description,
			})
		}
	}

	// Setup list
	delegate := newItemDelegate(delegateKeys)
	groceryList := list.New(items, delegate, 0, 0)
	groceryList.Title = "Project Templates"
	groceryList.Styles.Title = titleStyle

	return model{
		list:         groceryList,
		delegateKeys: delegateKeys,
	}
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return appStyle.Render(m.list.View())
}
