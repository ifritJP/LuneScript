;;; lns-mode.el --- a major-mode for editing LuneScript
;;
;; MIT License
;;
;; Copyright (c) 2018 ifritJP
;;
;; Permission is hereby granted, free of charge, to any person obtaining a copy
;; of this software and associated documentation files (the "Software"), to deal
;; in the Software without restriction, including without limitation the rights
;; to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
;; copies of the Software, and to permit persons to whom the Software is
;; furnished to do so, subject to the following conditions:
;;
;; The above copyright notice and this permission notice shall be included in all
;; copies or substantial portions of the Software.
;;
;; THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
;; IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
;; FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
;; AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
;; LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
;; OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
;; SOFTWARE.
;;
;; usage:
;; (add-to-list 'auto-mode-alist '("\\.lns$" . lns-mode))


(require 'lns-dispatch)
(require 'lns-inq)

(defalias 'lns--prog-mode
  (if (fboundp 'prog-mode) 'prog-mode 'fundamental-mode))


(defvar lns-proj-info-check-useing-indent-command
  (lambda (proj-dir)
    (let ((proj-info (lns-proj-info-get proj-dir)))
      (or (plist-get proj-info :indent-command-path)
	  (plist-get proj-info :indent-command-script))))
  "proj-dir のプロジェクトでインデントコマンドを利用するかどうか。"
  )
(defvar lns-get-formatter-path
  (lambda (lnsc-path)
    (expand-file-name (concat lnsc-path "../../../tools/formatter")))
  "formatter のパスを取得する関数を指定する"
  )


(defvar lns-indent-level 3)

(defvar lns-no-token-pattern "[^a-zA-Z0-9_]")
(defvar lns-no-space-pattern "[^\s \t]")


(defvar lns-mode-hook nil
  "Hooks called when Lns mode fires up.")



(defun lns-make-regex-or (list)
  (apply 'concat
	 (format "\\_<%s\\_>" (car list))
	 (mapcar (lambda (x)
		   (format "\\|\\_<%s\\_>" x))
		 (cdr list))))

(eval-and-compile
  (defconst
    lns-builtin (lns-make-regex-or
		 '("_VERSION" "assert" "collectgarbage" "dofile" "error"
		   "getfenv" "getmetatable" "ipairs" "load" "loadfile"
		   "loadstring" "next" "pairs" "pcall" "print"
		   "rawequal" "rawget" "rawlen" "rawset" "require" "select"
		   "setfenv" "setmetatable" "tonumber" "tostring"
		   "type" "unpack" "xpcall")))
  (defconst
    lns-keyword (lns-make-regex-or
		 '("self" "fn" "elseif" "else" "while" "repeat" "for"
		   "apply" "of" "foreach" "forsort" "in" "return" "class" "false"
		   "nil" "null" "true" "switch" "match" "case" "default" "_default"
		   "extend" "proto"
		   "override" "macro" "let" "unwrap" "if" "module" "subfile" "allmut"
		   "mut" "__init" "local" "pub" "pro" "pri" "form" "advertise" "provide" 
		   "wrap" "static" "global" "sync" "then" "do" "interface" "enum"
		   "glue" "when" "alge" "non" "alias" "dep"
		   "abstract" "final" "trust" "import" "import\\.l" "import\\.d"
		   "as" "not" "and" "or" "break" "new" )))
  (defconst
    lns-bloak-statement-head (concat (lns-make-regex-or
				      '("let" "let\!" "let\*" "if" "elseif" "else" "while"
					"repeat" "for" "apply" "foreach" "forsort"
					"class" "pub" "pro" "pri" "form" "advertise"
					"switch" "match" "proto" "case" "interface"
					"wrap" "static" "trust" "import" "''"))
				     "\\|\\_<fn[ \t]*[^(]" ))
  (defconst
    lns-type (lns-make-regex-or
	      '("int" "real" "sym" "stat" "stem"
		"Map" "Array" "List" "Set" "str" "bool")))
  )
  

(defvar lns-font-lock-keywords
  `((,lns-builtin . font-lock-builtin-face)
    (,lns-keyword . font-lock-keyword-face)
    ("\\<_exp\\>" . font-lock-warning-face)
    ("\\<_[a-zA-Z0-9_]+\\>" . font-lock-warning-face)
    ;;("\\<let\\|unwrap\\|if\\>" . font-lock-keyword-face)
    ("@@=\\|@@\\|@\\|\\?\\|&\\|\\.\\$\\|\\$[\\.\\[(]\\|#\\|\\!\\|\*\*\\|_lune_control" .
     font-lock-warning-face)
    (,lns-type . font-lock-type-face)
    ;; (":\\s-*\\(\\w+\\)" . (1 font-lock-type-face nil t))
    (,(format "\\(%s\\)\\s-*\\(\\w+\\)"
	      (lns-make-regex-or '("fn" "form"))) .
	      ( 2 font-lock-function-name-face nil t))    
    (,(format "\\(%s\\)\\s-*\\(\\w+\\)"
	      (lns-make-regex-or '("let" "for" "foreach" "forsort" "apply"))) .
	      ( 2 font-lock-variable-name-face nil t))
    ))



(defun lns-syntax-multi-string (char)
  "\""
  (let ((font-lock-syntactic-keywords nil)
	(loop-flag t)
	(syntax-code "|")
	in-flag)
    (save-excursion
      (if (not (lns-is-in-comment-string (point)))
	  ;; 文字列中でない場合
	  nil
	;; 文字列中の場合、前方の文字列開始位置を見つける
	(while loop-flag
	  (cond
	   ((not (re-search-backward "[\"`]" nil t))
	    ;; 見つからないのはありえない
	    (setq loop-flag nil))
	   (t
	    ;; 見つかった場合は、前の文字が文字列中か調べる
	    (cond
	     ((lns-is-in-comment-string (1- (point)))
	      ;; まだ文字列中なのでもっと前の文字を調べる
	      )
	     ;;(setq loop-flag nil))
	     ;; ((lns-is-in-comment-string (1- (point)))
	     ;;  ;; まだ文字列中なのでもっと前の文字を調べる
	     ;;  nil)
	     (t
	      ;; 文字列開始位置が見つかったので、その文字を調べる
	      (setq loop-flag nil)
	      (if (eq (char-after) char)
		  ;; 同じ文字なので文字列終了
		  (setq syntax-code "|")
		;; 違う文字なので文字列継続
		(setq syntax-code "\."))
	      ))))
	  )))
    syntax-code))


