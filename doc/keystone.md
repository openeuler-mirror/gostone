## keystone tables message


#### 1、 assignment

| 序号 |     名称     |  描述   |                              类型                              | 键  | 为空 | 额外 | 默认值 |
|:---:|:-----------:|:-------:|:-------------------------------------------------------------:|:---:|:---:|:---:|:-----:|
|  1  |   `type`    | 绑定类型 | enum('UserProject','GroupProject','UserDomain','GroupDomain') | PRI | NO  |      |       |
|  2  | `actor_id`  |  用户   |                          varchar(64)                          | PRI |  NO  |     |       |
|  3  | `target_id` |  项目   |                          varchar(64)                          | PRI |  NO  |     |       |
|  4  |  `role_id`  |         |                          varchar(64)                          | PRI | NO  |     |       |
|  5  | `inherited` |         |                          tinyint(1)                           | PRI | NO  |     |       |


#### 2、 endpoint

| 序号 |         名称          |        描述         |     类型      | 键  | 为空 | 额外 | 默认值 |
|:---:|:--------------------:|:-------------------:|:------------:|:---:|:---:|:---:|:-----:|
|  1  |         `id`         |                     | varchar(64)  | PRI | NO  |     |       |
|  2  | `legacy_endpoint_id` |                     | varchar(64)  |     | YES |     |       |
|  3  |     `interface`      |       节点类型       |  varchar(8)  |     |  NO  |     |       |
|  4  |     `service_id`     |       所属服务       | varchar(64)  | MUL |  NO  |     |       |
|  5  |        `url`         |         地址         |     text     |     | NO  |     |       |
|  6  |       `extra`        |       其他信息       |     text     |     | YES  |     |       |
|  7  |      `enabled`       | 是否可用 0不可用 1可用 |  tinyint(1)  |     | NO  |      |   1   |
|  8  |     `region_id`      |       所属区域       | varchar(255) | MUL | YES  |     |       |


#### 3、 local_user

| 序号 |         名称         |    描述    |     类型      | 键  | 为空 |      额外       | 默认值 |
|:---:|:-------------------:|:----------:|:------------:|:---:|:---:|:--------------:|:-----:|
|  1  |        `id`         |            |   int(11)    | PRI | NO  | auto_increment |       |
|  2  |      `user_id`      |   用户id    | varchar(64)  | UNI | NO  |                |       |
|  3  |     `domain_id`     | 所属domain  | varchar(64)  | MUL | NO  |                |       |
|  4  |       `name`        |   用户名    | varchar(255) |     | NO  |                |       |
|  5  | `failed_auth_count` | 验证失败次数 |   int(11)    |     | YES |                |       |
|  6  |  `failed_auth_at`   | 验证失败日期 |   datetime   |     | YES |                |       |


#### 4、 password

| 序号 |       名称        |      描述      |     类型      | 键  | 为空 |      额外       | 默认值 |
|:---:|:----------------:|:--------------:|:------------:|:---:|:---:|:--------------:|:-----:|
|  1  |       `id`       |                |   int(11)    | PRI | NO  | auto_increment |       |
|  2  | `local_user_id`  |                |   int(11)    | MUL | NO  |                |       |
|  3  |    `password`    |   国标加密密码   | varchar(128) |     | YES |                |       |
|  4  |   `expires_at`   |    过期时间     |   datetime   |     | YES |                |       |
|  5  |  `self_service`  |                |  tinyint(1)  |     | NO  |                |   0   |
|  6  | `password_hash`  | bycript加密密码 | varchar(255) |     | YES |                |       |
|  7  | `created_at_int` |                |  bigint(20)  |     | NO  |                |   0   |
|  8  | `expires_at_int` |                |  bigint(20)  |     | YES |                |       |
|  9  |   `created_at`   |                |   datetime   |     | NO  |                |       |


#### 5、 project

| 序号 |      名称      | 描述 |    类型     |  键  | 为空 | 额外 | 默认值 |
|:---:|:-------------:|:---:|:-----------:|:---:|:----:|:---:|:-----:|
|  1  |     `id`      |     | varchar(64) | PRI |  NO  |     |       |
|  2  |    `name`     |     | varchar(64) |     |  NO  |     |       |
|  3  |    `extra`    |     |    text     |     | YES  |     |       |
|  4  | `description` |     |    text     |     | YES  |     |       |
|  5  |   `enabled`   |     | tinyint(1)  |     | YES  |     |       |
|  6  |  `domain_id`  |     | varchar(64) | MUL |  NO  |     |       |
|  7  |  `parent_id`  |     | varchar(64) | MUL | YES  |     |       |
|  8  |  `is_domain`  |     | tinyint(1)  |     |  NO  |     |   0   |


#### 6、 region

| 序号 |        名称         | 描述 |     类型     |  键  | 为空 | 额外 | 默认值 |
|:---:|:------------------:|:---:|:------------:|:---:|:----:|:---:|:-----:|
|  1  |        `id`        |     | varchar(255) | PRI |  NO  |     |       |
|  2  |   `description`    |     | varchar(255) |     |  NO  |     |       |
|  3  | `parent_region_id` |     | varchar(255) |     | YES  |     |       |
|  4  |      `extra`       |     |     text     |     | YES  |     |       |


#### 7、 role

| 序号 |     名称     | 描述 |     类型     |  键  | 为空 | 额外 |  默认值   |
|:---:|:-----------:|:---:|:------------:|:---:|:----:|:---:|:--------:|
|  1  |    `id`     |     | varchar(64)  | PRI |  NO  |     |          |
|  2  |   `name`    |     | varchar(255) | MUL |  NO  |     |          |
|  3  |   `extra`   |     |     text     |     | YES  |     |          |
|  4  | `domain_id` |     | varchar(64)  |     |  NO  |     | <<null>> |


#### 8、 service

| 序号 |    名称    | 描述 |     类型     |  键  | 为空 | 额外 | 默认值 |
|:---:|:---------:|:---:|:------------:|:---:|:----:|:---:|:-----:|
|  1  |   `id`    |     | varchar(64)  | PRI |  NO  |     |       |
|  2  |  `type`   |     | varchar(255) |     | YES  |     |       |
|  3  | `enabled` |     |  tinyint(1)  |     |  NO  |     |   1   |
|  4  |  `extra`  |     |     text     |     | YES  |     |       |


#### 9、 user

| 序号 |         名称          | 描述 |    类型     |  键  | 为空 | 额外 | 默认值 |
|:---:|:--------------------:|:---:|:-----------:|:---:|:----:|:---:|:-----:|
|  1  |         `id`         |     | varchar(64) | PRI |  NO  |     |       |
|  2  |       `extra`        |     |    text     |     | YES  |     |       |
|  3  |      `enabled`       |     | tinyint(1)  |     | YES  |     |       |
|  4  | `default_project_id` |     | varchar(64) | MUL | YES  |     |       |
|  5  |     `created_at`     |     |  datetime   |     | YES  |     |       |
|  6  |   `last_active_at`   |     |    date     |     | YES  |     |       |
|  7  |     `domain_id`      |     | varchar(64) | MUL |  NO  |     |       |


