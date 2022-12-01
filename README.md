# Twitter Clone

An attempt to recreate one of the largest social networking application Twitter.

> Note: This is my first attempt at developing a distributed system, so any feedback would be greatly appreciated.

## Services

1. [Edge Service (GraphQL)](./edge/)
2. [Media Service](./media/)
3. [Notifications Service](./notifications/)
4. [Search Service](./search/)
5. [Timeline Service](./timeline/)
6. [Tweet Service](./tweet/)
7. [Web UI Service](./website/)

## Technologies Used

1. Golang (Internal Services)
2. Next.JS (Website UI)
3. Kafka (Pub/Sub)
4. Apache Spark (Analytics)
5. PostgreSQL (Main Database)
6. PgPool II (Database Load Balancer)
7. Redis (Caching)
8. Docker (Containerization)
9. Nginx (Load Balancer)
10. GraphQL (Edge Server / Frontend for Backend)
11. ElasticSearch (Searching & Indexing)
12. Firebase Cloud Messenger (Notification Service)

## System Design

### Functional Requirements
1. Create Tweets (text, images, videos, etc)
2. View Timeline
3. Like tweets
4. Retweet / Quote Retweet
5. Follow others

### Non-Functional Requirements
1. Scalable & Efficient
2. High Availability

### Optional Requirements
1. Metrics and analytics
2. Notifications
3. Observability & Monitoring (Prometheus, Grafana, Jaeger, etc)

### Database Schema

# How to run locally

Check [here](RUNNING_LOCALLY.md) on how to run locally