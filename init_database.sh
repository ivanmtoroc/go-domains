#!/usr/bin/bash

# 1. Delete previous certificates if it exists
rm -rf certs/ ca-dir/
mkdir certs/ ca-dir/

# 2. Create the CA (Certificate Authority)
cockroach cert create-ca --certs-dir=certs --ca-key=ca-dir/ca.key

# 3. Create the certificate and key pair for the node
cockroach cert create-node localhost $(hostname) --certs-dir=certs --ca-key=ca-dir/ca.key

# 4. Create a client certificate and key pair for the root user
cockroach cert create-client root --certs-dir=certs --ca-key=ca-dir/ca.key

# 5. Start the single-node cluster
cockroach start-single-node --certs-dir=certs --listen-addr=localhost:26257 --background

# 6. Create database table
cockroach sql --certs-dir=certs --host=localhost:26257 < models/config/create_db_tables.sql
