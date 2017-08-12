#include "gpio.h"

#include <sys/stat.h>
#include <sys/types.h>
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>



#define LED_HIGH 1
#define LED_LOW 0

const char mode_input[] = "in";
const char mode_output[] = "out";

const char gpio_export[] = "/sys/class/gpio/export";
const char gpio_direction[] = "/sys/class/gpio/gpio%d/direction";
const char gpio_value[] = "/sys/class/gpio/gpio%d/value";
const char gpio_unexport[] =  "/sys/class/gpio/unexport";


static GPIODriverRet gpio_writei(GPIODriver* thiz, const char* path, size_t value);
static GPIODriverRet gpio_writes(GPIODriver* thiz, char* path, const char* value);
static size_t gpio_read(GPIODriver* thiz, const char* path);
static GPIODriverRet led_ctrl(GPIODriver* thiz,size_t led_state);

struct _GPIODriver
{
	size_t pin;
	const char *pin_mode;
	size_t pin_value;
};

GPIODriver* gpio_creat(size_t ledpin, const char* mode, int defvalue)
{
	GPIODriver* dr = malloc(sizeof(GPIODriver));
	if(dr != NULL)
	{
		dr->pin = ledpin;
		dr->pin_mode = mode;
		dr->pin_value = defvalue;
		return dr;
	}
	else
	{
		return NULL;
	}
}

GPIODriverRet gpio_destroy(GPIODriver* thiz)
{
	if(thiz == NULL)
	{
		fprintf(stderr, "invailed input\n");
		return GPIODriver_RET_PARAMS;
	}
	thiz->pin = 0;
	thiz->pin_mode = NULL;
	thiz->pin_value = 0;
	free(thiz);
	thiz  = NULL;
	return GPIODriver_RET_OK;
	
}

GPIODriverRet gpio_init(GPIODriver* thiz)
{
	char* str;
	size_t len;
	
	if(thiz == NULL)
	{
		fprintf(stderr, "invailed params\n");
		return GPIODriver_RET_PARAMS;
	}
	
	//export
	if(gpio_writei(thiz, "/sys/class/gpio/export", thiz->pin) == GPIODriver_RET_FAIL)
	{
		fprintf(stderr, "gpio_init failed\n");
		return GPIODriver_RET_FAIL;
	}
	
	//Direction
	//要获取字符串的长度 需要使用strlen而不是sizeof
	len = strlen(gpio_direction);
	//fprintf(stderr, "len = %d\n", len);
	str = (char*)malloc(len);
	if(str == NULL)
		
	{
		fprintf(stderr, "malloc failed\n");
		return GPIODriver_RET_FAIL;
	}
	
	if((len = sprintf(str,gpio_direction,thiz->pin)) == 0)
	{
		fprintf(stderr, "sprintf failed\n");
		free(str);
		str = NULL;
		return GPIODriver_RET_FAIL;
	}
	//fprintf(stderr, "len = %d\n", len);
	if(gpio_writes(thiz, str, thiz->pin_mode) != GPIODriver_RET_OK)
	{
		fprintf(stderr, "gpio_init failed\n");
		free(str);
		str = NULL;	
		return GPIODriver_RET_FAIL;
	}

	free(str);
	str = NULL;	
	
	return GPIODriver_RET_OK;
}

GPIODriverRet led_close(GPIODriver* thiz)
{
	if(led_ctrl(thiz, LED_LOW) != GPIODriver_RET_OK)
	{
		fprintf(stderr, "led_close error\n");
		return GPIODriver_RET_FAIL;
	}
	//fprintf(stderr, "close\n");
	return GPIODriver_RET_OK;
}

GPIODriverRet led_open(GPIODriver* thiz)
{
	if(led_ctrl(thiz, LED_HIGH) != GPIODriver_RET_OK)
	{
		fprintf(stderr, "led_open error\n");
		return GPIODriver_RET_FAIL;
	}
	//fprintf(stderr, "open\n");
	return GPIODriver_RET_OK;
}

