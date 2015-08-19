# copyright-protector

usage:

```bash
git clone git@github.com:dlwr/copyright-protector.git
cd copyright-protector
heroku create -b https://github.com/kr/heroku-buildpack-go.git
git push heroku master
```

parameter:
- `url={url(encoded is better)}` image url
- `glitch=true` glitch effect
   - e.g. http://copyright-protector.herokuapp.com/protected.png?url=http://41.media.tumblr.com/PwCIEKd8N2oro4neq2zBCcLi_400.jpg&glitch=true
   - ![glitched](http://copyright-protector.herokuapp.com/protected.png?url=http://41.media.tumblr.com/PwCIEKd8N2oro4neq2zBCcLi_400.jpg&glitch=true)
- `mozaic=true` mozaic effect
   - e.g. http://copyright-protector.herokuapp.com/protected.png?url=http://41.media.tumblr.com/PwCIEKd8N2oro4neq2zBCcLi_400.jpg&mozaic=true
   - ![mozaiced](http://copyright-protector.herokuapp.com/protected.png?url=http://41.media.tumblr.com/PwCIEKd8N2oro4neq2zBCcLi_400.jpg&mozaic=true)
- `tile=true` tile line effect
   - e.g. http://copyright-protector.herokuapp.com/protected.png?url=http://41.media.tumblr.com/PwCIEKd8N2oro4neq2zBCcLi_400.jpg&mozaic=true&tile=true
   - ![tiled](http://copyright-protector.herokuapp.com/protected.png?url=http://41.media.tumblr.com/PwCIEKd8N2oro4neq2zBCcLi_400.jpg&mozaic=true&tile=true)

#### Products which use this protector

- [copy-right protect](http://let.hatelabo.jp/yuta25/let/hJmf5_3V9vtf) (bookmark let)
