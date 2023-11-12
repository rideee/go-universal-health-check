srcDir=./src
binDir=./bin

ifeq ($(OS),Windows_NT)     		# is Windows_NT on XP, 2000, 7, Vista, 10...
    detected_OS := Windows
	binName = unihc.exe
else
    detected_OS := $(shell uname)  	# same as "uname -s"
	binName = unihc
endif


### Tasks ###

# Run application.
devrun:
	@go run $(srcDir)

# Clean bin folder.
clean:
	@rm $(binDir)/* 2> /dev/nul || echo "clean: Nothing to do."

# Build application.
build:
	go build -o $(binDir)/$(binName) $(srcDir)

# Run the built application.
run:
	$(binDir)/$(binName)
