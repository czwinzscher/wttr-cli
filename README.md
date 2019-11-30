# wttr-cli
A cli program that prints the current temperature for a given location

## Installation
```bash
go get github.com/czwinzscher/wttr-cli
```

## Usage
```bash
# wttr-cli <scale-opts> <Location>
wttr-cli Leipzig # defaults to Celcius
wttr-cli -f Leipzig # for Fahrenheit
wttr-cli -c -f Leizig # both Celcius and Fahrenheit
wttr-cli # uses current location
```
