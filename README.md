# Overview

- practice with golang and graphql
- used [this site](https://future-architect.github.io/articles/20200609/) as a reference
- I'm still getting used to the GraphQL schema design

# メモ

#### docker 接続確認

```
$ PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -d postgres -c 'select * from station limit 10';
```

#### postgres の sslmode クエリがエラーになる

xo を実行すると以下のエラーが表示される

```
zsh: no matches found: pgsql://postgres:postgres@localhost/postgres?sslmode=disable
```

こちらで解消方法が記載されていた。zsh の問題で、ダブルクォートでクエリを囲むと解消できる。<br/>
https://github.com/golang-migrate/migrate/issues/290
