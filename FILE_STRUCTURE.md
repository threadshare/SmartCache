# SmartCache Project File Structure

## Introduction

This document provides a detailed overview of the file and directory structure of the SmartCache project. It explains the purpose of each directory and file, helping developers and contributors understand the organization of the project.

## Directory Structure

```
SmartCache/
├── api/
│ ├── controllers/
│ │ ├── cache_controller.go
│ │ ├── user_controller.go
│ │ └── ...
│ ├── middlewares/
│ │ ├── auth_middleware.go
│ │ └── ...
│ ├── routes/
│ │ ├── routes.go
│ │ └── ...
│ └── main.go
├── config/
│ ├── config.go
│ ├── database.go
│ └── ...
├── docs/
│ ├── api/
│ │ ├── swagger.yaml
│ │ └── ...
│ └── architecture/
│ ├── architecture_diagram.png
│ └── ...
├── internal/
│ ├── cache/
│ │ ├── cache.go
│ │ ├── cache_manager.go
│ │ └── ...
│ ├── models/
│ │ ├── user.go
│ │ ├── cache_entry.go
│ │ └── ...
│ ├── services/
│ │ ├── cache_service.go
│ │ ├── user_service.go
│ │ └── ...
│ ├── utils/
│ │ ├── logger.go
│ │ ├── response.go
│ │ └── ...
│ └── core/
│ ├── adapters/
│ │ ├── adapter.go
│ │ ├── redis_adapter.go
│ │ └── ...
│ ├── embeddings/
│ │ ├── embedding.go
│ │ ├── onnx_embedding.go
│ │ └── ...
│ ├── similarity/
│ │ ├── similarity.go
│ │ ├── distance_similarity.go
│ │ └── ...
│ ├── storage/
│ │ ├── storage.go
│ │ ├── sqlite_storage.go
│ │ ├── mysql_storage.go
│ │ └── ...
│ ├── vector_store/
│ │ ├── vector_store.go
│ │ ├── faiss_store.go
│ │ ├── milvus_store.go
│ │ └── ...
│ ├── eviction/
│ │ ├── eviction.go
│ │ ├── lru_eviction.go
│ │ └── ...
│ ├── data_structures/
│ │ ├── string.go
│ │ ├── list.go
│ │ ├── set.go
│ │ ├── hash.go
│ │ ├── zset.go
│ │ └── ...
│ └── interfaces/
│ ├── cache_interface.go
│ ├── storage_interface.go
│ ├── vector_store_interface.go
│ ├── eviction_interface.go
│ ├── data_structure_interface.go
│ └── ...
├── scripts/
│ ├── setup.sh
│ ├── deploy.sh
│ └── ...
├── web/
│ ├── public/
│ │ ├── index.html
│ │ └── ...
│ ├── src/
│ │ ├── assets/
│ │ │ ├── logo.png
│ │ │ └── ...
│ │ ├── components/
│ │ │ ├── CacheTable.vue
│ │ │ └── ...
│ │ ├── router/
│ │ │ ├── index.js
│ │ │ └── ...
│ │ ├── store/
│ │ │ ├── index.js
│ │ │ └── ...
│ │ ├── views/
│ │ │ ├── Home.vue
│ │ │ ├── CacheManagement.vue
│ │ │ └── ...
│ │ ├── App.vue
│ │ └── main.js
│ └── package.json
├── .gitignore
├── README.md
└── go.mod
```



## Directory and File Descriptions

### api/

Contains all API-related code.

- **controllers/**: Handles HTTP requests.
    - `cache_controller.go`: Handles cache-related requests.
    - `user_controller.go`: Handles user-related requests.

- **middlewares/**: Middleware such as authentication and logging.
    - `auth_middleware.go`: Middleware for handling authentication.

- **routes/**: Defines all routes.
    - `routes.go`: Route configuration file.

- `main.go`: Entry point of the application.

### config/

Contains configuration files and initialization code.

- `config.go`: Configuration file.
- `database.go`: Database connection and initialization.

### docs/

Contains project documentation.

- **api/**: API documentation, such as Swagger files.
    - `swagger.yaml`: Swagger API documentation.

- **architecture/**: Project architecture diagrams and related documents.
    - `architecture_diagram.png`: Project architecture diagram.

### internal/

Contains the internal logic of the project.

- **cache/**: Cache-related logic.
    - `cache.go`: Cache implementation.
    - `cache_manager.go`: Cache manager.

- **models/**: Data models.
    - `user.go`: User data model.
    - `cache_entry.go`: Cache entry data model.

- **services/**: Service layer containing business logic.
    - `cache_service.go`: Cache service.
    - `user_service.go`: User service.

- **utils/**: Utility classes such as logging and response formatting.
    - `logger.go`: Logging utility.
    - `response.go`: Response formatting utility.

- **core/**: Core component packages.
    - **adapters/**: Adapters for different inputs and outputs.
        - `adapter.go`: Adapter interface.
        - `redis_adapter.go`: Redis adapter implementation.

    - **embeddings/**: Embedding generators for semantic embeddings.
        - `embedding.go`: Embedding generator interface.
        - `onnx_embedding.go`: ONNX embedding generator implementation.

    - **similarity/**: Similarity calculation modules.
        - `similarity.go`: Similarity calculation interface.
        - `distance_similarity.go`: Distance-based similarity calculation implementation.

    - **storage/**: Storage modules supporting various backends.
        - `storage.go`: Storage interface.
        - `sqlite_storage.go`: SQLite storage implementation.
        - `mysql_storage.go`: MySQL storage implementation.

    - **vector_store/**: Vector storage modules.
        - `vector_store.go`: Vector storage interface.
        - `faiss_store.go`: FAISS storage implementation.
        - `milvus_store.go`: Milvus storage implementation.

    - **eviction/**: Cache eviction policies.
        - `eviction.go`: Eviction policy interface.
        - `lru_eviction.go`: LRU eviction policy implementation.

    - **data_structures/**: Data structure modules.
        - `string.go`: String data structure.
        - `list.go`: List data structure.
        - `set.go`: Set data structure.
        - `hash.go`: Hash data structure.
        - `zset.go`: Sorted set data structure.

    - **interfaces/**: Defines interfaces for all core components.
        - `cache_interface.go`: Cache interface.
        - `storage_interface.go`: Storage interface.
        - `vector_store_interface.go`: Vector store interface.
        - `eviction_interface.go`: Eviction policy interface.
        - `data_structure_interface.go`: Data structure interface.

### scripts/

Contains script files such as setup and deployment scripts.

- `setup.sh`: Setup script.
- `deploy.sh`: Deployment script.

### web/

Contains front-end code.

- **public/**: Public resources such as HTML files.
    - `index.html`: Entry HTML file.

- **src/**: Source code.
    - **assets/**: Static resources such as images.
        - `logo.png`: Project logo.

    - **components/**: Vue components.
        - `CacheTable.vue`: Cache table component.

    - **router/**: Router configuration.
        - `index.js`: Router configuration file.

    - **store/**: Vuex state management.
        - `index.js`: Vuex state management configuration.

    - **views/**: View components.
        - `Home.vue`: Home view.
        - `CacheManagement.vue`: Cache management view.

    - `App.vue`: Root component.
    - `main.js`: Entry file.

- `package.json`: Project dependency configuration file.

### Root Directory Files

- `.gitignore`: Git ignore file.
- `README.md`: Project description file.
- `go.mod`: Go module file.