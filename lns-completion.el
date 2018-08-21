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
	  info)
      (setq info (json-read-from-string (buffer-string)))
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
  (lns-json-val info :type))
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
  (let ((proj-dir (lns-get-proj-dir))
	command-list process)
    (with-temp-buffer
      (when input-txt
	(insert input-txt))      
      (setq default-directory proj-dir)
      (setq command-list (apply 'lns-command-get-command args ))
      (if async-callback
	  (progn
	    (setq process (apply 'start-process "lns"
				 out-buffer command-list))
	    (set-process-sentinel process async-callback)
	    (when input-txt
	      (process-send-string process (concat input-txt "\n"))))
	(apply 'call-process-region (point-min) (point-max)
	       (car command-list) nil out-buffer nil
	       (cdr command-list))
	(with-current-buffer out-buffer
	  (buffer-string))))
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



(defun lns-get-complete-list ( callback async &optional file-path analyze-module)
  (let ((out-buf (lns-get-buffer "*lns-process*" t))
	result process)
    (when (not file-path)
      (setq file-path buffer-file-name))
    (setq process
	  (lns-execute-command
	   (if async
	       (lexical-let ((lex-buf out-buf)
			     (lex-callback callback))
		 (lambda (process event)
		   (condition-case err
		       (let (json)
			 (setq json (lns-json-get lex-buf :candidateList))
			 (funcall lex-callback json))
		       (error nil))
		     ))
	     nil)
	   out-buf
	   (concat (buffer-substring-no-properties (point-min) (point)) "lune")
	   (lns-convert-path-2-proj-relative-path file-path) "comp" 
	   (or analyze-module (lns-convert-path-2-module file-path))
	   (number-to-string (lns-get-line))
	   (number-to-string (lns-get-column)) "-i"))
    (when (not async)
      (funcall callback
	       (lns-json-get out-buf :candidateList)))
    process
    ))

(defun lns-completion-get-candidate-list (callback &optional async)
  (condition-case err
      (let* ((command-info (lns-command-get-info))
	     (owner-file (plist-get command-info :owner))
	     (analyze-module (plist-get command-info :module)))
	(lns-get-complete-list callback async owner-file analyze-module))
    (error nil)))

(provide 'lns-completion)