(defvar lns-font-lock-syntactic-keywords
  `(
    ;;("'''" 0 (14 . 41) nil t) ;;; comment
    ;;("''" 0 (11 . 40) nil t) ;;; comment
    ;;("\n" 0 (12 . 40) nil t) ;;; comment

    
    ;;("''" 0 (12 . 41) nil t) ;;; comment
    ;; ("```" 0 (7 . 41) nil t) ;;; string
    ("\\(```\\)"
     (1 (lns-syntax-multi-string ?`) t t))
    ;; ("\\(\"\\)"
    ;;  (1 (lns-syntax-multi-string ?\") t t))
    ;; ("[^\?]\\(\'\\)"
    ;;  (1 "\"")) ;;; string
    ("\'" 0 (7 . 41) nil t) ;;; string
    ;;("?" 0 (10 . 40) nil t) ;; char-quote
    ;; ("\?" 0 (10 . 40) nil t) ;;; char-quote
    ;; ("\\?"
    ;;  (1 "(" t t))
    ))

(defvar lns-mode-syntax-table
  (with-syntax-table (copy-syntax-table)
    (modify-syntax-entry ?\\ "\\")
    (modify-syntax-entry ?_ "w")
    (modify-syntax-entry ?& ".")
    (modify-syntax-entry ?  " ")
    (modify-syntax-entry ?\t " ")
    ;;(modify-syntax-entry ?? ".")
    ;;(modify-syntax-entry ?\" ".")

    ;; /* */ //
    (modify-syntax-entry ?/ ". 124b")
    (modify-syntax-entry ?\* ". 23")
    (modify-syntax-entry ?\n "> b")
    (syntax-table))
  "`lns-mode' syntax table.")

(defvar lns-imenu-generic-expression
  '((nil "\\(\\_<fn\\_>\\|\\_<class\\_>\\)[ \t]*\\([A-Za-z0-9_\.]+\\)" 2)))

(defvar lns-proj-file "lune.js")

(defun lns-proj-search ()
  ""
  (let ((dir (expand-file-name default-directory)))
    (when buffer-file-name
      (setq dir (file-name-directory (buffer-file-name))))
    (while (and dir
		(not (file-exists-p (expand-file-name lns-proj-file dir))))
      (if (equal dir "/")
	  (setq dir nil)
	(setq dir (file-name-directory (directory-file-name dir)))))
    (or dir default-directory)
    ))

(defun lns-get-proj-dir ()
  (if (boundp 'lns-proj-dir)
      lns-proj-dir
    (lns-proj-search)))

(defun lns-get-proj-info (&optional proj-dir)
  (when (not proj-dir)
    (setq proj-dir (lns-get-proj-dir)))
  (let ((conf-path (expand-file-name lns-proj-file proj-dir)))
    (if (file-exists-p conf-path)
	(with-temp-buffer
	  (insert-file-contents conf-path)
	  (let ((json-object-type 'plist)
		(json-array-type 'list)
		obj)
	    (setq obj (json-read-from-string
		       (buffer-substring-no-properties (point-min) (point-max))))
	    obj
	    )
	  )
      nil)
    ))

(defun lns-proj-info-get-conf (proj)
  "proj 内に lune.js の静的情報を conf として持つ。"
  (plist-get proj :conf))

(defun lns-proj-info-set-lnsc-path (proj lnsc-path)
  "lnsc のパスをキャッシュする"
  (plist-put proj :lnsc-path lnsc-path)
  )
(defun lns-proj-info-get-lnsc-path (proj)
  "キャッシュしている lnsc のパス"
  (plist-get proj :lnsc-path)
  )
(defun lns-proj-info-get-use-indent-command-p (proj)
  "インデントに外部コマンドを利用するかどうか"
  (plist-get proj :use-indent-comamnd)
  )

(defun lns-proj-get-cmd-option (proj)
  (plist-get (lns-proj-info-get-conf proj) :cmd_option)
  )

(defun lns-proj-get-lnsc-path (proj)
  (plist-get (lns-proj-info-get-conf proj) :lnsc_path)
  )

(defvar lns-proj-info-list nil)


(defun lns-proj-info-new (proj-dir)
  (let ((info (list :conf (lns-get-proj-info proj-dir)))
	formatter-dir formatter-path)
    (add-to-list 'lns-proj-info-list (list proj-dir info))
    (plist-put info :proj-dir proj-dir)
    ;; lnsc の場所を更新
    (lns-command-get-lnsc)
    ;; formatter の場所を登録
    (setq formatter-dir (funcall lns-get-formatter-path
				 (lns-proj-info-get-lnsc-path info)))
    (cond
     ((and (setq formatter-path (expand-file-name "formatter" formatter-dir))
	   (file-exists-p formatter-path))
      (plist-put info :indent-command-path formatter-path))
     ((and (setq formatter-path (expand-file-name "lns/main.lns" formatter-dir))
	   (file-exists-p formatter-path))
      (plist-put info :indent-command-script formatter-path)))
    (plist-put info :use-indent-comamnd
	       (funcall lns-proj-info-check-useing-indent-command proj-dir))
    ))

(defun lns-proj-info-get (&optional proj-dir new)
  (when (not proj-dir)
    (setq proj-dir (lns-get-proj-dir)))
  (let ((info (assoc proj-dir lns-proj-info-list)))
    (cond
     ((and (not info) new)
      (setq info (lns-proj-info-new proj-dir)))
     (t
      (setq info (nth 1 info))))
    info
    ))

(defun lns-proj-reload (&optional proj-dir)
  (interactive)
  (let ((info (lns-proj-info-get proj-dir)))
    ;; lns-proj-info-list から proj-dir のプロジェクト情報を削除する
    (setq lns-proj-info-list
	  (delq nil (mapcar (lambda (X)
			      (when (not (equal (car X)
						(plist-get info :proj-dir)))
				X))
			    lns-proj-info-list)))
    (lns-proj-info-get proj-dir t)
    ))


;;;###autoload
(define-derived-mode lns-mode lns--prog-mode "Lns"
  "Major mode for editing Lns code."
  :syntax-table lns-mode-syntax-table
  :group 'lns

  (set (make-local-variable 'font-lock-defaults)
       `(lns-font-lock-keywords ;; keywords
         nil                    ;; keywords-only
         nil                    ;; case-fold
         nil                    ;; syntax-alist
         nil                    ;; syntax-begin
         (font-lock-syntactic-keywords . ,lns-font-lock-syntactic-keywords)
         (font-lock-extra-managed-props . (syntax-table))
	 ;; comment 内を対象外にする
	 (parse-sexp-ignore-comments . t )
         (parse-sexp-lookup-properties . t)
         (beginning-of-defun-function . lns-beginning-of-fn)
         (end-of-defun-function . lns-end-of-fn)
         (indent-line-function . lns-indent-line)
	 (indent-region-function . lns-indent-region)
         (comment-use-syntax . t)
         (comment-use-global-state . t)
	 (indent-tabs-mode . nil)
	 (comment-start . "//")
	 (imenu-generic-expression . ,lns-imenu-generic-expression)
         ))
  (set (make-local-variable 'isearch-wrap-function) 'lns-isearch-wrap-function)
  (set (make-local-variable 'lns-proj-dir) (lns-proj-search))
  (set (make-local-variable 'lns-proj-info) (lns-proj-info-get lns-proj-dir t))
  (set (make-local-variable 'flycheck-executable-find) 'lns-flycheck-executable-find)
  ;; 改行時、複数回 indent が走ると重いので、electric が動作しないようにする。
  ;; また、 ブロックで indent が走るように {} を登録する。
  (set (make-local-variable 'electric-indent-chars) `(?{ ?}))
  (when (boundp 'company-backends)
    (set (make-local-variable 'company-backends)
	 (delq 'company-dabbrev (copy-sequence company-backends))))


  (local-set-key (kbd "C-c C-e") 'lns-dispatch-mode)
  (local-set-key (kbd "C-c I") 'lns-show-inq)
  )


(defun lns-is-in-comment-string( pos )
  (save-excursion
    (if (>= (point-min) pos)
	nil
      (let ((syntax (syntax-ppss pos)))
	(or (elt syntax 3)
	    (elt syntax 4)
	    (eq (get-char-property pos 'face) 'font-lock-comment-face)
	    (eq (get-char-property pos 'face) 'font-lock-string-face)
	    (eq (char-after) ?')
	    )))))
(defun lns-is-in-comment( pos )
  (save-excursion
    (if (>= (point-min) pos)
	nil
      (let ((syntax (syntax-ppss pos)))
	(or (elt syntax 4)
	    (eq (get-char-property pos 'face) 'font-lock-comment-face)
	    )))))
(defun lns-is-in-string( pos )
  (save-excursion
    (if (>= (point-min) pos)
	nil
      (let ((syntax (syntax-ppss pos)))
	(or (elt syntax 3)
	    (eq (get-char-property pos 'face) 'font-lock-string-face)
	    )))))


(defun lns-indent-prev-eol ()
  "前の行の行末に移動する。
空行は飛ばす。
移動できた場合 nil 以外。
"
  (let (line-num)
    (if (eq (point-max) (point))
	(setq line-num (count-lines 1 (point)))
      (setq line-num (count-lines 1 (1+ (point)))))
    (if (eq 1 line-num)
	nil
      (beginning-of-line)
      (previous-line)
      (end-of-line)
      (if (and (eq (current-column) 1)
	       (string-match "[^\s \t]"
			     (buffer-substring-no-properties (1- (point)) (point))))
	  t
	(re-search-backward "[^\s \t]" (point-min) 'noerror)
	(when (eq (char-before) 10)
	  (re-search-backward "[^\s \t]" (point-min) 'noerror))
	(end-of-line)
	(re-search-backward "[^\s \t]" (point-min) 'noerror)
	(forward-char)
	t
	))))


(defun lns-indent-search-block (pattern)
  "
現在のカーソル位置から前方に pattern の文字列を見つける。

pattern は  {, }, {{, }} のいずれか。

前方方向に閉じたペアがある場合、そのペアは飛ばす。
"
  (let ((pattern-char (substring pattern 1))
	(count 0)
	(loop t)
	end-pos pos)
    (while loop
      (save-excursion
	(beginning-of-line)
	(setq end-pos (point)))
      (let ((find t)
	    (pattern-char (substring pattern 1))
	    find-len work)
	(while find
	  (setq find (re-search-backward "[{}]" end-pos 'noerror))
	  (when find
	    (when (not (lns-is-in-comment-string find))
	      (if (eq (char-after) (char-before))
		  (progn
		    (setq find-len 2)
		    (setq work (format "%c%c" (char-after) (char-after))))
		(setq find-len 1)
		(setq work (format "%c" (char-after))))
	      (if (not (eq (length pattern) find-len))
		  ;; 検索対象と文字数が異なるのでスキップ
		  (backward-char (1- find-len))
		;; 検索対象と文字数が同じ
		(if (eq count 0)
		    ;; 検索対象は閉じている
		    (if (equal work pattern)
			(progn
			  ;; 検索対象と等しい
			  (setq pos (- (point) (1- (length pattern))))
			  (setq find nil))
		      ;; 検索対象と異なる
		      (setq count (+ count 1))
		      (backward-char (1- find-len)))
		  ;; 検索対象は閉じていない
		  (if (equal work pattern)
		      ;; 検索対象と等しい
		      (setq count (- count 1))
		    ;; 検索対象と異なる
		    (setq count (+ count 1)))
		  (backward-char (1- find-len))))))
	  (if (< (point) end-pos)
	      (setq find nil))
	  ))
      (if pos
	  (setq loop nil)
	(when (not (lns-indent-prev-eol))
	  (setq loop nil)))
      )
    (when pos
      (goto-char pos))
    ))

(defun lns-indent-is-line-no-term ()
  "現在の行が終端されていないか調べる"
  (save-excursion
    (end-of-line)
    (let (pos)
      (save-excursion
	(beginning-of-line)
	(setq pos (point)))
      (if (not (re-search-backward "[^\s \t]" pos t))
	  nil
	(not (or (eq (char-after) ?\;) (eq (char-after) ?\})))
	))))
  
  

(defun lns-indent-search-open-pair ()
  "現在のカーソル位置から前方向で、閉じていない括弧を探す。
あるいは、行頭の文の開始位置を見つける。
"
  (let ((count-brace 0) ;; {
	(count-paren 0) ;; (
	(count-bracket 0) ;; [
	(org-pos (point))
	end-pos start-pos pos)
    (save-excursion
      (while (not pos)
	(setq start-pos (point))
	(save-excursion
	  (beginning-of-line)
	  (setq end-pos (point)))
	(let ((searchFlag t)
	      find-len work ignore)
	  (while searchFlag
	    (setq ignore nil)
	    (setq searchFlag (re-search-backward "\\([{}()]\\|\\[\\|\\]\\)"
						 end-pos 'noerror))
	    (if (and searchFlag
		     (lns-is-in-comment-string searchFlag ))
		(setq ignore t)
	      (when (or (eq (char-before) ??)
			(eq (char-before) ?\\))
		(save-excursion
		  (backward-char 1)
		  (re-search-backward "[^\\]" end-pos 'noerror)
		  (setq ignore (eq (% (- searchFlag (point)) 2) 0)))))
	    (when (and (not ignore)
		       searchFlag (not (lns-is-in-comment-string searchFlag)))
	      (cond
	       ((eq (char-after) ?})
		(setq count-brace (+ count-brace 1))
		)
	       ((eq (char-after) ?\))
		(setq count-paren (+ count-paren 1))
		)
	       ((eq (char-after) ?\])
		(setq count-bracket (+ count-bracket 1))
		)
	       
	       ((eq (char-after) ?{)
		(cond
		 ((eq count-brace 0)
		  (setq searchFlag nil) (setq pos (point)))
		 (t
		  (setq count-brace (1- count-brace))
		  )))
	       ((eq (char-after) ?\()
		(cond
		 ((eq count-paren 0)
		  (setq searchFlag nil) (setq pos (point)))
		 (t
		  (setq count-paren (1- count-paren))
		  )))
	       ((eq (char-after) ?\[)
		(cond
		 ((eq count-bracket 0)
		  (setq searchFlag nil) (setq pos (point)))
		 (t
		  (setq count-bracket (1- count-bracket))
		  )))
	       ))
	    (when (< (point) end-pos)
	      (setq searchFlag nil))
	    ))
	(when (not pos)
	  ;; 行内に該当位置を見つけられなかったので、行頭に文の先頭がないか確認する
	  (beginning-of-line)
	  (if (and
	       (eq count-brace 0)
	       (eq count-paren 0)
	       (eq count-bracket 0)
	       (re-search-forward
		(format "^[\s \t]*\\(%s\\)" lns-bloak-statement-head)
		start-pos t)
	       (not (lns-is-in-comment-string (point)))
	       )
	      ;; 文の先頭が見つかった場合
	       (progn
		 (goto-char (match-beginning 0))
		 (re-search-forward "[^\s \t]")
		 (if (re-search-backward "[\s \t]" end-pos t)
		     (forward-char)
		   (beginning-of-line))
		 (setq pos (point)))
	    (when (not (lns-indent-prev-eol))
	      (setq pos 0)))
	  )
	(when (eq org-pos (point))
	  (setq pos 0))
	))
    (when pos
      (goto-char pos))
    (if (eq pos 0)
	nil
      pos)
    ))

(defun lns-re-search-forward-eol (pattern)
  (let ((loopFlag t)
	(org-pos (point))
	find pos )
    (save-excursion
      (end-of-line)
      (setq pos (point)))
    (while loopFlag
      (setq find (re-search-forward pattern pos t))
      (if find
	  (if (lns-is-in-comment-string find)
	      (progn
		(setq find nil)
		(forward-char)
		(when (< pos (point))
		  (setq loopFlag nil))
		)
	    (setq loopFlag nil))
	(setq loopFlag nil)))
    (when (not find)
      (goto-char org-pos))
    find))

(defun lns-indent-goto-no-space-bol ()
  (beginning-of-line)
  (when (not (string-match "[^ \t]"
			   (buffer-substring-no-properties (point)
							   (1+ (point)))))
    (re-search-forward "[^ \t]")
    (backward-char)))

(defun lns-get-current-token ()
  (interactive)
  (save-excursion
    (let (symstart symend)
      (setq symstart (re-search-backward lns-no-token-pattern nil t))
      (when symstart
	(setq symstart (1+ symstart)))
      (forward-char)
      (setq symend (re-search-forward lns-no-token-pattern nil t))
      (when symend
	(setq symend (1- symend)))
      (when (and symstart symend)
	(buffer-substring symstart symend)))))

(defun lns-indent-to (indent)
  (let ((org-pos (point)))
    (if (not (lns-indent-search-open-pair))
	(setq column 0)
      (if (or (eq (char-after) ?{)
	      (eq (char-after) ?\()
	      (eq (char-after) ?\[))
	  (progn
	    (setq pos (point))
	    (forward-char)
	    (if (lns-re-search-forward-eol "[^ \t]")
		(progn
		  (goto-char pos)
		  (setq column (+ (current-column) 2)))
	      (lns-indent-goto-no-space-bol)
	      (setq column (+ (current-column) lns-indent-level))))
	(cond ((equal (lns-get-current-token) "case")
	       (setq column (current-column))
	       (save-excursion
		 (if (re-search-forward "}" org-pos t)
		     (setq column (+ column indent))
		   (setq column (+ column 5)))))
	      ((equal (lns-get-current-token) "fn")
	       (setq column (current-column)))
	      ;; ((and (lns-get-current-token)
	      ;; 	    (string-match lns-bloak-statement-head
	      ;; 			  (lns-get-current-token)))
	      ;;  (setq column (current-column)))
	      (t
	       (setq column (+ (current-column) indent))))
	))))

(defvar lns-indent-region-running nil)


(defun lns-indent-by-command-result (json-obj)
  ;;; インデント結果を反映
  (let ((cur-line (lns-get-line))
	(cur-marker (point-marker))
	return-point)
    (dolist (X (plist-get (plist-get json-obj :indent) :lines))
      (let* ((info (car (cdr X)))
	     (indent (1- (plist-get info :column)))
	     (lineNo (plist-get info :lineNo))
	     (same (plist-get info :same)))
	(when (>= cur-line lineNo) 
	  (goto-line lineNo)
	  (if (eq same :json-true)
	      (move-to-column indent)
	    (indent-line-to indent))
	  (when (equal cur-line (lns-get-line))
	    (setq return-point (point)))
	  )))
    (goto-char (cond
		(return-point
		 (if (< return-point cur-marker)
		     cur-marker
		   return-point))
		(t
		 cur-marker)))
    ;; marker の後処理
    (set-marker cur-marker nil)
    ))

(defun lns-process-line-by-command-main (newline-p tab-p &optional is-range start end)
  ;;; formater を使ったインデント
  (let ((proj (lns-proj-info-get))
	command-path command-list workbuf lns-code json-obj indent pos
	)

    (save-excursion
      (end-of-line)
      (setq lns-code (concat (buffer-substring-no-properties (point-min) (point))
			     " ___LNS___"
			     (buffer-substring-no-properties (point) (point-max))
			     " ___LNS___"
			     )))
    (setq pos (cond (newline-p
		     (save-excursion
		       (setq end (point))
		       (previous-line)
		       (setq start (point)))
		     (format "%d:%d"
			     (line-number-at-pos start)
			     (line-number-at-pos end)))
		    (is-range
		     (when (or (not start) (not end))
		       (setq start (region-beginning))
		       (setq end (region-end)))
		     (format "%d:%d"
			     (line-number-at-pos start)
			     (line-number-at-pos end)))
		    (t
		     (format "%d" (lns-get-line)))))
    (setq command-list (list "-i" pos "@-" ""))
    (setq workbuf (lns-get-buffer "*lns-process*" t))
    ;; formatter 実行
    (cond
     ((setq command-path (plist-get proj :indent-command-script))
      (lns-execute-command nil workbuf lns-code
			   (append (list command-path "-shebang") command-list)))
     ((setq command-path (plist-get proj :indent-command-path))
      (with-temp-buffer
	(setq default-directory (plist-get proj :proj-dir))
	(insert lns-code)
	(apply 'call-process-region (point-min) (point-max)
	       command-path nil workbuf nil command-list)))
     (t
      (error "not set the indent-command"))
     )

    (condition-case err
	(progn
	  ;; formatter 結果反映
	  (with-current-buffer workbuf
	    (setq json-obj
		  (let ((json-object-type 'plist)
			(json-array-type 'list))
		    (json-read-from-string
		     (buffer-substring-no-properties (point-min) (point-max))))))

	  (cond ((or newline-p tab-p)
		 (lns-indent-by-command-result json-obj))
		(t
		 (save-excursion
		   (lns-indent-by-command-result json-obj))))

	  (when (and transient-mark-mode mark-active)
	    (deactivate-mark))
	  )
      (error (princ (format "failed to detect the indent. check '*lns-process*' buffer")))
      )))
(defun lns-process-line-by-command (&optional is-range start end)
  ;; どの関数から実行されたか調べる
  (let (newline-p tab-p)
    (let ((backtrace-list (reverse (cdr (reverse (backtrace-frames))))))
      (dolist (frame backtrace-list)
	(let ((func (nth 1 frame)))
	  (cond
	   ((equal func 'newline-and-indent)
	    (setq newline-p t))
	   ((equal func 'indent-for-tab-command)
	    (setq tab-p t))
	   ))))
    (lns-process-line-by-command-main newline-p tab-p is-range start end)
    ))
(defun lns-indent-region-by-command ()
  (interactive)
  (lns-process-line-by-command t))




(defun lns-indent-line (&optional force)
  (cond
   ((lns-proj-info-get-use-indent-command-p (lns-proj-info-get))
    (lns-process-line-by-command))
   (t
    (if (and (not force)
	     (not lns-indent-region-running)
	     (eq (lns-get-column) 1))
	;; コメント改行時、
	(run-at-time 0.01 nil (lambda () (lns-indent-line-sub)))
      (lns-indent-line-sub)))))

(defun lns-indent-line-sub ()
  (let ((case-fold-search nil)
	(org-pos (point))
	pos end-block-flag start-block-flag start-pos column)
    (save-excursion
      (beginning-of-line)
      (cond ((and (lns-is-in-comment (point))
		  (lns-is-in-comment (1- (point))))
	     (if (or lns-indent-region-running
		     (eq (lns-get-line) 1))
		 (setq column -1)
	       ;;行頭がコメント、文字列の場合は直前の行の先頭に合せる
	       (previous-line)
	       (beginning-of-line)
	       (when (looking-at "[ \t]")
		 (re-search-forward "[^ \t]" nil t)
		 (backward-char))
	       (if (lns-is-in-comment (point))
		   (setq column (current-column))
		 (setq column 0))))
	    ((and (lns-is-in-string (point))
		  (lns-is-in-string (1- (point))))
	     ;; 文字列の場合は、インデント調整しない。
	     (setq column -1))
	    (t
	     (re-search-forward "[^ \t]" nil t)
	     (cond 
	      ((eq (char-before) ?})
	       (if (eq (char-after) ?})
		   (setq end-block-flag "{{")
		 (setq end-block-flag "{")))
	      ((eq (char-before) ?{)
	       (if (eq (char-after) ?{)
		   (setq start-block-flag "}}")
		 (setq start-block-flag "}")))
	      ((eq (char-before) ?\))
	       (setq end-block-flag ")"))
	      ((eq (char-before) ?\()
	       (setq start-block-flag ")"))
	      ((eq (char-before) ?\])
	       (setq end-block-flag "["))
	      ((eq (char-before) ?\[)
	       (setq start-block-flag "]"))
	      )
	     (when (not (lns-indent-prev-eol))
	       (setq column 0))
	     (cond
	      (column
	       nil)
	      (end-block-flag
	       ;; ブロック終了の場合、ブロック開始を見つける。
	       (when (lns-indent-search-open-pair)
		 (if (or (eq (char-after) ?{)
			 (eq (char-after) ?\()
			 (eq (char-after) ?\[))
		     (progn
		       ;; 括弧
		       (lns-indent-goto-no-space-bol)
		       (setq column (current-column)))
		   ;; 行の先頭が見つかった
		   (setq column (- (current-column) lns-indent-level))
		   )))
	      (start-block-flag
	       ;; ブロック開始の場合
	       (save-excursion
		 (goto-char org-pos)
		 (lns-indent-prev-eol)
		 (when (eq (char-before) ?})
		   (setq column (1- (current-column)))))
	       (when (not column)
		 (save-excursion
		   (lns-indent-search-open-pair)
		   (setq column (current-column))
		   (setq pos (point)))
		 (save-excursion
		   (when (lns-indent-search-open-pair)
		     (if (or (eq (char-after) ?{)
			     (eq (char-after) ?\()
			     (eq (char-after) ?\[))
			 (progn
			   (setq pos (point))
			   (forward-char)
			   (if (lns-re-search-forward-eol "[^\\s \t]")
			       (progn
				 (goto-char pos)
				 (setq column (+ (current-column) 2)))
			     (lns-indent-goto-no-space-bol)
			     (setq column (+ (current-column) lns-indent-level))))
		       (if (equal start-block-flag "}")
			   (setq column (current-column))
			 (setq column (+ 4 (current-column))))
		       )))))
	      (t
	       ;; ブロック開始、終了でない場合、一つ前のインデントに併せる
	       (cond ((lns-is-in-comment-string (1- (point)))
		      (lns-indent-to 0))
		     (t (save-excursion
			  (goto-char org-pos)
			  (lns-indent-prev-eol)
			  (when (eq (char-before) ?})
			    (if (and (> (point) 0)
				     (eq (char-before (1- (point))) ?{))
				;; {} の場合は { に合せる
				(setq column (- (current-column) 2))
			    (setq column (1- (current-column)))))
			  (when (not column)
			    (goto-char org-pos)
			    (lns-indent-prev-eol)
			    (if (lns-indent-is-line-no-term)
				(lns-indent-to lns-indent-level)
			      (lns-indent-to 0))))))
	       )
	      ))))
    (let ((marker (point-marker)))
      (when (>= column 0)
	(when (and (not start-block-flag) (not end-block-flag))
	  (move-to-column column t))
	(indent-line-to column))
      (when (> marker (point))
	(goto-char marker)))
    ))

(defun lns-indent-region (start end)
  (cond
   ((lns-proj-info-get-use-indent-command-p (lns-proj-info-get))
    (lns-process-line-by-command t start end))
   (t
    (let ((indent-region-function nil)
	  (lns-indent-region-running t))
      (indent-region start end nil) 
      ))))


(defun lns-beginning-of-fn (&optional arg)
  ""
  (interactive "P")
  (error "not support now"))


(defun lns-end-of-fn (&optional arg)
  (interactive "P")
  (error "not support now")
)

(defun lns-forward-sexp (&optional count)
  (interactive "p")
  (error "not support now")
)

(defun lns-get-subfile-list (mode)
  (let ((txt (buffer-substring-no-properties (point-min) (point-max)))
	list start-pos)
    (while (setq start-pos
		 (string-match (concat "^subfile[ \t]+" mode "[ \t]+.+\\.\\([^\.;]+\\);")
			       txt start-pos))
      (setq start-pos (1+ start-pos))
      (setq list (append list (list (concat (match-string 1 txt) ".lns")))))
    list))

(defun lns-isearch-wrap-function ()
  "subfile を辿って isearch"
  ;;(isearch-dehighlight)
  (let ((now-file (file-name-nondirectory (buffer-file-name)))
	file-list next-file owner-file find-flag)
    (cond
     ((setq file-list (lns-get-subfile-list "use"))
      ;; サブファイルがある場合、サブファイルを開いて検索する
      (when (not isearch-forward)
	(setq file-list (reverse file-list)))
      (setq owner-file now-file)
      )
     ((setq file-list (lns-get-subfile-list "owner"))
      ;; オーナーファイルがある場合、次のサブファイルを開いて検索する
      (setq owner-file (car file-list))
      (find-file owner-file)
      (setq file-list (lns-get-subfile-list "use"))
      (when (not isearch-forward)
	(setq file-list (reverse file-list)))
      (setq file-list (cdr (member now-file file-list)))
      ))
    (setq next-file (car file-list))
    (while (and (not find-flag) next-file)
      (find-file next-file)
      (if isearch-forward
	  (progn
	    (goto-char 1)
	    (if (isearch-search-string isearch-string (point-max) t)
		(setq find-flag t)
	      (setq file-list (cdr file-list))
	      (setq next-file (car file-list))))
	(goto-char (point-max))
	(if (isearch-search-string isearch-string nil t)
	    (setq find-flag t)
	  (setq file-list (cdr file-list))
	  (setq next-file (car file-list)))
	))
    (when (and (not find-flag) owner-file)
      (find-file owner-file))
    )
  (setq isearch-other-end nil)
  (cond
   (isearch-forward
    (goto-char 1)
    ;;(isearch-repeat-forward)
    )
   (t
    (goto-char (point-max)))
   ))

(provide 'lns-mode)
