# color-box
print colourful boxes
# Prerequisite
install go. Then setup the module
```
go mod tidy
go run main
```
# Usage
```
go run main
```
creates the colourful box as output with default size (20x20).
```
go run main 10 10
```
creates the colourful box as output. This time the size is 10x10
# Building
```
make
```
builds executables for linux, android and windows into the 'release' directory.
