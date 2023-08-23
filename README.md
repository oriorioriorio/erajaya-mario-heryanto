# erajaya

Ini adalah aplikasi sederhana untuk mengelola dan menampilkan product.

TECH STACK
- GO
- Mysql
- Redis
- Docker

Architecture : 
- Clean code architecture -> pemisahan antara controller, usecase/library, dan repository
- Hexagonal architecture -> dependency injection (repository, library, dll) menggunakan interface agar tidak bergantung pada tipe data tertentu. maupun jenis database yang dipakai, sehingga jika kedepannya ada perubahan jenis database yang dipakai tidak mengganggu dari sisi library maupun usecasenya (bisnis logic).
 
 
langkah untuk menjalankan aplikasi:
- import .env dari .env.example

- Manual Setting
1. sediakan mysql, redis, go
2. set config untuk connection db
3. run query pada migration.sql
4. jalankan aplikasi

- Docker-Compose
1. sediakan docker
2. docker compose build --no-cache  
3. docker compose up 


API Collection
- import postman collection