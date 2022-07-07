from phpserialize import loads
import redis

def ReadSession(Key):
    r = redis.Redis(host='localhost', port=6379, db=0)
    PHPSession = r.get("PHPREDIS_SESSION:"+Key)
    SessionData = loads(PHPSession)
    return SessionData