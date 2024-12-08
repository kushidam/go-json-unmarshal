
## 1.　map宣言
```
map[キーの型]値の型
```

`map`はキーと値のペアを持つコレクション（ハッシュマップ）を表す

### 1.1. []map[string]interface{}
[] は配列（スライス）を表す
配列の要素が map[string]interface{} 型であることを示す

### 1.2. map[string]interface{}
配列内の各要素は、再びキーが文字列で、値が任意の型（interface{}）となるマップ

### 1.3. map[string]string
キーが文字列で、値も文字列となるマップ
