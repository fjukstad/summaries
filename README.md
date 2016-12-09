# Summaries
Command line tool for retrieving NCBI gene summaries. Can get summaries for a
single gene or read gene symbols from a file. All the heavy lifting is done by
the [eutils](https://github.com/fjukstad/kvik/tree/master/eutils) and
[genenames](https://github.com/fjukstad/kvik/tree/master/genenames)  packages in
[Kvik](http://github.com/fjukstad/kvik). 

# Example
Single gene symbol: 
```
$ summaries -symbol BRCA1
This gene encodes a nuclear phosphoprotein that plays a role in maintaining
genomic stability, and it also acts as a tumor suppressor. The encoded protein
combines with other tumor suppressors, DNA damage sensors, and signal
transducers to form a large multi-subunit protein complex known as the
BRCA1-associated genome surveillance complex (BASC). This gene product
associates with RNA polymerase II, and through the C-terminal domain, also
interacts with histone deacetylase complexes. This protein thus plays a role in
transcription, DNA repair of double-stranded breaks, and recombination.
Mutations in this gene are responsible for approximately 40% of inherited breast
cancers and more than 80% of inherited breast and ovarian cancers. Alternative
splicing plays a role in modulating the subcellular localization and
physiological function of this gene. Many alternatively spliced transcript
variants, some of which are disease-associated mutations, have been described
for this gene, but the full-length natures of only some of these variants has
been described. A related pseudogene, which is also located on chromosome
17, has been identified. [provided by RefSeq, May 2009]
```

Read a file with gene symbols: 
```
$ summaries -i example/example.csv -o example/example-with-summaries.csv
```
The output is stored in `example/example-with-summaries.csv`.  Check
[example/](example/) for example input and output.



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

