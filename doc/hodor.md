# GoStone


**简介**: GoStoneAPI


**HOST**:127.0.0.1:8100

**Version**:0.1

# 通用参数说明

Links:

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|next|下一个的地址|string||
|previous|上一个地址|string||
|self|自己的查询地址|string||

# Role管理

Role

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|id|角色ID|string||
|description|描述信息|string||
|name|角色名|string||
|links|相关链接|schema|Links|
|domain_id|所属Domain|string||



## 查询role集合


**接口地址**:`/v3/roles`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|domainId|domain的ID|query|false|string||
|name|角色名称|query|false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Role|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|links|相关地址|schema|Links|
|roles|角色集合|array|Role|


**响应示例**:
```json
{  
    "links": {
        "next": null,
        "previous": null,
        "self": "http://example.com/identity/v3/roles"
    },
    "roles": [
        {
            "id": "5318e65d75574c17bf5339d3df33a5a3",
            "links": {
                "self": "http://example.com/identity/v3/roles/5318e65d75574c17bf5339d3df33a5a3"
            },
            "description": "My new role",
            "name": "admin"
        }
    ]
}
```


## 创建role


**接口地址**:`/v3/roles`


**请求方式**:`POST`


**请求数据类型**:`application/json`



**接口描述**:


**请求示例**:


```json
{
	"role": {
		 "description": "My new role",
          "name": "developer",
          "domain_id": "12345asdqweqwe"
	}
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|role|role|body|true|||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created|Role|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
    "role": {
        "description": "My new role",
        "domain_id": "92e782c4988642d783a95f4a87c3fdd7",
        "name": "developer",
        "links": {
            "self": "http://example.com/identity/v3/roles/1e443fa8cee3482a8a2b6954dd5c8f12"
        }
    }
}
```


## 根据id查询role


**接口地址**:`/v3/roles/{role_id}`


**请求方式**:`GET`


**请求数据类型**:`*`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|role_id|role_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Role|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:





**响应示例**:
```json
{
    "role": {
        "domain_id": "d07792fd66ac4ed881723ab9f1c9925f",
        "id": "1e443fa8cee3482a8a2b6954dd5c8f12",
        "links": {
            "self": "http://example.com/identity/v3/roles/1e443fa8cee3482a8a2b6954dd5c8f12"
        },
        "description": "My new role",
        "name": "Developer"
    }

}
```


## 删除role


**接口地址**:`/v3/roles/{role_id}`


**请求方式**:`DELETE`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|role_id|role_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


暂无



## 更新role


**接口地址**:`/v3/roles/{role_id}`


**请求方式**:`PATCH`


**请求数据类型**:`application/json`



**接口描述**:


**请求示例**:


```json
{
	"role": {
    		 "description": "My new role",
              "name": "developer",
              "domain_id": "12345asdqweqwe"
    }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|role|role|body|true|||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Role|
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


**响应示例**:
```json
{
	"role": {
            "description": "My new role",
            "domain_id": "92e782c4988642d783a95f4a87c3fdd7",
            "name": "developer"
    }
}
```


# User管理

User

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|id|用户ID|string||
|name|用户名|string||
|enable|是否可用|bool||
|links|相关链接|schema|Links|
|domain_id|所属Domain|string|
|description|描述信息|string||
|email|邮箱|string||


## 查询user集合


**接口地址**:`/v3/users`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|domainId|所属domain|query|false|string||
|enabled|是否可用|query|false|boolean||
|name|用户名|query|false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|用户|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|links|相关链接|schema|Links|
|users|用户列表|schema|User|



**响应示例**:
```json
{
    "links": {
        "next": null,
        "previous": null,
        "self": "http://example.com/identity/v3/users"
    },
    "users": [
        {
            "domain_id": "default",
            "enabled": true,
            "id": "2844b2a08be147a08ef58317d6471f1f",
            "links": {
                "self": "http://example.com/identity/v3/users/2844b2a08be147a08ef58317d6471f1f"
            },
            "name": "glance"
        }
    ]
}
```


## 创建user


**接口地址**:`/v3/users`


**请求方式**:`POST`


**请求数据类型**:`application/json`


**接口描述**:


**请求示例**:


```json
{
	"user": {
            "default_project_id": "263fd9",
            "domain_id": "1789d1",
            "enabled": true,
            "name": "James Doe",
            "password": "secretsecret",
            "description": "James Doe user",
            "email": "jdoe@example.com"
        }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|user|user|body|true|||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created|用户|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
	"user": {
            "default_project_id": "263fd9",
            "description": "James Doe user",
            "domain_id": "1789d1",
            "email": "jdoe@example.com",
            "enabled": true,
            "id": "ff4e51",
            "name": "James Doe",
            "links": {
               "self": "https://example.com/identity/v3/users/ff4e51"
            }
        }
}
```


## 根据id查询user


**接口地址**:`/v3/users/{user_id}`


**请求方式**:`GET`



**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|user_id|user_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|用户|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
	"user": {
            "default_project_id": "263fd9",
            "domain_id": "1789d1",
            "enabled": true,
            "id": "9fe1d3",
            "links": {
                "self": "https://example.com/identity/v3/users/9fe1d3"
            },
            "name": "jsmith"
        }
}
```


## 删除user


**接口地址**:`/v3/users/{user_id}`


**请求方式**:`DELETE`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|user_id|user_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


暂无



## 更新user


**接口地址**:`/v3/users/{user_id}`


**请求方式**:`PATCH`



**接口描述**:


**请求示例**:


```json
{
	"user": {
                "default_project_id": "263fd9",
                "domain_id": "1789d1",
                "enabled": true,
                "name": "James Doe",
                "password": "secretsecret",
                "description": "James Doe user",
                "email": "jdoe@example.com"
     }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|user|user|body|true|schema|user|
|user_id|user_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|用户|
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


**响应示例**:
```json
{
	"user": {
                "default_project_id": "263fd9",
                "domain_id": "1789d1",
                "enabled": true,
                "id": "9fe1d3",
                "links": {
                    "self": "https://example.com/identity/v3/users/9fe1d3"
                },
                "name": "jsmith"
       }
}
```


## 修改密码


**接口地址**:`/v3/users/{user_id}/password`


**请求方式**:`POST`


**请求数据类型**:`application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```json
{
	"user": {
            "password": "new_secretsecret",
            "original_password": "secretsecret"
        }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|password|新密码|body|true|string||
|original_password|老密码|body|true|string||
|user_id|user_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created||
|204|No Content||
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


暂无


**响应示例**:
```json

