# SmartCache Cache Types

SmartCache provides a flexible and powerful caching system designed to optimize the performance and cost-effectiveness of large language model (LLM) applications. This document outlines the key cache types supported by SmartCache, each tailored to address specific semantic search needs.

## Key Cache Types

### 1. Fixed QA Cache
**Description**: Stores fixed questions and their corresponding fixed text answers without any changes.
**Use Cases**: FAQ systems, common question-answering scenarios.
**Example**:
```json
{
  "type": "fixed_qa",
  "question": "What is SmartCache?",
  "answer": "SmartCache is a semantic cache for LLMs, optimizing performance and reducing costs."
}
```

### 2. Step-by-Step Cache
**Description**: Records the steps and thought processes for solving a problem.
**Use Cases**: Task planning, process guidance.
**Example**:
```json
{
  "type": "step_by_step",
  "task": "How to set up a web server?",
  "steps": [
    "Install the web server software.",
    "Configure the server settings.",
    "Deploy the web application.",
    "Monitor and maintain the server."
  ]
}
```

### 3. Code Cache
**Description**: Stores code snippets, functions, and other code-related content.
**Use Cases**: Programming assistants, code generation.
**Example**:
```json
{
  "type": "code",
  "function_name": "sort_array",
  "code": "def sort_array(arr): return sorted(arr)"
}
```

### 4. API Cache
**Description**: Stores API requests and their corresponding responses.
**Use Cases**: API gateways, interface caching.
**Example**:
```json
{
  "type": "api",
  "endpoint": "/get-user-info",
  "request": {"user_id": 123},
  "response": {"name": "John Doe", "age": 30}
}
```

### 5. Few-shot Cache
**Description**: Stores a few examples to facilitate model inference.
**Use Cases**: Few-shot learning, example-driven generation tasks.
**Example**:
```json
{
  "type": "few_shot",
  "examples": [
    {"input": "Translate 'hello' to French", "output": "bonjour"},
    {"input": "Translate 'goodbye' to French", "output": "au revoir"}
  ]
}
```

### 6. Entity-based Cache
**Description**: Caches data based on entity categories such as time, location, and person.
**Use Cases**: Entity recognition, context-aware caching.
**Example**:
```json
{
  "type": "entity_based",
  "entity_type": "location",
  "entity_value": "New York",
  "data": {
    "description": "New York is a state in the Northeastern United States."
  }
}
```

## Summary

SmartCache's cache types provide a comprehensive and versatile framework for optimizing LLM applications. By categorizing and storing data in these defined types, SmartCache ensures efficient retrieval and management of cached content, leading to improved performance and reduced operational costs. Each cache type is designed to address specific use cases, making SmartCache a robust solution for a wide range of semantic search needs.

