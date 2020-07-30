package webrtc

// ICEGathererState represents the current state of the ICE gatherer.
type ICEGathererState byte

const (
	// ICEGathererStateNew indicates object has been created but
	// gather() has not been called.
	ICEGathererStateNew ICEGathererState = iota + 1

	// ICEGathererStateGathering indicates gather() has been called,
	// and the ICEGatherer is in the process of gathering candidates.
	ICEGathererStateGathering

	// ICEGathererStateComplete indicates the ICEGatherer has completed gathering.
	ICEGathererStateComplete

	// ICEGathererStateClosed indicates the closed state can only be entered
	// when the ICEGatherer has been closed intentionally by calling close().
	ICEGathererStateClosed
)

func (s ICEGathererState) String() string {
	switch s {
	case ICEGathererStateNew:
		return "new"
	case ICEGathererStateGathering:
		return "gathering"
	case ICEGathererStateComplete:
		return "complete"
	case ICEGathererStateClosed:
		return "closed"
	default:
		return unknownStr
	}
}
