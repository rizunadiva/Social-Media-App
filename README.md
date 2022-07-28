
<div id="top"></div>
<!-- PROJECT LOGO -->
<br/>
<div align="center">

  <h3 align="center">Project-2 "Sosial Media App"</h3>

  <p align="center">
    Project ke 2 RESTful API menggunakan Framework Echo Golang dan Clean Architectures
    <br />
    <a href=""><strong>Kunjungi kami »</strong></a>
    <br />
  </p>
</div>

<!--- ABOUT THE PROJECT --->
### 💻 &nbsp;About The Project

   
Dilengkapi dengan berbagai fitur yang memungkinkan user untuk mengakses data yang ada didalam server, mulai dari membuat akun hingga melakukan postingan/news 
Adapun fitur yang ada dalam RESTful API kami antara lain :
<div>
      <details>
<summary>🙎 User</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
 User dapat membuat Akun dan Login, agar mendapat legalitas untuk mengakses berbagai fitur lain di aplikasi ini. 
 Terdapat juga fitur Update untuk mengedit data yang berkaitan dengan user, serta fitur delete jika user menginginkan akunnya dihapus.
 
<div>
  
| Feature User | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /users  | - | NO | Melakukan proses registrasi user |
| POST | /login | - | NO | Melakukan proses login user |
| GET | /users | - | YES | Mendapatkan informasi akun user yang sedang login |
| PUT | /users | - | YES | Melakukan update informasi akun user yang sedang login | 
| DEL | /users | - | YES | Menghapus akun user yang sedang login |

</details>  

<details>
<summary>🎉 &nbsp;News</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
  
User dapat memposting news sendiri yang bisa dikomentari oleh orang lain. Beberapa fitur yang lain ialah User dapat mengupdate dan menghapus news yang ia buat,
  
| Feature news | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| GET | /news  | - | NO | Mendapatkan informasi seluruh news yang ada |
| GET | /news/:idnews | idnews | NO | Mendapatkan informasi news secara detail melalui id news |
| GET | /news/mylists | - | YES | Mendapatkan informasi seluruh news user yang ia selenggarakan |
| POST | /news | - | YES | YES | Membuat news baru |
| DELETE | /news/:idnews | idnews | YES | Melakukan delete news yang diselenggarakan oleh user berdasarkan id news |
| PUT |  /news/:idnews | idnews | YES | Melakukan update news tertentu yang diselenggarakan oleh user berdasarkan id news |

</details>

<details>
<summary>💬 &nbsp;Comment</summary>
  
  <!---
  | Command | Description |
| --- | --- |
  --->
Comment merupakan fitur dimana user dapat memberikan comment/komentar pada event yang tersedia dan komentar tersebut dapat dilihat juga oleh user yang lain.

| Feature comment | Endpoint | Param | JWT Token | Fungsi |
| --- | --- | --- | --- | --- |
| POST | /comments | - | YES | Memberikan comment/komentar pada event yang ada |
| GET | /comments/:idEvent | idEvent | YES | Menampilkan comment/komentar yang ada pada suatu event berdasarkan id event |

</details>


<details>
<summary>📈&nbsp;ERD</summary>
<img src="https://github.com/rizunadiva/Social-Media-App-Group-5/blob/deployment/erd/erd.jpg">
</details>

<!-- CONTACT -->
### Contact

[[Ahmad Reski]](https://github.com/reski-id)

[[Rizuana]](https://github.com/rizunadiva)

[[Artby]](https://github.com/)


<p align="center">:copyright: 2022 | Group 5</p>
</h3>
