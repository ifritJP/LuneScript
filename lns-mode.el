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
		   "trust" "import" "as")))
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
    ("\"\"\"" 0 (7 . 41) nil t) ;;; string
    ("\"" 0 (7 . 41) nil t) ;;; string
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
	(elt syntax 4))))


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
	(previous-line)
	(end-of-line))
      )
    (when pos
      (goto-char pos))
    ))

(defun lns-indent-search-pattern-in-line (pattern &optional independ skip-pattern)
  (end-of-line)
  (let ((end-pos (point))
	(find t)
	pos work)
    (beginning-of-line)
    (when (not independ)
      (setq pattern (concat "\\s *" pattern)))
    (while find
      (setq find (re-search-forward pattern end-pos t))
      (when find
	(when (not (lns-is-in-comment-string find))
	  (if skip-pattern
	      (progn
		(setq work (buffer-substring-no-properties
			    find (+ find (length skip-pattern))))
		(if (equal work skip-pattern)
		    (forward-char (length skip-pattern))
		  (setq pos find)
		  (setq find nil)
		  ))
	    (setq pos find)
	    (setq find nil)
	    ))))
    pos))

(defun lns-indent-search-statement-head ()
  (when (re-search-backward
	 (format "^\\s *\\(}\\|{\\|%s\\)" lns-bloak-statement-head) nil t)
    (when (or (eq (char-after) ? ) (eq (char-after) ?\t))
      (when (re-search-forward "[^\t ]")
	(backward-char)
	)
  )))

(defun lns-indent-line ()
  (let (pos end-block-flag start-block-flag start-pos column)
    (save-excursion
      (setq end-block-flag (lns-indent-search-pattern-in-line "}" nil "}"))
      (when (not end-block-flag)
  	(setq start-block-flag (lns-indent-search-pattern-in-line "^\\s *{" t)))
      (previous-line)
      (end-of-line)
      (cond
       (end-block-flag
  	;; ブロック終了の場合、ブロック開始を見つける。
  	(setq pos (lns-indent-search-block "{"))
  	(when pos
  	  (save-excursion
  	    (beginning-of-line)
  	    (setq start-pos (point)))
  	  (backward-char)
  	  (if (or (< (point) start-pos)
  		  (not (re-search-backward "[^\\s ]" start-pos t)))
  	      ;; { と同じ場所にインデント
  	      (progn
  		(goto-char pos)
  		(setq column (current-column)))
	    ;;
	    (goto-char pos)
	    (lns-indent-search-statement-head)
	    ;;(setq column (1- (current-column)))
	    (setq column (current-column))
	    )
  	  ))
       (start-block-flag
	;; ブロック開始の場合
	(save-excursion
	  (lns-indent-search-statement-head)
	  ;;(setq column (1- (current-column)))
	  (setq column (current-column))
	  (setq pos (point)))
	(when (lns-indent-search-block "{")
	  (when (> (point) pos)
	    (setq column (+ column lns-indent-level)))))
       (t
	(setq pos (point))
	(lns-indent-search-statement-head)
	(cond ((eq (char-after) ?{)
	       ;;(setq column (+ (1- (current-column)) lns-indent-level)))
	       (setq column (+ (current-column) lns-indent-level)))
	      ((not (eq (char-after) ?}))
	       ;;(setq column (1- (current-column)))
	       (setq column (current-column))
	       (setq start-pos (point))
	       (goto-char pos)
	       (when (lns-indent-search-block "{")
		 (setq column (+ column lns-indent-level))))
	      (t
	       ;;(setq column (1- (current-column)))))
	       (setq column (current-column))))
	)
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
