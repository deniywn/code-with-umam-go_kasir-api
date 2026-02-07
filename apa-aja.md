# apa aja
## 3/2/2026
Toni 8:15 PM
Interface{} diganti any-any mas 🙈

Wildan Febrian 8:15 PM
any-any siapa mas?

deepseek-r1 8:15 PM
mau dong

Toni 8:16 PM
Harap fokus, bapak2 hadirin

Agung 8:16 PM
nah kek anu

Heru Oktafian, ST., CTT 8:16 PM
sundul Mas @Agung

Agung 8:16 PM
@Heru Oktafian, ST., CTT revo ku rong payu, gak wani due any2

deepseek-r1 8:17 PM
tanya: kenapa argsnya digabung?

Muhamad Putra Satria 8:17 PM
kenapa gapake printf aja kak?

deepseek-r1 8:17 PM
Di append maksudnya?

tommy h 8:17 PM
query param nya name ya..bukan search

Ahmad Eko Kurniawan 8:18 PM
like ini harus sama apa ada pembobotan search misal typo masih bisa kueri

M Rizal Hidayat 8:19 PM
pak sy masih bingung untuk GO jadi frontend wkwk

Rendy Putra 8:19 PM
$1 di query darimana ya?

ibnɘꟻ 8:17 PM
Kalo mau pake prepared statements, apakah perlu install library kah? Khawatir kena sql injection.

Muhammad Tri Wibowo 8:20 PM
aman mas, itu termasuk parameterized query, secara keamanan sudah seperti prepared statement

ibnɘꟻ 8:21 PM
Owh.. Makasih mas.

Zannns 8:20 PM
ada yang lebih suka pake sqlx ga?

Muhammad Tri Wibowo 8:21 PM
saya dulu pakai itu

Irwan Andriawan 8:19 PM
apakah ada semacam builder query di go? meminimalisir salah tulis query.

fatchurrosii 8:19 PM
Ada bisa pake goqu atau huandu

Eka Anre Sangsya 8:19 PM
coba gorm

fatchurrosii 8:20 PM
https://github.com/doug-martin/goqu 

https://github.com/huandu/go-sqlbuilder
@Eka Anre Sangsya kalo golang banyak yang bisa bagusnya pake sql builder aja daripada orm

Irwan Andriawan 8:21 PM
goqu jadi inget dragon ball :D

M Rizal Hidayat 8:16 PM
sorry itu ILIKE apa ya ?

Alvin Mantovani 8:17 PM
CMMIW kalau gak salah itu untuk nyari data sesuai kata yang dicari

bedanya LIKE itu harus exactly match
kalau ILIKE itu contain

Sukron Ch 8:19 PM
Itu keyword specific di postgre, di db engine lain belum tentu ada kek di mysql

alfi 8:19 PM
Apakah di go best practice menggunakan raw query seperti ini?

Alvin Mantovani 8:22 PM
mungkin bisa dibuat class nya bang, kaya buat fungsi sendiri untuk generate raw query.

