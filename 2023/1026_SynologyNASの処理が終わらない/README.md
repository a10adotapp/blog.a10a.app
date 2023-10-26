2時間弱動き続けている重たい処理があるようだ。動画のサムネイル作成っぽい。

``` shell
$ top
top - 19:11:40 up 1 day, 11:43,  1 user,  load average: 8.26, 9.40, 9.69 [IO: 7.02, 7.92, 8.10 CPU: 1.24, 1.48, 1.59]
Tasks: 276 total,   2 running, 274 sleeping,   0 stopped,   0 zombie
%Cpu(s):  3.1 us,  1.6 sy, 25.0 ni, 14.1 id, 56.2 wa,  0.0 hi,  0.0 si,  0.0 st
GiB Mem :    0.473 total,    0.078 free,    0.295 used,    0.100 buff/cache
GiB Swap:    2.000 total,    1.415 free,    0.585 used.    0.078 avail Mem

PID   USER  PR  NI   VIRT    RES   %CPU   %MEM      TIME+  S  COMMAND
9503  root  30  10  96.8m  36.2m  84.97  7.469  119:18.07  R  /usr/syno/bin/ffmpeg-thumb -threads 1 -i /volume1/photo/xxxxxxxxxx/IMG_3275.MOV -threads 1 -y -vcodec libx264 -preset superfast -vprofile baseline -level 30 -b:v 2500k -bt 2500k -r 15 -pix_fmt yuv420p -f mp4 -acodec copy -s 720x1280 -aspect 720:1280 /volume1/photo/xxxxxxxxxx/@eaDir/IMG_3275.MOV/SYNOPHOTO_FILM_M.mp4.tmp
```

ちなみに対象のファイルを見てみると7.9Gある。デカすぎる。

``` shell
~$ du -h /volume1/photo/xxxxxxxxxx/IMG_3275.MOV
7.9G	/volume1/photo/xxxxxxxxxx/IMG_3275.MOV
```

どうしたものか、、、

-----

> 尚、移行データに画像ファイルが多かったのが原因なのかサムネイルを作るプロセス「ffmpeg-thumb」がCPUリソースを消費しまくり、長時間CPU使用率99%の状況が続きました。1-2日（機種によっては1週間）で収まるという話ですが、画像の管理は見送りました。機会があれば少しづつ追加してみようと考えています。
> 
> - https://blog.8wired.jp/gadget/review-synology-diskstation-ds210/

**（機種によっては1週間）**

ええ、、、

-----

大量のメディアファイルをアップロードするときはSynology Photosの管理外に一旦置いた上で、
ストレージアナライザで重複をチェックして、重複ファイルがあるなら削除するなり、大きい動画ファイルはリサイズするなりして
Synology Photosの管理フォルダへ移動するのがいいかもしれない。

-----

そうこうしてるうちに先述のプロセスの実行時間が3時間を超えた
