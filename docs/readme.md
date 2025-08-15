# Mock Endpoints Service

This service is a mock endpoints service that provides mock endpoints for the following endpoints:

## HTTP REST

| Purpose                                           | Method | Path             | Description          |
| ------------------------------------------------- | ------ | ---------------- | -------------------- |
| List                                              | GET    | `/<path>`        | Get a list of items  |
| [Search](#search-pathsearch-pagination-supported) | GET    | `/<path>/search` | Search for items     |
| Create                                            | POST   | `/<path>`        | Create a new item    |
| Get                                               | GET    | `/<path>/:id`    | Get an item by ID    |
| Update                                            | PUT    | `/<path>/:id`    | Update an item by ID |
| Delete                                            | DELETE | `/<path>/:id`    | Delete an item by ID |

### Search `/<path>/search` (pagination supported)

When using the search endpoint, the following parameters are supported:

| Parameter | Type     | Description  | Required     |
| --------- | -------- | ------------ | ------------ |
| `query`   | `string` | Search query | **Required** |
| `page`    | `number` | Page number  |              |
| `size`    | `number` | Page size    |              |

Request body:
```json
{
  "query": "string",
  "page": 0,
  "limit": 0
}
```

## gRPC 

## WebSocket

## WebRTC