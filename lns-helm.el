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
;;   (require 'lns-helm)

(defvar lns-anything nil)

(defvar lns-max-size-search-subfile 10000)

(require 'lns-completion)

(when (not lns-anything)
  (require 'helm))


(defun lns-helm-wrap (src keymap preselect)
  (if lns-anything
      (let ((anything-candidate-number-limit 9999))
	(anything :sources '(src) :keymap keymap :preselect preselect))
    (let ((helm-candidate-number-limit 9999))
      (helm :sources src :keymap keymap :preselect preselect)))
  )



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
  (let ((candidate-list (lns-completion-get-candidate-list))
	helm-candidate-list
	helm-params)
    (setq helm-candidate-list
	  (mapcar (lambda (candidate)
		    (let* ((info (lns-json-val candidate :candidate))
			   (item-type (lns-candidate-get-type info))
			   (item-txt (lns-candidate-get-displayTxt info))
			   (item-name item-txt)
			   (item-name-append "")
			   face)
		      (setq face (cond ((equal item-type "Typ")
					'font-lock-type-face)
				       ((or (equal item-type "Mbr")
					    (equal item-type "Var"))
					'font-lock-variable-name-face)
				       ((or (equal item-type "Mtd")
					    (equal item-type "Fun"))
					'font-lock-function-name-face)))
		      (when (string-match "\\([^ \t(]+\\)" item-txt)
			(setq item-name (substring item-txt
						   (match-beginning 1)
						   (match-end 1)))
			(setq item-name-append (substring item-txt
							  (match-end 1))))
		      (cons (format "(%s) %s%s"
				    (propertize
				     item-type 'face face)
				    (propertize
				     item-name 'face face)
				    item-name-append)
			    info)))
		  candidate-list))
    (setq helm-params
	  `((name . ,(format "comp-at:%s:%d:%d"
			     (file-name-nondirectory buffer-file-name)
			     (lns-get-line) (lns-get-column)))
	    (candidates . ,helm-candidate-list)
	    (action . lns-helm-select)))
    (lns-helm-wrap helm-params nil nil)
    ))



(add-hook 'lns-mode-hook
      '(lambda ()
         (local-set-key (kbd "C-c C-/") 'lns-helm-complete-at)))


(provide 'lns-helm)
