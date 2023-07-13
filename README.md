## Агульная інфармацыя

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
Калі вы жадаеце пабудаваць свае блогі на базе нашай версыі рухавіка, то пры адрозных ад letapis.VKL.world наладах файла config.ini мы не гарантуем стабільнасць яго працы і раім выкарыстоўваць апошюю версыю афіцыйнага [WriteFleery](https://github.com/writefreely/writefreely).

Праект напісаны на Go і выкарыстоўвае базу дадзеных MySQL (ці SQLLight).
Для стылей мы выкарыстоўваем LESS, таму вам спатрэбіцца Node.js, каб скапміляваць іх.

## Лакальнае разгортванне

### Прэрэквізіты

* Go 1.13+
* Node.js

### Build аплікацыі

1. Кланіруем рэпазіторый сабе лакальна ды пераходзім у папку
```
git clone https://github.com/it-beard/writefreely-vkl.git
cd writefreely
```

2.  Усталёўваем пакет `go-bindata` каб скампіляваць фалй `bindata.go` (ён патрабуецца для паспяховага білда далей).
```
go get -u github.com/jteeuwen/go-bindata/go-bindata
go-bindata -pkg writefreely -ignore=\\.gitignore schema.sql sqlite.sql
```
2. 1. Вам можа спатрэбіцца самастойна скампіляваць `go-bindata` з зыходнікаў і занесці скампіляваны выканаўчы файл у папку са спісу пераменных асяродзя PATH, каб каманда `go-bindata` запрацавала, бо каманда з пункту 2 працуе не ва ўсіх асяродзях адразу. Вось прыкладны спіс каманд для кампіляцыі `go-bindata`:
   ```
   go get github.com/jteeuwen/go-bindata
   cd $GOPATH/src/github.com/jteeuwen/go-bindata/go-bindata
   go build
   ```
3. Далей трэба адкрыць атрыманы `bindata.go` файл і крыху падправіць у ім памылкі (чамусці ў `bindata.go` з'яўляецца некалькі канфліктуючык імён функцій з аснаўным файлам аплікацыіі `app.go`). Будзе дастаткова проста змяніць назвы дзвух функцый і два месцы іх выклікання ў `bindata.go`:
```
schemaSql змяняем на schemaSql1
sqliteSql змяняем на sqliteSql1
змяняем назвы функцый і у выкліках ніжэй (два месцы)
захоўваем файл
``` 

4. Білдзім выканаўчы файл `writefreely` з падтрымкай SQLite. (Выдаліце `-tags='sqlite'` калі вам не патрэбна падтрымка SQLite.)
```
go build -v -tags='sqlite' ./cmd/writefreely/
```

Цяпер вы можаце запусціць WriteFreely! Але вам спатрэбіцца яшчэ адзін крок, каб стварыць некаторыя ассэты і паспяхова запусціць свой лакальны інстанс.

### Канфігураванне (будзе дапоўнена)
### Запуск (будзе дапоўнена)


