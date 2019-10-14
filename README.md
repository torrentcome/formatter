# formatter
small study case about xlxs, csv, docker

### docker
docker on it, base on archlinux
```bash
docker build -t go_image .
docker run -it go_image
```

### why ?
have a routine who fix pain head for every PO who want to transform xlsx file to csv file

### what ?
parse, read, build, give it a xlxs, get out a xlxs clean with valid csv files for each sheet
```bash
go run myroutine.go my.xlsx
```
### how ?
golang is a great language, so we use it...

#### thanks to @tealeg for the great lib
