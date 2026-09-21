# IT APP
It app adalah aplikais untuk melakukan monitoring terhadap para pegawai dengan cara mengelola CR mulai dari pembuatan dan update dokumen,review dokumen,revisi dokumen,approve dokumen,update status dan juga delete

# arsitektur Domain Driven Design
domain/
berisi aturan bisnis inti,entity ,value object ,request object,response object dan interface.framework dan database tidak berada di layer ini

application/
berisi aturan usecase aplikasi yang menerima request dari handler dan menjalankan validasi dan alur bisnis 

infrastructure/
berisi implemetasi seperti koneksi postgreSql,repository gorm,logger dan notification

presentation/
berisi http handler dan routing untuk membaca dan memanggil lalu mengubah nya menjadi hasil json

container/
berisi depedency injection yang menyusun semua repository,service,storage,handler dan notifier

# alur 
1. request dimulai dari presentation handler yang mengubah request http menjadi request domain ke service
2. file application doc_cr.service menjalankan usecase 
3. service memanggil repository melalui interface domain 
4. service menyimpan data menggunakan gorm ke postgre


# pengembangan
1. aaturan bisnis di letakan di domain
2. handler hanya berisi request dan response
3. application tidak boleh langsung menyentuh gorm
4. database hanya boleh berada di infrastructure/persistance
5. depedency baru harus di tulis melalui container.go


# instalasi
1. cloning
```bash
git clone
```
cd itapp
2. download depedency
```bash
go mod download
```
3. sesuaikan env.example
4. jalankan 
```bash
go run main.go
```
5. server berjalan pada 
```bash
http://localhost:8280
```
6. dokumentasi API berjlana pada
```bash
http://localhost:8280/swagger/index.html
```