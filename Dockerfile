FROM archlinux/base

MAINTAINER torrentcome

COPY . /app

RUN pacman -Syyu --noconfirm
RUN pacman -S go --noconfirm
RUN pacman -S git --noconfirm

WORKDIR /app

RUN go get github.com/tealeg/xlsx

# CMD ["/bin/bash"]
