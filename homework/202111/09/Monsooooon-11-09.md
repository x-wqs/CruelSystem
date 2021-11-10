design a chat system
# design requirement
- 1-on-1 or group chat
- mobile / web
- dau
- group size limit
- contacts num limit
- message size limit
- message type (text/media)
- online/offline status
- encryption
- store chat history
- add/rm contacts
- multi-device login
- sign up
- user profile
- TODO: functional vs non-functional

# high level design
## comm protocol:
- comm pattern: sender -> chat srv -> receiver
- sender -> chat srv: HTTP, keep-alive (less TCP handshakes)
- char srv -> receiver:
    - polling: keep asking (inefficient if frequently)
    - long polling: polling, but srv takes more time to wait and resp
        - sender & receiver need take to the same chat srv
        - receiver may have gone but srv still waits
        - inefficient
    - **websocket (BEST)**
        - send async updates from srv -> cli
        - persistent
        - bi-directional
        - based on http
        - use port 80 or 443 so it works with firewall
        - efficient
## component categories
- stateless:
    - login / sign up / change profile
    - LB -> app servers
- stateful
    - chat service
    - makes sure consistent connection btw sender/chat server/receiver
- 3rd party
    - notify a new message has arrived to user even when app is not running
    - just doing integration is ok
## capacity & scalability
- 1 M concurrent user
- 10 K per user connection
- 10 GB mem for all connections
- start from single srv but finally cluster
## arch
- chat servers
- presence servers
- LB
- api servers
- notification servers
- storage (KV? MySQL?)
## data model
- kinds of data
    - relational
        - profiles
        - user setting
        - contacts list
    - no-sql
        - chat history
        - easy horizontal scaling
        - low latency to access data
        - used by tech giants
## data-model
- 1-on-1
- group chat
- how: msg id
    - unique in one chat is ok
    - snowflake

# design deep dive
## service discovery
- why:
> “recommend the best chat server for a client based on the criteria like geographical location, server capacity, etc.”
- how:

> “Apache Zookeeper is a popular open-source solution for service discovery. It registers all the available chat servers and picks the best chat server for a client based on predefined criteria.”

find server: user -> lb -> api server -> service discovery
connect server: user -> chat server (via ws)

## msg flows
### 1-on-1
- basic flow: user -> chat srv1 -> mq -> chat srv2 -> receiver
- chat srv1 use idGen srv to get msgID and store msg to persistent kv db.
- chat srv2 may use push notification to inform receiver
### multi-device msg sync
- diff devices maintains a local cur_max_msg_id
- device talks to the chat server
- if there is message with larger msg_id in the kv store, chat srv send them to the lagged device
### group chat
- group size matters, should be < 500
- every user has a message sync queue(inbox) to store latest message
- user checks its inbox to get new messages
- redundant data but ok for small group

## online presence
### user login/logout
api server is responsible for login authentication & send status updates to all chat servers
### user disconnect
use heartbeat message with interval(30s) to mark user status
