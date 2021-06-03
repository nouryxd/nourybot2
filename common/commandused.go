package common

var (
	tempCommands = 0
)

func CommandUsed() {
	tempCommands++
}

func GetCommandsUsed() int {
	return tempCommands
}
