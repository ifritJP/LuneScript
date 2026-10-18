/* lnsrt.c と lnslua.c の共通の定義 */

#ifndef LNSRT_H
#define LNSRT_H

#define EXPORT(name) __attribute__((export_name(#name))) name

/*
  main() の初期化の最後に呼ぶ、追加の初期化 (Lua の初期化など)。
  lnsrt.c には何もしない weak の定義があり、lnslua.c で置き換える。
*/
void lnsr_initExt( void );

#endif
