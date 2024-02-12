# Devnet

This directory defines a docker-compose file alongside a set of scripts to generate a 3 node devnet.

For each full run, the nodes are initialized, the validator accounts are created, the genesis transaction is built and the chain is started.

This setup makes use of docker volumes to persist the chain state between runs.

To fully wipe a chain use `make clean` or `docker-compose down --volumes`.

To spin up a localnet run `make localnet` or `docker-compose up` in this directory.
