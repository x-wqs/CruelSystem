web crawler

# Requirements
- Scalability: needs to be scalable such that it can crawl the entire Web and can be used to fetch hundreds of millions of Web documents.
- Extensibility: should be designed in a modular way with the expectation that there could be newer document types that needs to be downloaded and processed in the future.
  
# Design Considerations
- **HTML only or not?**
If we are writing a general-purpose crawler to
download different media types, we might want to break down the parsing module
into different sets of modules: one for HTML, another for images, or another for
videos, where each module extracts what is considered interesting for that media
type.

- **Http only or not?**
it shouldn’t be hard to extend the design to use FTP and
other protocols later.

- **Data size?** 1 billion websites, 15 billion web pages

# Capacity Estimation
- QPS: 15 Billion pages / (4 weeks) = 15 / 28 days = 15 billion / (28 * 0.1 million) ~= 0.5 billion / 0.1 million = 5000 pages / sec
- Storage: 15 billion * (100KB + 500B) -> 15 billion * 100 KB = 1500 Tb = 1.5 Pb, reserve capacity 70% -> 2 Pb

# High Level Design
## steps
- pick up a url from unvisised url pool
- use dns to find ip
- use http (or ftp) to fetch content
- parse content based on file type
- filter new urls and check dedup
- add remain url to unvisited url pool
- index contents

## components
1. URL frontier: To store the list of URLs to download and also prioritize which
URLs should be crawled first.
2. HTTP Fetcher: To retrieve a web page from the server.
3. Extractor: To extract links from HTML documents.
4. Duplicate Eliminator: To make sure the same content is not extracted twice
unintentionally.
5. Datastore: To store retrieved pages, URLs, and other metadata.

# Detailed Design
- modular: Based on the URL’s
scheme, the worker calls the appropriate protocol module to download the
document.

- politeness constraint: do not crawl a server too frequently

- URL distribution: consistent hashing based on hostname to determine crawler server, then map url to specific worker to make sure only one worker is reponsible for that url

- URL frontietr: enqueue && dequeue buffer, use disk to store url and load to cache for dequeuing

- Document input stream: avoid downloading a document
multiple times, we cache the document locally

- A DIS is an input stream that caches the entire contents of the document read from
the internet. It also provides methods to re-read the document.

- Document Dedupe test: calc checksum for content & filter url based on bloom filter
