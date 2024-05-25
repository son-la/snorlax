# snorlax

## TODO
1. Testing, testing and more testing! 
2. CI pipeline
3. Develop snorlax UI with registration/ login functionality
4. Integrate snorlax with Okta for SSO



## Development environment

1. Setup local k8s for dependencies
* Refer to `https://github.com/son-la/snorlax-devsetup`
* Port forwarding mysql
2. Create `config.yaml` 
```
---
kafka: 
  brokers:
  - "localhost:30002"

  useTLS: true
  caFile: "ca.crt"
  version: "3.7.0"

  authentcation:
    username: KAFKA_USER
    password: KAFKA_PASSWORD
    algorithm: sha512

  topic: test
  
database:
  host: MYSQL_HOST
  port: FORWARDED_PORT
  database: USER_DB
  username: USERNAME
  password: PASSWORD
```