Sans 8:22 PM
CREATE TABLE IF NOT EXISTS transactions (
id SERIAL PRIMARY KEY, total_amount INT NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transaction_details (
id SERIAL PRIMARY KEY,
transaction_id INT REFERENCES transactions (id) ON DELETE CASCADE, product_id INT REFERENCES products (id), quantity INT NOT NULL, subtotal INT NOT NULL
);

aziz 8:20 PM
kenapa []interface{}, maksudnya kenapa pake type interface mas?

Muhammad Tri Wibowo 8:21 PM
karena parameter itu bisa banyak tipe, boolean, string, integer, dll

Zannns 8:22 PM
tambahin, dikarenakan type parameter itu ga nentu, jadinya pake interface{}, btw ini sama kayak any di ts

aziz 8:23 PM
baru tau, thanks guys

Sihar William Jubilant Simbolon 8:23 PM
where itu tadi sudah binding parameter kah?
biar mencegah sql injection?

Muhammad Tri Wibowo 8:23 PM
sudah

ariefbayu 8:24 PM
apa pertimbangan Transaction dan TransactionDetails dalam 1 file model/transaction.go?

Eka Anre Sangsya 8:24 PM
biar gak banyak file, karena masih 1 konteks :D
mau dipisah juga gak masalah

Sihar William Jubilant Simbolon 8:25 PM
lebih ke selera sih,
Plusnya dengan pendekatan header dan detail
kalau mau nambah info lain, cukup tambahkan di header, tidak perlu duplicate sebanyak detailnya

HONOR ABR-NX1 8:29 PM
Ctrl + d aja mas
Ntar stiap yg text nya sama ke detect
Klo banyak tinggal ctrl + d trs menerus sampe blink nya nge cover smua 

Ata 8:25 PM
Kenapa modelnya TransactionDetails bukan TransactionDetail?

Muhammad Tri Wibowo 8:30 PM
nah, ini setuju biar konsisten

Arif 8:28 PM
Ngubah banyak sekaligus tadi pakai shortcut apa ya?

Rendy Putra 8:28 PM
ctrl + d

aziz 8:28 PM
CTRL + D

deepseek-r1 8:28 PM
Cmd + D

Arif 8:29 PM
Sip thank you

Rido Rachmatan 8:29 PM
select wordnya, lalu ctrl d

Syamsul Arifin 8:32 PM
https://code.visualstudio.com/docs/reference/default-keybindings

Aghnat 8:31 PM
Maaf ada link youtube nya? Saya kelempar barusan

Putu Nadya Putri Astina 8:32 PM
https://www.youtube.com/watch?v=5rJ5g8knuRU ini kak

Aghnat 8:33 PM
Makasih

Rido Rachmatan 8:34 PM
kalau tadi di models transaction ada banyak type, dengan kita tulis returnnya (*models.Transaction) semua types ini bisa keakses kah?

M Rizal Hidayat 8:31 PM
btw mau nanya, sy masih blm ada gambaran klw untuk di hostingkan / di server itu apakah seperti php folder projectnya atau bagaimana ya untuk go ini ? soalnya di zeabur langsung dari github

deepseek-r1 8:32 PM
Go perlu vps, yang visa ssh. Nanti go projectnya di build dan akan ada binary atau .exe nya, ini di run di vps

M Rizal Hidayat 8:32 PM
oo ada buildnya bentuk exe toh, berarti tiap ada perubahan di build lagi ya ?

deepseek-r1 8:33 PM
iya

winyata 8:33 PM
go tidak bisa disimpan di VPS. mesti  VPS karena dia hasil compile nya binary...

Heru Oktafian, ST., CTT 8:33 PM
@M Rizal Hidayat kalo gak mau build bisa langsung jalankan go run main.go

deepseek-r1 8:33 PM
Atau bisa pake docker, tapi tetap ngerun perlu dibuild dulu

HONOR ABR-NX1 8:33 PM
Klo pengen gak build build trs bisa pake air
Stiap ada prubahan lgsg ke build ulang

fatchurrosii 8:33 PM
@M Rizal Hidayat iya betul pake go build -c  Nama-build  
Nanti untuk jalaninnya bisa pake systemd / supervisor
Atau kalo pake docker bisa dari dockernya untuk jalanin golangnya

HONOR ABR-NX1 8:35 PM
Dulu sih aku buat jalanin di background pake pm2, manggil air nya. 

Sehingga jalan di bg tapi auto rebuild

M Rizal Hidayat 8:35 PM
wah mantap terima kasih semua jawabannya sy catat ❤️

winyata 8:35 PM
@M Rizal Hidayat kalau develop di windows jadi .exe file. Kalau di Linux atau mac jadi binary file yg tinggal di panggil dari cmd line

HONOR ABR-NX1 8:35 PM
Atau pake docker jg bisa. 
Jalanin di bg nya pake docker compose up -d
---
btw mau nanya, sy masih blm ada gambaran klw untuk di hostingkan / di server itu apakah seperti php folder projectnya atau bagaimana ya untuk go ini ? soalnya di zeabur langsung dari github

deepseek-r1 8:32 PM
Go perlu vps, yang visa ssh. Nanti go projectnya di build dan akan ada binary atau .exe nya, ini di run di vps

M Rizal Hidayat 8:32 PM
oo ada buildnya bentuk exe toh, berarti tiap ada perubahan di build lagi ya ?

deepseek-r1 8:33 PM
iya

winyata 8:33 PM
go tidak bisa disimpan di VPS. mesti  VPS karena dia hasil compile nya binary...

Heru Oktafian, ST., CTT 8:33 PM
@M Rizal Hidayat kalo gak mau build bisa langsung jalankan go run main.go

deepseek-r1 8:33 PM
Atau bisa pake docker, tapi tetap ngerun perlu dibuild dulu

HONOR ABR-NX1 8:33 PM
Klo pengen gak build build trs bisa pake air
Stiap ada prubahan lgsg ke build ulang

fatchurrosii 8:33 PM
@M Rizal Hidayat iya betul pake go build -c  Nama-build  
Nanti untuk jalaninnya bisa pake systemd / supervisor
Atau kalo pake docker bisa dari dockernya untuk jalanin golangnya

HONOR ABR-NX1 8:35 PM
Dulu sih aku buat jalanin di background pake pm2, manggil air nya. 

Sehingga jalan di bg tapi auto rebuild

M Rizal Hidayat 8:35 PM
wah mantap terima kasih semua jawabannya sy catat ❤️

winyata 8:35 PM
@M Rizal Hidayat kalau develop di windows jadi .exe file. Kalau di Linux atau mac jadi binary file yg tinggal di panggil dari cmd line

HONOR ABR-NX1 8:35 PM
Atau pake docker jg bisa. 
Jalanin di bg nya pake docker compose up -d

Romi Alief Rahman 8:42 PM
jadi ingat pascal ..

Indra A 8:42 PM
Ijin leave duluan ya mas @Muhammad Zuhrul Umam , ada kepentingan lain.. 🙏
Nanti tak lanjut dari recording

deepseek-r1 8:43 PM
mas itu apa gk jadi n+1 query?

deepseek-r1 8:44 PM
Setauku kalo db query jangan sampe ditaruh di dalam for/loop

ibnɘꟻ 8:44 PM
Mungkin nanti di task kita, bisa implementasi whereIn() query. 😅

21022012_SHOFYAN 8:46 PM
@deepseek-r1 ini masih naive dulu, bentar lagi juga dibenerin

Gede Etika 8:44 PM
Ud mulai puyeng

M Rizal Hidayat 8:45 PM
iya betul wkwk

Romi Alief Rahman 8:46 PM
lulus dari sini, langsung punya portofolio ...

Rizky Hidayattulloh 8:42 PM
bang tanya, misalkan ada transaction repository dan payment repository, nah saat create transaction kan butuh panggil kedua nya, itu cara handle db tx nya gimana ya ?

21022012_SHOFYAN 8:45 PM
bang tanya, misalkan ada transaction repository dan payment repository, nah saat create transaction kan butuh panggil kedua nya, itu cara handle db tx nya gimana ya ?
bisa dinaikin ke level service, jadi layer service yang buka tx.begin dan tx.commit nya

Rizky Hidayattulloh 8:46 PM
brti instance tx nya di passing ke repo nya ya dari service ?

Rido Rachmatan 8:46 PM
%d dari mana ya

aziz 8:47 PM
%d itu nanti nampung item.ProductID

M Rizal Hidayat 8:47 PM
iya betul, %d, $i, make itu dari mana ya tau nya ??

gabriel 8:47 PM
%d itu untuk output angka int dari productID itu

Rido Rachmatan 8:48 PM
thanks mas aziz & gabriel

ibnɘꟻ 8:48 PM
%d untuk formatting integer, kalo string pake %s. Tergantung id nya pake tipe data apa.

Z.ai 8:48 PM
Kalau stok nya 0 saat transaksi gimana tuh?

harusnya sih sblum update harus cek stocknya apakah lebih/sama dengan stock yg dibutuhkan untuk transaksi

Pandu Satria Nur Ananda 8:50 PM
Bisa ditambahin pengecekan stok per item waktu looping. Kalo nol, return error

bang tanya, misalkan ada transaction repository dan payment repository, nah saat create transaction kan butuh panggil kedua nya, itu cara handle db tx nya gimana ya ?

21022012_SHOFYAN 8:45 PM
bang tanya, misalkan ada transaction repository dan payment repository, nah saat create transaction kan butuh panggil kedua nya, itu cara handle db tx nya gimana ya ?
bisa dinaikin ke level service, jadi layer service yang buka tx.begin dan tx.commit nya

Rizky Hidayattulloh 8:46 PM
brti instance tx nya di passing ke repo nya ya dari service ?

Muhammad Tri Wibowo 8:52 PM
@Rizky Hidayattulloh iyes

bisa passing lewat context, tapi di latihan ini gak bahas context, jadi mungkin jadiin parameter tx saja

Pandu Satria Nur Ananda 8:56 PM
Kalo menurut saya, tx tetep di level repo karena tx dipake untuk handle db..untuk trx service akan handle bisnis logicnya, akan panggil trx repo dan payment repo.

Zannns 8:59 PM
kenpa ga pake http.handleFunc("POST /api/..") aja mas?

Gede Etika 8:56 PM
Saya sebagai pengguna c# pusing liat syntax golang

fatchurrosii 8:58 PM
:= deklarasi sama assignment
= assignment aja

Sukron Ch 8:59 PM
Saya dari java apalagi 🗿, kebiasaan disiplin make dto per layer (model/domain harus bersih) 

Muhammad Tri Wibowo 8:59 PM
@Gede Etika saya mah nubie om, cuma ini bahasa masih 1 rumpun masih enak kok
kalau beda rumpun baru berasa

Gede Etika 8:59 PM
Mesti banyak belajar lagi 😂

Ohh, saya kira dev go ngindarin dto karena terlalu verbose 😅 @Muhammad Tri Wibowo 

Muhammad Tri Wibowo 9:05 PM
hexagonal-go-grpc/internal/adapters/primary/http/dto/customer.go at main · brainplusplus/hexagonal-go-grpc
@Sukron Ch go mau dibikin complex jg bisa, lagi2 balik yang bikin boilerplate pertama kali

Zannns 8:59 PM
kenpa ga pake http.handleFunc("POST /api/..") aja mas?

Boy 9:00 PM
banyak jalan menuju Roma
pasti ada trade-off nya, cuma saya juga gak tahu

wow... no error langsung jalan

Kurniawan Sudirman 9:04 PM
kenapa logic prosesnya di hanlder ya bukan di service?

faqih yugo susilo 9:03 PM
Loh ada pak Wisnu

Wisnu | PT Webhade MDS 9:04 PM
ada dong hehe

faqih yugo susilo 9:05 PM
Saya murid php nya hehe

faqih yugo susilo 9:06 PM
Ada project golang pak?

Wisnu | PT Webhade MDS 9:07 PM
@faqih yugo susilo belum, kan masih belajar ini, abis ini ntr siap nyari project golang haha

faqih yugo susilo 9:08 PM
Pak Umam, ngga ada handle response nya? Karena return ga selalu sama dengan model

Catur 9:08 PM
Dah mo maghrib leave duluan ya

Muhamad Putra Satria 9:09 PM
ijin nanya kak, kalau ada 2 user yang transaksi bersamaan kak?

Pak Umam, ngga ada handle response nya? Karena return ga selalu sama dengan model

Catur 9:08 PM
Dah mo maghrib leave duluan ya

Muhamad Putra Satria 9:09 PM
ijin nanya kak, kalau ada 2 user yang transaksi bersamaan kak?

Agung 9:09 PM
ini challenge nya nanti hahahaa

Muhamad Putra Satria 9:09 PM
uwaw

faqiah mudo 9:09 PM
Izin leave dulu mas umam, nnti saya masuk lg

aran mas, kalo baca komentar suaranya yang keras mas, kita juga pengen tau mas umam lagi baca komen yang mana biar gak bingung :D
makasih

Catur 9:08 PM
Dah mo maghrib leave duluan ya

Pandu Satria Nur Ananda 9:10 PM
Om catur dimana? Di sini sudah jam 9 malam

Fadli 9:10 PM
udah di handle sama db transaction bukan ya itu?

Frenki Herlambang 9:10 PM
Belajar GO Lang dengan Hanan Attaki

Wisnu | PT Webhade MDS 9:11 PM
ya kan dari sesi 1 begitu feel nya

aryandi 9:11 PM
programmer hijrah

Boy 9:11 PM
busyet

Zannns 9:11 PM
lah iya baru ngeh mirip suaranya ustadz hanan attaki

Zai 9:11 PM
Jadi kalau stock tinggal 1, siapa yang dapat?

Eka Anre Sangsya 9:11 PM
orang dalam :D

Boy 9:12 PM
war kalau ini

Muhamad Putra Satria 9:12 PM
alamakk

Agung 9:12 PM
kalo mie di seduh aja

Hafid Nur 9:12 PM
ketinggalan jauh ngodingnya, Mas. Nyimak tanya jawabnya aja

Sukron Ch 9:12 PM
Mas nanti ditambahin middleware sendiri g? 


udah di handle sama db transaction bukan ya itu?

Frenki Herlambang 9:10 PM
Belajar GO Lang dengan Hanan Attaki

Wisnu | PT Webhade MDS 9:11 PM
ya kan dari sesi 1 begitu feel nya

aryandi 9:11 PM
programmer hijrah

Boy 9:11 PM
busyet

Zannns 9:11 PM
lah iya baru ngeh mirip suaranya ustadz hanan attaki

Zai 9:11 PM
Jadi kalau stock tinggal 1, siapa yang dapat?

Eka Anre Sangsya 9:11 PM
orang dalam :D

Boy 9:12 PM
war kalau ini

Muhamad Putra Satria 9:12 PM
alamakk

Agung 9:12 PM
kalo mie di seduh aja

Hafid Nur 9:12 PM
ketinggalan jauh ngodingnya, Mas. Nyimak tanya jawabnya aja

Sukron Ch 9:12 PM
Mas nanti ditambahin middleware sendiri g? 

M Rizal Hidayat 9:12 PM
ini kelas advance isinya wkwk

Bang Fadli 9:12 PM
udah cukup mas,
cukup ngebul..

Zannns 9:12 PM
bang hafid molor scrolling fesnuk kyknya

Romi Alief Rahman 9:13 PM
masih baru di go ...

Dimas Rafi Ramaharmuzi 9:13 PM
Banyak anggota yang sudah expert kayaknya

Boy 9:14 PM
iya bang, pada nyamar semua pura2 nubie

Rizky Hidayattulloh 9:13 PM
mas bisa minta contoh nya cara handle db tx di 2 repo yang berbeda pakai context ?

Alan muqilul atsarot 9:14 PM
Mas mau nanya, saya mencoba melihat structure clean code go project milik PZN, disitu ada entity dan model. Saya lihat isinya hampir sama.. kira2 perbedaannya apa ya antara keduanya?

Maftuhriza Yumida 9:14 PM
performance wise, kalau query di loop bakal membebani db kayaknya ya kalau requestnya banyak

Darm 9:14 PM
Mas Umam mau tanya, ketika create transaction ini business logicnya memang di repository layer ya mas? Seperti cek stok, hitung total, update stok

Zannns 9:15 PM
keyword for update itu cuman ada du postgre?

Wisnu | PT Webhade MDS 9:13 PM
ini baru create transaksi, belum bikin laporan / reporting transaksinya wkwk

winyata 9:14 PM
blm export excel :)

