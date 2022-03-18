## 需求拆分

### 任务列表

1. 常规Option
2. 列表Option
3. 可扩展支持其他Option

### 测试列表

1. Single Option
    - case 1
      - `-l` bool
      - `-p 8080` int
      - `-d /usr/logs` string
    - case 2
      - `-g this is a list` []stirng
      - `-d 1 3 2 4` []int
2. Multiple Options
   - `-l -p 8080 -d /usr/logs`
   - `-g this is a list -d 1 3 2 4`
3. Sad path
   - bool `-l t` `-l t f`
   - int `-p` `-p 8080 8081`
   - string `-d /usr/logs /usr/vars`
4. Default Values
   - bool `false`
   - int `0`
   - string `""`
5. Handle Error
