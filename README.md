# Url shortener

## Description

This project is a URL shortener built using Go, Redis, PostgreSQL, and NGINX. It is built without many external framework, relying solely on the Go standard library . NGINX handles load balancing across three input services, ensuring a scalable and efficient architecture.

### System Overview:

- **Go Standard Library**: The core of the project is written using the Go standard library, which handles routing and HTTP requests.
- **PostgreSQL**: Used for persistent storage of URLs and metadata, with indexes implemented to optimize database query performance.
- **Redis**: Acts as a caching layer to enhance read performance and reduce load on PostgreSQL.
- **NGINX**: Load balances requests across multiple services, distributing traffic efficiently.

The system is designed for scalability and optimized performance, making it capable of handling high traffic while minimizing latency.

![System Design](./images/design.drawio.svg)

## How to Start the Server 

1. **Clone the repository**

   ```bash
   git clone https://github.com/Sarath191181208/url_shortner
   ```

2. **Install dependencies**

   ```bash
   docker-compose up 
   ```

**NOTE:** The database migrations are handled automatically.

## Optimization Techniques

- **Database Indexing**: The PostgreSQL database uses indexes on key columns (like the short URL and long URL) to speed up lookups and optimize query performance.
- **Redis Caching**: Frequently accessed URLs are cached in Redis to reduce database load and improve read performance, ensuring that subsequent requests are faster.

## System Design Highlights

- **NGINX**: Handles load balancing to distribute incoming traffic across three URL input services.
- **PostgreSQL**: Provides durable and consistent data storage with indexing strategies to enhance query efficiency.
- **Redis**: Acts as a high-speed cache for quick URL lookups, reducing the load on the database.

The system is designed to handle large volumes of traffic efficiently while maintaining high availability and performance.

---

