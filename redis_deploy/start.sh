/usr/local/redis/bin/redis-server redis_master.conf &
/usr/local/redis/bin/redis-server redis_slave.conf &
/usr/local/redis/bin/redis-sentinel sentinel_1.conf &
/usr/local/redis/bin/redis-sentinel sentinel_2.conf &
/usr/local/redis/bin/redis-sentinel sentinel_3.conf &
