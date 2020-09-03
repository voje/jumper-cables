# jumper-cables
Use: Jumper Cables will sometimes be able to shock a dead docker container back to life.   

```bash
docker-compose up
```
Check browser: `localhost:8000`.   
Click the `Kill` button to kill the unhealthy server.   
Jumper-cables should see that the server's in a bad state and restart the container.   

TODO: build the actual jumper-cables container

ENV: container_to_watch, timeout