deepseek-r1 9:15 PM
Belom cara monetise nya hehe

winyata 9:15 PM
@deepseek-r1 nah ini yg penting 😄

Denny Rachmadi 9:16 PM
trigger ?

Agung 9:16 PM
mas umam minta di elaborate flow nya tadi  dari awal bikin urut nya

Muhammad Tri Wibowo 9:15 PM
pakai column version juga bisa

Muhammad Tri Wibowo 9:18 PM
optimistic locking

Nandang Gozali 9:18 PM
Sekedar saran mas @Muhammad Zuhrul Umam, bagi saya yg pemula saya lebih prefer untuk ngoding nya manual kaya tadi, sehingga setidaknya sedikit paham fungsi2 kode nya, kalau di sesi 2 kemarin yg tinggal copas jadinya ga ngerti
(curhatan pemula) 😅

Agung 9:18 PM
cocok bang

Muhammad Tri Wibowo 9:18 PM
iya lebih enak buat belajar

Wisnu | PT Webhade MDS 9:18 PM
API belajar sama mas Umam, frontend nya serahkan sama mas AI 🤣 sambil nyimak sambil vibe coding juga ini wkwk

pakai apaan ini om?

Google AI Studio

Sukron Ch 9:22 PM
Bagi prompt nya pak, saya belum mahir pakai AI 😅

Wisnu | PT Webhade MDS 9:29 PM
@Sukron Ch cuma gini doang kok 😅

create point of sales application, tablet view first, containing category, products, new transaction, search a product
the api will be using Golang

Ramond Julio 9:30 PM
Terima kasih Mas Umam

Irwan Andriawan 9:30 PM
izin leave ya. ada keperluan lain. terima kasih mas umam kelasnya hari ini.

Rido Rachmatan 9:30 PM
mas nanti bikin case dimana kita harus handle logic di service dongg

deepseek-r1 9:30 PM
Mas umam, trus service fungsinya buat apa ya? Selain konek ke repo?

Asep Sae Puloh 9:31 PM
Mas, kenapa logic nya ada di repository ya ?
Kenapa Begin Transactionnya gak di Service ya ?
Bukannya Repository itu harus menganut Patern SRP ya ?

feri 9:31 PM
contoh business logic paling simple apa mas umam ?

Sukron Ch 9:31 PM
Mas kalo buat respon apakah perlu bikin pointer ke model(kan mapping manual/buat baru)?, walaupun paham pointer tapi saya masih bingung implementasi dan kapan harus pakainya

Wahyu Candra Tama 9:31 PM
AFK dulu mas, besok lanjut via youtube, maturnuwun

jalil 9:31 PM
oke mas, masih agak bingung sama arsitektur nya aja sih
dulu saya ngerjain java, yang gendut itu di service, handler/controller nya bersih

Zannns 9:31 PM
mas, ini kan kode repositori nya n+1, nah di tulisannya nanti ada versi best practices nya gak?



Mas umam, trus service fungsinya buat apa ya? Selain konek ke repo?

M Rifki Aulia P 9:32 PM
buat bisnis logicnya mas, repo cuman bagian ngambil bahan, service yang mengolahnya

Wisnu | PT Webhade MDS 9:18 PM
API belajar sama mas Umam, frontend nya serahkan sama mas AI 🤣 sambil nyimak sambil vibe coding juga ini wkwk

Muhammad Tri Wibowo 9:21 PM
pakai apaan ini om?

Wisnu | PT Webhade MDS 9:21 PM
@Muhammad Tri Wibowo Google AI Studio

Sukron Ch 9:22 PM
Bagi prompt nya pak, saya belum mahir pakai AI 😅

Wisnu | PT Webhade MDS 9:29 PM
@Sukron Ch cuma gini doang kok 😅

create point of sales application, tablet view first, containing category, products, new transaction, search a product
the api will be using Golang

Sukron Ch 9:32 PM
Ohh, ok pak langsung dicoba @Wisnu | PT Webhade MDS 

Muhammad Tri Wibowo 9:32 PM
@Wisnu | PT Webhade MDS wah belum pernah pakai google ai studio, mantab sepertinya


external itu maksudnya apa mas Umam?

Asep Sae Puloh 9:34 PM
Oww berarti Repo itu untuk yang berhubungan ke Eksternal ya

Wildan Febrian 9:36 PM
mas bagaimana menggunakan fungsi yang ada di module/dependency lain? misalkan fungsi existsCategory() dari repo/service category namun mau dipake oleh domain product

Denny Rachmadi 9:36 PM
lebih ke external dependencies yah @Muhammad Zuhrul Umam ?

kalo notification kaya semacam buat kirim email gitu pantas nya taruh di folder apa ya mas? repo juga kah karena external?

Fitriningtyas 9:37 PM
mas, yang update_version tadi berarti nambahin kolom di tabel product ya mas?

Romi Alief Rahman 9:37 PM
jadi lebih mendingan manfaatin autocomplete nya kali ya mas ... daripada suruh ngerjain/vibe ...

Wisnu | PT Webhade MDS 9:37 PM
bener

Sukron Ch 9:38 PM
Autocomplete + commit message generator enak tu

ibnɘꟻ 9:39 PM
Kalau saya biasanya dipake buat bantu bikin go struct dari format json, agak bosen juga kalo manual. He he..

Zai 9:39 PM
kalau untuk helper, util, dan semacamnya yang bisa dipakai bersama, struktur folder nya bgmn mas?

Asep Sae Puloh 9:41 PM
request

Denny Rachmadi 9:41 PM
pernah baca kemarin. internal hanya diconsume di project ini saja. diluar folder internal itu reusable bisa dipake di project lain @Muhammad Zuhrul Umam CMIIW

Sirojul 9:41 PM
Mas Umam, kadang di arsitektur java, repository modul lain tidak boleh di inject ke service lain, karena nantinya tercecer, jadi dependencynya antar service, itu bagaimana mas?

5200411407-Ristu Aji Wijayanto 9:42 PM
Kalau saya analoginya repo sebagai keluar masuk data/ jual-beli lah, usecase sebagai logic businessnya, handler tetap sebagai resepsionis menerima request dan memberikan response

Sirojul 9:42 PM
repository A hanya boleh di inject ke service A, mas

Wildan Febrian 9:39 PM
brrti kita inject dependency satu persatu ke semua handler yg butuh fungsi tersebut mas? saya dari bahasa yg OOP bgt jdi agak kurang srek disini hehe

deepseek-r1 9:39 PM
Wah jadi penasaran, kalo OOP trus approach nya gimana?

jalil 9:40 PM
ada constructor dll

deepseek-r1 9:41 PM
berarti koneksi Dari handler -> service -> repo leeway constructor gitu? Atau gin?

jalil 9:42 PM
func New... itu kalo di OOP buat kayak gitu juga mas sebenere cuma agak beda aja

Wisnu | PT Webhade MDS 9:42 PM
mas klo unit test di golang kaya gimana ya?

Asep Sae Puloh 9:42 PM
Serahin AI

Asep Sae Puloh 9:43 PM
Biar coveragenya 90 %

Wildan Febrian 9:43 PM
kalo OOP konsepnya kita bisa manggil object sesuka hati, tpi klo di GO nih kyak "lu mo ngelakuin ini ya ini aja gausah macem-macem" gini yg saya tangkep wkw

Rido Rachmatan 9:43 PM
gambaran tugas selanjutnya mas umam, spill dikit hahah

bhre astrajingga 9:43 PM
om, ijin ya.. dah pedes nih mata...🥹

Wisnu | PT Webhade MDS 9:43 PM
oh nanti ada materi unit test nya ya, oke2 mas siap

mas, di golang pakai dto juga ga mas di layer service ? atau tetep nembak langsung ke  model ?

Agung 9:50 PM
aku pakai dto

Sukron Ch 9:50 PM
Gw pake bang

jalil 9:50 PM
okee, cari temen aja wkwk

Dimas Ady 9:50 PM
dto paling aman, klo ada many to many bisa puyeng golang nampilin datanya

Angga VB 9:51 PM
Ulang mas, putus putus di awal tadj

Syamsul Arifin 9:51 PM
ulang rekamannya aja di yt

danirosady zoom 9:52 PM
nanti di record putuus2 gak ya

Habibie (paidy) 9:52 PM
engga kok
hmm

Wisnu | PT Webhade MDS 9:52 PM
kalo di golang ada macem query builder gak mas kaya eloquent di laravel

Asep Sae Puloh 9:52 PM
GORM

kalo di golang ada macem query builder gak mas kaya eloquent di laravel

Asep Sae Puloh 9:52 PM
GORM

Wisnu | PT Webhade MDS 9:53 PM
@Asep Sae Puloh oh ada ya, mantap juga, lgsg riset deh

sy beneran nol besar id golang nih, kelamaan di php & python

fatchurrosii 9:53 PM
Ada kok 
https://github.com/huandu/go-sqlbuilder 
https://github.com/doug-martin/goqu 
Bisa pake salah satunya

Asep Sae Puloh 9:53 PM
Keluarga Lambo

deepseek-r1 9:54 PM
Gimana? Lambo lu warna apa?

Asep Sae Puloh 9:54 PM
Warna Sawit

Agung 9:55 PM
@Wisnu | PT Webhade MDS mencair

Rama Fajar Fadhillah 9:56 PM
bruh di laravel itu jatuhnya ORM bukan query builder

Wildan Febrian 9:43 PM
kalo OOP konsepnya kita bisa manggil object sesuka hati, tpi klo di GO nih kyak "lu mo ngelakuin ini ya ini aja gausah macem-macem" gini yg saya tangkep wkw

