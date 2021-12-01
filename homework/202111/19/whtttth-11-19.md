# Design Google Drive

## Functional Requirement

* Add files by dragging or dropping file into google drive

* Download files

* Sync files across multiple devices

* See file revisions

* Share with friends and other users

* Send a notification when file is updated/deleted/shared

## Non Functional Requirement

* Highly available, should not be down anytime

* Scalability, handle user growth, data replication etc.

* Low latency, service should act in real time

* Reliability, data loss is not acceptable

* Fast sync speed acrpss different machines


### Back of the envolop calculation

* 50M new users and 10M DAU, 10GB free space for each user

* assume 2 files per day with 500KB

* 500PB in storage and 240-480 QPS

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




