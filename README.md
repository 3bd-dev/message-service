
# **Message Service**

### 1. Overview

The **Message Service** is a microservice that manages the creation, approval, and rejection of messages through a **maker-checker** workflow. It allows makers to create messages and checkers to approve or reject them. A message is sent to the recipient only if it receives the required number of approvals. If any checker rejects the message, it is not sent.

**Development Time**: Approximately 8 hours.

Key capabilities include:
- **Message Creation**: Makers can create messages with content and recipient details.
- **Approval Workflow**: Checkers can approve or reject messages. A message is approved only if it meets the required number of approvals.
- **Rejection Handling**: If any checker rejects the message, it is immediately marked as rejected.
- **Extensible Architecture**: The service is designed to be easily extended with additional features such as dynamic approval requirements or event-driven workflows.


### 2. Project Structure

The project is organized into various directories, each with a clear responsibility, making it easy to maintain and extend. Below is the directory structure:

```bash
├── cmd/                         # Command entry points for the services
│   └── api/
│       └── messages/            # Main message service entry point
├── config/                      # Service configuration (env variables, etc.)
│   └── config.go
├── docker/                      # Docker setup for building and running services
│   └── dockerfile.api
├── docker-compose.yml           # Docker Compose file to run all services
├── docs/                        # OpenAPI/Swagger documentation
│   ├── docs.go
│   ├── message.go
│   └── swagger.yaml
├── internal/                    # Core business logic and internal packages
│   ├── db/                      # Database migration scripts
│   │   ├── migrations/          # Database migration files
│   │   └── seed/                # Database seed files
│   ├── handlers/                # HTTP/gRPC layer handlers
│   │   ├── checkapi/            # Health check API handlers
│   │   └── messageapi/          # Message API handlers
│   ├── models/                  # Entities (e.g., user, message, message_review)
│   ├── repos/                   # Repositories for interacting with the database
│   │   └── postgres/            # PostgreSQL repository implementations
│   ├── services/                # Business logic / Use case services
│   │   └── message/             # Message service logic
│   └── web/                     # Middleware, routing, and logging utilities
├── pkg/                         # Reusable packages and utilities
│   ├── database/                # Database connection helpers
│   ├── errs/                    # Custom error handling package
│   ├── logger/                  # Logging package
│   ├── page/                    # Pagination utilities
│   ├── query/                   # Query building utilities
│   ├── queue/                   # Simple in-memory queue for async processing
│   ├── rest/                    # HTTP client with retry functionality
│   └── web/                     # Web response helpers
├── go.mod                       # Go module file
├── go.sum                       # Go dependencies checksum file
├── makefile                     # Makefile for building and running the service
└── README.md                    # Project README file
└── message.postman_collection.json # Postman collection for the api with examples
```

### 3\. Features

*   **Maker-Checker Workflow**: Makers create messages, and checkers approve or reject them.
    
*   **Dynamic Approval Requirements**: Messages require a configurable number of approvals to be sent.
    
*   **Rejection Handling**: A single rejection from any checker marks the message as rejected.
    
*   **Role-Based Access Control**: Users are assigned roles (admin, maker, checker) to control access to specific actions.
    
*   **Database Migrations**: Uses dbmate for database migrations and seeding.
    
*   **Extensibility**: The service is designed to be easily extended with additional features such as event-driven workflows or dynamic approval rules.
    

### 4\. Architecture Overview

The **Message Approval Service** is built with a modular and extensible architecture to handle message creation, approval, and rejection workflows. The design separates business logic from data access, making the system easy to extend and scale.

#### **Architecture**:

* API Layer: Entry point for the service in cmd.

* Handler Layer: Handles HTTP/gRPC requests and responses.

* Service Layer: Contains business logic and use cases, with clear contracts for dependencies.

* Repository Layer: Manages data layer (e.g., PostgreSQL).


### 5\. Technical Design and Decisions

1.  **Role-Based Access Control**:
    
    *   Users are assigned roles (admin, maker, checker) to separate responsibilities. Makers create messages, and checkers approve or reject them.
        
2.  **Configurable Approvals**:
    
    *   The approvals\_required field is configurable, allowing flexibility to set different approval thresholds based on the message or task (for now its hardcoded and >= 2).
        
3.  **Approval and Rejection Logic**:
    
    *   A message requires the specified number of approvals to be marked as approved. However, if any checker rejects the message, it is immediately marked as rejected.
        
4.  **PostgreSQL Database**:
    
    *   PostgreSQL is used as the relational database to handle the structured data and relationships between users, messages, and approvals.
        
5.  **Database Migrations with dbmate**:
    
    *   dbmate is used for database migrations and seeding, ensuring a consistent and version-controlled database schema.
        
6.  **HTTP Handling with Gorilla Mux**:
    
    *   While the standard library is preferred, Gorilla Mux is used for HTTP routing due to time constraints. Its signature is similar to the standard library, making it easy to transition to a custom router in the future.
    

### 6\. Flow of Message Approval

The message approval flow is designed to handle message creation, approval, and rejection efficiently.

#### **1\. Message Creation**:

*   A maker creates a message with content and recipient details.
    
*   The message is saved in the database with a status of pending.
    

#### **2\. Approval/Rejection**:

*   Checkers review the message and either approve or reject it.
    
*   If the message receives the required number of approvals, its status is updated to approved.
    
