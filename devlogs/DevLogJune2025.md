[Delta vs snapshot streaming](https://docs.anthropic.com/en/docs/build-with-claude/streaming#delta-vs-snapshot-streaming)

Maybe a buffer is enoug? Get the stream, push the data from the stream to the buffer, then send the data from the buffer to a channel of CustomResponse? -> Go for this, dont overthink

content_block_start events for tool use blocks
content_block_delta events with accumulated partial JSON

`ContentBlock` as a unified interface for different content block types

sqlite3 as a lightweight option to store conversation

Flag to choose between snapshot response and streaming response?