#define MAX_SIEZ 100
static GPIODriverRet gpio_writei(GPIODriver* thiz, const char* path, size_t value)
{
	size_t fd;
	size_t len;
	char* str;
	//fprintf(stderr, "%s\n", path);
	fd = open(path, O_RDWR);
	if(fd < 0)
	{
		fprintf(stderr, "failed to open file in path: %s\n", path);
		return GPIODriver_RET_FAIL;
	}
	
	str = (char *)malloc(MAX_SIEZ);
	if(str == NULL)
	{
		close(fd);
		fprintf(stderr, "malloc failed\n");
		return GPIODriver_RET_FAIL;
	}
	
	len = snprintf(str, 10, "%d", value);
	//fprintf(stderr, "len1: %d\n", len);
	//fprintf(stderr, "%s\n", str);
	
	if(len == 0)
	{
		free(str);
		str = NULL;
		close(fd);
		fprintf(stderr, "sprintf failed\n");
		return GPIODriver_RET_OOM;
	}
	
	//fprintf(stderr, "str: %s, len: %d\n", str, len);
	
	// if(write(fd, str, len) < 0)
	//{
/* 		fprintf(stderr, "failed to write file in path: %s\n", path);
		free(str);
		str = NULL;
		close(fd);
		return GPIODriver_RET_FAIL; */

	//} 
	
	len = write(fd, str, len);
	//fprintf(stderr, "len2: %d\n", len);
	
	free(str);
	str = NULL;
	close(fd);
	return GPIODriver_RET_OK;
}

static GPIODriverRet gpio_writes(GPIODriver* thiz, char* path, const char* value)
{
	size_t fd;
	size_t len;
	
	fd = open(path, O_RDWR);
	if(fd < 0)
	{
		fprintf(stderr, "failed to open file in path: %s\n", path);
		return GPIODriver_RET_FAIL;
	}
	
	len = strlen(value);
	//fprintf(stderr, "value: %s\n", value);
	if(write(fd, value, len) < 0)
	{
		fprintf(stderr, "failed to write file in path: %s\n", path);
		close(fd);
		return GPIODriver_RET_FAIL;
	}
	close(fd);
	
	return GPIODriver_RET_OK;
}

static size_t gpio_read(GPIODriver* thiz, const char* path)
{
	size_t flen;
	char* p;
	FILE* fp;
	size_t value;
	
	fp = fopen(path, "r");
	if(fp == NULL)
	{
		fprintf(stderr, "failed to open file in path: %s\n", path);
		return -1;
	}
	
	fseek(fp, 0L, SEEK_END);
	
	flen = ftell(fp);
	
	p = (char *)malloc(flen);
	if(p == NULL)
	{
		fprintf(stderr, "malloc failed\n");
		fclose(fp);
		return -1;
	}
	
	fseek(fp, 0L, SEEK_SET);
	
	if(!(fread(p, flen, 1, fp)))
	{
		fprintf(stderr, "failed to read file in path: %s\n", path);
		fclose(fp);
		return -1;
	}
	
	value = atoi(p);
	free(p);
	p = NULL;
	
	fclose(fp);
	return value;
	
}

static GPIODriverRet led_ctrl(GPIODriver* thiz,size_t led_state)
{
	char* str;
	size_t len;
	
	if(thiz == NULL || !(led_state == 0 || led_state == 1))
	{
		fprintf(stderr, "led_ctrl error\n");
		return GPIODriver_RET_PARAMS;
	}
	
	//len = strlen(gpio_value);
	//fprintf(stderr, "len: %d\n", len);
	str = (char*)malloc(MAX_SIEZ);
	if(str == NULL)
	{
		fprintf(stderr, "malloc failed\n");
		return GPIODriver_RET_FAIL;
	}
	
	if((len = sprintf(str,gpio_value,thiz->pin)) == 0)
	{
		fprintf(stderr, "sprintf failed\n");
		free(str);
		str = NULL;
		return GPIODriver_RET_FAIL;
	}
	
 	if(gpio_writei(thiz, str, led_state) != GPIODriver_RET_OK)
	{
		fprintf(stderr, "led_open failed\n");
		free(str);
		str = NULL;
		return GPIODriver_RET_FAIL;
	}
	 
	thiz->pin_value = led_state;
	
	free(str);
	str = NULL;
	
	return GPIODriver_RET_OK;
	
}
