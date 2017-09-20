Mikael Blomstrand (mbloms@kth.se), Filip Johansson (fij@kth.se)

# Project Plan IK 2218 - The rumor protocol
The rumor protocol is a decetralized peer-to-peer gossip protocol that spreads rumors from peer-to-peer in the network. The idea is that it should be minimal and general purpose, rather than addressing a specific use.

### Functionality
The Rumor Protocol should have the following functionality.
Peers should periodically send IP addresses and data to peers in the network.

##### Peer Exchange
A peer in the network should be able to send IP addresses from its IP list to any other peer in their list.

##### Send Data
A peer should be able to send data to any other peer in its list of peers. A hash of the data is included to ensure data integrity.

##### Request Data
A peer should be able to request data from any other peer in its list of peers.

##### Data
Data will be stored in a hash table with (key, value) where key is a hash of the data and value is the data.

### Technology
* The protocol will be based on UDP.
* All communication should be simplex.

### Possible extensions
* Group peer exchange, send data and request data together in one datagram.
* Onion routing.
* Signatures instead of hashes.

### Possible usage
* Share git objects
* Twitter like social network (GUI)
