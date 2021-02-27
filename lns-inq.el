;;
;; MIT License
;;
;; Copyright (c) 2020 ifritJP
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

(require 'lns-completion)

(defun lns-get-inq (&optional file-path analyze-module buf-name)
  (let ((out-buf (lns-get-buffer (or buf-name "*lns-inq-process*") t))
	process txt output-txt)
    (when (not file-path)
      (setq file-path buffer-file-name))
    (setq txt (buffer-substring-no-properties (point-min) (point-max)) )
    (setq process
	  (lns-execute-command
	   nil
	   out-buf
	   txt
	   (lns-convert-path-2-proj-relative-path file-path) "inq" 
	   (or analyze-module (lns-convert-path-2-module file-path))
	   (number-to-string (lns-get-line))
	   (number-to-string (lns-get-column)) "-i"))
    (with-current-buffer out-buf
      (setq output-txt (buffer-string)))
    (condition-case err
	(lns-json-get out-buf :inquire)
      (error
       nil))
    ))

(defun lns-show-inq ()
  (interactive)
  (let ((info (lns-get-inq)))
    (save-selected-window
      (delete-other-windows)
      (split-window-vertically)
      (other-window 1)
      (switch-to-buffer (lns-get-buffer "*lns-inq*" t))
      (if (not info)
	  (insert "not support")
	(insert (format "name: %s\n" (lns-json-val info :name)))
	(insert (format "type: %s\n" (lns-json-val info :type)))
	(insert (format "typeKind: %s\n" (lns-json-val info :typeKind)))
	(insert (format "static: %s\n" (lns-json-val info :static)))
	(insert (format "display: %s" (lns-json-val info :display))))
      (beginning-of-buffer)
      (fit-window-to-buffer)
      )
  ))

(provide 'lns-inq)
