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


# pengembangan
