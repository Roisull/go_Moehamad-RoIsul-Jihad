# Redis
<!-- SELECT * FROM users; -->

<!-- dalam bentuk redis -->
KEYS users:*

# Neo4j
<!-- SELECT * FROM users; -->

<!-- dalam bentuk Neo4j -->
MATCH (u:User)
RETURN u.*;

# Cassandra
<!-- SELECT * FROM users; -->

<!-- dalam bentuk Cassandra -->
SELECT * FROM users;