Remo 9:52 PM
kalo orang jawa bilang 

Go supports most OOP concepts using a different, more compositional approach. Go promotes a "favor object composition over class inheritance" design philosophy


---



kalo di golang ada macem query builder gak mas kaya eloquent di laravel

Asep Sae Puloh 9:52 PM
GORM

Wisnu | PT Webhade MDS 9:53 PM
@Asep Sae Puloh oh ada ya, mantap juga, lgsg riset deh

sy beneran nol besar id golang nih, kelamaan di php & python

fatchurrosii 9:53 PM
Ada kok 
https://github.com/huandu/go-sqlbuilder 
https://github.com/doug-martin/goqu 
Bisa pake salah satunya

Asep Sae Puloh 9:53 PM
Keluarga Lambo

deepseek-r1 9:54 PM
Gimana? Lambo lu warna apa?

Asep Sae Puloh 9:54 PM
Warna Sawit

Agung 9:55 PM
@Wisnu | PT Webhade MDS mencair

Rama Fajar Fadhillah 9:56 PM
bruh di laravel itu jatuhnya ORM bukan query builder

Yuristian Ardianto 9:56 PM
makasih mas

Zai 9:57 PM
kafaratul majlis

Agung 9:57 PM
belajar bareng dan ghibah bareng

tofumaulana 9:57 PM
terimakasih mas umam

Romi Alief Rahman 9:57 PM
jazakumullah khair, Mas Umam

Agung 9:57 PM
makasih mas umam

Romi Alief Rahman 9:57 PM
barakallahu fiikum

Dimas Ady 9:57 PM
yowess dapet one piece boy klo udah kelar pertemuan

MUH. IQRAM BAHRING 9:57 PM
Mas ada sertificate gk ya?

Rama Fajar Fadhillah 9:57 PM
bikin sendiri awkak

Setiawan 9:57 PM
GolangkanIndonesia

Rama Fajar Fadhillah 9:57 PM
tinggal 1 pertemuan lagi yak ?

Agung 9:57 PM
siapin baju toga

----
## XX-01-2025
Langsung aja mas
Gas

Syamsul Arifin 8:02 PM
gas

Reza Ali Nirwansyah_Teknik Komputer 8:02 PM
gas

Muchammad Rivari 8:03 PM
waalaikumsalam

Heru Oktafian, ST., CTT 8:03 PM
Bismillah....,

Ahmad Fadilah 8:03 PM
gass

M Rizal Hidayat 8:03 PM
waalaikumsalam

toni 8:03 PM
Waalaikumsalam wr.wb

Rizki A. 8:02 PM
halo all (mengapa saya disini yah? 🤔)

Syamsul Arifin 8:03 PM
lah rizky

Azuwir 8:03 PM
Walkmslm mas Umam

Remo 8:03 PM
volumenya kekecilan (kayaknya)

Syamsul Arifin 8:03 PM
kenceng di aku

toni 8:03 PM
Gasskeuun jadi bocah golang

Vivin Wahyudi 8:05 PM
ada live youtube-nya nggk zoom saya ngebug kadang nggk muncul suara

Rifki 8:05 PM
https://www.youtube.com/watch?v=47BLJ3EPNAw

Vivin Wahyudi 8:06 PM
thanks mas

Eka Anre Sangsya 8:06 PM
pake zoom web suaranya kenceng, zoom aplikasi suaranya kecil :D

toni 8:09 PM
Wah udah bikin ini di tugas 1 lagi 😅, kecepetan saya 😬

toni 8:09 PM
Service biasanya di golang pake Use Case

Agung 8:10 PM
model isinya struct sepertinya

Achmad Hadi Kurnia 8:10 PM
hemat bw YT atau zoom sih

toni 8:11 PM
Tinggal apply dependency injection aja atau pake poor man injection

M Rizal Hidayat 8:10 PM
handlers itu seperti controller ya ?

Muhammad Tri Wibowo 8:11 PM
iya, tapi gak berisi logic

M Rizal Hidayat 8:11 PM
oo di pisah lagi di services ya

Agung 8:11 PM
@Muhammad Tri Wibowo isinya request ya

Muhammad Tri Wibowo 8:11 PM
iyes, services itu business logic, repository untuk access data/database

toni 8:11 PM
Yes, handler == controller, bisa http handler, grpc handler, dkk

M Rizal Hidayat 8:12 PM
i see, terima kasih semua

Bayu 8:12 PM
izin ini materi tentang apa yaa baru gabung soalnya

Muhammad Tri Wibowo 8:13 PM
layering aplikasi kak

Rizki A. 8:12 PM
teman-teman disini team MVCS atau clean?

toni 8:12 PM
Team yg penting jalan 😀

Muhammad Tri Wibowo 8:12 PM
lebih enak clean, buat maintenance lebih mudah

Zannns 8:12 PM
MVCS lebih gampang dipahami

Akbar Hidayatuloh 8:13 PM
biasa pake mvc kalo ngoding js

Arif Setyo 8:13 PM
clean code lebih ke gimana codenya gampang di baca, bisa dengan mvc pattern atau service repository pattern, dll

Agung 8:14 PM
@Arif Setyo like it

Muhamad Putra Satria 8:14 PM
yahh udah limit

__-_-__ 8:15 PM
Supabase free tier dapat berapa projek ya, udah pernah buat satu tapi masih dipakai

Rizki 8:15 PM
atau d pause dl satu mas?
kyknya maksimal dua yg aktif kalo free

Rama AS 8:15 PM
Harus dimatikan 1 itu, max 3 yang aktif mas

__-_-__ 8:15 PM
Mak tiga ya mas?

Rizki 8:16 PM
project di supabase ini sama dengan schema ya kalo di mysql/postgres sendiri?
atau dalam project bisa banyak schema?

Agung 8:17 PM
 mudah2an di endorse sumopod hhehehe

Asep Sae Puloh 8:17 PM
products

toni 8:17 PM
Ga pake db migration aja?

M Rizal Hidayat 8:17 PM
ada doc tutorialnya nanti ga ya ?

Hendra Gunarsa 8:17 PM
biasanya ada

Djibrani Yuda 8:16 PM
kenapa choose supabase ketimbang neon mas?

Muhammad Tri Wibowo 8:17 PM
neon cenderung postgres only, kalau supabase bisa sebagai postgres maupun BaaS (backend as a service)

Denny Rachmadi 8:17 PM
biar satset :)

Djibrani Yuda 8:17 PM
aaa, okee

Putu Nadya Putri Astina 8:16 PM
supabase itu sql/no sql ya?

FARRELY FIRENZA 8:16 PM
postgresql

Arif Setyo 8:16 PM
no sql

Hilmi Fawwaz Saad 8:16 PM
supabase -> sql
firebase -> no sql

Putu Nadya Putri Astina 8:16 PM
i see thank you infonya kak

Arif Setyo 8:17 PM
saya kira like firebase 🤭

Hilmi Fawwaz Saad 8:17 PM
tidak hehe. Kalau firebase basis collection dan document ya. Kalau supabase ya normal, db, table gitu

Ikhsan 8:18 PM
Maaf baru gabung 🙏🏼

Agung 8:18 PM
kalo ada yang nanya kenapa ini kenapa itu, jawabanya semua karena .....

toni 8:18 PM
Tidak punya ….

Imal 8:18 PM
pemula T_T

toni 8:18 PM
Pelan2 saja mas

Imal 8:19 PM
kotak mas :D

Ariaseta Alam 8:21 PM
Youtube live: https://www.youtube.com/live/47BLJ3EPNAw

toni 8:17 PM
Ini bikin table dan kolom2nya manual mas? :Sw

toni 8:20 PM
Kalau deploy ke beda environment (development, staging, production) mesti dibuat manual 😅

Muhammad Tri Wibowo 8:20 PM
@Agung di go ada juga kok

Muhammad Hasan 8:21 PM
@Muhammad Tri Wibowo iya betul pake GORM

E_Akmal Amilunnizar 8:20 PM
ni gada youtube livenya kah min? biar bisa rewind kalo ada kelewat gitu

Syamsul Arifin 8:20 PM
ada tuh lg live di kanan atas

E_Akmal Amilunnizar 8:20 PM
oh iya bang thanks infonya

Ariaseta Alam 8:21 PM
https://www.youtube.com/live/47BLJ3EPNAw

ndra 8:22 PM
udah dengerin saja penjelasan nya

M Rizal Hidayat 8:20 PM
oooh GO ini bisa buat FE juga ya ? sy pikir untuk BE/API aja

Muhammad Tri Wibowo 8:21 PM
bisa, dulu saya awal pernah fullstack pakai go

fatah 8:22 PM
contohnya apa bang kalo FE pake go?

M Rizal Hidayat 8:22 PM
iya sy masih blm ada gambaran go jadi FE

Muhammad Hasan 8:22 PM
@Muhammad Tri Wibowo pake html css juga kan?

toni 8:22 PM
Tugas 2 kayaknya bakal bikin clean code ya? Wah udah dikerjain di tugas 1 lagi 🤣, apa dibikin makin clean aja ya

Ariaseta Alam 8:22 PM
Kalau udah disapu, lanjut dipel

Agung 8:22 PM
lanjut bantu ortu dulu brow

Sukron Ch 8:23 PM
Mending ngopi bang

toni 8:24 PM
Shiapp masih ada PR ngepel, nyuci piring, bantu ortu 😄

Wildan Febrian 8:21 PM
karena supabase ketat bgt menurut saya terkait password (password paling kuat yg saya ingat di kepala msih kurang kuat bagi supabase) apakah ada rekomendasi tools password manager buat nginget temen"?? :)

Ariaseta Alam 8:22 PM
Bitwarden

Wildan Febrian 8:23 PM
wah siap, makasih mas :)

