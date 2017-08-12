#ifndef __GPIO_H
#define __GPIO_H

#include <stdlib.h>

#define LED_PIN 26

extern const char mode_input[2];
extern const char mode_output[3];

typedef enum _GPIODriverRet
{
	GPIODriver_RET_OK,
	GPIODriver_RET_OOM,
	GPIODriver_RET_PARAMS,
	GPIODriver_RET_FAIL,
}GPIODriverRet;

struct _GPIODriver;
typedef struct _GPIODriver GPIODriver;

GPIODriver* gpio_creat(size_t ledpin, const char* mode, int defvalue);
GPIODriverRet gpio_destroy(GPIODriver* thiz);
GPIODriverRet gpio_init(GPIODriver* thiz);
GPIODriverRet led_open(GPIODriver* thiz);
GPIODriverRet led_open(GPIODriver* thiz);

#endif
