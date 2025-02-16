# 🚀 URL Shortener in Golang  

A simple URL shortener built using **Golang** and **Gin**, storing short URLs in memory using a Go map.

## 📌 Features  
✅ Shorten long URLs into unique short links  
✅ Redirect users from short links to the original URLs  
✅ Store URLs in memory using a **Go map**  
✅ REST API built with **Gin**  

---

## 🛠️ Installation  
### 1️⃣ Clone the Repository  
```sh
git clone https://github.com/yourusername/go-url-shortener.git  
cd go-url-shortener
```

### 2️⃣ Install Dependencies
```sh
go mod tidy
```

### 3️⃣ Run the Application
```sh
go run main.go
```

---

## 📡 API Endpoints
### 1️⃣ Shorten a URL
Endpoint: `POST /shorten`
Request Body:
```sh
{
    "original_url": "https://example.com"
}
```
Response:
```sh
{
    "short_url": "http://localhost:8080/abc123"
}
```
### 2️⃣ Redirect to Original URL
Endpoint: `GET /:shortCode`
Example: Visiting `http://localhost:8080/abc123` redirects to `https://example.com`.

## 📝 Notes
URLs are stored in memory, meaning they disappear when the app stops.
For persistent storage, consider PostgreSQL, SQLite, or a file-based database.