toni 8:23 PM
Pake ini mas https://github.com/tonidy/pa-cli  #self-promo 🙈

Wildan Febrian 8:23 PM
apa nih mas?? otw ngecek

toni 8:24 PM
Simple CLI password manager mas

toni 8:25 PM
Bitwarden based on electron (berat)

fatah 8:25 PM
1password

Akbar Hidayatuloh 8:25 PM
nanti ada materi text lagi mas seperti kemarin di docs.kodingworks.io?

ndra 8:26 PM
mending dimatiin sih komen nya

Agung 8:27 PM
setubuh brow,

Sukron Ch 8:27 PM
Kenapa saat install viper tiba-tiba ada go.sum? 

MUH. IQRAM BAHRING 8:27 PM
Moment pada berisik, ada waktunya bertanya.

toni 8:28 PM
Hati2 .env jangan di-push ke github

Hafid Nur 8:28 PM
Model?

Rivaldi Putra 8:28 PM
repo

Valdo 8:28 PM
repositories

Ndlr 8:28 PM
model

M Rizal Hidayat 8:28 PM
repositories

Dio Saputra 8:28 PM
Repository

Z.ai 8:28 PM
Repositories

Muhamad Putra Satria 8:28 PM
repo mas

Hilmi Fawwaz Saad 8:28 PM
Repo

Fachry Ikhsal 8:28 PM
repo

Maulana 8:28 PM
Service

Asep Sae Puloh 8:28 PM
Repo

Wahyu Kurnia Prambudi 8:28 PM
repositories

Kurniawan Sudirman 8:28 PM
controller

Jefri Saputra 8:28 PM
models

Fajar setiawan Siagian 8:29 PM
define database di repositories layer. kalau service biasanya udah di business layer


Suteja 8:29 PM
config

fatah 8:29 PM
db

M Rizal Hidayat 8:29 PM
o beda

Agung 8:29 PM
revo 2015 Harga 7.500.000 pajak mati

Muhammad Ziad 8:29 PM
Revository

Ahmad Fajrul Falaah 8:28 PM
Best practicenya, kapan kita define config di file .env, kapan harus di ENV OS Level?

Sukron Ch 8:29 PM
.env file biasanya buat development aja

MUH. IQRAM BAHRING 8:30 PM
underline

ariefbayu 8:30 PM
underscore

Kurniawan Sudirman 8:31 PM
apa bisa konek ke 2 db berbeda?
--
apa bisa konek ke 2 db berbeda?

Ahmad Fajrul Falaah 8:28 PM
Best practicenya, kapan kita define config di file .env, kapan harus di ENV OS Level?

Sukron Ch 8:29 PM
.env file biasanya buat development aja

Kurniawan Sudirman 8:31 PM
biasanya pakai env variable buat env lain, sit ,uat prod, jadi pakai env variable OS aja.

Ndlr 8:32 PM
waw keren - bisa setup db pool di kodenya langsung

M Rizal Hidayat 8:31 PM
kalau pengen tau dimana lib2 nya atau kumpulan lib2 go itu dimana ? misal klw flutter kan di pub.dev

fatah 8:31 PM
Ada di pkg.go.dev

M Rizal Hidayat 8:32 PM
naiss, noted thanks mas

Rio Permana 8:32 PM
go.sum

Muhammad Ziad 8:32 PM
Go and go sum kalau mau cek di src mas

toni 8:32 PM
Go.mod

dfc 8:32 PM
mending dengarkan dulu, biar nggak ganggu yg lain. Nanti ada Q/A biasanya

Ferry Hasan 8:32 PM
btw ada link dokumentasi kodingworks untuk pertemuan ke-2 ini kah?

Maulana 8:32 PM
Selesai kelas baru di share bang

Ferry Hasan 8:32 PM
owh... okok thanks infonya bg

fatchurrosii 8:34 PM
koneksi.php

Muhammad Hasan 8:34 PM
mbah

karrang 8:34 PM
ada link YouTube-nya gk mas telat ini masuk karna jaringan

fatah 8:35 PM
https://www.youtube.com/watch?v=47BLJ3EPNAw

karrang 8:35 PM
terimakasih

Agung 8:37 PM
mas mau nanya dulu mas

ndra 8:38 PM
nanti ada sesi nya

anggavb 8:38 PM
notes sesi 2 ini udah di share belum ya ges?

dfc 8:38 PM
blm, itu lagi dibuat

Syamsul Arifin 8:39 PM
belum paham magic yg pake "_"

Rivaldi Putra 8:39 PM
Bleh di ijinin open mic gak mas

Hafid Nur 8:39 PM
Bagian import _ "github.com/lib/pg" kok ga bisa mass?

anggavb 8:39 PM
di install dulu mas
bisa jalanin go mod tidy

Muhammad Shahwal 8:39 PM
mapstructure itu bawaan viper kah mas ?

Alfian 8:40 PM
iya

Romi Alief Rahman 8:43 PM
ga camelCase, ya Mas Umam ?

Arif Pamungkas 8:43 PM
Izin buat pelajari recordnya aja 

Muhammad Hasan 8:43 PM
suaranya kok lama2 makin mengecil ya, apa laptopq

Hendra Gunarsa 8:39 PM
mas tadi kan autoenv baca dari OS. berarti cuma ke PORT aja ya?

Imal 8:41 PM
Harusnya di akan baca semua content yang ada di env mas

Hendra Gunarsa 8:42 PM
betul yg di define .env kan mas ya, tp tadi kan confignya autobaca dari OS dlu, nah itu hanya port saja ya berarti missal kita di os hanya buka port 8080, itu akan kepakai ya instead of yg ada di .env folder project?

Imal 8:44 PM
CMIIW
Harusnya bakal prio baca yang di OS sih

Alfian 8:44 PM
sebenernya kita ga buka port 8080 di os mas, maksud dari baca env di OS tuh hanya untuk membaca value port yang ingin kita gunakan di aplikasinya tanpa harus hard code di code nya

Muhammad Hasan 8:44 PM
kalau di nest JS service itu wajib kalau di golang wajib juga kah?

Dio Saputra 8:45 PM
dependency inject itu yang biasanya kayak func NewSomething gitu bukan ya?

toni 8:47 PM
Itu poor man DI mas, ada beberapa library di golang yg bisa automatic DI (cmiiw)


kalau di nest JS service itu wajib kalau di golang wajib juga kah?

Alfian 8:50 PM
ga wajib mas, di golang bebas aja logic bisa taro di handlers langsung, tapi kita ngikutin layered arch

Dio Saputra 8:45 PM
dependency inject itu yang biasanya kayak func NewSomething gitu bukan ya?

toni 8:47 PM
Itu poor man DI mas, ada beberapa library di golang yg bisa automatic DI (cmiiw)

Z.ai 8:49 PM
product_model

Gustio Nusamba 8:52 PM
kapan kita menggunakan pointer di go?

Hendra Gunarsa 8:39 PM
mas tadi kan autoenv baca dari OS. berarti cuma ke PORT aja ya?

Imal 8:41 PM
Harusnya di akan baca semua content yang ada di env mas

Hendra Gunarsa 8:42 PM
betul yg di define .env kan mas ya, tp tadi kan confignya autobaca dari OS dlu, nah itu hanya port saja ya berarti missal kita di os hanya buka port 8080, itu akan kepakai ya instead of yg ada di .env folder project?

Imal 8:44 PM
CMIIW
Harusnya bakal prio baca yang di OS sih

Alfian 8:44 PM
sebenernya kita ga buka port 8080 di os mas, maksud dari baca env di OS tuh hanya untuk membaca value port yang ingin kita gunakan di aplikasinya tanpa harus hard code di code nya

Hendra Gunarsa 8:52 PM
ahh i se thank you mas penjelasannya

Call Vin 8:51 PM
Mas *repository &repository itu apa ya?

Dio Saputra 8:51 PM
Pointer

Muhammad Ziad 8:52 PM
Sama reference

Riki Ridwan 8:54 PM
Handler itu sama kayak controller ya?

Handler itu sama kayak controller ya?

fatchurrosii 8:54 PM
Iya bener

Muhammad Ziad 8:54 PM
Yes

Kurniawan Sudirman 8:54 PM
user java

Riki Ridwan 8:55 PM
jawascript

Gustio Nusamba 8:52 PM
kapan kita menggunakan pointer di go?

Rivaldi Putra 8:55 PM
Biasanya Ketika butuh nilai aslinya bukan nilai copy

Z.ai 8:56 PM
saat butuh modifikasi objek aslinya -cmiiw

Rido Rachmatan 8:59 PM
nama tablenya singular

Muchammad Rivari 9:00 PM
nama routenya sekalian diubah english mas

Ewink Cess 9:01 PM
Bang klo bisa jangan di copy koding nya pusing bacanya aahahaha

san 9:02 PM
Stuju bang jadi kecepatan agak pusing wkwk

Ewink Cess 9:02 PM
Iya gak paham alurnya gimana

toni 9:03 PM
Tanya AI aja gaes

Ewink Cess 9:03 PM
Solusi catat yg penting2 aja atar tanya AI

Z.ai 9:03 PM
kalo ada DI jadi muter2 ya

Jefri Saputra 9:02 PM
ORM golang biasanya pakai apa guys?

Heru Oktafian, ST., CTT 9:02 PM
saya biasa pake gorm

fatah 9:03 PM
Saya baru belajar pake uptrace bun

Heru Oktafian, ST., CTT 9:03 PM
cuman jangan samakan orm di golang sama laravel yah hehehe

toni 9:03 PM
Golang pake raw query

Jefri Saputra 9:04 PM
aku biasa JPA

Fendi 9:03 PM
Library untuk form validation di Golang, biasanya pakai apa ya?

fatah 9:03 PM
govalidator
atau pake ozzo-validation bisa juga
--

