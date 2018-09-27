(require 'lns-command)

(defvar lns-dispatch-prev-window-conf nil)

(defvar lns-dispatch-buf-name "lns dispatcher")

(defvar lns-dispatch-menu-info
  '((:name "compile" :bind "c"
	   :submenu
	   ((:name "lua" :bind "l" :action lns-command-compile)
	    (:name "&meta" :bind "m" :action lns-command-compile-meta)
	    (:name "glue" :bind "g" :action lns-command-glue)
	    ))
    ))


(defvar lns-dispatch-menu-info-current nil)


(defun lns-dispatch-menu-get-name (info)
  (plist-get info :name))

(defun lns-dispatch-menu-get-bind (info)
  (plist-get info :bind))

(defun lns-dispatch-menu-get-submenu (info)
  (plist-get info :submenu))

(defun lns-dispatch-menu-get-action (info)
  (plist-get info :action))

(defun lns-dispatch-menu-get-keymap (info)
  (plist-get info :keymap))

(defun lns-dispatch-menu-set-keymap (info keymap)
  (plist-put info :keymap keymap))

(defun lns-dispatch-mode-exit (&optional action param)
  (set-window-configuration lns-dispatch-prev-window-conf)
  (kill-buffer (get-buffer lns-dispatch-buf-name))
  (when action
    (if param
	(funcall action param)
      (funcall action)))
  )

(defun lns-dispatch-build-keymap (menu-info)
  (let ((map (make-sparse-keymap)))
    (define-key map (kbd "C-g")
      (lambda ()
	(interactive)
	(lns-dispatch-mode-exit nil)
	))
    (mapcar
     (lambda (menu)
       (define-key map (eval `(kbd ,(lns-dispatch-menu-get-bind menu)))
	 (if (lns-dispatch-menu-get-submenu menu)
	     (eval `(lambda ()
		      (interactive)
		      (lns-dispatch-redraw ',(lns-dispatch-menu-get-submenu menu))))
	   (eval `(lambda (param)
		    (interactive "P")
		    (lns-dispatch-mode-exit
		     ',(lns-dispatch-menu-get-action menu) param))))))
       menu-info)
    map
    ))


(defun lns-dispatch-redraw (menu-info)
  (font-lock-mode nil)
  (let ((buffer-read-only nil)
        (old-point (point))
        (is-first (zerop (buffer-size)))
        (actions-p nil))
    (erase-buffer)
    (use-local-map (lns-dispatch-build-keymap menu-info))

    (lns-dispatch-draw-menu lns-dispatch-menu-info menu-info)
    
    (delete-trailing-whitespace)
    (setq mode-name "lns-dispatcher" major-mode 'lns-dispatcher)
    )
  (setq buffer-read-only t)
  (fit-window-to-buffer)
  (beginning-of-buffer))


(defun lns-dispatch-mode ()
  ""
  (interactive)
  (setq lns-dispatch-prev-window-conf
        (current-window-configuration))
  (let ((buf (get-buffer-create lns-dispatch-buf-name)))
    (delete-other-windows)
    (split-window-vertically)
    (other-window 1)
    (switch-to-buffer buf)
    (setq lns-dispatch-menu-info-current lns-dispatch-menu-info)
    (lns-dispatch-redraw lns-dispatch-menu-info))
  )

(defface lns-dispatch-valid-bind-face
  '((t
     :foreground "red"))
  "valid bind face")
(defvar lns-dispatch-valid-bind-face 'lns-dispatch-valid-bind-face)

(defface lns-dispatch-invalid-bind-face
  '((t
     :foreground "gray30"))
  "valid bind face")
(defvar lns-dispatch-invalid-bind-face 'lns-dispatch-invalid-bind-face)


(defun lns-dispatch-draw-menu (menu-info current-info &optional prefix )
  (mapcar (lambda (menu)
	    (when (> (+ (current-column)
			(length prefix)
			(length (lns-dispatch-menu-get-name menu)) 5)
		     (window-width))
	      (insert "\n"))
	    (insert (format "%s%s: %s   "
			    (or prefix "")
			    (if (eq menu-info current-info)
				(propertize 
				 (lns-dispatch-menu-get-bind menu)
				 'face 'lns-dispatch-valid-bind-face)
			      (propertize 
			       (lns-dispatch-menu-get-bind menu)
			       'face 'lns-dispatch-invalid-bind-face))
			    (lns-dispatch-menu-get-name menu)
			    ))
	    (when (lns-dispatch-menu-get-submenu menu)
	      (insert "\n")
	      (lns-dispatch-draw-menu
	       (lns-dispatch-menu-get-submenu menu)
	       current-info
	       (concat prefix "    "))
	      (insert "\n")
	      ))
	  menu-info))


(provide 'lns-dispatch)
