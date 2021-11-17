# Design Tiny URL

## Functional Requirement

* Given any URL generate a shorter url that act as an alias for it

* The short link should redirect user request to original link

* User can optionally pick a customized short link for their URL

* Linked could be expired for a default timespam, the expiration time could be specified by user

## Non Functional Requirement

* Highly available, should not be down anytime

* Scalability, handle user growth, data replication etc.

* Low latency, service should act in real time

* Fault tolerance, tailover handling

* Short link should not be predictable


### Back of the envolop calculation

* 500M new URL per month and 50B redirections per month

* ~200 QPS for new URLs and 20K QPS for redireciton

* Storage for 15 years 15TB 

## System APIs

createURL(api_dev_key, original_url, custom_alias=None, user_name=None, expire_date=None)

## Data model

### User Metadata

key: userID\
value: { name, email, DOB, location, etc.}

### URL Metadata 

key: Hash\
value: { OriginalURL, Creation date, userID, expirationDate, etc.}

## URL Encoding Algorithm

We can use Base64 encoding to compose shortened URL, with 6 letters long we can have 64^6 = 68.7 billion possible strings

To avoid collision with same input url we can append an increasing sequence number and make them unique and then generate the hash.

We can also append userID to the url since its guarateed to be unique.

## Key Generation Service

We can have a stand alone key generation service that generates random unique keys beforehand such that we can use them when a request come in.




