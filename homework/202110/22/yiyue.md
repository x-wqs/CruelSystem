 - url redirecting
 ```java
 GET api/v1/shortUrl
 ```
 return: longURL for http redirection

### URL redirecting
client -> visit short URL -> tinyurl server
       <- status code: 301 <- 
          location: long URL
       -> visit long URL  -> the redirected server

- *301 redirect*
 - the requested URL is "permanently" moved to the long URL
 - browser caches the response
 - subsequent requests for the same URL will not be sent to the URL shortening service
 - requests are redirected to the long URL server directly

- *302 redirect*
 - the URL is "temporarily" moved to the long URL
 - subsequent requests for the same URL will be sent to the URL shortening service

- implementation - hash tables <shortURL, longURL>
 - Get longURL: longURL = hashTable.get(shortURL)
 - perform URL redirect

### URL shortening
- find a hash function fx - maps a long URL to hashValue
- requirements
 - longURL hashed to one hashValue
 - hashValue mapped back to the longURL

### Step 3 - design deep dive

#### Data model

- relational database
- Primary Key - id(auto increment)
- shortURL, longURL

#### hash function
- hash value length, containing [0-9, a-z, A-Z] 62 characters, find smallest n such that 62^n $\geq $ 365 billion

#### hash + collision resolution
- bloom filter

#### base 62 conversion
- base conversion
- convert the same number between its different number representation systems

#### URL shortening deep dive
1. input: longURL -> 2. longURL in DB? -> yes -> 3. return shortURL
				  					   -> no  -> 4. generate a new ID -> 5. convert ID to shortURL -> 6. save ID, short URL, long URL in DB
