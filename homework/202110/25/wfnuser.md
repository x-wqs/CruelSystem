# Chapter8: Design a URL shortener

* traffic volume: 100 million urls / day
* as short as possible
* (0-9) | (a-z, A-Z)
* urls won't be deleted

storage 10 years: 365 TB.

301: permanently redirect
302: termporarily redirect

## URL shortening
hashValue -> logn URL

* each longURL hashed to one hashValue
* hashValue mapped back to longURL

### hash function
CRC hash function is too long.
->
First multiple characters. 
->
Hash Collisions
->
Append predefined string to url... Util no collision.

-> 
Expensive to query database for checking collision.
Bloom filter.💡

### Base62
generate id -> convert to base62
generate id should reference chapter 7.

## redirect
it's easy. Just check the database.