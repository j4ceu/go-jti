# go-jti
Repository untuk tugas Test Tech JTI

## How To Run
1. Sesuaikan .ENV untuk menyesuaikan PostgreSQL di local
DB_HOST = 
DB_PORT = 
DB_NAME = 
DB_USER = 
DB_PASSWORD = 

2. Buka folder root dan jalankan dengan command line berikut
```
go run cmd/main.go
```

3. Buka pada browser http://localhost:8000

## Notes
1. Database menggunakan PostgreSql
2. Menggunakan Encrypt dan Decrypt dengan aes-256-cbc
3. Menggunakan Sign in with Google
