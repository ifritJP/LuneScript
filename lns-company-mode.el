;; -*- lexical-binding: t -*-

(require 'cl-lib)

(require 'company)
(require 'lns-completion)

(defvar company-lns-buf-name "*company-lns-process*")
(defvar company-lns-buf-name-work "*company-lns-process-work*")


(defun lns-company-pick-type (txt)
  (when (string-match ": \\([a-zA-Z_0-9<>!]+\\)$" txt)
    (setq txt (match-string 1 txt)))
  txt
  )


(defun lns-company-exclude-type (txt)
  (when (string-match "\\(.*\\): [a-zA-Z_0-9<>!]+$" txt)
    (setq txt (match-string 1 txt)))
  txt
  )

(defun lns-company-make-candidate (candidate)
  (let* ((displayTxt (lns-candidate-get-displayTxt candidate))
	 (type (lns-candidate-get-type candidate)))
    (when displayTxt
      (cond
       ;; 型情報を補完文字列から除外
       ((or (equal type "Mtd")
	    (equal type "Fun")
	    (equal type "Var")
	    (equal type "Arg")
	    (equal type "Mbr"))
	(setq displayTxt (lns-company-exclude-type displayTxt)))
       )
      ;; getter アクセスの $ への変換
      (when (string-match "^get_\\([a-zA-Z_0-9<>!]+\\)()$" displayTxt)
	(setq displayTxt (format "$%s" (match-string 1 displayTxt))))
      (propertize displayTxt 'meta candidate))
    ))

(defun company-lns--candidates (prefix callback)
  ;; 補完候補のリストを生成。補完候補は文字列。

  ;; lns-mode 以外は動作させない。
  (if (not (eq major-mode 'lns-mode))
      (funcall callback nil)

    (message (format "prefix:%s" prefix) )
    
    ;; ピックアップした候補はバッファに出力するが、非同期に補完候補をピックアップするので、
    ;; バッファがそれぞれで必要なため generate-new-buffer で生成し、
    ;; ピックアップ終了時に kill する。
    (let ((tmp-buf (generate-new-buffer company-lns-buf-name-work)))
      (lns-completion-get-candidate-list
       (lambda (candidate-list err)
	 (cond
	  (candidate-list
	   ;; ピックアップした候補から company に表示する情報に変換し、callback に情報を渡す。
	   (let (local-candidate-list)
	     (setq local-candidate-list
		   (mapcar (lambda (X)
			     (lns-company-make-candidate (lns-json-val X :candidate)))
			   candidate-list))
	     (funcall callback (delq nil local-candidate-list))
	     )
	   )
	  (t
	   (funcall callback '())))
	 (with-current-buffer (get-buffer-create company-lns-buf-name)
	   (erase-buffer)
	   (insert (with-current-buffer tmp-buf
		     (buffer-string)))
	   )
	 (kill-buffer tmp-buf)
	 )
       t tmp-buf)
      )))

(defun company-lns--meta (candidate)
  ;; 補完候補リスト表示時の、補足説明文字列取得。 mini-buffer に表示される。
  "")
  ;; (format "This will use %s of %s"
  ;;         (get-text-property 0 'meta candidate)
  ;;         (substring-no-properties candidate)))

(defun company-lns--annotation (item)
  ;; company-lns--candidates で生成したリストのアイテムを表示する処理
  ;; candidate は、リストのアイテム
  (let* ((candidate (get-text-property 0 'meta item))
  	 (displayTxt (lns-candidate-get-displayTxt candidate)))
    (format ": %s" (lns-company-pick-type displayTxt))))


(defun company-lns (command &optional arg &rest ignored)
  (interactive (list 'interactive))
  (cl-case command
    (interactive (company-begin-backend 'company-lns))
    ;; バージョンコマンドが動作しない場合は、
    ;; lunescript がインストールされていないと判断する。
    (init (when (not (eq (lns-command-sync "--version") 0))
	    (error "not found lns")))
    ;; 補完を開始する prefix の定義。 . で区切った場合に開始する。
    (prefix (company-grab-symbol-cons "\\." 1))
    ;; 候補を生成する。 非同期。
    (candidates (cons :async
    		      (lambda (callback)
    			(company-lns--candidates arg callback))))
    ;; 候補のリストを表示する
    (annotation (company-lns--annotation arg))
    (meta (company-lns--meta arg))))

(add-to-list 'company-backends 'company-lns)

;;(add-to-list 'company-safe-backends '(company-lns . "LuneScript"))

(provide 'lns-company-mode)
