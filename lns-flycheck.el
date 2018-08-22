;;; lns-mode.el --- a major-mode for editing LuneScript
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

(require 'lns-command)

(require 'flycheck)

(flycheck-define-checker lunescript
  "A LuneScript syntax checker.

See URL `https://github.com/ifritJP/LuneScript'."
  :command ("lua5.3"
	    (eval (nth 1 (lns-command-get-command)))
	    (eval (nth 2 (lns-command-get-command)))
	    (eval (nth 3 (lns-command-get-command)))
	    (eval (lns-convert-path-2-proj-relative-path
		   (plist-get (lns-command-get-info) :owner)))
	    "diag" "--nodebug")
  :error-patterns
  ((error line-start (file-name) ":" line ":" column ":" (message) line-end))
  :modes lns-mode
  :working-directory (lambda (checker)
		       (plist-get (lns-command-get-info) :dir))
  )

(defun lns-flycheck-after-save-hook-func ()
  (flycheck-mode)
  (remove-hook 'after-save-hook 'lns-flycheck-after-save-hook-func)
  )

(defun lns-flycheck-lns-mode-hook-func ()
  (flycheck-select-checker 'lunescript)
  (if (file-exists-p (buffer-file-name))
      (flycheck-mode)
    (add-hook 'after-save-hook 'lns-flycheck-after-save-hook-func)
    ))
(add-hook 'lns-mode-hook 'lns-flycheck-lns-mode-hook-func)

(provide 'lns-flycheck)
