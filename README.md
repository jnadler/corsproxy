# corsproxy
Simple golang proxy for local development against APIs that do not support CORS

Just `go install` and `corsproxy 9200 http://elasticsearch.whateverdomain.com:9200/`

Now you can point your local browser-based application at localhost:9200 for development, and access the remote API even though it may not serve the appropriate CORS headers. 

I don't need to tell you that this ain't for production use, do I?

I mean really.   Behave, children.
