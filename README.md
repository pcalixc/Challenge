# Challenge: Enron e-mail indexer 📧


This application allows you to search the ENRON Corp. e-mail database. It is indexed in ZincSearch search engine.


## Technologies Used

- Backend: [Go (Golang)](https://go.dev/)
- API Router: [Chi ](https://github.com/go-chi/chi)
- Search Engine: [ZincSearch](https://zincsearch.com/)
- Frontend: [Vue](https://vuejs.org/) + [JavaScript](https://developer.mozilla.org/es/docs/Web/JavaScript)


## ZincSearch

### Steps

- Download [Enron Corp's email database](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz )

- Download and install the [ZincSearch](https://github.com/zinclabs/zinc/releases) tool.


**`Version used:`**  https://github.com/zinclabs/zinc/releases/download/v0.3.5/zinc_0.3.5_Linux_x86_64.tar.gz

- Configure [ZincSearch](https://docs.zincsearch.com/installation/) locally. 

```bash
    set ZINC_FIRST_ADMIN_USER=admin
    set ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123
    mkdir data
    zinc.exe
```

## Project implementation

- Backend 

```bash
    cd Challenge\\Backend\\cmd
    go run main.go
```

- Frontend 

```bash
    cd Challenge\\frontend
    npm install
    npm run serve
```



