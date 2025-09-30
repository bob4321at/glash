export GTK_IM_MODULE=fcitx  # or ibus
export XMODIFIERS=@im=fcitx
export QT_IM_MODULE=fcitx
export GDK_BACKEND=x11
env -u WAYLAND_DISPLAY ~/golang/flash/main
