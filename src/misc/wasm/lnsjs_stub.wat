;; WASI のランタイム (wasmtime など) 用の、Node.js でだけ使える関数 (モジュール "lnsjs") の代わり。
;; lnswasi.sh で、変換した wasm と C のランタイムに合わせて 1 つの wasm にする。
(module
  ;; Lua の値の参照を、GC で回収された時に解放するように登録する (WASI では解放しない)
  (func (export "luaval_new") (param anyref i32))
  ;; io.popen (WASI ではコマンドを実行できないので失敗する)
  (func (export "popen") (param i32 i32) (result i32)
    i32.const -1)
  (func (export "popen_data") (param i32))
)
