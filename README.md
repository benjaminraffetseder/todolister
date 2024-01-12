# TODO lister

This is a command-line tool written in Go that scans a given directory for TODO comments in the code. It generates a CSV file with the file name and line number of each TODO found.

## Installation

First, clone this repository to your local machine using `git clone`.

Then, navigate to the project directory and build the project using `go build`.

## Usage

To use this tool, run the following command:

```bash
./todofinder [directory]
```

Replace [directory] with the path to the directory you want to scan.

The tool will create a stats folder in the given directory and generate a todos.csv file in this folder. This file will contain the file name and line number of each TODO found.
