# auth-session-redis

Redis session store for togo [auth](https://github.com/togo-framework/auth)
(`SESSION_DRIVER=redis`). Revocable, shared across instances.

```bash
togo install togo-framework/auth
togo install togo-framework/auth-session-redis
```

Env: `REDIS_URL` (e.g. redis://localhost:6379/0).
