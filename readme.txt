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


// Run api server by Nodemon
nodemon --exec go run main.go






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
//example ::: Adaptive bitrate
ffmpeg -y -i filename.mp4 \
  -vf scale=w=640:h=360:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod  -b:v 800k -maxrate 856k -bufsize 1200k -b:a 96k -hls_segment_filename 360p_%03d.ts 360p.m3u8 \
  -vf scale=w=842:h=480:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 1400k -maxrate 1498k -bufsize 2100k -b:a 128k -hls_segment_filename 480p_%03d.ts 480p.m3u8 \
  -vf scale=w=1280:h=720:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 2800k -maxrate 2996k -bufsize 4200k -b:a 128k -hls_segment_filename 720p_%03d.ts 720p.m3u8 \
  -vf scale=w=1920:h=1080:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 5000k -maxrate 5350k -bufsize 7500k -b:a 192k -hls_segment_filename 1080p_%03d.ts 1080p.m3u8

//https://www.createwithswift.com/converting-video-files-for-hls-streaming/

ffmpeg -y -i hotd.mp4 -vf scale=w=640:h=360:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod  -b:v 800k -maxrate 856k -bufsize 1200k -b:a 96k -hls_segment_filename 360p_%03d.ts 360p.m3u8 -vf scale=w=1920:h=1080:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 5000k -maxrate 5350k -bufsize 7500k -b:a 192k -hls_segment_filename 1080p_%03d.ts 1080p.m3u8


---- คำสั่ง --- แปลงไฟล์ สร้าง thumbnail =>jpeg
//example ::: แปลงในทุกๆ 20 วินาที
ffmpeg -i test.mp4 -vf fps=1/20 thumb%04d.png
//example ::: แปลงในทุกๆ 5 วินาที
ffmpeg -i hotd.mp4 -vf fps=1/5 hotd_thumb%d.jpeg;  