#########################################################################
# 輸入 Leetcode 題目名稱，自動建立相對應的資料夾及檔案
# 範例: bash create_leetcode.sh 1. Two Sum
#########################################################################
#!/bin/bash

if [ "$#" -lt 1 ]; then
  echo "請提供【{數字}.{空格}{題目名稱}】格式的輸入，例如 \"1. Two Sum\""
  exit 1
fi

# 設定參數
## 1. Two Sum
problem_name="$@"
## 1
problem_num=$(echo "$1" | tr -dc '0-9')
## 0001
problem_4digit_num=$(printf "%04d" "$problem_num")
## TwoSum
problem_func_name=$(echo "$@" | sed -E 's/[0-9]+[.]\s//; s/\s+//g')
## ./leetcode/0001.Two-Sum
## 由於資料夾名稱不允許空白，因此進行過濾
dest_folder="./leetcode/${problem_4digit_num}.$(echo "$@" | sed -E 's/[0-9]+[.]\s//; s/\s+/-/g')" 

# 建立目錄
mkdir -p "${dest_folder}"

# 在目錄中建立 README.md
cat << EOF > "${dest_folder}/README.md"
# ${problem_name}

<>

## Takeaway


EOF

# 在目錄中建立 .go
cat << EOF > "${dest_folder}/${problem_name}.go"
package leetcode

func ${problem_func_name}() {}
EOF

# 在目錄中建立 _test.go
cat << EOF > "${dest_folder}/${problem_name}_test.go"
package leetcode

import (
  "reflect"
  "testing"
)

// 填入 function input type
type parameters${problem_4digit_num} struct {
  
}

func Test${problem_func_name}(t *testing.T) {
  tests := []struct {
    parameters${problem_4digit_num}
    // 填入 function output type
    ans 
  }{
    // 填入 test case
    {
      parameters${problem_4digit_num}{},
      ,
    },
  }
  
  for _, test := range tests {
    t.Run("Test ${problem_func_name}", func(t *testing.T) {
      // 完整輸入參數
      result := ${problem_func_name}(test.parameters${problem_4digit_num})
      // compare 的方式需視情況調整
      if !reflect.DeepEqual(result, test.ans) {
        t.Errorf("${problem_func_name}(%+v) got %+v, want %+v", test.parameters${problem_4digit_num}, result, test.ans)
      }
    })
  }
}
EOF

# End
echo "已成功建立目錄 ${dest_folder}。"
