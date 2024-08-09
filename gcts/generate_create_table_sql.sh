#!/bin/bash
# Wrote by yijian on 2024/08/03
# 生成建表语句

# 参数检查
if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <table_name> <delimiter> <input_file>"
    echo "Example: $0 't_user',' 'input_file.txt'"
    exit 1
fi

table_name="$1" # 表名
delimiter="$2"  # 分隔符（不能为空格、TAB符、单引号、双引号和反引号）
input_file="$3" # 输入文件（只支持单个表）

# delimiter 为单引号、双引号或反引号则报错
if [[ "X$delimiter" == "X" || $delimiter == "\t" || $delimiter == "'" || $delimiter == '"' || $delimiter == '`' ]]; then
    echo "Error: delimiter must not be a space, TAB, single quote, double quote or back quote"
    exit 1
fi
if test ${#delimiter} -ne 1; then
    echo "Error: delimiter must be a single character"
    exit 1
fi

# 生成建表语句
echo "DROP TABLE IF EXISTS \`$table_name\`;"
echo "CREATE TABLE \`$table_name\` ("
while IFS= read -r line; do
    # 使用指定的分隔符分割每一行
    IFS="$delimiter" read -ra columns <<< "$line"
    if [[ ${#columns[@]} -eq 3 ]]; then
        # 移除字段名、数据类型和字段注释的前后空格
        f_name=$(echo "${columns[0]}" | sed 's/^[[:blank:]]*//;s/[[:blank:]]*$//')
        f_type=$(echo "${columns[1]}" | sed 's/^[[:blank:]]*//;s/[[:blank:]]*$//')
        f_comment=$(echo "${columns[2]}" | sed 's/^[[:blank:]]*//;s/[[:blank:]]*$//')

        # 生成列定义，并添加到建表语句中
        echo "  \`$f_name\` $f_type COMMENT '$f_comment',"
    fi
done < "$input_file"

# 替换最后一个列定义中的逗号为分号，并添加结束括号
echo ");" | sed '$s/,$//'
