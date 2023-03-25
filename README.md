# Simple Scholarship App

Repo ini memuat code dari wesbite pendaftaran beasiswa sederhana. Untuk dapat menjalankan program ini dapat mengikuti intruksi berikut
## Run Locally For Setup Server


#### Clone the project

```bash
  git clone https://github.com/rezairfanwijaya/simple-scholarship-registration.git
```

#### Go to the project directory

```bash
  cd simple-scholarship-registration/backend
```


#### Setup ENV
##### Rename .env.example file menjadi .env dan sesuaikan dengan database yang dipilih
##### contoh :
```bash
DATABASE_USERNAME = "root"
DATABASE_PASSWORD = ""
DATABASE_HOST = "127.0.0.1"
DATABASE_PORT = "3306"
DATABASE_NAME = "serkom"
SECRET_KEY = "3465837DVBJHFB3764SHNBVN_89"
```

#### Run application
##### Terdapat executable file bernama `main.exe` dan jalankan file tersebut sehingga terbuka terminal. Jangan menutup terminal selama program akan dijalankan.


![tutor1](https://user-images.githubusercontent.com/87264553/227724073-3f2d6a94-5de9-4461-9f65-4d49ef1dc099.png)
![tutor2](https://user-images.githubusercontent.com/87264553/227724108-cd45bc91-7538-4de4-a849-3a94830488e4.png)



## Data Admin 
##### Untuk masuk kedalam dashboard silahkan masukan credentials berikut
```bash
username : "admin@gmail.com"
password : "12345"
``` 
## Data Email 
##### Terdapat beberapa email yang dapat digunakan untuk melakukan registrasi beasiswa yaitu
```bash
agung@gmail.com
angelica@gmail.com
babas@gmail.com
rustam@gmail.com
mawar@gmail.com
``` 

## Run Locally For Setup Client
#### Go to the project directory

```bash
  cd simple-scholarship-registration/frontend
```

#### Get Dependency
```bash
  npm install
```

#### Serve Application
```bash
  npm start
```
![tutor3](https://user-images.githubusercontent.com/87264553/227725765-c979ec7b-4deb-47dd-b8e4-603bfc34a6ae.png)

