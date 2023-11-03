set datafile separator ","
set grid
# X 軸のデータを時間に設定
set xdata time
# 日付データの入力フォーマット
set timefmt "%Y-%m-%d"
# x 軸のフォーマット
set format x "%Y/%m"
set mxtics 5
set nomxtics

set ylabel "ビルド時間(sec)"
set y2label "ソースコード行数"
set y2tics 0, 5000, 80000

# 凡例位置
#set key right top
set key below

# # サイズ比
# set size ratio 0.5

# # svg 保存
# set term svg
# set output 'history.svg'

plot "history.csv.org" using 1:4 with linespoints linewidth 2 axes x1y1 title "実時間"
replot "history.csv.org" using 1:5 with linespoints linewidth 2 axes x1y1 title "ユーザ時間"
replot "history.csv.org" using 1:3 with linespoints linewidth 2 axes x1y2 title "行数"
replot "history.csv.org" using 1:($3/$4) with linespoints linewidth 2 axes x1y2 title "1秒間に処理した行数"

# 保存する場合は pause の意味がない
pause -1
