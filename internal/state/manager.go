package state

import "time"

type State struct {
	Step     string
	Title    string
	Priority string
	Urgency  string
	DueDate  string
	Status   bool
}

type Manager struct {
	states map[int64]*State
}

func NewManager() *Manager {
	return &Manager{
		states: make(map[int64]*State),
	}
}

func (m *Manager) StartCreation(chatID int64) {
	m.states[chatID] = &State{Step: "title"}
}

func (m *Manager) Get(chatID int64) *State {
	return m.states[chatID]
}

func (m *Manager) SetTitle(chatID int64, title string) {
	if state := m.states[chatID]; state != nil {
		state.Title = title
		state.Step = "priority"
	}
}

func (m *Manager) SetPriority(chatID int64, priority string) {
	if state := m.states[chatID]; state != nil {
		state.Priority = priority
		state.Step = "urgency"
	}
}

func (m *Manager) SetUrgency(chatID int64, urgency string) {
	if state := m.states[chatID]; state != nil {
		state.Urgency = urgency
		state.Step = "due_date"
	}
}

func (m *Manager) SetDueDate(chatID int64, date string) bool {
	if _, err := time.Parse("2006-01-02", date); err != nil {
		return false
	}

	if state := m.states[chatID]; state != nil {
		state.DueDate = date
		state.Step = "status"
		return true
	}
	return false
}

func (m *Manager) SetStatus(chatID int64, status bool) {
	if state := m.states[chatID]; state != nil {
		state.Status = status
	}
}

func (m *Manager) Complete(chatID int64) {
	delete(m.states, chatID)
}
