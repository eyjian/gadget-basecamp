# gcts

生成建表 SQL 的工具（generate create table sql），本工具生成的建表 SQL 不包含创建索引和分区等。

# gogcts

提供一个生成建表 SQL 的 Go 函数，可基于它实现一个在线生成建表 SQL 工具。

# 应用场景

使用文档，如 Word、腾讯文档、wiki 等以表格方式维护表的定义，需要一个工具维护对应的建表 SQL 。注意的字段类型和字段注释前加上前导的分隔符，以便于生成建表 SQL 。注意分隔符不能为空格、TAB符、单引号、双引号和反引号三种字符。

| 字段名           | 字段类型                                                            | 字段注释   |
|---------------|-----------------------------------------------------------------|--------|
| f_id          | ;INT UNSIGNED                                                   | ;自增 ID |
| f_name        | ;VARCHAR(20)                                                    | ;姓名    |
| f_age         | ;INT UNSIGNED                                                   | ;年龄    |
| f_sex         | ;VARCHAR(1)                                                     | ;性别    |
| f_birthday    | ;DATE                                                           | ;生日    |
| f_address     | ;VARCHAR(100)                                                   | ;地址    |
| f_phone       | ;VARCHAR(11)                                                    | ;电话    |
| f_email       | ;VARCHAR(100)                                                   | ;邮箱    |
| f_qq          | ;VARCHAR(10)                                                    | ;QQ    |
| f_wechat      | ;VARCHAR(100)                                                   | ;微信    |
| f_weibo       | ;VARCHAR(100)                                                   | ;微博    |
| f_create_time | ;DATETIME DEFAULT CURRENT_TIMESTAMP                             | ;创建时间  |
| f_update_time | ;DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP | ;更新时间  |
| f_delete_time | ;DATETIME                                                       | ;删除时间  |

使用的时候，从 Word 等将字段名、字段类型、字段注释三列复制到文本文件，然后执行工具就可以生成建表 SQL 。

# generate_create_table_sql.sh 

一个简单的 shell 脚本，可以生成建表 SQL 。

## 使用示例

脚本方式:

```shell
% sh generate_create_table_sql.sh 't_user' ';' 'input_file.txt'
CREATE TABLE `t_user` (
  `f_id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增 ID',
  `f_name` VARCHAR(20) NOT NULL COMMENT '姓名',
  `f_age` INT UNSIGNED NOT NULL COMMENT '年龄',
  `f_sex` VARCHAR(1) NOT NULL COMMENT '性别',
  `f_birthday` DATE NOT NULL COMMENT '生日',
  `f_address` VARCHAR(100) NOT NULL COMMENT '地址',
  `f_phone` VARCHAR(11) DEFAULT '' COMMENT '电话',
  `f_email` VARCHAR(100) DEFAULT '' COMMENT '邮箱',
  `f_qq` VARCHAR(10) DEFAULT '' COMMENT 'QQ',
  `f_wechat` VARCHAR(100) DEFAULT '' COMMENT '微信',
  `f_weibo` VARCHAR(100) DEFAULT '' COMMENT '微博',
  `f_create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `f_update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `f_delete_time` DATETIME COMMENT '删除时间',
);
```

命令方式：

```shell
$ gcts 't_user' ';' 'input_file.txt'
```

## 参数说明

一共三个必要参数：

* 参数1：表名
* 参数2：分隔符（不能为空格、TAB符、单引号、双引号和反引号）
* 参数3：输入文件

## 输入文件说明

输入文件是一个文本文件，每行三列组成，列之间用分隔符分隔。分隔符不能出现在列值中。三列依次为：

* 字段名
* 字段类型
* 字段注释

**注意**一个输入文件只能定义一张表。
