const phpSerializer = require('php-serialize');


function session(req, res, next) {
    if (!req.cookies) {
      console.error('must use cookie-parser middleware');
      return next();
    }
  
    const sessionId = req.cookies['PHPSESSID'];
    if (!sessionId) {
      return next();
    }
  
    const sessionKey = 'PHPREDIS_SESSION:' + sessionId;
    redis.getAsync(sessionKey)
      .then((sessionStr) => {
        if (sessionStr) {
          req.session = new Session(sessionKey, sessionStr);
        }
        return next();
      })
      .catch((err) => {
        return next(new Error('Server error'));
      });
    };
  
  function Session(sessionId, sessionStr) {
    this.id = sessionId;
  
    let data;
    try {
      data = phpSerializer.unserialize(sessionStr);
    } catch(e) {
      return;
    }
  
    for (const prop in data) {
      if (!(prop in this)) {
        this[prop] = data[prop];
      }
    }
  }