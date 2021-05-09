package memento

import "strings"

type Runner struct {
	in *InputText
	sh *SnapshotHolder
}

func (r *Runner)Run(in string) string  {
	switch in {
	case ":list":
		return r.in.GetText()
	case ":undo":
		// restore
		r.in.RestoreSnapshot(r.sh.Pop())
	default:
		// backup
		r.sh.Push(r.in.CreateSnapshot())
		r.in.Append(in)
	}
	return ""
}

// InputText
type InputText struct {
	bs strings.Builder
}

func (t *InputText)GetText() string {
	return t.bs.String()
}

func (t *InputText)Append(s string)  {
	t.bs.WriteString(s)
}

func (t *InputText)CreateSnapshot() *Snapshot {
	return NewSnapshot(t.GetText())
}

func (t *InputText)RestoreSnapshot(s *Snapshot)  {
	t.bs.Reset()
	t.bs.WriteString(s.GetText())
}

func NewSnapshot(text string) *Snapshot {
	return &Snapshot{text: text}
}

// Snapshot
type Snapshot struct {
	text string
}

func (s *Snapshot)GetText() string {
	return s.text
}

func NewSnapshotHolder() *SnapshotHolder {
	return &SnapshotHolder{snapshots: []*Snapshot{}}
}

// SnapshotHolder
type SnapshotHolder struct {
	snapshots []*Snapshot
}

func (h *SnapshotHolder) Pop() *Snapshot {
	lastIdx := len(h.snapshots) - 1
	last := h.snapshots[lastIdx]
	h.snapshots = h.snapshots[:lastIdx]
	return last
}

func (h *SnapshotHolder) Push(snapshot *Snapshot)  {
	h.snapshots = append(h.snapshots, snapshot)
}
