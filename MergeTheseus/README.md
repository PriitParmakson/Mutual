### MergeTheseus

This program merges authors.json files produced by
[https://github.com/erikbern/git-of-theseus](git-of-theseus),
so that LOC evolution
chart can be produced over multiple repos.
Output is written to file mergedAuthors.json.

There is a Go version of the program - `main.go` - and a Python version - `merge.py`. 

Usage:

````
go run . <input file>...
````

````
python merge.py <input file>...
````

