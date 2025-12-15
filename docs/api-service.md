## API Service

The API Service is responsible for accepting client requests, validating input,
persisting request data, and publishing domain events.

Once an event is successfully published, the service immediately returns an
event identifier to the client. This identifier can be used to query the
processing status asynchronously.

The API Service does not perform business processing itself. Instead, it
guarantees that accepted requests are safely recorded and will be processed
by downstream services that subscribe to the published events.
