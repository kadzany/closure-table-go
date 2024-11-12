-- Enable the UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the nodes table
CREATE TABLE nodes
(
    id          UUID         NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    title       VARCHAR(255) NOT NULL,
    type        VARCHAR(50)  NOT NULL,
    description TEXT,
    created_at  TIMESTAMP(0) WITH TIME ZONE,
    updated_at  TIMESTAMP(0) WITH TIME ZONE,
    deleted_at  TIMESTAMP(0) WITH TIME ZONE
);

-- Create the node_closure table
CREATE TABLE node_closure
(
    ancestor   UUID REFERENCES nodes (id),
    descendant UUID REFERENCES nodes (id),
    depth      INT
);
