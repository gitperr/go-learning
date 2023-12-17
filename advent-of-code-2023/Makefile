ifeq ($(OS),Windows_NT)
    SHELL := powershell.exe
else ifeq ($(shell uname),Linux,Darwin)
    SHELL := zsh
else
    $(error Unsupported operating system)
endif

ifeq ($(OS),Windows_NT)
	MKDIR := $(SHELL) -Command New-Item -ItemType Directory
	NEWFILE := $(SHELL) -Command New-Item -ItemType File
	ECHO := $(SHELL) -Command Write-Output
	CURL := curl.exe
else
	MKDIR := mkdir -p
	NEWFILE := touch
	ECHO := echo
	CURL := curl
endif

ifdef cookie
	cookie := $(cookie)
endif

.PHONY: all
all: create_directories mod download_input
	@echo "Using SHELL: $(SHELL)"

.PHONY: create_directories
create_directories:
ifeq ($(DAY),)
	$(error DAY is not set. Usage: make create_directories DAY=<day number e.g. 1,2,3>)
endif
	$(MKDIR) day$(DAY)/firsthalf
	$(MKDIR) day$(DAY)/secondhalf
	@echo "Created directories: day$(DAY)/firsthalf and day$(DAY)/secondhalf"
	$(NEWFILE) day$(DAY)/firsthalf/day$(DAY).go
	$(NEWFILE) day$(DAY)/firsthalf/day$(DAY)_test.go
	$(NEWFILE) day$(DAY)/firsthalf/main.go
	$(NEWFILE) day$(DAY)/secondhalf/day$(DAY).go
	$(NEWFILE) day$(DAY)/secondhalf/day$(DAY)_test.go
	$(NEWFILE) day$(DAY)/secondhalf/main.go
	
mod:
	@cd day$(DAY); go mod init github.com/gitperr/advent-of-code-2023/day$(DAY)

.PHONY: download_input
download_input:
ifeq ($(DAY),)
	$(error DAY is not set. Usage: make download_input DAY=<day_number> cookie=<cookie>)
endif
	$(CURL) -H "Cookie: $(cookie)" https://adventofcode.com/2023/day/$(DAY)/input -o day$(DAY)/input.txt
	@echo "Downloaded input for $(DAY) and saved to $(DAY)/input.txt"