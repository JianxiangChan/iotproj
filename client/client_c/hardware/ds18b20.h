#ifndef __DS18B20_H
#define __DS18B20_H


struct _Ds18b20Driver;
typedef struct _Ds18b20Driver Ds18b20Driver;

typedef enum _Ds18b20DriverRet
{
	Ds18b20Driver_RET_OK,
	Ds18b20Driver_RET_OOM,
	Ds18b20Driver_RET_PARAMS,
	Ds18b20Driver_RET_FAIL,
}Ds18b20DriverRet;

Ds18b20Driver* ds18b20_creat();
Ds18b20DriverRet ds18b20_destory(Ds18b20Driver* thiz);
Ds18b20DriverRet ds18b20_init(Ds18b20Driver* thiz);
Ds18b20DriverRet ds18b20_read(Ds18b20Driver* thiz);

#endif