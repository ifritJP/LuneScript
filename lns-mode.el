;;; lns-mode.el --- a major-mode for editing LuneScript


;; (add-to-list 'auto-mode-alist '("\\.lns$" . lns-mode))


(defalias 'lns--prog-mode
  (if (fboundp 'prog-mode) 'prog-mode 'fundamental-mode))


(defvar lns-indent-level 4)
(defvar lns-command-path "lns" )

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
		   "loadstring" "module" "next" "pairs" "pcall" "print"
		   "rawequal" "rawget" "rawlen" "rawset" "require" "select"
		   "setfenv" "setmetatable" "tonumber" "tostring"
		   "type" "unpack" "xpcall")))
  (defconst
    lns-keyword (lns-make-regex-or
		 '("self" "let" "fn" "if" "elseif" "else" "while" "repeat" "for"
		   "apply" "of" "foreach" "in" "return" "class" "false" "nil" "true"
		   "mut" "pub" "pro" "pri" "form" "advertise" "wrap" "static"
		   "trust" "import" "as" "not" "and" "or")))
  (defconst
    lns-bloak-statement-head (lns-make-regex-or
			      '("let" "fn" "if" "elseif" "else" "while" "repeat" "for"
				"apply" "foreach" "class" 
				"pub" "pro" "pri" "form" "advertise" "wrap" "static"
				"trust" "import")))
  (defconst
    lns-type (lns-make-regex-or
	      '("int" "real" "int_" "real_" "stem" "Map" "Array" "List" "str")))
  )
  

(defvar lns-font-lock-keywords
  `((,lns-builtin . font-lock-builtin-face)
    (,lns-keyword . font-lock-keyword-face)
    ("@@\\?\\|@@\\|@\\|\\?\\|&" . font-lock-keyword-face)
    (,lns-type . font-lock-type-face)
    ;; (":\\s-*\\(\\w+\\)" . (1 font-lock-type-face nil t))
    (,(format "\\(%s\\)\\s-*\\(\\w+\\)"
	      (lns-make-regex-or '("fn" "form"))) .
	      ( 2 font-lock-function-name-face nil t))    
    (,(format "\\(%s\\)\\s-*\\(\\w+\\)"
	      (lns-make-regex-or '("let" "for" "foreach" "apply"))) .
	      ( 2 font-lock-variable-name-face nil t))
    ))

(defvar lns-font-lock-syntactic-keywords
  `(("'''" 0 (14 . 41) nil t) ;;; comment
    ("''" 0 (11 . 40) nil t) ;;; comment
    ("\n" 0 (12 . 40) nil t) ;;; comment
    ;;("''" 0 (12 . 41) nil t) ;;; comment
    ("```" 0 (7 . 41) nil t) ;;; string
    ("\"" 0 (7 . 41) nil t) ;;; string
    ("\'" 0 (7 . 41) nil t) ;;; string
    ;;(,(lambda (limit) (re-search-forward "\"\"\"" limit t)) 0 (7 . 41) nil t) ;;; string
    ))


(defvar lns-mode-syntax-table
  (with-syntax-table (copy-syntax-table)
    (modify-syntax-entry ?\\ "\\")
    (modify-syntax-entry ?_ "w")
    (modify-syntax-entry ?& ".")
    (modify-syntax-entry ?  " ")
    (modify-syntax-entry ?\t " ")
    (syntax-table))
  "`lns-mode' syntax table.")


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
         (font-lock-syntactic-keywords  . ,lns-font-lock-syntactic-keywords)
         (beginning-of-defun-function   . lns-beginning-of-fn)
         (end-of-defun-function         . lns-end-of-fn)
         (indent-line-function          . lns-indent-line)
         (comment-use-syntax            . t)
         (comment-use-global-state      . t)
         ))
  )


(defun lns-is-in-comment-string( pos )
  (let ((syntax (syntax-ppss pos)))
    (or (elt syntax 3)
	(elt syntax 4)
	(eq (get-char-property pos 'face) 'font-lock-comment-face)
	(eq (char-after) ?')
	)))

(defun lns-indent-prev-eol ()
  (let (line-num)
    (if (eq (point-max) (point))
	(setq line-num (count-lines 1 (point)))
      (setq line-num (count-lines 1 (1+ (point)))))
    (if (eq 1 line-num)
	nil
      (beginning-of-line)
      (previous-line)
      (end-of-line)
      t
      )))


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



(defun lns-indent-search-open-pair ()
  "現在のカーソル位置から前方向で、閉じていない括弧を探す。
あるいは、行頭の文の開始位置を見つける。
"
  (let ((count-brace 0) ;; {
	(count-paren 0) ;; (
	(count-bracket 0) ;; [
	end-pos start-pos pos)
    (save-excursion
      (while (not pos)
	(setq start-pos (point))
	(save-excursion
	  (beginning-of-line)
	  (setq end-pos (point)))
	(let ((searchFlag t)
	      find-len work)
	  (while searchFlag
	    (setq searchFlag (re-search-backward "\\([{}()]\\|\\[\\|\\]\\)"
						 end-pos 'noerror))
	    (when (and searchFlag (not (lns-is-in-comment-string searchFlag)))
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
		(format "^[\s \t]*\\(%s\\)" lns-bloak-statement-head) start-pos t))
	      (progn
		(if (re-search-backward "[\s \t]" end-pos t)
		    (forward-char)
		  (beginning-of-line))
		(setq pos (point)))
	    (when (not (lns-indent-prev-eol))
	      (setq pos 0)))
	  )))
    (when pos
      (goto-char pos))
    pos
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

(defun lns-indent-line ()
  (let ((org-pos (point))
	pos end-block-flag start-block-flag start-pos column)
    (save-excursion
      (beginning-of-line)
      (re-search-forward "[^\\s \t]")
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
      (lns-indent-prev-eol)
      (cond
       (end-block-flag
  	;; ブロック終了の場合、ブロック開始を見つける。
  	(when (lns-indent-search-open-pair)
	  (if (or (eq (char-after) ?{)
		  (eq (char-after) ?\()
		  (eq (char-after) ?\[))
	      (progn
		;; 括弧
		(beginning-of-line)
		(re-search-forward "[^\\s \t]")
		(backward-char)
		(setq column (current-column)))
	    ;; 行の先頭が見つかった
	    (setq column (- (current-column) lns-indent-level))
	    )))
       (start-block-flag
	;; ブロック開始の場合
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
		  (beginning-of-line)
		  (re-search-forward "[^\\s \t]")
		  (backward-char)
		  (setq column (+ (current-column) lns-indent-level)))
	      (setq column (current-column))
	      ))))
       (t
	;; ブロック開始、終了でない場合、
	(if (not (lns-indent-search-open-pair))
	    (setq column 0)
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
		  (beginning-of-line)
		  (re-search-forward "[^\\s \t]")
		  (backward-char)
		  (setq column (+ (current-column) lns-indent-level))))
	    (setq column (current-column)))))
	    ;;(setq column (+ (current-column) lns-indent-level)))))
       ))
    (when column
      (move-to-column column t)
      (indent-line-to column))
    )
  )


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

(provide 'lns-mode)
