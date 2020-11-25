ENEMIES = 1

build:
	go build

run: build
	./MultithreadedPacman -n $(ENEMIES)

clean:
	rm ./MultithreadedPacman