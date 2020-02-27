package core

import "github.com/oneconcern/datamon/pkg/model"

// LabelOption is a functor to build labels
type LabelOption func(*Label)

// LabelDescriptor sets the descriptor for this label
func LabelDescriptor(r *model.LabelDescriptor) LabelOption {
	return func(l *Label) {
		if r != nil {
			l.Descriptor = *r
		}
	}
}