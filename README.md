# go-prrrrr

[![.github/workflows/test.yml](https://github.com/tjhu/go-prrrrr/actions/workflows/test.yml/badge.svg)](https://github.com/tjhu/go-prrrrr/actions/workflows/test.yml)

## How to run

To test

```bash
go test ./stream
```

To benchmark

```bash
go test ./benchmark -bench=. -cpu 1 
```

## How to build 

Install stuff(come on js people, be nice to others).

```bash
npm install --global mermaid-filter
npm install -g mermaid.cli
sudo apt install -y gconf-service libasound2 libatk1.0-0 libc6 libcairo2 libcups2 libdbus-1-3 libexpat1 libfontconfig1 libgcc1 libgconf-2-4 libgdk-pixbuf2.0-0 libglib2.0-0 libgtk-3-0 libnspr4 libpango-1.0-0 libpangocairo-1.0-0 libstdc++6 libx11-6 libx11-xcb1 libxcb1 libxcomposite1 libxcursor1 libxdamage1 libxext6 libxfixes3 libxi6 libxrandr2 libxrender1 libxss1 libxtst6 ca-certificates fonts-liberation libnss3 lsb-release xdg-utils wget
```

Then run

```
make
```

## Known issues

* This repo is full of hacks because of the tight deadline. Although there're some cool stuff here, don't take the architecture of this codebase too seriously.
* Forking a stream(a list instead of a DAG) is not supported.
* The source stream will be modified.
