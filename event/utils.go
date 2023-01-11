package event

func (ae ArkEvent) GetColor() int {
	switch ae.Kind {
	case JoinEvent:
		return 0x4db329
	case LeaveEvent:
		return 0xb36f6f
	case DefaultEvent:
		fallthrough
	default:
		return 0xdedede
	}
}

func (ae ArkEvent) GetEventTitle() string {
	switch ae.Kind {
	case JoinEvent:
		return " hat diese ARK betreten."
	case LeaveEvent:
		return " hat diese ARK verlassen."
	case DefaultEvent:
		fallthrough
	default:
		return " "
	}
}
