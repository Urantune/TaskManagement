#  Simple URL Shortener

A lightweight and fast URL Shortener service built with **Go** and **Gin**, designed to create, manage, and track shortened links via RESTful APIs.

---

##  Tech Stack

| Component | Technology |
|--------|-------------|
| Programming Language | **Go 1.25** |
| Web Framework | **Gin** |
| API Type | **RESTful API** |
| Use Case | **Shorten long URLs into short codes** |

---

##  Features

- Create short links from long URLs  
- Redirect using short codes  
- View all created links  
- Track statistics for each link  
- Create short links via browser or API  

---

##  API Endpoints

###  Create a Short Link (API)

PowerShell:
```powershell
curl http://localhost:8080/api/short `
-Method POST `
-Headers @{ "Content-Type" = "application/json" } `
-Body '{ "url": "https://your-long-link.com" }'
