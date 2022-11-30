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
ffmpeg -i filename.mp4 -codec: copy -start_number 0 -hls_time 10 -hls_list_size 0 -f hls filename.m3u8