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