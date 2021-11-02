### Notification Type
- Provider ---> Apple Push Notification Services ---> iOS
- Provider ---> Google Firebase Cloud Messaing ---> Android
- Provider ---> SMS Service / Twilio / Nexmo ---> SMS

### Contact Info gathering
- user --> install/sign up --> load balancer --> application server ---> ContactInfoDB

### Notification sending/receiving
- Service1 -+-> Notification System N -+-> iOS Queue --> WorkerN --> APNS ---> iOS
- Service2 -+           |              +-> And Queue --> WorkerN --> FCM  ---> Android
- Service3 -+         Cache            +-> SMS Queue --> WorkerN --> SmsSvc -> SMS
-                       |              +-> EML Queue --> WorkerN --> EmailSvc -> Email
-                     MsgDB						
- ServiceN: MicroServices, cron jobs
- Notification System
	- provide API for service to send notification
	- basic validation
	- query database and form message
	- put notification to message queue	
- 3rd-party services, extensibility
- user devices:
- Cache: user info, device info, notification template
- Message Queue: worker pull events and push to 3rd-party services

## Detail Design
- reliability, persist notification log and retry
- notification template/settings, {Hi - Subject}, OptIn
- rate limiting
- retry mechanism, add back to message queue
- security with token
- monitor message queue, increase worker if traffic too high
- analysis, open/click/engament rate

## Update Design
-                     +---> Analysis Service <------------------------------+
-                     |                           +-<--- retry on error ----+
- ServiceN --> NotifSystem(Auth,RateLimit) ---> iOS PN ------------------> Worker ----> APNS ----> iOS
-                     |                                                     +---> Notification Template
-                   Cache                                                   +---> Notification Log
-                     |
-                DB(device/user/setting info) 
