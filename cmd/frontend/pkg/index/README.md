# Notification Index

This package provides Bleve-based search indexing for notifications in the trading system.

## Features

- **Full-text search** on notification messages
- **Filtered search** by notification type, target, broker, account, stage
- **Date range queries** for time-based searches
- **Efficient indexing** with Bleve search engine

## Components

### Index Mapping
- `NewIndexMapping()` - Creates the Bleve index mapping configuration
- `NewNotificationDocumentMapping()` - Defines field mappings for notification documents

### Notification Entry
- `NotificationEntry` - Represents a notification document in the index
- Contains all searchable fields: type, message, target, broker, account, stage, timestamps

### Indexing
- `NotificationIndexer` - Adds/removes notifications from the index
- `NewNotificationIndexer(index)` - Creates a new indexer instance

### Searching
- `NotificationIndexSearcher` - Searches the index with various filters
- `NewNotificationIndexSearcher(bleveIndex)` - Creates a new searcher instance

### Finding
- `NotificationFinder` - High-level interface that combines searching with data retrieval
- `NewNotificationFinder(getter, searcher)` - Creates a finder that returns full notification objects

## Usage Example

```go
// Create index
mapping := index.NewIndexMapping()
bleveIndex, err := bleve.New("notifications.bleve", mapping)
if err != nil {
    return err
}

// Create components
indexer := index.NewNotificationIndexer(libindex.NewIndex(bleveIndex))
searcher := index.NewNotificationIndexSearcher(bleveIndex)
finder := index.NewNotificationFinder(notificationGetter, searcher)

// Add notification to index
err = indexer.Add(ctx, notification)

// Search notifications
notifications, err := finder.Find(ctx, tx, 
    nil, // from date
    nil, // until date
    core.NotificationTypes{core.NotificationType("signal")}, // types
    nil, // targets
    nil, // brokers
    nil, // accounts
    nil, // stages
    "EURUSD", // query string
    10, // limit
)
```

## Search Capabilities

- **Date Range**: Filter by creation date
- **Type Filter**: Filter by notification type (signal, test, etc.)
- **Target Filter**: Filter by notification target (Discord channel, etc.)
- **Broker Filter**: Filter by broker identifier
- **Account Filter**: Filter by account identifier
- **Stage Filter**: Filter by trading stage (demo, live, etc.)
- **Text Search**: Full-text search in notification messages
- **Sorting**: Results sorted by creation date (newest first)

## Testing

Run tests with:
```bash
go test ./pkg/index/...
```

Tests include:
- Basic indexing and searching
- Type-based filtering
- Message content search
- Date range queries
- Index removal