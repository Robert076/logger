# 🚀 logger app
Contains 2 main services: One spams the other while the other logs the messages it received.

🌐 First service: **Logger**
- Every time the route `/ping` is being called with an *http post* request, it adds a log of the message in the request body.
- Displays all the logs from its database.

💻 Second service: **Pinger**
- Pings the *logger service* and sends http post requests with a random message every 5 seconds


