#include <stdlib.h>
#include <stdio.h>

#include "JSON_pkg.h"
#include "../lib/json/cJSON.h"
#include "../hardware/ds18b20.h"
#include "../hardware/gpio.h"
#include "../config/config.h"



//上行数据
typedef  struct _TempPacket
{
	char* mac;
	float temperature;
	cJSON_bool led;
}TempPacket;

//下行数据
typedef  struct _LedPacket
{
	char* mac;
	cJSON_bool led;
}LedPacket;


struct _JSONPkg
{
	TempPacket* temp_pkg;
	LedPacket* led_pkg;
	cJSON* JSON_temp_pkg;
	cJSON* JSON_led_pkg;
};

static TempPacket* JSON_temp_pkg_creat();
static LedPacket* JSON_led_pkg_creat();
static JSONPkgRet JSON_temp_pkg_destory(TempPacket* tkpg);
static JSONPkgRet JSON_led_pkg_destory(LedPacket* lkpg);

static TempPacket* JSON_temp_pkg_creat()
{
	TempPacket* temp_pkg = (TempPacket*)malloc(sizeof(TempPacket));
	if(temp_pkg == NULL)
	{
		fprintf(stderr,"Func: %s, Line: %d, malloc error\n",__func__, __LINE__);
		return NULL;
	}
	temp_pkg->mac = mac;
	temp_pkg->temperature = 0.0;
	temp_pkg->led = false;
	return temp_pkg;
}

static JSONPkgRet JSON_temp_pkg_destory(TempPacket* tkpg)
{
	if(tkpg = NULL)
	{
		fprintf(stderr,"Func: %s, Line: %d, invalid params\n",__func__, __LINE__);
		return JSONPkgRet_RET_PARAMS;
	}
	free(tkpg);
	tkpg = NULL;
	return JSONPkgRet_RET_OK;
}

static LedPacket* JSON_led_pkg_creat()
{
	LedPacket* led_pkg = (LedPacket*)malloc(sizeof(LedPacket));
	if(led_pkg == NULL)
	{
		fprintf(stderr,"Func: %s, Line: %d, malloc error\n",__func__, __LINE__);
		return NULL;
	}
	led_pkg->mac = mac;
	led_pkg->led = false;
	return led_pkg;
}

static JSONPkgRet JSON_led_pkg_destory(LedPacket* lkpg)
{
	if(lkpg = NULL)
	{
		fprintf(stderr,"Func: %s, Line: %d, invalid params\n",__func__, __LINE__);
		return JSONPkgRet_RET_PARAMS;
	}
	free(lkpg);
	lkpg = NULL;
	return JSONPkgRet_RET_OK;
}

JSONPkg* JSON_pkg_creat()
{
	JSONPkg* json_pkg = (JSONPkg*)malloc(sizeof(JSONPkg));
	if(json_pkg == NULL)
	{
		fprintf(stderr,"Func: %s, Line: %d, malloc error\n",__func__, __LINE__);
		return NULL;
	}
	json_pkg->JSON_temp_pkg = cJSON_CreteObject();
	json_pkg->temp_pkg = JSON_temp_pkg_creat();
	if(json_pkg->temp_pkg == NULL)
	{
		fprintf(stderr,"Func: %s, Line: %d, JSON_temp_pkg_creat error\n",__func__, __LINE__);
		JSON_pkg_destory(json_pkg);
		return NULL;
	}
	
	json_pkg->led_pkg = JSON_led_pkg_creat();
	if(json_pkg->led_pkg == NULL)
	{
		fprintf(stderr,"Func: %s, Line: %d, JSON_temp_pkg_creat error\n",__func__, __LINE__);
		JSON_pkg_destory(json_pkg);
		return NULL;
	}
	
	return json_pkg;
}

JSONPkgRet JSON_pkg_destory(JSONPkg* thiz)
{
	if(thiz == NULL)
	{
		fprintf(stderr,"Func: %s, Line: %d, invalid params\n",__func__, __LINE__);
		return JSONPkgRet_RET_PARAMS;
	}
	JSON_led_pkg_destory(thiz->led_pkg);
	JSON_temp_pkg_destory(thiz->temp_pkg);
	free(thiz);
	thiz = NULL;
	return JSONPkgRet_RET_OK;
}

