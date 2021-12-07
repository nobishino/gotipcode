![Go Test](https://github.com/nobishino/gotipcode/actions/workflows/test.yml/badge.svg)
# gotipcode

- gotip用のコードを集める
- goのバージョンが上がったらgocodeリポジトリに移動する

## test

```sh
docker pull nobishino/gotip
docker run --volume `pwd`:/src nobishino/gotip /bin/bash -c "gotip test ./..."
```