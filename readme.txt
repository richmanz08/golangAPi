การลง Redis on macOs
1.mkdir redis && cd redis
2.curl -O http://download.redis.io/redis-stable.tar.gz
3.ar xzvf redis-stable.tar.gz
4.cd redis-stable
5.make
6.make test
7.sudo make install



การ Start Server 
cd เข้าไป redis/redis-stable แล้ว redis-server










การติดตั้ง ffmpeg 
1. Download zip https://www.gyan.dev/ffmpeg/builds/ffmpeg-git-full.7z
2. แตกไฟล์ และแก้ไขชื่อ folder ให้ง่ายเป็น ffmpeg
3. ใช้ cmd Adminstrator :> setx /m PATH "C:\ffmpeg\bin;%PATH%"
4. Restart computer
5. check version =>  ใช้ cmd Adminstrator :> ffmpeg -version
---- คำสั่ง --- แปลงไฟล์ mp4=> to m3u8 and ts
//example ::: แบบ ultra quality from video
ffmpeg -i filename.mp4 -codec: copy -start_number 0 -hls_time 10 -hls_list_size 0 -f hls filename.m3u8 
//example ::: แบบ เลือก quaility 144,360p, 720p,1080p,
ffmpeg -i input.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls index.m3u8
//example ::: แบบ fix 144p
ffmpeg -i godzilla-video.mp4 -profile:v baseline -level 3.0 -s 144x100 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls godquality_hd.m3u8

---- คำสั่ง --- แปลงไฟล์ สร้าง thumbnail =>jpeg
//example ::: แปลงในทุกๆ 20 วินาที
ffmpeg -i test.mp4 -vf fps=1/20 thumb%04d.png
//example ::: แปลงในทุกๆ 5 วินาที
ffmpeg -i hotd.mp4 -vf fps=1/5 hotd_thumb%d.jpeg;  