```


# domain管理

Domain

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|id|项目ID|string||
|name|项目名|string||
|enable|是否可用|bool||
|links|相关链接|schema|Links|
|domain_id|所属Domain|string|
|description|描述信息|string||


## 查询domain集合


**接口地址**:`/v3/domains`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|enabled|是否可用|query|false|boolean||
|name|domain名|query|false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Project|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|links|相关链接|schema|Links|
|domains|domain列表|schema|Domain|



**响应示例**:
```json
{
"domains": [
        {
            "description": "Owns users and tenants (i.e. projects) available on Identity API v2.",
            "enabled": true,
            "id": "default",
            "links": {
                "self": "http://example.com/identity/v3/domains/default"
            },
            "name": "Default"
        }
    ],
    "links": {
        "next": null,
        "previous": null,
        "self": "http://example.com/identity/v3/domains"
    }
}
```


## 创建domain


**接口地址**:`/v3/domains`


**请求方式**:`POST`

**接口描述**:


**请求示例**:


```json
{
	 "domain": {
            "description": "Domain description",
            "enabled": true,
            "name": "myDomain"
      }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|domain|domain|body|true||
|&emsp;&emsp;description|描述信息|body|N|string|
|&emsp;&emsp;enabled|是否可用|body|N|bool|
|&emsp;&emsp;name|domain名|body|N|string|




**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created|Project|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:



**响应示例**:
```json
{
    "domain": {
        "description": "Owns users and tenants (i.e. projects) available on Identity API v2.",
        "enabled": true,
        "id": "default",
        "links": {
            "self": "http://example.com/identity/v3/domains/default"
        },
        "name": "Default"
    }
}
```


## 根据id查询domain


**接口地址**:`/v3/domains/{domain_id}`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|domain_id|domain_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Project|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
    "domain": {
        "description": "Owns users and tenants (i.e. projects) available on Identity API v2.",
        "enabled": true,
        "id": "default",
        "links": {
            "self": "http://example.com/identity/v3/domains/default"
        },
        "name": "Default"
    }
}
```


## 更新domain


**接口地址**:`/v3/domains/{domain_id}`


**请求方式**:`PATCH`


**接口描述**:


**请求示例**:


```json
{
	 "domain": {
            "description": "Domain description",
            "enabled": true,
            "name": "myDomain"
      }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|domain|domain|body|true|||
|domain_id|domain_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Project|
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


**响应示例**:
```json
{
    "domain": {
        "description": "Owns users and tenants (i.e. projects) available on Identity API v2.",
        "enabled": true,
        "id": "default",
        "links": {
            "self": "http://example.com/identity/v3/domains/default"
        },
        "name": "Default"
    }
}
```


## 删除domain


**接口地址**:`/v3/domains/{id}`


**请求方式**:`DELETE`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id|id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK||
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


暂无



# token-controller


## adminGetToken


**接口地址**:`/v3/admin/tokens`


**请求方式**:`POST`


**请求数据类型**:`application/json`


**响应数据类型**:`application/json;charset=UTF-8`


**接口描述**:


**请求示例**:


```json
{
	"project_domain_id": "",
	"project_id": "",
	"project_name": "",
	"user_id": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|adminGetToken|adminGetToken|body|true|AdminGetToken|AdminGetToken|
|&emsp;&emsp;project_domain_id|项目所属Domain(如果projectId不为空时可不填写)||false|string||
|&emsp;&emsp;project_id|项目ID(项目名二选一)||false|string||
|&emsp;&emsp;project_name|项目名(和项目ID二选一)||false|string||
|&emsp;&emsp;user_id|用户ID||false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created||
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:

| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|X-Subject-Token|生成的用户token|header|true|string||



## validateToken


**接口地址**:`/v3/auth/tokens`


**请求方式**:`GET`


**请求数据类型**:`*`


**响应数据类型**:`application/json;charset=UTF-8`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|X-Subject-Token|要验证的token|header|true|string||
|X-Auth-Token|带有验证权限的token|header|true|string||
|allow_expired|allow_expired|query|false|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Token|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- |
|token||schema|Token|
|&emsp;&emsp;catalog||array|Catalog|
|&emsp;&emsp;&emsp;&emsp;endpoints||array|节点|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;enabled|是否可用||false|boolean||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;extra|其他信息||false|object||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;id|id||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;interface|节点类型||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;region|||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;region_id|所属区域||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;service_id|所属服务||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;url|访问url||false|string||
|&emsp;&emsp;&emsp;&emsp;id||string||
|&emsp;&emsp;&emsp;&emsp;name||string||
|&emsp;&emsp;&emsp;&emsp;type||string||
|&emsp;&emsp;expires_at||string(date-time)|string(date-time)|
|&emsp;&emsp;is_domain||boolean||
|&emsp;&emsp;issued_at||string(date-time)|string(date-time)|
|&emsp;&emsp;methods||array||
|&emsp;&emsp;project||Project|Project|
|&emsp;&emsp;&emsp;&emsp;domain||Domain|Domain|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;id|||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;name|||false|string||
|&emsp;&emsp;&emsp;&emsp;id||string||
|&emsp;&emsp;&emsp;&emsp;name||string||
|&emsp;&emsp;roles||array|Role|
|&emsp;&emsp;&emsp;&emsp;id||string||
|&emsp;&emsp;&emsp;&emsp;name||string||
|&emsp;&emsp;user||User|User|
|&emsp;&emsp;&emsp;&emsp;domain||Domain|Domain|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;id|||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;name|||false|string||
|&emsp;&emsp;&emsp;&emsp;email||string||
|&emsp;&emsp;&emsp;&emsp;id||string||
|&emsp;&emsp;&emsp;&emsp;name||string||
|&emsp;&emsp;&emsp;&emsp;password_expires_at||string(date-time)||


**响应示例**:
```json
{
  "token": {
    "catalog": [
		{
			"endpoints": [
				{
					"enabled": true,
					"extra": {},
					"id": "",
					"interface": "",
					"region": "",
					"region_id": "",
					"service_id": "",
					"url": ""
				}
			],
			"id": "",
			"name": "",
			"type": ""
		}
	],
	"expires_at": "",
	"is_domain": true,
	"issued_at": "",
	"methods": [],
	"project": {
		"domain": {
			"id": "",
			"name": ""
		},
		"id": "",
		"name": ""
	},
	"roles": [
		{
			"id": "",
			"name": ""
		}
	],
	"user": {
		"domain": {
			"id": "",
			"name": ""
		},
		"email": "",
		"id": "",
		"name": "",
		"password_expires_at": ""
	}

}
}
```


## issueToken


**接口地址**:`/v3/auth/tokens`


**请求方式**:`POST`



**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|auth|验证信息|body|true|schema|Auth|
|&emsp;&emsp;identity|验证信息|body|true|schema|Identity|

# GoStone


**简介**: GoStoneAPI


**HOST**:127.0.0.1:8100

**Version**:0.1

#通用参数说明

Links:

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|next|下一个的地址|string||
|previous|上一个地址|string||
|self|自己的查询地址|string||

# Role管理

Role

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|id|角色ID|string||
|description|描述信息|string||
|name|角色名|string||
|links|相关链接|schema|Links|
|domain_id|所属Domain|string||



## 查询role集合


**接口地址**:`/v3/roles`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|domainId|domain的ID|query|false|string||
|name|角色名称|query|false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Role|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|links|相关地址|schema|Links|
|roles|角色集合|array|Role|


**响应示例**:
```json
{  
    "links": {
        "next": null,
        "previous": null,
        "self": "http://example.com/identity/v3/roles"
    },
    "roles": [
        {
            "id": "5318e65d75574c17bf5339d3df33a5a3",
            "links": {
                "self": "http://example.com/identity/v3/roles/5318e65d75574c17bf5339d3df33a5a3"
            },
            "description": "My new role",
            "name": "admin"
        }
    ]
}
```


## 创建role


**接口地址**:`/v3/roles`


**请求方式**:`POST`


**请求数据类型**:`application/json`



**接口描述**:


**请求示例**:


```json
{
	"role": {
		 "description": "My new role",
          "name": "developer",
          "domain_id": "12345asdqweqwe"
	}
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|role|role|body|true|||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created|Role|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
    "role": {
        "description": "My new role",
        "domain_id": "92e782c4988642d783a95f4a87c3fdd7",
        "name": "developer",
        "links": {
            "self": "http://example.com/identity/v3/roles/1e443fa8cee3482a8a2b6954dd5c8f12"
        }
    }
}
```


## 根据id查询role


**接口地址**:`/v3/roles/{role_id}`


**请求方式**:`GET`


**请求数据类型**:`*`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|role_id|role_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Role|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:





**响应示例**:
```json
{
    "role": {
        "domain_id": "d07792fd66ac4ed881723ab9f1c9925f",
        "id": "1e443fa8cee3482a8a2b6954dd5c8f12",
        "links": {
            "self": "http://example.com/identity/v3/roles/1e443fa8cee3482a8a2b6954dd5c8f12"
        },
        "description": "My new role",
        "name": "Developer"
    }

}
```


## 删除role


**接口地址**:`/v3/roles/{role_id}`


**请求方式**:`DELETE`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|role_id|role_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


暂无



## 更新role


**接口地址**:`/v3/roles/{role_id}`


**请求方式**:`PATCH`


**请求数据类型**:`application/json`



**接口描述**:


**请求示例**:


```json
{
	"role": {
    		 "description": "My new role",
              "name": "developer",
              "domain_id": "12345asdqweqwe"
    }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|role|role|body|true|||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Role|
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


**响应示例**:
```json
{
	"role": {
            "description": "My new role",
            "domain_id": "92e782c4988642d783a95f4a87c3fdd7",
            "name": "developer"
    }
}
```


# User管理

User

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|id|用户ID|string||
|name|用户名|string||
|enable|是否可用|bool||
|links|相关链接|schema|Links|
|domain_id|所属Domain|string|
|description|描述信息|string||
|email|邮箱|string||


## 查询user集合


**接口地址**:`/v3/users`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|domainId|所属domain|query|false|string||
|enabled|是否可用|query|false|boolean||
|name|用户名|query|false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|用户|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|links|相关链接|schema|Links|
|users|用户列表|schema|User|



**响应示例**:
```json
{
    "links": {
        "next": null,
        "previous": null,
        "self": "http://example.com/identity/v3/users"
    },
    "users": [
        {
            "domain_id": "default",
            "enabled": true,
            "id": "2844b2a08be147a08ef58317d6471f1f",
            "links": {
                "self": "http://example.com/identity/v3/users/2844b2a08be147a08ef58317d6471f1f"
            },
            "name": "glance"
        }
    ]
}
```


## 创建user


**接口地址**:`/v3/users`


**请求方式**:`POST`


**请求数据类型**:`application/json`


**接口描述**:


**请求示例**:


```json
{
	"user": {
            "default_project_id": "263fd9",
            "domain_id": "1789d1",
            "enabled": true,
            "name": "James Doe",
            "password": "secretsecret",
            "description": "James Doe user",
            "email": "jdoe@example.com"
        }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|user|user|body|true|||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created|用户|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
	"user": {
            "default_project_id": "263fd9",
            "description": "James Doe user",
            "domain_id": "1789d1",
            "email": "jdoe@example.com",
            "enabled": true,
            "id": "ff4e51",
            "name": "James Doe",
            "links": {
               "self": "https://example.com/identity/v3/users/ff4e51"
            }
        }
}
```


## 根据id查询user


**接口地址**:`/v3/users/{user_id}`


**请求方式**:`GET`



**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|user_id|user_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|用户|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
	"user": {
            "default_project_id": "263fd9",
            "domain_id": "1789d1",
            "enabled": true,
            "id": "9fe1d3",
            "links": {
                "self": "https://example.com/identity/v3/users/9fe1d3"
            },
            "name": "jsmith"
        }
}
```


## 删除user


**接口地址**:`/v3/users/{user_id}`


**请求方式**:`DELETE`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|user_id|user_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


暂无



## 更新user


**接口地址**:`/v3/users/{user_id}`


**请求方式**:`PATCH`



**接口描述**:


**请求示例**:


```json
{
	"user": {
                "default_project_id": "263fd9",
                "domain_id": "1789d1",
                "enabled": true,
                "name": "James Doe",
                "password": "secretsecret",
                "description": "James Doe user",
                "email": "jdoe@example.com"
     }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|user|user|body|true|schema|user|
|user_id|user_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|用户|
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


**响应示例**:
```json
{
	"user": {
                "default_project_id": "263fd9",
                "domain_id": "1789d1",
                "enabled": true,
                "id": "9fe1d3",
                "links": {
                    "self": "https://example.com/identity/v3/users/9fe1d3"
                },
                "name": "jsmith"
       }
}
```


## 修改密码


**接口地址**:`/v3/users/{user_id}/password`


**请求方式**:`POST`


**请求数据类型**:`application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```json
{
	"user": {
            "password": "new_secretsecret",
            "original_password": "secretsecret"
        }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|password|新密码|body|true|string||
|original_password|老密码|body|true|string||
|user_id|user_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created||
|204|No Content||
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


暂无


**响应示例**:
```json

```


# domain管理

Domain

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|id|项目ID|string||
|name|项目名|string||
|enable|是否可用|bool||
|links|相关链接|schema|Links|
|domain_id|所属Domain|string|
|description|描述信息|string||


## 查询domain集合


**接口地址**:`/v3/domains`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|enabled|是否可用|query|false|boolean||
|name|domain名|query|false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Project|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|links|相关链接|schema|Links|
|domains|domain列表|schema|Domain|



**响应示例**:
```json
{
"domains": [
        {
            "description": "Owns users and tenants (i.e. projects) available on Identity API v2.",
            "enabled": true,
            "id": "default",
            "links": {
                "self": "http://example.com/identity/v3/domains/default"
            },
            "name": "Default"
        }
    ],
    "links": {
        "next": null,
        "previous": null,
        "self": "http://example.com/identity/v3/domains"
    }
}
```


## 创建domain


**接口地址**:`/v3/domains`


**请求方式**:`POST`

**接口描述**:


**请求示例**:


```json
{
	 "domain": {
            "description": "Domain description",
            "enabled": true,
            "name": "myDomain"
      }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|domain|domain|body|true||
|&emsp;&emsp;description|描述信息|body|N|string|
|&emsp;&emsp;enabled|是否可用|body|N|bool|
|&emsp;&emsp;name|domain名|body|N|string|




**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created|Project|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:



**响应示例**:
```json
{
    "domain": {
        "description": "Owns users and tenants (i.e. projects) available on Identity API v2.",
        "enabled": true,
        "id": "default",
        "links": {
            "self": "http://example.com/identity/v3/domains/default"
        },
        "name": "Default"
    }
}
```


## 根据id查询domain


**接口地址**:`/v3/domains/{domain_id}`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|domain_id|domain_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Project|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
    "domain": {
        "description": "Owns users and tenants (i.e. projects) available on Identity API v2.",
        "enabled": true,
        "id": "default",
        "links": {
            "self": "http://example.com/identity/v3/domains/default"
        },
        "name": "Default"
    }
}
```


## 更新domain


**接口地址**:`/v3/domains/{domain_id}`


**请求方式**:`PATCH`


**接口描述**:


**请求示例**:


```json
{
	 "domain": {
            "description": "Domain description",
            "enabled": true,
            "name": "myDomain"
      }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|domain|domain|body|true|||
|domain_id|domain_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Project|
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


**响应示例**:
```json
{
    "domain": {
        "description": "Owns users and tenants (i.e. projects) available on Identity API v2.",
        "enabled": true,
        "id": "default",
        "links": {
            "self": "http://example.com/identity/v3/domains/default"
        },
        "name": "Default"
    }
}
```


## 删除domain


**接口地址**:`/v3/domains/{id}`


**请求方式**:`DELETE`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id|id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK||
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


暂无



# token-controller


## adminGetToken


**接口地址**:`/v3/admin/tokens`


**请求方式**:`POST`


**请求数据类型**:`application/json`


**响应数据类型**:`application/json;charset=UTF-8`


**接口描述**:


**请求示例**:


```json
{
	"project_domain_id": "",
	"project_id": "",
	"project_name": "",
	"user_id": ""
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|adminGetToken|adminGetToken|body|true|AdminGetToken|AdminGetToken|
|&emsp;&emsp;project_domain_id|项目所属Domain(如果projectId不为空时可不填写)||false|string||
|&emsp;&emsp;project_id|项目ID(项目名二选一)||false|string||
|&emsp;&emsp;project_name|项目名(和项目ID二选一)||false|string||
|&emsp;&emsp;user_id|用户ID||false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created||
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:

| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|X-Subject-Token|生成的用户token|header|true|string||



## validateToken


**接口地址**:`/v3/auth/tokens`


**请求方式**:`GET`


**请求数据类型**:`*`


**响应数据类型**:`application/json;charset=UTF-8`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|X-Subject-Token|要验证的token|header|true|string||
|X-Auth-Token|带有验证权限的token|header|true|string||
|allow_expired|allow_expired|query|false|boolean||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Token|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:

注:返回的issuedAt和ExpiresAt格式必须为:yyyy-MM-dd'T'HH:mm:ss.SSSSS'Z' 即 [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- |
|token||schema|Token|
|&emsp;&emsp;catalog||array|Catalog|
|&emsp;&emsp;&emsp;&emsp;endpoints||array|节点|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;enabled|是否可用||false|boolean||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;extra|其他信息||false|object||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;id|id||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;interface|节点类型||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;region|||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;region_id|所属区域||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;service_id|所属服务||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;url|访问url||false|string||
|&emsp;&emsp;&emsp;&emsp;id||string||
|&emsp;&emsp;&emsp;&emsp;name||string||
|&emsp;&emsp;&emsp;&emsp;type||string||
|&emsp;&emsp;expires_at||string(date-time)|string(date-time)|
|&emsp;&emsp;is_domain||boolean||
|&emsp;&emsp;issued_at||string(date-time)|string(date-time)|
|&emsp;&emsp;methods||array||
|&emsp;&emsp;project||Project|Project|
|&emsp;&emsp;&emsp;&emsp;domain||Domain|Domain|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;id|||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;name|||false|string||
|&emsp;&emsp;&emsp;&emsp;id||string||
|&emsp;&emsp;&emsp;&emsp;name||string||
|&emsp;&emsp;roles||array|Role|
|&emsp;&emsp;&emsp;&emsp;id||string||
|&emsp;&emsp;&emsp;&emsp;name||string||
|&emsp;&emsp;user||User|User|
|&emsp;&emsp;&emsp;&emsp;domain||Domain|Domain|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;id|||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;name|||false|string||
|&emsp;&emsp;&emsp;&emsp;email||string||
|&emsp;&emsp;&emsp;&emsp;id||string||
|&emsp;&emsp;&emsp;&emsp;name||string||
|&emsp;&emsp;&emsp;&emsp;password_expires_at||string(date-time)||


**响应示例**:
```json
{
  "token": {
    "catalog": [
		{
			"endpoints": [
				{
					"enabled": true,
					"extra": {},
					"id": "",
					"interface": "",
					"region": "",
					"region_id": "",
					"service_id": "",
					"url": ""
				}
			],
			"id": "",
			"name": "",
			"type": ""
		}
	],
	"expires_at": "",
	"is_domain": true,
	"issued_at": "",
	"methods": [],
	"project": {
		"domain": {
			"id": "",
			"name": ""
		},
		"id": "",
		"name": ""
	},
	"roles": [
		{
			"id": "",
			"name": ""
		}
	],
	"user": {
		"domain": {
			"id": "",
			"name": ""
		},
		"email": "",
		"id": "",
		"name": "",
		"password_expires_at": ""
	}

}
}
```


## issueToken


# **接口地址**:`/v3/auth/tokens`


**请求方式**:`POST`



**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|auth|验证信息|body|true|schema|Auth|
|&emsp;&emsp;identity|验证信息|body|true|schema|Identity|
|&emsp;&emsp;&emsp;&emsp;methods|验证方式password/token|body|true|array||
|&emsp;&emsp;&emsp;&emsp;password|与token二选一|body|false|schema|Password|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;user|用户信息|body|false|schema|User|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;name|用户名|body|false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;domain|domain信息|body|false|schema|Domain|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;name|domain名称(与id二选一)|body|false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;id|domain ID(与name二选一)|body|false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;password|密码|body|false|string||
|&emsp;&emsp;&emsp;&emsp;token|与password二选一|body|false|schema|Token|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;id|token|body|false|string||
|&emsp;&emsp;scope|scope信息|body|true|schema|Scope|
|&emsp;&emsp;&emsp;&emsp;project|所项目信息|body|true|schema|Project|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;name|项目名称(与ID二选一)|body|false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;id|项目ID(与name二选一)|body|false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;domain|所属Domain信息|body|false|schema|Domain
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;name|domain名称(与ID二选一)|body|false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;id|domain ID(与name二选一)|body|false|string||

**请求示例**

````json
     {
     "auth":{
         "identity":{
         "methods":[
         "password"
         ],
         "password":{
         "user":{
             "name":"admin",
             "domain":{
             "name":"Default" // or "id": "a6944d763bf64ee6a275f1263fae0352"
             },
             "password":"*******"
           }
         }
        },
     "scope":{
         "project":{
         "domain":{
            "name":"Default" //or "id": "ee4dfb6e5540447cb3741905149d9b6e"
         },
         "name":"services" // or "id": "e56944d76bf64ee6a275f1263fae0352"
         }
       }
     }
   }
````

````json
{
 "auth": {
         "identity": {
             "methods": [
                 "token"
             ],
             "token": {
                 "id": "'$OS_TOKEN'"
             }
         },
         "scope": {
             "domain": {
                 "id": "default"
             }
         }
     }
}
````


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created|Token|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:

注:返回的issuedAt和ExpiresAt格式必须为:yyyy-MM-dd'T'HH:mm:ss.SSSSS'Z' 即 [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- |
|X-Subject-Token|Header|string||
|token||schema|Token|
|&emsp;&emsp;catalog||array|Catalog|
|&emsp;&emsp;&emsp;&emsp;endpoints||array|节点|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;enabled|是否可用||false|boolean||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;extra|其他信息||false|object||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;id|id||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;interface|节点类型||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;region|||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;region_id|所属区域||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;service_id|所属服务||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;url|访问url||false|string||
|&emsp;&emsp;&emsp;&emsp;id||string||
|&emsp;&emsp;&emsp;&emsp;name||string||
|&emsp;&emsp;&emsp;&emsp;type||string||
|&emsp;&emsp;expires_at||string(date-time)|string(date-time)|
|&emsp;&emsp;is_domain||boolean||
|&emsp;&emsp;issued_at||string(date-time)|string(date-time)|
|&emsp;&emsp;methods||array||
|&emsp;&emsp;project||Project|Project|
|&emsp;&emsp;&emsp;&emsp;domain||Domain|Domain|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;id|||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;name|||false|string||
|&emsp;&emsp;&emsp;&emsp;id||string||
|&emsp;&emsp;&emsp;&emsp;name||string||
|&emsp;&emsp;roles||array|Role|
|&emsp;&emsp;&emsp;&emsp;id||string||
|&emsp;&emsp;&emsp;&emsp;name||string||
|&emsp;&emsp;user||User|User|
|&emsp;&emsp;&emsp;&emsp;domain||Domain|Domain|
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;id|||false|string||
|&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;name|||false|string||
|&emsp;&emsp;&emsp;&emsp;email||string||
|&emsp;&emsp;&emsp;&emsp;id||string||
|&emsp;&emsp;&emsp;&emsp;name||string||
|&emsp;&emsp;&emsp;&emsp;password_expires_at||string(date-time)||


**响应示例**:
```json
{
  "token": {
    "catalog": [
		{
			"endpoints": [
				{
					"enabled": true,
					"extra": {},
					"id": "",
					"interface": "",
					"region": "",
					"region_id": "",
					"service_id": "",
					"url": ""
				}
			],
			"id": "",
			"name": "",
			"type": ""
		}
	],
	"expires_at": "",
	"is_domain": true,
	"issued_at": "",
	"methods": [],
	"project": {
		"domain": {
			"id": "",
			"name": ""
		},
		"id": "",
		"name": ""
	},
	"roles": [
		{
			"id": "",
			"name": ""
		}
	],
	"user": {
		"domain": {
			"id": "",
			"name": ""
		},
		"email": "",
		"id": "",
		"name": "",
		"password_expires_at": ""
	}

}
}
```


# version-controller


## getAllVersion


**接口地址**:`/`


**请求方式**:`GET`


**接口描述**:



**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|300|Multiple Choices||



**响应示例**:
```json
{
"versions":{
   "values": [
        {
              "id": "v2.0",
              "status": "deprecated",
              "updated": "2016-08-04T00:00:00Z",
              "links": [
                 { 
                     "rel": "self",
                     "href": "%s/v2.0/"
                 }, {
                     "rel": "describedby",
                     "type": "text/html",
                     "href": "https://docs.openstack.org/"
                 }],
              "media-types": [
                 {
                     "base": "application/json",
                     "type": "application/vnd.openstack.identity-v2.0+json"
                 }
                ]
        },
       {
               "id": "v3.10",
               "status": "stable",
               "updated": "2018-02-28T00:00:00Z",
               "links": [
                  { 
                      "rel": "self",
                      "href": "%s/v3/"
                  }],
               "media-types": [
                  {
                      "base": "application/json",
                      "type": "application/vnd.openstack.identity-v3+json"
                  }
             ]
        }
  ]
}

}
```


## getVersion3


**接口地址**:`/v3`


**请求方式**:`GET`



**接口描述**:


**请求参数**:




**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK||
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应示例**:
```json
{
               "id": "v3.10",
               "status": "stable",
               "updated": "2018-02-28T00:00:00Z",
               "links": [
                  { 
                      "rel": "self",
                      "href": "%s/v3/"
                  }],
               "media-types": [
                  {
                      "base": "application/json",
                      "type": "application/vnd.openstack.identity-v3+json"
                  }
             ]
}
```


# 区域管理

Region

| 参数名称 | 参数说明 | 类型 | schema |
|description|区域描述|string||
|id|id|string||
|parent_region_id|父区域id|string||



## 查询区域集合


**接口地址**:`/v3/regions`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|parentRegionId|父区域id|query|false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|区域|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|links|相关链接|schema|Links|
|regions|区域集合|schema|Region|



**响应示例**:
```json
{
	"links": {
            "next": null,
            "previous": null,
            "self": "http://example.com/identity/v3/regions"
        },
        "regions": [
            {
                "description": "",
                "id": "RegionOne",
                "links": {
                    "self": "http://example.com/identity/v3/regions/RegionOne"
                },
                "parent_region_id": null
            }
        ]
}
```


## 新建区域


**接口地址**:`/v3/regions`


**请求方式**:`POST`


**接口描述**:


**请求示例**:


```json
{
	"region": {
            "description": "My subregion",
            "id": "RegionOneSubRegion",
            "parent_region_id": "RegionOne"
    }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|region|region|body|true|||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created|区域|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:

**响应示例**:
```json
{
	 "region": {
            "parent_region_id": "RegionOne",
            "id": "RegionThree",
            "links": {
                "self": "http://example.com/identity/v3/regions/RegionThree"
            },
            "description": "My subregion 3"
        }
}
```


## 根据id查询区域


**接口地址**:`/v3/regions/{region_id}`


**请求方式**:`GET`



**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|region_id|region_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|区域|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
	 "region": {
            "parent_region_id": "RegionOne",
            "id": "RegionThree",
            "links": {
                "self": "http://example.com/identity/v3/regions/RegionThree"
            },
            "description": "My subregion 3"
       }
}
```


## 删除区域


**接口地址**:`/v3/regions/{region_id}`


**请求方式**:`DELETE`



**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|region_id|region_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


暂无




## 更新区域


**接口地址**:`/v3/regions/{region_id}`


**请求方式**:`PATCH`


**请求数据类型**:`application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```json
{
	"region": {
                "description": "My subregion",
                "id": "RegionOneSubRegion",
                "parent_region_id": "RegionOne"
        }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|region|region|body|true|||
|region_id|region_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|区域|
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


**响应示例**:
```json
{
	 "region": {
            "parent_region_id": "RegionOne",
            "id": "RegionThree",
            "links": {
                "self": "http://example.com/identity/v3/regions/RegionThree"
            },
            "description": "My subregion 3"
       }
}
```


# 服务管理

Service

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|enabled|是否可用|boolean||
|description|描述信息|string||
|id|id|string||
|name|服务名称|string||
|type|服务类型|string||


## 查询服务集合


**接口地址**:`/v3/services`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|type|服务类型|query|false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|服务|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
"links": {
        "next": null,
        "previous": null,
        "self": "http://example.com/identity/v3/services"
    },
    "services": [
        {
            "description": "Cinder Volume Service",
            "enabled": true,
            "id": "cdda3bea0742407f95e70f4758f46558",
            "links": {
                "self": "http://example.com/identity/v3/services/cdda3bea0742407f95e70f4758f46558"
            },
            "name": "cinder",
            "type": "volume"
        }
    ]
}
```


## 新建服务


**接口地址**:`/v3/services`


**请求方式**:`POST`


**接口描述**:


**请求示例**:


```json
{
	 "service": {
        "type": "compute",
        "name": "compute2",
        "description": "Compute service 2"
     }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|service|service|body|true|||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created|服务|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
    "service": {
        "name": "cinderv2",
        "links": {
            "self": "http://example.com/identity/v3/services/5789da9864004dd088fce14c1c626a4b"
        },
        "enabled": true,
        "type": "volumev2",
        "id": "5789da9864004dd088fce14c1c626a4b",
        "description": "Block Storage Service V2"
    }
}
```


## 根据id查询服务


**接口地址**:`/v3/services/{service_id}`


**请求方式**:`GET`


**请求数据类型**:`*`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|service_id|service_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|服务|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:



**响应示例**:
```json
{
    "service": {
        "name": "cinderv2",
        "links": {
            "self": "http://example.com/identity/v3/services/5789da9864004dd088fce14c1c626a4b"
        },
        "enabled": true,
        "type": "volumev2",
        "id": "5789da9864004dd088fce14c1c626a4b",
        "description": "Block Storage Service V2"
    }
}
```


## 删除服务


**接口地址**:`/v3/services/{service_id}`


**请求方式**:`DELETE`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|service_id|service_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


暂无


## 更新服务


**接口地址**:`/v3/services/{service_id}`


**请求方式**:`PATCH`


**接口描述**:


**请求示例**:


```json
{
	"service": {
            "type": "compute",
            "name": "compute2",
            "description": "Compute service 2"
    }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|service|service|body|true|||
|service_id|service_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|服务|
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:

**响应示例**:
```json
{
	"service": {
        "name": "cinderv2",
        "links": {
            "self": "http://example.com/identity/v3/services/5789da9864004dd088fce14c1c626a4b"
        },
        "enabled": true,
        "type": "volumev2",
        "id": "5789da9864004dd088fce14c1c626a4b",
        "description": "Block Storage Service V2"
    }
}
```


# 权限管理


## 权限查询


**接口地址**:`/v3/projects/{project_id}/users/{user_id}/roles`


**请求方式**:`GET`


**请求数据类型**:`*`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|project_id|项目id|path|true|string||
|user_id|用户id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|AssignResponse|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|links|相关链接|object||
|roles|角色列表|array|Role|
|&emsp;&emsp;id|角色id|string||
|&emsp;&emsp;name|角色名|string||
|&emsp;&emsp;links|相关链接|schema|Links|


**响应示例**:
```json
{
	 "links": {
            "self": "http://example.com/identity/v3/projects/9e5a15e2c0dd42aab0990a463e839ac1/users/b964a9e51c0046a4a84d3f83a135a97c/roles",
            "previous": null,
            "next": null
        },
        "roles": [
            {
                "id": "3b5347fa7a144008ba57c0acea469cc3",
                "links": {
                    "self": "http://example.com/identity/v3/roles/3b5347fa7a144008ba57c0acea469cc3"
                },
                "name": "admin"
            }
        ]
}
```


## 新建关联权限


**接口地址**:`/v3/projects/{project_id}/users/{user_id}/roles/{role_id}`


**请求方式**:`PUT`


**请求数据类型**:`application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|project_id|project_id|path|true|string||
|role_id|role_id|path|true|string||
|user_id|user_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created||
|204|No Content||
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


暂无




## 删除关联权限


**接口地址**:`/v3/projects/{project_id}/users/{user_id}/roles/{role_id}`


**请求方式**:`DELETE`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|project_id|project_id|path|true|string||
|role_id|role_id|path|true|string||
|user_id|user_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


暂无



## 验证关联权限


**接口地址**:`/v3/projects/{project_id}/users/{user_id}/roles/{role_id}`


**请求方式**:`HEAD`



**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|project_id|project_id|path|true|string||
|role_id|role_id|path|true|string||
|user_id|user_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


暂无



## 查询所有权限


**接口地址**:`/v3/role_assignments`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|actorId|用户id|query|false|string||
|roleId|角色id|query|false|string||
|targetId|项目id|query|false|string||
|type|枚举类型,可用值:UserProject,UserDomain|query|false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK||
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
"role_assignments": [
        {
            "links": {
                "assignment": "http://example.com/identity/v3/domains/161718/users/313233/roles/123456"
            },
            "role": {
                "id": "123456"
            },
            "scope": {
                "project": {
                    "id": "161718"
                }
            },
            "user": {
                "id": "313233"
            }
        }
    ],
    "links": {
        "self": "http://example.com/identity/v3/role_assignments",
        "previous": null,
        "next": null
    }
}
```


# 节点管理

Endpoint

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|enabled|是否可用|boolean||
|id|id|string||
|interface|节点类型|string||
|region|所属区域(为适配低版本keystone)|string||
|region_id|所属区域|string||
|service_id|所属服务|string||
|url|访问url|string||

## 查询节点列表


**接口地址**:`/v3/endpoints`


**请求方式**:`GET`


**请求数据类型**:`*`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|interfaceType|节点类型|query|false|string||
|serviceId|所属服务id|query|false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|节点|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
	"endpoints": [
            {
                "enabled": true,
                "id": "ea74f9771dec475eabfc2cdff5364413",
                "interface": "admin",
                "links": {
                    "self": "http://example.com/identity/v3/endpoints/ea74f9771dec475eabfc2cdff5364413"
                },
                "region": "RegionOne",
                "region_id": "RegionOne",
                "service_id": "ef6b15e425814dc69d830361baae0e33",
                "url": "http://23.253.211.234:8080"
            }
        ],
        "links": {
            "next": null,
            "previous": null,
            "self": "http://example.com/identity/v3/endpoints"
        }
}
```


## 创建节点


**接口地址**:`/v3/endpoints`


**请求方式**:`POST`


**请求数据类型**:`application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```json
{
	"endpoint": {
            "interface": "public",
            "region_id": "RegionOne",
            "url": "http://example.com/identity/v3/endpoints/828384",
            "service_id": "9242e05f0c23467bbd1cf1f7a6e5e596"
        }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|endpoint|endpoint|body|true|||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created|节点|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:



**响应示例**:
```json
{
	"endpoint": {
            "id": "828384",
            "interface": "internal",
            "links": {
                "self": "http://example.com/identity/v3/endpoints/828384"
            },
            "region_id": "north",
            "service_id": "686766",
            "url": "http://example.com/identity/v3/endpoints/828384"
        }
}
```


## 根据id查询节点


**接口地址**:`/v3/endpoints/{endpoint_id}`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|endpoint_id|endpoint_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|节点|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:



**响应示例**:
```json
{
	"endpoint": {
            "id": "828384",
            "interface": "internal",
            "links": {
                "self": "http://example.com/identity/v3/endpoints/828384"
            },
            "region_id": "north",
            "service_id": "686766",
            "url": "http://example.com/identity/v3/endpoints/828384"
        }
}
```


## 删除节点


**接口地址**:`/v3/endpoints/{endpoint_id}`


**请求方式**:`DELETE`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|endpoint_id|endpoint_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


暂无


## 更新节点


**接口地址**:`/v3/endpoints/{endpoint_id}`


**请求方式**:`PATCH`


**请求数据类型**:`application/json`


**响应数据类型**:`*/*`


**接口描述**:


**请求示例**:


```json
{
	"endpoint": {
        "interface": "public",
        "region_id": "RegionOne",
        "url": "http://example.com/identity/v3/endpoints/828384",
        "service_id": "9242e05f0c23467bbd1cf1f7a6e5e596"
     }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|endpoint|endpoint|body|true|||
|endpoint_id|endpoint_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|节点|
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:

**响应示例**:
```json
{
	"endpoint": {
            "id": "828384",
            "interface": "internal",
            "links": {
                "self": "http://example.com/identity/v3/endpoints/828384"
            },
            "region_id": "north",
            "service_id": "686766",
            "url": "http://example.com/identity/v3/endpoints/828384"
        }
}
```


# 项目管理

Project

| 参数名称 | 参数说明 | 类型 | schema |
| -------- | -------- | ----- |----- | 
|id|项目id|string||
|name|项目名|string||
|description|描述信息|string||
|is_domain|是否为domain|bool||
|enabled|是否可用|bool||
|parent_id|父id|string||

## 查询项目集合


**接口地址**:`/v3/projects`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Project|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
  "links": {
        "next": null,
        "previous": null,
        "self": "http://example.com/identity/v3/projects"
    },
    "projects": [
        {
            "is_domain": false,
            "description": null,
            "domain_id": "default",
            "enabled": true,
            "id": "fdb8424c4e4f4c0ba32c52e2de3bd80e",
            "links": {
                "self": "http://example.com/identity/v3/projects/fdb8424c4e4f4c0ba32c52e2de3bd80e"
            },
            "name": "alt_demo",
            "parent_id": null
        }
    ]
}
```


## 新建项目


**接口地址**:`/v3/projects`


**请求方式**:`POST`




**接口描述**:


**请求示例**:


```json
{
	"project": {
            "description": "My new domain",
            "enabled": true,
            "is_domain": true,
            "name": "myNewDomain"
     }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|project|project|body|true|||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|201|Created|Project|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
	"project": {
            "description": "My updated project",
            "domain_id": null,
            "links": {
                "self": "http://example.com/identity/v3/projects/93ebbcc35335488b96ff9cd7d18cbb2e"
            },
            "enabled": true,
            "id": "93ebbcc35335488b96ff9cd7d18cbb2e",
            "is_domain": true,
            "name": "myUpdatedProject",
            "parent_id": null,
        }
}
```


## 查询项目分页集合


**接口地址**:`/v3/projects/pages`


**请求方式**:`GET`


**请求数据类型**:`*`


**响应数据类型**:`*/*`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|domainId|所属domain_id|query|false|string||
|enabled|是否可用|query|false|boolean||
|isDomain|是否为domain|query|false|boolean||
|name|项目名|query|false|string||
|pageNo||query|false|integer(int32)||
|pageSize||query|false|integer(int32)||
|parentId|父项目id|query|false|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|PageResult«Project»|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
	"contents": [
		{
			"description": "My updated project",
            "domain_id": null,
            "links": {
                "self": "http://example.com/identity/v3/projects/93ebbcc35335488b96ff9cd7d18cbb2e"
            },
            "enabled": true,
            "id": "93ebbcc35335488b96ff9cd7d18cbb2e",
            "is_domain": true,
            "name": "myUpdatedProject",
            "parent_id": null,
		}
	],
	"total_count": 0
}
```


## 删除项目


**接口地址**:`/v3/projects/{id}`


**请求方式**:`DELETE`



**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|id|id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK||
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


暂无



## 根据id查询项目


**接口地址**:`/v3/projects/{project_id}`


**请求方式**:`GET`


**接口描述**:


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|project_id|project_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Project|
|401|Unauthorized||
|403|Forbidden||
|404|Not Found||


**响应参数**:


**响应示例**:
```json
{
	"project": {
            "description": "My updated project",
            "domain_id": null,
            "links": {
                "self": "http://example.com/identity/v3/projects/93ebbcc35335488b96ff9cd7d18cbb2e"
            },
            "enabled": true,
            "id": "93ebbcc35335488b96ff9cd7d18cbb2e",
            "is_domain": true,
            "name": "myUpdatedProject",
            "parent_id": null,
            "tags": [],
            "options": {}
     }
}
```


## 更新项目


**接口地址**:`/v3/projects/{project_id}`


**请求方式**:`PATCH`


**接口描述**:


**请求示例**:


```json
{
	"project": {
        "description": "My new domain",
        "enabled": true,
        "is_domain": true,
        "name": "myNewDomain"
      }
}
```


**请求参数**:


| 参数名称 | 参数说明 | in    | 是否必须 | 数据类型 | schema |
| -------- | -------- | ----- | -------- | -------- | ------ |
|project|project|body|true|||
|project_id|project_id|path|true|string||


**响应状态**:


| 状态码 | 说明 | schema |
| -------- | -------- | ----- | 
|200|OK|Project|
|204|No Content||
|401|Unauthorized||
|403|Forbidden||


**响应参数**:


**响应示例**:
```json
{
	"project": {
            "description": "My updated project",
            "domain_id": null,
            "links": {
                "self": "http://example.com/identity/v3/projects/93ebbcc35335488b96ff9cd7d18cbb2e"
            },
            "enabled": true,
            "id": "93ebbcc35335488b96ff9cd7d18cbb2e",
            "is_domain": true,
            "name": "myUpdatedProject",
            "parent_id": null,
            "tags": [],
            "options": {}
     }
}
```
