
|URL|请求方式|说明|
|---|-------|----|
|/users/register|   Post| 创建一个新的User|
|/users/login |Post|用户登陆，如果登陆成功就返回一个JWT
|/tasks|Post|创建一个Task|
|/tasks/{id}|         Put| 更新一个存在的Task|
|/tasks |   Get| 获取所有的Task|
|/tasks/{id} |   Get| 根据Id获取对应的Task|
|/tasks/users/{id}   |Get|获取相应User的Task|
|/tasks/{id}     |    Delete|删除对应Id的Task|
|/notes|Post|为一个存在的Task创建Note|
|/notes/{id}|         Put| 更新一个存在的Note|
|/notes |   Get| 获取所有的Note|
|/notes/{id} |   Get| 根据Id获取对应的Note|
|/notes/tasks/{id}   |Get|获取相应Task的Note|
|/notes/{id}     |    Delete|删除对应Id的Note|