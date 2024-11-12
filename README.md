Closure Table Example with Go
===============
Initiator : Anhar Solehudin [anhsbolic]

### How it works:

This approach uses a separate "closure" table to store paths between nodes, allowing flexible and efficient querying.
Each row in the closure table represents a path between an ancestor and a descendant node.

Best for highly dynamic trees with frequent updates, as it provides more scalable insertions, deletions, and moves.

### Advantages:

1. Dynamic Connections: The closure table can store paths between any nodes, not just ancestors and descendants.
2. Efficient Queries for Hierarchies and Paths: Finding all descendants of a node is a simple query.
3. Ease of Reorganization: Moving a node to a different parent is a simple update to the closure table.

### Drawbacks:

This model requires maintaining the closure table, which adds complexity, though it can be more efficient for updates.

### INSTALLATION

#### Run Docker Compose

```
docker-compose up -d --build
```

#### Run Migration

```
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```

#### Run Application

```
air
```