Ga pake uuid mas?


Muhammad Tri Wibowo 9:07 PM
itu lanjutan, yang penting basicnya mengerti, cmiiw

Kurniawan Sudirman 9:07 PM
tsid :v

toni 9:08 PM
ksuid

Jefri Saputra 9:07 PM
go bisa auto create field table juga gak?

Muhammad Tri Wibowo 9:07 PM
kalau pakai GORM bisa

Jefri Saputra 9:08 PM
Thanks
framework GO yang biasa di gunakan apa guys?

M Rizal Hidayat 9:09 PM
mas sy pusing itu ada 2 service

Denny Rachmadi 9:10 PM
iya harusnya product_model.go kah?

Z.ai 9:10 PM
😄

lilo 9:10 PM
apakah di go juga ada prepared statement?

Kurniawan Sudirman 9:11 PM
kenapa logic prosesnya itu dilakukan di handler ya mas ? bukan di service?

Rido Rachmatan 9:14 PM
iya juga ya, kenapa logicnya ga ditaro di service

darm 9:13 PM
Kalau pake db transaction itu harus buat func gk ya

deeant 9:11 PM
   ini nanti kita tinggal copy script?

Maulana 9:11 PM
Iyaa
Sudah2 begitu

deeant 9:12 PM
mau ngejar ngetik codingan dan memahami gak kekejar

anggavb 9:13 PM
samaa, nyimak doang akhir nya ini.
moga terserap dengan baik setelah ini

Rifki 9:13 PM
pahami dulu aja, ngoding nya nanti sambil nonton rekaman nya, speed nya di set ke 0,5

Maulana 9:13 PM
Ini saya juga menyimak full . Nanti sambil cek2 video dan code

Ewink Cess 9:13 PM
Biasanya di share meteri code beserta penjelasannya bang

Dimas Ady 9:14 PM
Sql Injection : Hola amigos !!

san 9:06 PM
Itu kadang ada fungsi yang func () NewFunc sama func NewFunc() apa bedanya ya dan kapan dipakainya?

anggavb 9:11 PM
di go ada method ada function
kalo func () NamaFunc ==> method
dipake untuk membuat argumen receiver khusus yang biasanya buat bikin kontrak ke struct/tipe data tertentu
 
func NewFunc() ===> function
 function pada umumnya (sepertinya hehe)

cmiww, baru belajar juga '-'

san 9:12 PM
Wah nice insight mas, berarti buat DI harus selalu pake func () NamaFunc ya buat inject si contract structnya itu?

kimi k2.5 9:14 PM
Itu kayak konstruktor kalo di TS

Maulana 9:14 PM
Cara nambah kolom baru bang di codingan dan supabase gimana?

Fitriningtyas 9:16 PM
mas, itu yg di dalam folder  models itu product_model.go?

Z.ai 9:14 PM
aduh udah deploy aja

Ewink Cess 9:15 PM
Fast moving wkwwkw

toni 9:15 PM
Sat..set..satt..settt

Z.ai 9:16 PM
wkwk

kimi k2.5 9:16 PM
Tiba2 di udara

toni 9:18 PM
Ini ada z.ai terus kimi k2.5, deepseek sama qwen di mana? 😂

Z.ai 9:16 PM
source : https://github.com/zuhrulumam/kasir-api

anggavb 9:18 PM
izin menyelam puh

Dzikran Azka Sajidan 9:18 PM
pake deploy yang di kiri atas

Muchammad Rivari 9:18 PM
lognya apa mas?

Imal 9:19 PM
apply changes di kiri atas mas

Rama AS 9:19 PM
Mungkin harus init env nya di setting dulu bang Umam?

Muchammad Rivari 9:19 PM
belum masukin supabase ke env railway mas

Dimas Ady 9:18 PM
sidenote:
"Dasar" means "orang yang pernah ngoding" in 
"Jago Golang Dasar"

toni 9:19 PM
Dasar != pemula 😅

ARN182-23160 Arif Zein 9:19 PM
yg master aja msh error...gmn nasib beginner

Maulana 9:19 PM
Railways open all port ga ya ?

Fajar setiawan Siagian 9:19 PM
kalau kita commit k main dia akan jalankan ci/cd railways untuk otomatis build dan deploy ya mas

Dzikran Azka Sajidan 9:19 PM
coba klik tombol deploy warna ungu di kiri atas mas umam

Rivaldi Putra 9:19 PM
Apply changes dull mas

Jefri Saputra 9:19 PM
Apply 2 changes -> Deploy

Imal 9:19 PM
apply changes dulu om
---
kiri atas apply chanhe

faqih yugo susilo 9:20 PM
Database nya ga connect

Rido Rachmatan 9:20 PM
apply 2 change

Ardi 9:20 PM
kiri atas apply changes

E_Akmal Amilunnizar 9:20 PM
plss liat komen bang apply

Ferry Hasan 9:20 PM
oakokwaw

E_Akmal Amilunnizar 9:20 PM
buset

ARN182-23160 Arif Zein 9:19 PM
yg master aja msh error...gmn nasib beginner

toni 9:20 PM
Kalo ga ada error justru aneh mas
Mau master pun kalau pas ngoding selalu ada problem yg mesti di-solve karena itu tugasnya kita sbg engineer

M Rizal Hidayat 9:21 PM
apakah type Product struct, baik di models atau handlers, dll itu harus sama dengan nama file ?

2. DI dependency injection, itu berarti setiap handlers, services, dan repositories itu wajib ada DI nya itu ya termasuk main.go?

Jefri Saputra 9:21 PM
Mantap

Bayu 9:21 PM
debug in production ♥️

Rama AS 9:21 PM
connection pool di Supabasenya reset atau tambah pool connection nya maybe?

san 9:06 PM
Itu kadang ada fungsi yang func () NewFunc sama func NewFunc() apa bedanya ya dan kapan dipakainya?

san 9:12 PM
Wah nice insight mas, berarti buat DI harus selalu pake func () NamaFunc ya buat inject si contract structnya itu?

kimi k2.5 9:14 PM
Itu kayak konstruktor kalo di TS

Dio Saputra 9:21 PM
kalo misal ada func () NamaFunc ini berarti function tersebut adalah method dari sebuah struct.
Semisal masnya punya struct product lalu mas buat func seperti:
func (p product) NamaFunction
function tersebut berarti merupakan method dari struct product, nah di dalam function tersebut mas nya bisa akses field struct product lewat variable p yang ada di (p product) sebelum nama function itu tadi. (cmiiw)

Dimas Ady 9:18 PM
sidenote:
"Dasar" means "orang yang pernah ngoding" in 
"Jago Golang Dasar"

toni 9:19 PM
Dasar != pemula 😅

