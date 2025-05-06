# 🚀 logger app
Contains 2 main services: One *spams* the other while the other *logs* the messages it received.

🌐 First service: **Logger**
- Every time the route `/ping` is being called with an *http post* request, it adds a log of the message in the request body.
- Displays all the logs from its database.

💻 Second service: **Pinger**
- Pings the *logger service* and sends *http post* requests with a random message every 5 seconds to that `/ping` route.

---

🏁 Running the app ***(locally, using minikube and macOS)***

1. Clone the repository:

```bash
git clone https://github.com/Robert076/logger.git
```

2. Start your local *minikube* server:

```bash
minikube start
```

3. Run the *Kubernetes* deployment:

```bash
kubectl apply -f kubernetes/
```

4. Port forward the logger port from inside the minikube

```bash
kubectl port-forward service/logger-service 8080:8080
```

5. Visit [http://localhost:8080/pings](localhost:8080/pings) on your browser.
