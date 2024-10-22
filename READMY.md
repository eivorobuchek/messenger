           +-------------+
           |  Auth       |<---> PostgreSQL
           |  Service    |
           +-------------+
                 |
                 v
          RabbitMQ (events)
                 |
     +-----------+------------+
     |                        |
+--------------+         +-------------+
| User Profile |  <---> MongoDB        |
| Service      |         +-------------+
+--------------+                       | 
                                       |
                        +--------------+
                        |   Chat       |
                        |   Service    |
                        +--------------+
                        |
                        v
                        Cassandra