Agung 9:22 PM
@toni if Dasar != nill {

ARN182-23160 Arif Zein 9:22 PM
kalo "dari nol " apa artinya mass

Agung 9:22 PM
@ARN182-23160 Arif Zein cuci tangan cuci kaki bobo
---

olong tim infra hahah

san 9:06 PM
Itu kadang ada fungsi yang func () NewFunc sama func NewFunc() apa bedanya ya dan kapan dipakainya?

anggavb 9:11 PM
di go ada method ada function
kalo func () NamaFunc ==> method
dipake untuk membuat argumen receiver khusus yang biasanya buat bikin kontrak ke struct/tipe data tertentu
 
func NewFunc() ===> function
 function pada umumnya (sepertinya hehe)

cmiww, baru belajar juga '-'

san 9:12 PM
Wah nice insight mas, berarti buat DI harus selalu pake func () NamaFunc ya buat inject si contract structnya itu?

kimi k2.5 9:14 PM
Itu kayak konstruktor kalo di TS

Dio Saputra 9:21 PM
kalo misal ada func () NamaFunc ini berarti function tersebut adalah method dari sebuah struct.
Semisal masnya punya struct product lalu mas buat func seperti:
func (p product) NamaFunction
function tersebut berarti merupakan method dari struct product, nah di dalam function tersebut mas nya bisa akses field struct product lewat variable p yang ada di (p product) sebelum nama function itu tadi. (cmiiw)

san 9:23 PM
Wahh makasi mas mas semuanya, sangat menjawab

Zannns 9:23 PM
kalo saya lebih suka notasi :
func () functionName() {}
jadi jelas aja mana params sebagai penghubung ke struct, mana params yang emng punya si func itu

Frenki Herlambang 9:23 PM
harus SSL mungkin ya?

Otniel Abiezer 9:23 PM
Gak allow mungkin IP address

Ewink Cess 9:24 PM
Semaking menyala

Agung 9:24 PM
segera support dan deploy ke sumopod

Hafid Nur 9:24 PM
Apakah karena ini? Supabase No IPv4 compatible?

dfc 9:24 PM
apa port-nya bentrok apa ya, sudah dipake yg sblmnya kali

Agung 9:25 PM
chat gpt adalah kounjhy

mochamad iqbal 9:25 PM
iya mas atau pake port yg agak random mgkn

Asep Sae Puloh 9:25 PM
Debug di Production

Ndlr 9:25 PM
Kok bisa kedetek localhost IPv6

mochamad iqbal 9:25 PM
(jarang dipake 3002

Muchammad Rivari 9:25 PM
supabase dashboardnya udah disetting kah?
Settings → Database → Connection string → URI

hijri hadi 9:25 PM
pakek database srtring yang ipv4 itu gak si mas ?

mochamad iqbal 9:25 PM
si railway mgkn ga bisa connect db luar mas(?)

ariefbayu 9:25 PM
IPv6 localhost

fatah 9:25 PM
coba di print dulu mas env nya Sebelum start server

ariefbayu 9:25 PM
belum pernah pakai supabase🙏

ARN182-23160 Arif Zein 9:26 PM
sama...kita tim sup ayam aja

kimi k2.5 9:26 PM
Add services coba mas

Rizki 9:26 PM
Connection stringnya ngga kebaca ya kyknya?

Syamsul Arifin 9:26 PM
butuh allow di supabasenya kayanya mas

kimi k2.5 9:26 PM
Kalo ada supabase alhamdulillah

Muchammad Rivari 9:26 PM
lalu pastiin PORTnya 5432, dan sslmode require: postgresql://postgres:<PASSWORD>@db.<PROJECT_REF>.supabase.co:5432/postgres?sslmode=require

Ewink Cess 9:26 PM
Saya bikin kopi dlu

Zannns 9:26 PM
railway support Neon gak ya?

FITRA FADILANA 9:26 PM
coba pake supabase bawaan railway mas

Alfian 9:26 PM
railway kayaknya gamau outbound connection

kimi k2.5 9:26 PM
Tinggal setting2 nantinya di env services nya

Hendra Gunarsa 9:26 PM
1. The "IPv6" Problem (Most Common)
By default, Supabase's direct connection (port 5432) uses IPv6. Railway will fail to reach this.
The Fix: Use the Supabase Connection Pooler string instead of the direct connection string.
In your Supabase Dashboard, go to Settings > Database.
Look for the Connection Pooler section.
Ensure Mode is set to Transaction (or Session depending on your needs).
Use the hostname that ends in *.pooler.supabase.com (usually port 6543 or 5432). This pooler supports IPv4, which Railway requires.

ini gak mas karena supabase pake ipv6 sedangkan railway  butuh ipv4?

Sukron Ch 9:26 PM
Coba buat log disekitar db conn nya mas, barang kali bisa dipersempit code yang buat error

Denny Rachmadi 9:27 PM
pake pooler mas di supabasenya coba

Rivaldi Putra 9:27 PM
Coba cek kodę bagian env mas

Imal 9:27 PM
btw pakai direct connection ya ?

FITRA FADILANA 9:26 PM
coba pake supabase bawaan railway mas

Z.ai 9:27 PM
oh ada juga ya

Rama AS 9:27 PM
Conection pool di supabase nya mas umam

faqih yugo susilo 9:27 PM
Newbie ngga paham ini

ariefbayu 9:26 PM
belum sampe supabase sih ini kalo liat tcp [::1]

Rido Rachmatan 9:27 PM
tcp [::1] tuh ip6?

Ferry Hasan 9:27 PM
"Supavisor Session Pooler string instead of the direct connection string"

toni 9:27 PM
Klik connect mas

Fadli 9:27 PM
coba di log env nya mas dapat atau engga

Hendra Gunarsa 9:27 PM
samplenya gini mas katanya
DATABASE_URL = postgres://postgres.[user]:[password]@aws-0-[region].pooler.supabase.com:6543/postgres?sslmode=verify-full

toni 9:27 PM
Boleh ngomong ga sih, wkwk

Asep Sae Puloh 9:28 PM
Jangan

Hafid Nur 9:28 PM
Coba alih-alih pake yang Direct Connection, pakai yang Transaction Pooler di "Connect" Supabase mas

Denny Rachmadi 9:28 PM
coba project overview -> connect (button)

RagilPratama 9:28 PM
salah copy connection stringnya kali

Imal 9:28 PM
coba ke supabase ganti connection string nya mas

Syamsul Arifin 9:29 PM
Use the Supabase Pooler (Recommended): The session pooler connection string is IPv4 compatible and the recommended approach for serverless environments like Railway.In your Supabase project dashboard, navigate to Project Settings > Database and under the "Connection Strings" section, find the "Connection pooler" tab.Use this connection string in your Railway environment variables, ensuring you use the correct port (usually 5432 or 6543).Remember to replace the [YOUR-PASSWORD] placeholder with your actual database password.

Agung 9:29 PM
DATABASE_URL=postgresql://postgres.xxxxx:password@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres?sslmode=require

Jefri Saputra 9:29 PM
sslmode=required

Agung 9:29 PM
wajib ada require mas

ariefbayu 9:29 PM
gpp, biar dapat pengalaman problem solving production

Asep Sae Puloh 9:29 PM
DEBUG Mas, Environment nya di Log

Ardi 9:29 PM
di suruh deploy supabase di railways nya

Raka Nugraha 9:29 PM
Smangaat Maass,, Bantu do'a aja saya.,,,,

Sukron Ch 9:29 PM
Bikin postgres langsung dari railway aja mas

M Rizal Hidayat 9:21 PM
apakah type Product struct, baik di models atau handlers, dll itu harus sama dengan nama file ?

2. DI dependency injection, itu berarti setiap handlers, services, dan repositories itu wajib ada DI nya itu ya termasuk main.go?

M Rizal Hidayat 9:29 PM
takut lupa sama materi utamanya wkwk focus ke supabase

Imal 9:29 PM
dibagian navbar mas
ada bacaan connect

Denny Rachmadi 9:29 PM
coba project overview -> connect (button)

Erwin Rianto 9:30 PM
- Go mencoba connect ke localhost:5432, padahal Supabase ada di luar.
- Solusi: pastikan env DB_HOST bukan localhost, tapi db.supabase.co.


kata mbah GPT sih gini

toni 9:30 PM
Nah ituu
Dicoret2 c

hijri hadi 9:30 PM
pakai string yang ipv 4 biar bisa di akses public

Rama AS 9:30 PM
methodnya mas

Boy 9:29 PM
sabar ges, itu mas umam sedang nunjukin juga problem solving

Syamsul Arifin 9:30 PM
dibantu gapapa harusnya

Imal 9:30 PM
yoi pak

Hendra Gunarsa 9:30 PM
nah ketemu

Muchammad Rivari 9:30 PM
mantab

ARN182-23160 Arif Zein 9:28 PM
pentingnya hilirisasi digital .... pelajari AI , pelajari coding sejak dini dari PAUD

Bayu 9:29 PM
sesuai instruksi mas wapres ya

ARN182-23160 Arif Zein 9:29 PM
😆

Boy 9:29 PM
siap mas wapres

Asep Sae Puloh 9:30 PM
Dari Kandungan

toni 9:30 PM
Coret2annya tolong dihapus

Boy 9:30 PM
itu yang coret2 board tolong ditertibkan pak

Z.ai 9:31 PM
mohon di clear lg ya 😄

fatah 9:31 PM
tolong yang coret2 dihapus gan

Jefri Saputra 9:31 PM
mantap

Z.ai 9:31 PM
ayo bismillah

Hendra Gunarsa 9:31 PM
mantap live debug kwkwkwkw
goks timnya

Ferry Hasan 9:32 PM
letsgoo

mikel 9:32 PM
yeayy

fatah 9:32 PM
gogogo

Rama AS 9:32 PM
http mas

kimi k2.5 9:32 PM
Yg public coba mas

MUH. IQRAM BAHRING 9:32 PM
Methodnya mas

ionig 9:32 PM
https

toni 9:32 PM
Internal bisa diakses ya?

Jefri Saputra 9:32 PM
asyik bgt debug jamaah

Hendra Gunarsa 9:32 PM
enak debug jamaah otaknya jadi banyak ya mas kwkwkw

Jefri Saputra 9:32 PM
iyaks
gara gara motor revo udah laku

Rifki 9:33 PM
pake public mas, jangan internal

Denny Rachmadi 9:33 PM
salah kopi === teh :)

fatah 9:33 PM
Ngopi dulu ges

Jefri Saputra 9:31 PM
mana tadi yang jual  motor revo

toni 9:31 PM
Udah laku bang

Jefri Saputra 9:32 PM
cepet bgt

fatah 9:32 PM
7.5jt bang

ARN182-23160 Arif Zein 9:33 PM
kok jadi kayak olx gini

M Rizal Hidayat 9:33 PM
di supabase tadi yg gambar listrik colokan ya ?

Kurniawan Sudirman 9:33 PM
kok lama bngt ya loadingnya

FITRA FADILANA 9:33 PM
namanya jg gratisan

__-_-__ 9:32 PM
Jadi tadi yang salah di apa bang aku ke logout

Hendra Gunarsa 9:33 PM
summarnya kalo dploy di railway, dipastikan pake mode transaction pooler di supabasenya mas, supaya dapat ipv4 . Supabase direct defaultny apake ipv6 sedangkan railway cuma bisa baca ipv4, kurang lebih begitu

mochamad iqbal 9:33 PM
itu yg lama loadingnya coba ganti region

Jefri Saputra 9:33 PM
collection nya mas umambanyak banget

fatah 9:34 PM
banyak projekannya

mochamad iqbal 9:33 PM
ke singapore
soalnya default ke US apa yah

Maftuhriza Yumida 9:34 PM
https://

MUH. IQRAM BAHRING 9:34 PM
Methodnya mas

Hafid Nur 9:34 PM
Kemarin issuenya karena ga pake https:// mas

Muchammad Rivari 9:34 PM
pake https

toni 9:34 PM
Aduuh ada coret2 itu

Z.ai 9:34 PM
haha sampe di doodle

Rivaldi Putra 9:34 PM
https mas

---


Internal bisa diakses ya?

Jefri Saputra 9:32 PM
asyik bgt debug jamaah

Hendra Gunarsa 9:32 PM
enak debug jamaah otaknya jadi banyak ya mas kwkwkw

Jefri Saputra 9:32 PM
iyaks
gara gara motor revo udah laku

Rifki 9:33 PM
pake public mas, jangan internal

Denny Rachmadi 9:33 PM
salah kopi === teh :)

fatah 9:33 PM
Ngopi dulu ges

Jefri Saputra 9:31 PM
mana tadi yang jual  motor revo

toni 9:31 PM
Udah laku bang

Jefri Saputra 9:32 PM
cepet bgt

