# Formatter
small study about xlxs, csv, docker

## Parse, Read, reBuild
give it a xlxs, get out a xlxs clean with valid Csv files for each sheet
```bash
go run myroutine.go my.xlsx
```

### docker
docker on it, base on archlinux
```bash
docker build -t go_image .
docker run -it go_image
```

#### thanks to @tealeg
