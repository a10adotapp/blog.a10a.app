`blog.a10a.app`へのアクセスをこちらのリポジトリへリダイレクトするようにした。

DNSのCNAMEレコードに書くのが正しいのかわからなかったのでリダイレクト用のサーバーを立てた。

nginxの設定書くよりGoでサーバー書いてしまった方が早いのありがたい。

https://github.com/a10adotapp/blog.a10a.app/blob/2e95a172470282eee823e0799b38adae208a6ecc/2023/1025_redirect/main.go#L1-L9

Cloudflareにもリダイレクト設定があるようだということには全て終わってから気づいた。

そちらに変えるかもしれないし変えないかもしれない。