fatah 9:32 PM
7.5jt bang

ARN182-23160 Arif Zein 9:33 PM
kok jadi kayak olx gini

M Rizal Hidayat 9:33 PM
di supabase tadi yg gambar listrik colokan ya ?

Kurniawan Sudirman 9:33 PM
kok lama bngt ya loadingnya

FITRA FADILANA 9:33 PM
namanya jg gratisan

__-_-__ 9:32 PM
Jadi tadi yang salah di apa bang aku ke logout

Hendra Gunarsa 9:33 PM
summarnya kalo dploy di railway, dipastikan pake mode transaction pooler di supabasenya mas, supaya dapat ipv4 . Supabase direct defaultny apake ipv6 sedangkan railway cuma bisa baca ipv4, kurang lebih begitu

mochamad iqbal 9:33 PM
itu yg lama loadingnya coba ganti region

Jefri Saputra 9:33 PM
collection nya mas umambanyak banget

fatah 9:34 PM
banyak projekannya

mochamad iqbal 9:33 PM
ke singapore
soalnya default ke US apa yah

Maftuhriza Yumida 9:34 PM
https://

MUH. IQRAM BAHRING 9:34 PM
Methodnya mas

Hafid Nur 9:34 PM
Kemarin issuenya karena ga pake https:// mas

Muchammad Rivari 9:34 PM
pake https

toni 9:34 PM
Aduuh ada coret2 itu

Z.ai 9:34 PM
haha sampe di doodle

Rivaldi Putra 9:34 PM
https mas

Jefri Saputra 9:34 PM
yoi

faqih yugo susilo 9:35 PM
👍

toni 9:35 PM
Apus lagi dong coretannya

Rido Rachmatan 9:35 PM
zeabur jg harus gitu ya pake pooler?

E_Akmal Amilunnizar 9:35 PM
buset gambar apa itu

MUH. IQRAM BAHRING 9:35 PM
Oala bukan method tpi protocol ternyata namanya ehehhe

ARN182-23160 Arif Zein 9:35 PM
udah solved,,,,hapus mas

Z.ai 9:35 PM
https://kasir-api-production-ecc4.up.railway.app/api/produk

Kurniawan Sudirman 9:35 PM
mas kalau unit test biasanya pakai apa mas? dan untuk test quality code di SonarQube biasanya gim ana ?

Ndlr 9:35 PM
Mas bisa dijabarin ttg Pointer, &, interface, dan method di Golang :D

Ricky Sutanto 9:35 PM
Gentian yg ngajar

Syamsul Arifin 9:35 PM
gara2 apa tadi?

E_Akmal Amilunnizar 9:35 PM
ditambahin https:// protocol didepan railway addressnya

Bayu 9:35 PM
clear doodle mas bro

M Rizal Hidayat 9:35 PM
solved nya tadi gimana mas umam wkwk

Hafid Nur 9:35 PM
Ganti connection string dari Direct ke Transaction Pooler

toni 9:35 PM
Tengok di YouTube bang

Eka Anre Sangsya 9:36 PM
tanya mas,
untuk kolom table yang NULLABLE, saat define struct prefer menggunakan tipe data pointer (*string, *int, dll) atau sql.NullString, sql.NullInt, dll?

menurut pengalaman mas Umam lebih nyaman pake yang mana?

Jefri Saputra 9:35 PM
gara gara revo laku

toni 9:36 PM
lol
Siapa yg beli yak?

Jefri Saputra 9:36 PM
aji mumpung sales

Kurniawan Sudirman 9:36 PM
hati hati kalua max connection besar, berat di db, kecuali dbnya kuat. apa lagi kalau banyak microservices.

Sukron Ch 9:36 PM
Bandingin sama php connect db secara default mas, yg ngga ada pooling

Z.ai 9:37 PM
"https://" nya dihapus tolong 😁
---
"https://" nya dihapus tolong 😁

ARN182-23160 Arif Zein 9:38 PM
orgnya udah zzzz....😴

ARN182-23160 Arif Zein 9:38 PM
orgnya nyiapin revo lain

toni 9:39 PM
Jualan supra gak?

ARN182-23160 Arif Zein 9:39 PM
gak bang,....yg lagi viral revo...byk yg cari

toni 9:40 PM
Cari supra bapak
Bapaknya siapa yg punya supra

ARN182-23160 Arif Zein 9:40 PM
nama bapaknya supri...bukan supra

MUH. IQRAM BAHRING 9:40 PM
Gk usah si, karena sudah di file product, rugi dong nerapin konsep modular

faqih yugo susilo 9:40 PM
Kenapa ngga SQLite?

Andre Tansil 9:41 PM
aku lupa join wkwkwk, ga ada di calendarrrrr. btw yg join cuma 250 an ya

toni 9:41 PM
Tiap company beda2 rule coding-nya 😅, jadi clean code tiap company bakalan beda2

dwirez 9:41 PM
jadi ini tadi clean code ya guys?

Ricky Sutanto 9:40 PM
Link YouTube nya apa ya??

Erwin Rianto 9:41 PM
https://www.youtube.com/watch?v=47BLJ3EPNAw

Ricky Sutanto 9:41 PM
👍🏼

Jefri Saputra 9:41 PM
ingat guys ini zero to hero

toni 9:42 PM
From nil to nol 😆

Gk usah si, karena sudah di file product, rugi dong nerapin konsep modular

Sukron Ch 9:42 PM
Masalah nya package masih ngikut nama general dari folder mas, jadi kalau ada file lain dengan nama function tersebut nanti udah ngga bisa dibuat lagi

darm 9:42 PM
Ada menu realtime itu

--
Ada concern security ga mas kalau pake .env? Misal isi file nya bisa kebaca dibanding pake OS variabel.

Andre Tansil 9:45 PM
jgn lupa masukin ke .gitignore

Denny Rachmadi 9:45 PM
concern klo ikut kepublish ke git repo sepertinya CMIIW

Z.ai 9:37 PM
"https://" nya dihapus tolong 😁

toni 9:40 PM
Bapaknya siapa yg punya supra

ARN182-23160 Arif Zein 9:40 PM
nama bapaknya supri...bukan supra

Z.ai 9:45 PM
https://kasir-api-production-ecc4.up.railway.app/api/produk/7 

udah di restock revo nya

Jefri Saputra 9:45 PM
toni lagi bae

fatah 9:45 PM
hahaha

Asep Sae Puloh 9:45 PM
Saya\

Asep Sae Puloh 9:46 PM
Thank you Mas

Agung 9:46 PM
abis nanya tadi langsung ke unmute mas mic nya

--terima kasih mas

atqa 9:50 PM
cmiiw

Setiawan 9:50 PM
Mas nanti di share lagi videonya mas? 

fatah 9:51 PM
Udh ada di yutub https://www.youtube.com/watch?v=47BLJ3EPNAw

naufal.nabbih 9:51 PM
di private

{%hermanJaro%} 9:51 PM
Izin left, masuk angin 😄 WITA

atqa 9:50 PM
handler urusin route, nerima req dan kirim res
service itu business logic
repo itu ngobrol sama database

Agung 9:51 PM
repo 2015 masih dengan Harga yang sama mas, 7.500.000

atqa 9:52 PM
surat surat lengkap pak? 😁

MUH. IQRAM BAHRING 9:40 PM
Gk usah si, karena sudah di file product, rugi dong nerapin konsep modular

fatah 9:44 PM
Yep bener, dalam 1 package walaupun beda file tetep gabisa bikin function din nama yg sama
*din = dgn

Gustio Nusamba 9:52 PM
klo saya pake interface, jdi nanti masing2 entity(product, category) punya interface sendiri2. interface ini memuat definisi function/method tanpa implementasi. nah nanti baru kita implementasi method/function berdasarkan definisi yang sudah ditulis di interface.

MUH. IQRAM BAHRING 9:53 PM
@Sukron Ch @fatah  wah benar, ternyata go berbeda. Terimakasih.

Eka Anre Sangsya 9:52 PM
terima kasih sudah mewakili pertanyaan saya mas febri :D

E_Akmal Amilunnizar 9:53 PM
sql apa tadi mas? sql null?
sorry kurang denger

Eka Anre Sangsya 9:53 PM
sql.NullString, sql.NullInt, dll

mochamad iqbal 9:56 PM
race conditino?

fatah 9:56 PM
race condition

Agung 9:56 PM
decline race

Muhammad Tri Wibowo 9:57 PM
crawling

Muhammad Hizbullah 9:57 PM
queue system

fatah 9:57 PM
parallel prosesing

Muhammad Tri Wibowo 9:57 PM
atau bulk process

fatah 9:57 PM
Merge n conquer

dwirez 9:58 PM
bedanya kodingan pakek goroutine ama gak pakek apa ya bang?
dalam syntaxnya

dfc 9:58 PM
izin leave dulu mas umam, nanti nunggu rekaman aja, trims

Denny Rachmadi 9:58 PM
Goﾠ                Node / NestJS
goroutineﾠ    async function
channelﾠ       EventEmitter / RxJS / queue

Muhammad Tri Wibowo 9:59 PM
ambil data tracking kurir

fatah 9:59 PM
Background proses

---
ambil data tracking kurir

fatah 9:59 PM
Background proses

Sukron Ch 9:59 PM
Update saldo? 

Z.ai 9:59 PM
pamit dulu, billing habis 🙏

dwirez 9:58 PM
bedanya kodingan pakek goroutine ama gak pakek apa ya bang?

Muhammad Tri Wibowo 9:59 PM
misal kita ada fungsi namanya "ambilData()"
kalau pakai go routine sesederhana

go ambilData() / ini gak blocking

kalau tanpa goroutine

ambilData saja / ini blocking

Agung 10:00 PM
buat queue mas