*   If any checker rejects the message, its status is updated to rejected.
    

#### **3\. Message Delivery**:

*   Once a message is approved, it should be sent to the recipient (future enhancement).

![Flow Diagram](https://uml.planttext.com/plantuml/png/ZPHDJuGm48Rl_HMJzR3hmNZdODm5JKmaCKlqFfPPhNOfJhkh_di7AA3YgXn8Q9--yyv7EVS1EUosLYdKmJ7a-4mCw26F2-aZwRYKgdlt1hDbNLkZCHXNmGfbiofdS7JHFWhagKNIQ5DZ5OI-kO9hV2UUDHBBAP2Jz_X0K13lZQOMouifbs70U_GTaAsKX0NnerGE9sUmJ21bma20KD4hR3jLjMmsyoA1IyUlo6M_0RgXcnIONdRaWygB7hr3Qyf6Qc1T5KWye8Ha4aD6lkA1R9L0WViHNaX9myudK8j3KpqSHr0ShYdik1h8g-m7k3FyFrQbKVYIWdxir_ipVan-tUKu9SOS91uRaZmECADzQqJ9RGw4InF8ZAyjla-1mQ-C4BizdQxfYNGut2niERIdt89_j4l7KMLdgUrLNxokSDyRnhGtCYaBR8MJ68pA68pAuqydRkil4z4fmi9iMaE6QJJhXP82YOzlEjC9Y1PafXoNG_r-XHJQ0FsTYNKiszlobzdWODxdE5TaFGqww-ZqtpfzQlFc6YfrBc_vYNm0)

### 7\. Entities

Here’s the Entities diagram representing the tables:

![Flow Diagram](https://uml.planttext.com/plantuml/png/nLHDIyD04BtlhrXwgGgv57e99LJy095AgVKSbkoqNKqswNvK9VM_EvlORP9-W21u2PFDUpFlpSnch0rJncOfwTDRM0W95AGHHe0cLV19hGQbQKaelH8yfDOAJaUZwMCKuL4GuCCX9CiWf1kcaXLJ5uFXy191LegrXIZ2KEQ6Ifkc5KlbAO8TBNPXztEYW1dWCJCXDI83r9WL2EAmO3OrzEPbDhjxciVpXydTy_nwCgrOjk2_PdsHdKl4QxQ4yqQJN1h4OtxuC5tLMCzO7UuonVLxGq2_0Cc5NFRh7a0Y2b5bwtIjdHKvsAQ6eFjQ42iAbMzOgcC5QoiK8498qwWt-EzcehADWFVpFTrX3ntMxEFc4oDo6NOAnVNnNwmJUcbS29kvss2pniewIundGZeBx0sKAxI2n4LL8n3TsGH7-La1Fp_WbT5NI53X0R5jgppDybeLfCo1zKeKcjHNTxiDWhpSRtX8UtwhUiGVjs5-TeZUZwXtD6KJsnYQp-s_ErfvrDn96NNlscNz50WXOtI2lwzl)

### 8\. API Endpoints

The service exposes the following API endpoints:

*   **GET /api/v1/messages**: List all messages.
    
*   **POST /api/v1/messages**: Create a new message.
    
*   **POST /api/v1/messages/{id}/reviews**: Approve or reject a message.
    

For detailed API documentation, refer to the Swagger documentation at http://localhost:8080/swagger.

### 9\. How to Run the Project

1.  `https://github.com/3bd-dev/message-service && cd message-service`
    
2.  `docker-compose up --build ` The message service will be available at http://localhost:8080/api/v1.
    
3.  **API Documentation**:
    
    *   Access the OpenAPI (Swagger) documentation at http://localhost:8080/swagger.
        
    *   Postman collection is available in `message.postman_collection.json`.
        

### 10\. Technologies Used

*   **Go**: Version 1.22
    
*   **PostgreSQL**: Database for user, message, and approval data.
    
*   **Docker**: Containerization of services.
    
*   **Docker Compose**: Manage multi-container applications.
    
*   **dbmate**: Database migration tool.
    
*   **Swagger (OpenAPI)**: API documentation.
    
*   **Postman**: API testing and collection.
    

### 11\. Future Enhancements

1.  **Unit and End-to-End Testing**:
    * Add comprehensive unit and end-to-end tests. Currently, testing is skipped due to time constraints, but it is planned to use a real database instance for testing, requiring a Docker setup for database management.
    
4.  **Event-Driven Architecture**:
    
    *   Implement an event-driven architecture using a message broker like Kafka or RabbitMQ. When a message is approved, an event will trigger a subscriber to send the message. Additional statuses (e.g., sent, delivered) can be added to track the message lifecycle after approval.
        
5.  **Dynamic Approval Rules**:
    
    *   Allow the approvals\_required field to be dynamically configured per message or task, providing more flexibility in the approval process (now it 2).
        
6.  **Comprehensive Logging and Monitoring**:
    
    *   Integrate logging and monitoring tools (e.g., ELK stack, Prometheus) to improve observability and debugging.
        
7.  **Data Security**:
    
    *   Add encryption and hashing for sensitive data, such as message content and user details, to enhance security.
    

This **Message Approval Service** is designed with **scalability**, **extensibility**, and **performance** in mind, enabling easy integration with additional features and providing a reliable platform for message approval workflows.