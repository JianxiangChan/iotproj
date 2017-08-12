#ifndef __JSON_PKG_H
#define __JSON_PKG_H

typedef enum _JSONPkgRet
{
	JSONPkgRet_RET_OK,
	JSONPkgRet_RET_OOM,
	JSONPkgRet_RET_PARAMS,
	JSONPkgRet_RET_FAIL,
}JSONPkgRet;


struct _JSONPkg;
typedef struct _JSONPkg JSONPkg;

JSONPkg* JSON_pkg_creat();
JSONPkgRet JSON_pkg_destory(JSONPkg* thiz);
#endif