# 🚀 logger app
Contains 2 main services: One *spams* the other while the other *logs* the messages it received.

🌐 First service: **Logger**
- Every time the route `/ping` is being called with an *http post* request, it adds a log of the message in the request body.
- Displays all the logs from its database.

💻 Second service: **Pinger**
- Pings the *logger service* and sends *http post* requests with a random message every 5 seconds to that `/ping` route.

---

### 🏁 Running the app ***(locally, using minikube and macOS)***

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

5. Visit [localhost:8080/pings](http://localhost:8080/pings) on your browser.

---

### Project structure

```bash
📁 .
├── 📄 LICENSE
├── 📄 README.md
├── 📁 kubernetes
│   ├── 📄 db-deployment.yml                # deployment for the postgres container
│   ├── 📄 db-pvc.yml                       # persistent volume claim so we don't lose the data from the database
│   ├── 📄 db-service.yml                   # k8s service for the database
│   ├── 📄 logger-config.yml                # config map for the project's env vars
│   ├── 📄 logger-deployment.yml            # deployment for the logger service
│   ├── 📄 logger-service.yml               # k8s service for the logger
│   ├── 📄 pg-deployment.yml                # deployment for the pg container (optional, we don't use pg anyways)
│   └── 📄 pg-service.yml                   # k8s for the pg service (optional, we don't use pg anyways)
├── 📁 logger-service
│   ├── 🐳 Dockerfile                       # Dockerfile for building the logger image
│   ├── 📁 db
│   │   └── 📄 db.go                        # database module for the logger app
│   ├── 🐙 docker-compose.yml               # docker-compose for local development, run with this if you don't want to use k8s
│   ├── 📄 go.mod                           # go module
│   ├── 📄 go.sum                           # imports for the logger app
│   ├── 📄 main.go                          # main for the logger service
│   └── 📁 ping            
│       └── 📄 ping.go                      # data structure for the ping
└── 📁 pinger-service
    ├── 📄 go.mod                           # go module
    ├── 📄 main.go                          # main for the pinger service
    └── 📁 messages
        └── 📄 messages.go                  # 20 random hardcoded messages
```
