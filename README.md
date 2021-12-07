# gotopcode

- gotip用のコードを集める
- goのバージョンが上がったらgocodeリポジトリに移動する

## test

```sh
docker pull nobishino/gotip
docker run --volume `pwd`:/src nobishino/gotip /bin/bash -c "gotip test ./..."
```