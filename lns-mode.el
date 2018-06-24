;;; lns-mode.el --- a major-mode for editing LuneScript


;; (add-to-list 'auto-mode-alist '("\\.lns$" . lns-mode))


(defalias 'lns--prog-mode
  (if (fboundp 'prog-mode) 'prog-mode 'fundamental-mode))


(defvar lns-indent-level 3)
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


(defun lns-indent-line ()
  0
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
