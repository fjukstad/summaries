# Summaries
Command line tool for retrieving NCBI gene summaries. Can get summaries for a
single gene or read gene symbols from a file. 

# Example
Single gene symbol: 
```
$ summaries -gene BRCA1
```

Read a file with gene symbols: 
```
$ summaries -i example.csv -o example-with-summaries.csv
```
The output is stored in `example-with-summaries.csv`. 


# Input 
If you want to read a file with gene names into summaries, you'll need to put
one gene symbol per line and nothing else. E.g: 

```
BRCA1
ESR1
TP53
```

# Install
- Install [go](http://golang.org).
- `go get github.com/fjukstad/summaries`

