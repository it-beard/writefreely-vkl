<p align="center">
	<a href="[https://hub.docker.com/r/writeas/writefreely/](https://letapis.vkl.world)">
		<img src="https://s3-eu-central-1.amazonaws.com/vklworld/site_uploads/files/000/000/002/@2x/034a7c3804534449.png" />
	</a>
</p>
<hr/>

## Агульныя звесткі

Гэта рэпазіторый з кодавай базай праекта https://letapis.vkl.world - невялічкай блогавай сістэмы, пабудаванай на базе open-source рухавіка [WriteFleery](https://github.com/writefreely/writefreely).

Рухавік падтрымлівае: 
* Мнагакарыстальніцкі рэжым
* Некалькі блогаў на адзін акаўнт
* Прыватныя, запароленыя ды публічныя блогі
* Добры інтэрфейс рэдактара з падтрымкай Markdown & HTML
* Асабістыя стылі блогаў прац карыстальніцкія CSS-стылі
* Падпіскі на блогі праз RSS і ActivityPub (Fediverse)
* Стужку з усімі публічнымі блогамі
* І шмат чаго яшчэ.

ВАЖНА: У гэтым рэпазіторыі знаходзіцца кодавая база менавіта той версіі рухавіка, якая патрэбна нам. Мы не займаемся падтрымкай усіх яго магчымасцей, а гэты форк з'яўляецца хардфоркам WriteFleery 0.13.2 з несумяшчальнымі змяненнямі.
Калі вы жадаеце пабудаваць свае блогі на базе нашай версіі рухавіка, то пры адрозных ад letapis.VKL.world наладах файла `config.ini` мы не гарантуем стабільнасць яго працы і раім выкарыстоўваць апошнюю версію афіцыйнага [WriteFleery](https://github.com/writefreely/writefreely).

Праект напісаны на Go і выкарыстоўвае базу дадзеных MySQL (ці SQLLight).
Для стылей мы выкарыстоўваем LESS, таму вам спатрэбіцца Node.js, каб скапміляваць іх.

## Лакальнае разгортванне

### Прэрэквізіты

* Go 1.13+
* Node.js

### Кампіляцыя аплікацыі

1. Кланіруем рэпазіторый сабе лакальна ды пераходзім у папку
```
git clone https://github.com/it-beard/writefreely-vkl.git
cd writefreely
```

2.  Усталёўваем пакет `go-bindata` каб скампіляваць файл `bindata.go` (ён патрабуецца для паспяховага білда далей).
```
go get -u github.com/jteeuwen/go-bindata/go-bindata
go-bindata -pkg writefreely -ignore=\\.gitignore schema.sql sqlite.sql
```
2. 1. Вам можа спатрэбіцца самастойна скампіляваць `go-bindata` з зыходнікаў і занесці скампіляваны выканаўчы файл у папку са спісу пераменных асяроддзя PATH, каб каманда `go-bindata` запрацавала, бо каманда з пункту 2 працуе не ва ўсіх асяроддзях адразу. Вось прыкладны спіс каманд для кампіляцыі `go-bindata`:
   ```
   go get github.com/jteeuwen/go-bindata
   cd $GOPATH/src/github.com/jteeuwen/go-bindata/go-bindata
   go build
   ```
3. Далей трэба адкрыць атрыманы `bindata.go` файл і крыху падправіць у ім памылкі (чамусьці ў `bindata.go` з'яўляецца некалькі канфліктуючых імён функцый з аснаўным файлам аплікацыі `app.go`). Будзе дастаткова проста змяніць назвы дзвух функцый і два месцы іх выклікання ў `bindata.go`:
```
schemaSql змяняем на schemaSql1
sqliteSql змяняем на sqliteSql1
змяняем назвы функцый і у выкліках ніжэй (два месцы)
захоўваем файл
``` 

4. Кампілюем выканаўчы файл `writefreely` з падтрымкай SQLite. (Выдаліце `-tags='sqlite'` калі вам не патрэбна падтрымка SQLite.)
```
go build -v -tags='sqlite' ./cmd/writefreely/
```

Цяпер вы можаце запусціць WriteFreely! Але вам спатрэбіцца яшчэ адзін крок, каб стварыць некаторыя асеты і паспяхова запусціць свой лакальны інстанс.

### Першы запуск
#### Кампіляцыя асетаў
Па першае, трэба скампіяваць праз Node.js LESS-файлы, каб атрымаць выніковыя стылявыя css-файлы. Для гэтага патрабуецца LESS-кампілятар:
```
npm install -g less
```
Далей камілюем less, і кладзём вынікі ў папку `static/css/`:
```
cd less
export CSSDIR=../static/css/
lessc app.less --clean-css="--s1 --advanced" $CSSDIR/write.css
lessc fonts.less --clean-css="--s1 --advanced" $CSSDIR/fonts.css
lessc icons.less --clean-css="--s1 --advanced" $CSSDIR/icons.css
```

#### Канфігураванне ды першы запуск
Канфігураванне робіцца адзін раз, у выніку чаго будзе створаны файл `config.ini`. Менавіта ў ім знаходзяцца ўсялякія налады, у тым ліку налады MySQL і Let's Encrypt.

Стварыць пусты `config.ini` можна камандай:
```
./writefreely config generate
```

Стварыць `config.ini` праз кансольны аўта-канфігуратар можна камандай: 
```
./writefreely config start
```

Прыкладна так выглядае `config.ini` letapis.vkl.world, за выключэнням налад базы дадзеных і сертыфікатаў Let's Encrypt (вельмі раю наладзіць свой `config.ini` прыкладна падобным чынам, каб пазбегнуць памылак рантайма):
```ini
[server]
hidden_host          = 
port                 = 8080
bind                 = localhost
tls_cert_path        = 
tls_key_path         = 
autocert             = false
templates_parent_dir = 
static_parent_dir    = 
pages_parent_dir     = 
keys_parent_dir      = 
hash_seed            = 
gopher_port          = 0

[database]
type     = sqlite3
filename = writefreely.db
username = 
password = 
database = writefreely
host     = localhost
port     = 3306
tls      = false

[app]
site_name             = Летапісы ВКЛ
site_description      = Летапісы ВКЛ
host                  = http://localhost:8080
theme                 = write
editor                = 
disable_js            = false
webfonts              = true
landing               = 
simple_nav            = false
wf_modesty            = true
chorus                = true
forest                = false
disable_drafts        = false
single_user           = false
open_registration     = true
open_deletion         = true
min_username_len      = 2
max_blogs             = 3
federation            = true
public_stats          = true
monetization          = false
notes_only            = false
private               = false
local_timeline        = true
user_invites          = 
default_visibility    = public
update_checks         = false
disable_password_auth = false

[oauth.slack]
client_id          = 
client_secret      = 
team_id            = 
callback_proxy     = 
callback_proxy_api = 

[oauth.writeas]
client_id          = 
client_secret      = 
auth_location      = 
token_location     = 
inspect_location   = 
callback_proxy     = 
callback_proxy_api = 

[oauth.gitlab]
client_id          = 
client_secret      = 
host               = 
display_name       = 
callback_proxy     = 
callback_proxy_api = 

[oauth.gitea]
client_id          = 
client_secret      = 
host               = 
display_name       = 
callback_proxy     = 
callback_proxy_api = 

[oauth.generic]
client_id          = 
client_secret      = 
host               = 
display_name       = 
callback_proxy     = 
callback_proxy_api = 
token_endpoint     = 
inspect_endpoint   = 
auth_endpoint      = 
scope              = 
allow_disconnect   = false
map_user_id        = 
map_username       = 
map_display_name   = 
map_email          = 

```

Далей трэба стварыць прыватныя ключы, якія будуць выкарыстоўвацца для хэшыравання пароляў. **ВАЖНА**: гэтыя ключы ствараюцца адзін раз для бягучай базы дадзеных! Калі вы іх пераўтворыце, то старыя карыстальнікі страцяць магчымасць лагініцца (праз тое, што старыя паролі хэшыравалісь другімі ключамі). Самі ключы пасля генерацыі будуць знаходзіцца ў папцы `/keys`. Каманда на генерацыю:
```
./writefreely keys generate
```

Калі вы будзеце выкарыстоўваць MySQL у якасці базы дадзеных, то трэба наладзіць яе самастойна (усталяваць і стварыць базу дадзеных, унесці звесткі аб ёй у файл `config.ini`. Калі вы абралі карыстацца SQLight, то база створыцца аўтаматычна пры запуску каманды `writefreely config generate`.

Віншую! Цяпер вы можаце запусціць лакальны інстанс праз прамы выклік `./writefreely` з кансолі. Ваш сайт будзе даступны па спасылцы `http://localhost:8080/`. 


### Другі і наступныя запускі
Пасля змяненняў кода трэба перакампіляваць код, перкампіляваць less, накаціць міграцыі базы дадзеных і запусціць праект:
```
go build -v ./cmd/writefreely/
cd less
export CSSDIR=../static/css/
lessc app.less --clean-css="--s1 --advanced" $CSSDIR/write.css
lessc fonts.less --clean-css="--s1 --advanced" $CSSDIR/fonts.css
lessc icons.less --clean-css="--s1 --advanced" $CSSDIR/icons.css
cd ../
./writefreely db migrate
./writefreely

```

## Папкі ды файлы, якія патрэбны для працы аплікацыі (для сервераў)
Папкі ды файлы, якія нельга чапаць/перапісваць (ці рабіць гэта пільна):

* `/keys` - сгенераваныя для сервера унікальныя прыватныя ключы
* `/certs` - папка з сертыфікатамі Let’s Encrypt
* `config.ini` - файл канфігурацый сервера

Папкі ды файлы, якія трэба замяняць пры змяненні кода:

* `/static` - статычныя асеты (css, выявы, js-скрыпты, шрыфты)
* `/pages` - шаблоны службовых старонак (Go templates)
* `/templates` - шаблоны старонак (Go templates)
* `writefreely` - выканаўчы файл аплікацыі

## Лакалізацыя
Лакалізацыя праекту захардкожана у макетах "go templates" (`.tmpl` файлы) і ў go-кодзе. Нажаль, у арыгінальным WriteFreely не было механізмаў лакалізацыі на момант нашага форку, таму жывем з тым, што маем.

## Распрацоўка
* Каб далучыцца да распрацоўкі, пішыце [Аляксею АйЦіБарадзе](https://vkl.world/@itbeard)
* Распрацоўка вядзецца ў галіне `develop`
* Сайт https://letapis.vkl.world збіраецца з галіны `main`
* [Абмеркаванне праекта](https://github.com/it-beard/writefreely-vkl/discussions)
* [Бягучы спіс задач на распрацоўку](https://github.com/it-beard/writefreely-vkl/issues)

## Карысныя спасылкі

* Працуюцы інстанс, сабраны з `mine` галіны гэтага рэпазіторыя - https://letapis.vkl.world
* [Арыгінальная інструукцыя па наладах і запуску WriteFreely](https://writefreely.org/docs/latest/developer)
* [Арыгінальная інструкцыя па запуску на серверах](https://writefreely.org/start)
