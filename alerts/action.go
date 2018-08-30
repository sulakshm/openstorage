package alerts

// Action is something that runs for an alerts manager.
type Action interface {
	Run(manager Manager) error
}

// ActionType defines categories of actions.
type ActionType int

// ActionType constants
const (
	// Deletes alert entries.
	DeleteAction ActionType = iota
	// Alerts get marked for deletion.
	ClearAction
	// Custom user action.
	CustomAction
)

// action implements Action interface.
type action struct {
	action  ActionType
	filters []Filter
	f       func(manager Manager, filter ...Filter) error
}

func (a *action) Run(manager Manager) error {
	return a.f(manager, a.filters...)
}

func deleteAction(manager Manager, filters ...Filter) error {
	return manager.Delete(filters...)
}

// clearAction first enumerates, then changes Cleared flag to true,
// then updates it.
// Raise method determines if TTL needs to be applied based on clear flag.
func clearAction(manager Manager, filters ...Filter) error {
	myAlerts, err := manager.Enumerate(filters...)
	if err != nil {
		return err
	}

	for _, myAlert := range myAlerts {
		myAlert.Cleared = true
		if err := manager.Raise(myAlert); err != nil {
			return err
		}
	}

	return nil
}