echo:
	echo "hello"

gen:
	rsrc -manifest assets/Reminder.manifest -ico assets/favicon.ico -o assets/Reminder.syso

build:
	go build -ldflags '-w -s -H=windowsgui' -o Reminder.exe
