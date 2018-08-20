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
  (add-to-list 'ac-sources 'ac-source-lns)
  (set (make-local-variable 'ac-ignore-case) nil)
  (set (make-local-variable 'ac-auto-show-menu) 0.3)
  (auto-complete-mode))

(add-hook 'lns-mode-hook 'lns-hook-for-ac)

(defun lns-ac-prefix ()
  (when (and (not (lns-is-in-comment-string (point)))
	     (re-search-backward "\\.\\(\\(?:[a-zA-Z0-9][_a-zA-Z0-9]*\\)?\\)\\=" nil t))
    (match-beginning 1)))

(defun lns-ac-action ()
  (let ((end-pos (point-marker))
	(start-pos (car ac-last-completion))
	move-pos)
    (save-excursion
      (goto-char start-pos)
      (if (re-search-forward "get_\\(.*\\)():[^,]+" end-pos t)
	  (progn
	    (replace-match "$\\1")
	    (setq move-pos end-pos)
	    )
	(goto-char start-pos)
	(when (re-search-forward "([^)]\\|:" end-pos t)
	  (setq move-pos (1- (point))))))
    (when move-pos
      (goto-char move-pos))
    ))

(defun lns-ac-candidates ()
  (let ((candidate-list (lns-completion-get-candidate-list)))
    (mapcar (lambda (candidate)
	      (let* ((info (lns-json-val candidate :candidate))
		     (item-txt (lns-candidate-get-displayTxt info)))
		item-txt))
	    candidate-list)))

(ac-define-source lns
  '((candidates . lns-ac-candidates)
    (prefix . lns-ac-prefix)
    (action . lns-ac-action)
    (requires . 0)
    (cache)
    ;;(symbol . "lns")
    ))

(provide 'lns-auto-complete)
