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

(require 'lns-command)

(defun lns-json-get (buf symbol)
  (with-current-buffer buf
    (let ((json-object-type 'plist)
	  (json-array-type 'list)
	  info
	  json-str)
      (goto-char (point-min))
      (if (search-forward "{\"lunescript\":" nil t)
	  (progn
	    (search-backward "{\"lunescript\":" nil t)
	    (setq json-str (buffer-substring-no-properties (point) (point-max))))
	(setq json-str (buffer-string)))
      (setq info (json-read-from-string json-str))
      (plist-get (plist-get info :lunescript) symbol))))

(defun lns-json-val (json symbol)
  (plist-get json symbol))


(defun lns-get-line ()
  (interactive)
  (if (eq (point) (point-max))
      (if (eq (current-column) 0)
	  (1+ (count-lines 1 (point)))
	(count-lines 1 (point)))
    (count-lines 1 (1+ (point)))))

(defun lns-get-column ()
  (interactive)
  (1+ (string-bytes (buffer-substring (point-at-bol) (point)))))

(defun lns-get-buffer (name &optional init)
  (let ((buffer (get-buffer-create name)))
    (when init
      (with-current-buffer buffer
	(setq buffer-read-only nil)
	(erase-buffer)))
    buffer))


(defun lns-candidate-get-type (info)
  (let ((type-txt (lns-json-val info :type)))
    (string-match ".+\\.\\([^\\.]+\\)" type-txt)
    (replace-match "\\1" nil nil type-txt)))
(defun lns-candidate-get-displayTxt (info)
  (lns-json-val info :displayTxt))


;; (defun lns-execute-command ( out-buffer input-txt &rest args )
;;   (let ((proj-dir (lns-get-proj-dir))
;; 	command-list)
;;     (with-temp-buffer
;;       (when input-txt
;; 	(insert input-txt))
;;       (setq default-directory proj-dir)
;;       (setq command-list (apply 'lns-command-get-command args ))
;;       (apply 'call-process-region (point-min) (point-max)
;; 	     (car command-list) nil out-buffer nil
;; 	     (cdr command-list)))
;;     (with-current-buffer out-buffer
;;       (buffer-string))))

(defun lns-execute-command ( async-callback out-buffer input-txt &rest args )
  (let ((proj-dir (or (lns-get-proj-dir) default-directory))
	command-list process)
    (setq command-list (apply 'lns-command-get-command args ))
    (with-temp-buffer
      (setq default-directory proj-dir)
      (when input-txt
	(insert input-txt))
      (if async-callback
	  (progn
	    (setq process (apply 'start-process "lns"
				 out-buffer command-list))
	    (set-process-sentinel process async-callback))
	(setq process (apply 'call-process-region (point-min) (point-max)
			     (car command-list) nil out-buffer nil
			     (cdr command-list)))
	(with-current-buffer out-buffer
	  (buffer-string))))
    (when (and async-callback input-txt)
      (process-send-string process (concat input-txt "\n"))
      (process-send-eof process))
    process))



(defun lns-convert-path-2-proj-relative-path (path)
  (file-relative-name (expand-file-name path) (lns-get-proj-dir)))

(defun lns-convert-path-2-module ( path )
  (let ((module (file-relative-name (expand-file-name path) (lns-get-proj-dir))))
    (setq module (file-name-sans-extension module))
    (setq module (apply 'concat
			(mapcar (lambda (X)
				  (concat "." X ))
				(split-string module "/"))))
    (substring module 1)
    ))

(defun lns-convert-module-2-path ( module )
  (let ((path module))
    (setq path (apply 'concat
		      (mapcar (lambda (X)
				(concat "/" X ))
			      (split-string path "\\."))))
    (concat (expand-file-name (substring path 1) (lns-get-proj-dir)) ".lns")
    ))


(defun lns-complete-from-json (buf)
  (let ((candidate-list (lns-json-get buf :candidateList)))
    (setq candidate-list
	  (sort candidate-list
		(lambda (obj1 obj2)
		  (let* ((info1 (lns-json-val obj1 :candidate))
			 (item-txt1 (lns-candidate-get-displayTxt info1))
			 (info2 (lns-json-val obj2 :candidate))
			 (item-txt2 (lns-candidate-get-displayTxt info2)))
		    (string< item-txt1 item-txt2)))))
    candidate-list))

(defun lns-get-complete-list ( callback async &optional file-path analyze-module buf-name)
  (let ((out-buf (lns-get-buffer (or buf-name "*lns-process*") t))
	result process txt)
    (when (not file-path)
      (setq file-path buffer-file-name))
    (setq txt (concat (buffer-substring-no-properties
		       (point-min) (point)) "lune"))
    (setq process
	  (lns-execute-command
	   (if async
	       (lexical-let ((lex-buf out-buf)
			     (lex-callback callback))
		 (lambda (process event)
		   (condition-case err
		       (let (json)
			 ;;(setq json (lns-json-get lex-buf :candidateList))
			 (setq json (lns-complete-from-json lex-buf))
			 (funcall lex-callback json nil))
		     (error nil
			    (funcall lex-callback nil t)))
		   ))
	     nil)
	   out-buf
	   txt
	   (lns-convert-path-2-proj-relative-path file-path) "comp" 
	   (or analyze-module (lns-convert-path-2-module file-path))
	   (number-to-string (lns-get-line))
	   (number-to-string (lns-get-column)) "-i"))
    (when (not async)
      (condition-case err
	  (funcall callback
		   ;;(lns-json-get out-buf :candidateList) nil)
		   (lns-complete-from-json out-buf) nil)
	(error nil
	       (funcall callback nil t))))
    process
    ))

(defun lns-completion-get-candidate-list (callback &optional async out-bufname)
  (let* ((command-info (lns-command-get-info))
	 (owner-file (plist-get command-info :owner))
	 (analyze-module (plist-get command-info :module)))
    (lns-get-complete-list callback async owner-file analyze-module out-bufname)))

(provide 'lns-completion)
