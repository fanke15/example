## Ghost Blog Themes




### Ghost Dcoker

```yaml

version: '1.0'

services:
  ghost:
    image: ghost:5.49.3-alpine
    restart: always
    volumes:
      - /www/wwwroot/blog/content:/var/lib/ghost/content
    ports:
      - 9000:2368
    environment:
      database__client: mysql
      database__connection__host: 127.0.0.1
      database__connection__user: blog
      database__connection__password: fanke
      database__connection__database: blog

```






