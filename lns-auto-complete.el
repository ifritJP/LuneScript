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
;;   (require 'lns-auto-complete)

(require 'lns-completion)
(require 'lns-mode)
(require 'auto-complete)

(defun lns-hook-for-ac ()
  (add-to-list 'ac-modes 'lns-mode)
  (add-to-list 'ac-sources 'ac-source-lns-field)
  (add-to-list 'ac-sources 'ac-source-lns-symbol)
  (delq 'ac-source-words-in-same-mode-buffers ac-sources)
  (delq 'ac-source-abbrev ac-sources)
  (delq 'ac-source-dictionary ac-sources)
  (set (make-local-variable 'ac-ignore-case) nil)
  (set (make-local-variable 'ac-auto-show-menu) 0.3)
  (auto-complete-mode))

(add-hook 'lns-mode-hook 'lns-hook-for-ac)

(defun lns-ac-prefix ()
  (when (and (not (lns-is-in-comment-string (point)))
	     ;;(re-search-backward "\\.\\(\\(?:[a-zA-Z0-9][_a-zA-Z0-9]*\\)?\\)\\=" nil t))
	     ;;(re-search-backward "\\.\\(\\(?:[a-zA-Z0-9]\\)?\\)\\=" nil t))
	     (re-search-backward "\\.\\(\\(?:[a-zA-Z0-9]\\)?\\)\\=" nil t))
	     ;;(re-search-backward "\\.\\(.\\)\\=" nil t))
    (match-beginning 1)))

(defun lns-ac-action-symbol ()
  (let ((end-pos (point-marker))
	(start-pos (car ac-last-completion))
	move-pos)
    (save-excursion
      (goto-char start-pos)
      (cond
       ((eq (char-before start-pos) ?.)
	;; フィールド補完
	(if (re-search-forward "get_\\(.*\\)():[^,]+" end-pos t)
	    (progn
	      (replace-match "$\\1")
	      (setq move-pos end-pos)
	      )
	  (goto-char start-pos)
	  (when (re-search-forward "([^)]\\|:" end-pos t)
	    (setq move-pos (1- (point))))))
       (t
	;; シンボル補完
	(if (re-search-forward "([^)]\\|:" end-pos t)
	    ;; 関数呼び出しの '(' の位置に移動
	    (setq move-pos (1- (point)))
	  (goto-char start-pos)
	  (when (re-search-forward "\\(.*\\):[^:]+" end-pos t)
	    (replace-match "\\1")
	    (setq move-pos end-pos)
	    )
	  )))
      (goto-char start-pos)
      (when (re-search-forward ":.*" end-pos t)
	(replace-match ""))
      )
    (when move-pos
      (goto-char move-pos))
    ))


(defvar lns-ac-process-state-field :idle)
(defvar lns-ac-process-state-symbol :idle)
(defvar lns-ac-candidate-list-field nil)
(defvar lns-ac-candidate-list-symbol nil)

(defun lns-ac-candidates-symbol ()
  (lns-ac-candidates :symbol)
  )

(defun lns-ac-candidates-field ()
  (lns-ac-candidates :field)
  )


(defun lns-ac-candidates (mode)
  (lexical-let ((let-list (if (eq mode :symbol)
			      'lns-ac-candidate-list-symbol
			    'lns-ac-candidate-list-field))
		(process-state (if (eq mode :symbol)
				   'lns-ac-process-state-symbol
				 'lns-ac-process-state-field)))
    (cond
     ((eq (symbol-value process-state) :done)
      (set process-state :idle)
      (symbol-value let-list))
     ((eq (symbol-value process-state) :processing)
      nil)
     ((eq (symbol-value process-state) :idle)
      (set process-state :processing)
      (lns-completion-get-candidate-list
       (lambda (candidate-list err)
	 (cond
	  (candidate-list
	   (set let-list
		(mapcar (lambda (candidate)
			  (let* ((info (lns-json-val candidate :candidate))
				 (item-txt (lns-candidate-get-displayTxt info)))
			    item-txt))
			candidate-list))
	   ;; field の内容を symbol にコピーする。
	   ;; 現状は field も symbol も同じ結果になるため。
	   (setq lns-ac-candidate-list-symbol (symbol-value let-list))
	   (set process-state :done)
	   (ac-stop)
	   (ac-start)
	   (ac-update)
	   )
	  (t
	   (set process-state :idle)))
	 )
       t (format "*%s-lns-process*" mode))
      nil)
     )))

(ac-define-source lns-field
  '((candidates . lns-ac-candidates-field)
    (prefix . lns-ac-prefix)
    (action . lns-ac-action-symbol)
    (requires . 0)
    (cache)
    ;;(symbol . "lns")
    ))

(ac-define-source lns-symbol
  '((candidates . lns-ac-candidates-symbol)
    ;;(prefix . lns-ac-prefix-symbol)
    (action . lns-ac-action-symbol)
    (requires . 0)
    (cache)
    ;;(symbol . "lns")
    ))


(provide 'lns-auto-complete)
