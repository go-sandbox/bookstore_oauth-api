# OAuth キースペース

```
create keyspace oauth with replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
```

## OAuthテーブル

```
create table access_token(access_token varchar primary key, user_id bigint, client_id bigint, expires bigint);
```
