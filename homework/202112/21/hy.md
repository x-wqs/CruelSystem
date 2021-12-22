**Nice properties of binary encoding based on schemas:**

- can be much more compact since they omit the field name
- can be documantation that is up to date
- forward/backward compatibility can be checked 
- for static typed languages type checking can be done at compile time

**Modes of dataflow**

- Through database
- Through services 
  - The API is application-specific
  - service oriented architecture / microservice architecture
  - compatibility is important
  - web services
    - HTTP is used as the underlying protocol
    - It can be used more than web
      - client app
      - service making request to another service 
    - REST
      - not a protocol, but a philosophy
      - build upon the principles of HTTP
      - ten to favor simpler approaches
    - SOAP
      - XML-based protocol for making network API requests
      - aim to be independent of HTTP, though it's most commonly used over HTTP
      - still used in many large enterprises 



*TODO:* RPC