## Content delivery network (CDN)

- A network of geographically dispersed servers to deliver static content
- mages, videos, css, javascript files
- dynamic content caching - caching of HTML pages based on request path, query strings, cookies, request headers
- How CDN works
	- user visits a website 
	- server closest to the user delivers static contents. The further the user, the slower
- workflow
	- user tries to get image -> CDN
	- if CDN doesn't have image in the cache, request file from origin (web server / online storage, e.g. Amazon s3)
	- origin -> image -> CDN
	- CDN -> image -> user
	- user B requests image
	- image returned from CDN

## considerations for CDN
- cost
- Appropriate cache expirty
- CDN fallback - CDN outage, clients able to request resources from origin
- Invalidating files
	- use APIs provided by CDN vendors
	- use object versioning to serve different version. Can add parameter to URL

## stateless web tier
- move state (e.g. user session data) out of the web tier
- store in persistent storage
- web server in the cluster can access state data

## Stateful architecture
- Difference stateful vs stateless
  - stateful server remembers client data from one rquest to the next
  - stateless keeps no state information
- every request from same client must be routed to the same server

## stateless architecture
- HTTP requests from users can be sent to any web servers
- autoscaling - adding or removing web servers automatically based on the traffic load
- can be achieved