all:
	@echo make exec
	@echo make plot

exec:
	../../src/lnsc -shebang main.lns > schedule.csv

plot:
	gnuplot plot.txt


raw_prof:
# SKIP_PROF=y を指定すると、プログラムの最後でプロファイル用の処理が入り、
# 不能な処理まで解析結果に入ってしまうため、
# 一度 SKIP_PROF=y を指定せずに実行して、
# SKIP_PROF=y を指定して実行する。
	make -C ../../src goprof PROF_SCHE=5 2>&1
	make -C ../../src goprof PROF_SCHE=5 SKIP_PROF=y 2>&1
