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
    

(defun lns-command-get-command (&rest args)
  (let (command)
    (if (functionp lns-lua-command)
	(setq command (funcall lns-lua-command))
      (setq command lns-lua-command))
    (lns-command-add-command
     (list command "-e" "require( 'lune.base.base' )" " ")
     (append args (lns-proj-get-cmd-option (lns-get-proj-info)) ))))

(defun lns-command-sync (&rest arg-list)
  (let (ver end-code)
    (with-temp-buffer
      (setq buffer-file-name "test")
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
