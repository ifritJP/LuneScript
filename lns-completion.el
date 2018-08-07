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

(defvar lns-anything nil)

(defvar lns-max-size-search-subfile 10000)

(when (not lns-anything)
  (require 'helm))

(defun lns-json-get (buf symbol)
  (with-current-buffer buf
    (let ((json-object-type 'plist)
	  (json-array-type 'list)
	  info)
      (setq info (json-read-from-string (buffer-string)))
      (plist-get (plist-get info :lunescript) symbol))))

(defun lns-json-val (json symbol)
  (plist-get json symbol))


(defface lns-candidate-face
  '((t
     :foreground "gold"))
  "candidate face")
(defvar lns-candidate-face 'lns-candidate-face)

(defface lns-candidate-path-face
  '((t
     :foreground "green"))
  "candidate path face")
(defvar lns-candidate-path-face 'lns-candidate-path-face)

(defface lns-expandable-candidate-face
  '((t
     :foreground "dark orange"))
  "expandable candidate face")
(defvar lns-expandable-candidate-face 'lns-expandable-candidate-face)

(defvar lns-lua-command "lua5.3"
  "lua command")

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

(defun lns-execute-command ( out-buffer input-txt &rest args )
  (let ((proj-dir lns-proj-dir))
    (with-temp-buffer
      (when input-txt
	(insert input-txt))
      (setq default-directory proj-dir)
      (apply 'call-process-region (point-min) (point-max)
	     lns-lua-command nil out-buffer nil
	     "-e" "require( 'lune.base.base' )" " " args ))
    (with-current-buffer out-buffer
      (buffer-string))))

(defun lns-convert-path-2-proj-relative-path (path)
  (file-relative-name (expand-file-name path) lns-proj-dir))

(defun lns-convert-path-2-module ( path )
  (let ((module (file-relative-name (expand-file-name path) lns-proj-dir)))
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
    (concat (expand-file-name (substring path 1) lns-proj-dir) ".lns")
    ))



(defun lns-get-complete-list ( &optional file-path analyze-module)
  (let ((out-buf (lns-get-buffer "*lns-process*" t))
	result )
    (when (not file-path)
      (setq file-path buffer-file-name))
    (lns-execute-command
     out-buf
     (concat (buffer-substring-no-properties (point-min) (point)) "lune")
     (lns-convert-path-2-proj-relative-path file-path) "comp" 
     (or analyze-module (lns-convert-path-2-module file-path))
     (number-to-string (lns-get-line))
     (number-to-string (lns-get-column)) "-i")
    (lns-json-get out-buf :candidateList)
    ))


(defun lns-helm-wrap (src keymap preselect)
  (if lns-anything
      (let ((anything-candidate-number-limit 9999))
	(anything :sources '(src) :keymap keymap :preselect preselect))
    (let ((helm-candidate-number-limit 9999))
      (helm :sources src :keymap keymap :preselect preselect)))
  )



(defun lns-candidate-get-type (info)
  (lns-json-val info :type))
(defun lns-candidate-get-displayTxt (info)
  (lns-json-val info :displayTxt))

(defun lns-helm-select (item)
  (let ((item-type (lns-candidate-get-type item))
	(item-txt (lns-candidate-get-displayTxt item)))
    (if (or (equal item-type "Fun")
	    (equal item-type "Mtd"))
	(progn
	  (if (and (equal item-type "Mtd")
		   (string-match "^get_\\(.*\\)():[^,]*" item-txt))
	      ;; accessor
	      (progn
		(setq item-txt (concat "$" (replace-match "\\1" t nil item-txt)))
		(insert item-txt))
	    (save-excursion
	      (insert item-txt))
	    (search-forward "(")))
      (save-excursion
	(insert item-txt))
      (search-forward ":")
      (backward-char)
      )))
    
    
(defun lns-helm-complete-at ()
  (interactive)
  (let (candidate-list
	helm-candidate-list
	helm-params owner-file analyze-module)
    (save-excursion
      (goto-char (point-min))
      (when (re-search-forward "^[ \t]*subfile[ \t]+owner[ \t]+"
			       lns-max-size-search-subfile t)
	(when (looking-at "\\([^;]+\\);")
	  (setq owner-file (buffer-substring-no-properties (match-beginning 1)
							   (match-end 1)))
	  (setq owner-file (lns-convert-module-2-path owner-file))
	  (setq analyze-module (lns-convert-path-2-module buffer-file-name))
	  )))
    (setq candidate-list (lns-get-complete-list owner-file analyze-module))
    (setq helm-candidate-list
	  (mapcar (lambda (candidate)
		    (let* ((info (lns-json-val candidate :candidate))
			   (item-txt (lns-candidate-get-displayTxt info)))
		      (cons item-txt info)))
		  candidate-list))
    (setq helm-params
	  `((name . ,(format "comp-at:%s:%d:%d"
			     (file-name-nondirectory buffer-file-name)
			     (lns-get-line) (lns-get-column)))
	    (candidates . ,helm-candidate-list)
	    (action . lns-helm-select)))
    (lns-helm-wrap helm-params nil nil)
    ))


    


(provide 'lns-completion)
