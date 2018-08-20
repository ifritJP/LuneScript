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

(require 'flymake)
;;c++のflymakeでmakefileを不要にする
(defun flymake-lns-init ()
  (let* ((command-info (lns-command-get-info))
	 (command-list (lns-command-get-command
			(lns-convert-path-2-proj-relative-path
			 (plist-get command-info :owner)) "diag" "--nodebug")))
    (list (car command-list) (cdr command-list)
	  (plist-get command-info :dir))))

(defun lns-flymake-get-real-file-name (file-name)
  ""
  (expand-file-name file-name (lns-get-proj-dir))
  )

(add-to-list 'flymake-allowed-file-name-masks
	     '("\\.lns$" flymake-lns-init nil lns-flymake-get-real-file-name) )

(add-hook 'lns-mode-hook '(lambda ()
			    (flymake-mode t)
			    ))

(add-to-list
 'flymake-err-line-patterns
 '( "^error:.*\\([^/]+\\.lns\\):\\([0-9]+\\):\\([0-9]+\\):\\(.+\\)" 1 2 3 4))



(provide 'lns-flymake)
