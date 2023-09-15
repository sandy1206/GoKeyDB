# GoKeyDB

GoKeyDB is a distributed, scalable key-value database built in Go. It is designed to offer high availability, fault tolerance, and easy scalability for modern applications.

## Features

- **Distributed**: Horizontally scalable architecture to handle large datasets and high traffic.
- **High Availability**: Ensures data availability even in the event of node failures.
- **Fault Tolerant**: Uses replication and consensus algorithms to maintain data integrity.
- **Sharding**: Efficiently partitions data across nodes for balanced load and performance.
- **Master-Slave Architecture**: Ensures read scalability and fault tolerance.
- **Scalable**: Easily add or remove nodes in powers of 2 (2, 4, 8, 16, etc.) to scale the database as needed.
- **Performance**: Optimized for high-speed read and write operations.

## Installation

1. **Clone the repository**
    ```bash
    git clone https://github.com/sandy1206/GoKeyDB.git
    cd GoKeyDB
    ```

2. **Build the project**
    ```bash
    go build
    ```

3. **Run the project**
    ```bash
    ./GoKeyDB
    ```

## Getting Started

### Configuration

Before running the database, configure the settings in `config.yaml` to suit your environment.

#### Sample `config.yaml`

```yaml
cluster:
  nodes: 4 # Number of nodes (must be a power of 2)
  master_slaves:
    - master: node1
      slaves: [node2, node3]
    - master: node4
      slaves: [node5, node6]

storage:
  path: /data/gokeydb
network:
  port: 8080
  bind_ip: 0.0.0.0
```

### Usage

-   **Starting a node**

    bash

    Copy code

    `./GoKeyDB -config=config.yaml`

-   **Adding a key-value pair**

    bash

    Copy code

    `curl -X POST http://localhost:8080/set -d 'key=mykey&value=myvalue'`

-   **Getting a value by key**

    bash

    Copy code

    `curl http://localhost:8080/get?key=mykey`

Architecture
------------

GoKeyDB uses a distributed hash table (DHT) for partitioning and replicating data across nodes. It employs a consensus algorithm to ensure data consistency and availability.

### Sharding

Data is partitioned into shards, each managed by a different node. Sharding ensures that the load is evenly distributed across all nodes.

### Master-Slave Replication

Each master node is responsible for handling write operations and synchronizing data with its slave nodes. Slave nodes handle read operations, ensuring high availability and load distribution.

### Components

-   **Node**: The basic unit of storage and processing in GoKeyDB.
-   **Coordinator**: Manages the cluster, handles requests, and ensures data consistency.
-   **Storage Engine**: Manages data storage, retrieval, and persistence.

Detailed Steps to Set Up a Cluster
----------------------------------

1.  **Prepare the Configuration File**

    Create a `config.yaml` file with your desired cluster settings. Ensure the number of nodes is a power of 2.

2.  **Start Each Node**

    Start each node with the following command:

    bash

    Copy code

    `./GoKeyDB -config=config.yaml -node=nodeX`

3.  **Monitor the Cluster**

    Use monitoring tools or built-in APIs to check the status of each node and ensure they are functioning correctly.

4.  **Scale the Cluster**

    To scale the cluster, update the `config.yaml` file to include the new nodes (in powers of 2) and restart the affected nodes.

Contributing
------------

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

1.  Fork the repository
2.  Create a new branch (`git checkout -b feature-branch`)
3.  Make your changes
4.  Commit your changes (`git commit -m 'Add new feature'`)
5.  Push to the branch (`git push origin feature-branch`)
6.  Open a pull request

License
-------

This project is licensed under the MIT License. See the LICENSE file for details.
