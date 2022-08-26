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

(defvar lns-lua-command "lua5.3"
  "lua command. This var can set string and function.
function must return string.")

(defvar lns-lnsc-command nil
  "lnsc command. This var can set string and function.
function must return string.")

(defvar lns-lnsc-args '("--enableTestBlock")
  "lnsc arguments list for lnsc command. This var can set list and function.
function must return list.")
  


(defvar lns-target-lua-ver nil
  "target lua version. 51 or 52 or 53.")


(defvar lns-max-size-search-subfile 10000)

(defun lns-command-get-info ()
  (let ((owner-file (buffer-file-name))
	analyze-module)
    (save-excursion
      (goto-char (point-min))
      (when (re-search-forward "^[ \t]*subfile[ \t]+owner[ \t]+"
			       lns-max-size-search-subfile t)
	(when (looking-at "\\([^;]+\\);")
	  (setq owner-file (buffer-substring-no-properties (match-beginning 1)
							   (match-end 1)))
	  (setq owner-file (lns-convert-module-2-path owner-file))
	  (setq analyze-module (lns-convert-path-2-module (buffer-file-name)))
	  )))
    (list :owner owner-file
	  :module analyze-module
	  :script (buffer-file-name)
	  :dir (lns-get-proj-dir))))

(defun lns-command-add-command (command-list args)
  (dolist (arg args)
    (if (listp arg)
	(setq command-list (lns-command-add-command command-list arg))
      (setq command-list (append command-list (list arg)))))
  (delq nil command-list))

(defun lns-command-get-lnsc ()
  "lune.js 内の lnsc_path 指定を確認し、
そこで指定されているコマンドを lnsc として返す。
lune.js 内に lnsc_path 指定がなければ、
lns-lnsc-command からパスを取得する。
"
  (let (proj lnsc-path)
    (when (boundp 'lns-proj-dir)
      ;; lune.js 内の lnsc-path 指定から確認
      (let ((proj-dir lns-proj-dir))
	(setq proj (lns-proj-info-get))
	(when proj
	  (setq lnsc-path (lns-proj-info-get-lnsc-path proj))
	  (when (not lnsc-path)
	    (dolist (work-path (lns-proj-get-lnsc-path proj))
	      (when (not lnsc-path)
		(setq work-path (expand-file-name work-path proj-dir))
		(when (file-exists-p work-path)
		  (setq lnsc-path work-path)
		  (lns-proj-info-set-lnsc-path proj lnsc-path)
		  ))))
	  )))
    (when (not lnsc-path)
      ;; lnsc-path が見つからない場合、lns-lnsc-command から確認
      (cond ((not lns-lnsc-command)
	     )
	    ((functionp lns-lnsc-command)
	     (setq lnsc-path (funcall lns-lnsc-command)))
	    ((stringp lns-lnsc-command)
	     (setq lnsc-path lns-lnsc-command))
	    (t
	     (error "lns-lnsc-command is illegal")))
      )
    (when proj
      ;; lnsc-path をキャッシュ
      (lns-proj-info-set-lnsc-path proj lnsc-path))
    lnsc-path))


(defun lns-command-get-path ()
  "lune.js 内の lnsc_path 指定を確認し、
そこで指定されているコマンドを lnsc として返す。
lune.js 内に lnsc_path 指定がなければ、
lns-command-get-lnsc からパスを取得する。
"
  (let* ((proj-dir lns-proj-dir)
	 (proj (lns-proj-info-get))
	 lnsc-path)
    (when proj
      (dolist (work-path (lns-proj-get-lnsc-path proj))
	(when (not lnsc-path)
	  (setq work-path (expand-file-name work-path proj-dir))
	  (when (file-exists-p work-path)
	    (setq lnsc-path work-path))))
      lnsc-path)
    (when (not lnsc-path)
      (setq lnsc-path (car (lns-command-get-command nil))))))


(defun lns-command-get-command (&rest args)
  (let ((lnsc t)
	(command (lns-command-get-lnsc))
	command-list)
    (when (not command)
      (setq lnsc nil)
      (if (functionp lns-lua-command)
	  (setq command (funcall lns-lua-command))
	(setq command lns-lua-command)))
    (if lnsc
	(setq command-list (list command))
      (setq command-list (list command "-e" "require( 'lune.base.base' )" " ")))
    (lns-command-add-command
     command-list (append args
			  (cond ((functionp lns-lnsc-args)
				 (funcall lns-lnsc-args))
				(t
				 lns-lnsc-args))))))

(defun lns-command-sync (&rest arg-list)
  (let ((dir default-directory)
	ver end-code)
    (with-temp-buffer
      (setq buffer-file-name (expand-file-name "test" dir))
      (setq end-code (apply 'lns-execute-command nil (current-buffer) nil arg-list))
      (setq ver (buffer-string))
      (setq buffer-file-name nil))
    end-code))
;; (lns-command-sync "--version")

(defun lns-command-exec (callback async owner-file input outbuf mode op-list)
  (let ((out-buf (lns-get-buffer (or outbuf "*lns-process*") t))
	result process)
    (when (not owner-file)
      (setq owner-file buffer-file-name))
    (setq process
	  (apply 'lns-execute-command
	   (if async
	       (lexical-let ((lex-buf out-buf)
			     (lex-callback callback))
		 (lambda (process event)
		   (when lex-callback
		     (condition-case err
			 (let (json)
			   ;;(setq json (lns-json-get lex-buf :candidateList))
			   (setq json (lns-complete-from-json lex-buf))
			   (funcall lex-callback json nil))
		       (error nil
			      (funcall lex-callback nil t))))
		   ))
	     nil)
	   out-buf
	   input
	   (lns-convert-path-2-proj-relative-path owner-file) mode op-list))
    (when (and callback (not async))
      (condition-case err
	  (funcall callback
		   ;;(lns-json-get out-buf :candidateList) nil)
		   (lns-complete-from-json out-buf) nil)
	(error nil
	       (funcall callback nil t))))
    process
    ))


(defun lns-command-compile-meta ()
  (lns-command-compile t))


(defun lns-command-compile (&optional meta)
  (let* ((command-info (lns-command-get-info))
	 (owner-file (plist-get command-info :owner))
	 (analyze-module (plist-get command-info :module)))
    (lns-command-exec nil nil nil "" nil (if meta "SAVE" "save") nil)))

(defun lns-command-glue ()
  (let* ((command-info (lns-command-get-info))
	 (owner-file (plist-get command-info :owner))
	 (analyze-module (plist-get command-info :module)))
    (lns-command-exec nil nil nil "" nil "glue" nil)))


(provide 'lns-command)
