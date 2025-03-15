#!/bin/sh

echo Extractions Started at `date`
# remove the active threads output file
rm -f input.txt

# run the Go script 
go run main/main.go

echo Extractions Completed at `date`