RUNNERFILES=\
	koan_runner.go
KOANSFILES=\
	about_asserts.go
KOANSLIB=\
	koans.go
SOURCES=$(KOANSLIB) $(KOANSFILES) $(RUNNERFILES)
OBJECTS=$(SOURCES:.go=.8)

%.8:%.go
	8g $<

all: koan_runner

koan_runner: $(OBJECTS)
	8l -o koan_runner koan_runner.8

clean:
	rm -f $(OBJECTS)
	
