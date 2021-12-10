# Scailing LInkedin

## Leo--Starting Point

Single monolithic web application host web servlets for all pages and business logic with connection to databases.

## Member Graph

To handle member connections and scale it independently, separated the system into a member graph. Using Java RPC for communication and let it scale independent from Leo. Using Lucene as service for member data.

## Read DB replication

To maintain and scale member profile database which involves both write and read traffic, introduced replica slave DB. The slave DB will be in sync with the main using databus and handle read traffic when it is consistent.

## Kill Leo

As the site scailing up to more traffic, the monolithic solution breanks down and hard to triage, so need to break it into many smaller functional stateless services.
