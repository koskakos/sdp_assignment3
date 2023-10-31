package main

func main() {
	brain := &Brain{}

	onCommand := &OnCommand{
		device: brain,
	}

	offCommand := &OffCommand{
		device: brain,
	}

	onButton := &Switcher{
		command: onCommand,
	}
	onButton.press()

	offButton := &Switcher{
		command: offCommand,
	}
	offButton.press()
}
