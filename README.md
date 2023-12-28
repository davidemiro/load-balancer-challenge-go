# The Challenge - Building a Load Balancer

A load balancer performs the following functions:

Distributes client requests/network load efficiently across multiple servers
Ensures high availability and reliability by sending requests only to servers that are online
Provides the flexibility to add or subtract servers as demand dictates
Therefore our goals for this project are to:

Build a load balancer that can send traffic to two or more servers.
Health check the servers.
Handle a server going offline (failing a health check).
Handle a server coming back online (passing a health check).

