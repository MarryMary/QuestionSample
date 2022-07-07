import redis

def ReadSession(Key):
    r = redis.Redis(host='localhost', port=6379, db=0)
    PHPSession = r.smembers("PHPREDIS_SESSION:"+Key)
    print(PHPSession)