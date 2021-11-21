Multiple live Videos:
    subscription request to inform the server who is watching what
    store in an in-memory(local, tied to machine lifecycle) subscriptions table

Concurrent Viewers:
    real-time dispatcher
    needs to know which frontend machine has connections: frontend machine tell dispatcher
    entry in subscriptions table (independent key-value store)
    Any dispatcher node should be able to access that subscriptions table
 
 
subscription:
 
   viewer subscribe to frontend node -> frontend node store in memory subscription table->
   frontend node subscribe to dispatcher -> entry created in key value store

publication:
  
  viewer like -> Likes backend -> dispatcher -> frontend node -> client 
  



 
 
