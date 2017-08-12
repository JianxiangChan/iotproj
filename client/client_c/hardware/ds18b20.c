#include "ds18b20.h"
#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>
#include <dirent.h>
#include <string.h>
#include <time.h>
#include <fcntl.h>

char path_head[50] = "/sys/bus/w1/devices/";
const char path_mid[] = "28-";
const char path_tail[] = "/w1_slave";

struct _Ds18b20Driver
{
	float temp;
	char * path;
};

Ds18b20Driver* ds18b20_creat()
{
	Ds18b20Driver* thiz = malloc(sizeof(Ds18b20Driver*));
	if(thiz == NULL)
	{
		fprintf(stderr,"Func: %s, Line: %d, malloc error\n",__func__, __LINE__);
		return NULL;
	}
	thiz->temp = 0;
	thiz->path = NULL;
	
}

Ds18b20DriverRet ds18b20_destory(Ds18b20Driver* thiz)
{
	if(thiz == NULL)
	{
		fprintf(stderr,"Func: %s, Line: %d, invailded params\n",__func__, __LINE__);
		return Ds18b20Driver_RET_PARAMS;
	}
	thiz->temp = 0;
	thiz->path = NULL;
	free(thiz);
	thiz = NULL;
	return Ds18b20Driver_RET_OK;
}

Ds18b20DriverRet ds18b20_init(Ds18b20Driver* thiz)
{
	DIR *dirp;
	char rom[20];
	struct dirent *direntp;
	size_t read_g = 0;
	
		
	if(thiz == NULL)
	{
		fprintf(stderr,"Func: %s, Line: %d, invailded params\n",__func__, __LINE__);
		return Ds18b20Driver_RET_PARAMS;
	}

	if(system("sudo modprobe w1-gpio") != 0 || system("sudo modprobe w1-therm") != 0)
	{
		fprintf(stderr,"Func: %s, Line: %d, system error\n",__func__, __LINE__);
		return Ds18b20Driver_RET_FAIL;
	}
	

	if((dirp = opendir(path_head)) == NULL)
	{
		fprintf(stderr,"Func: %s, Line: %d, open dir error\n",__func__, __LINE__);
		return Ds18b20Driver_RET_FAIL;
	}
	
	//遍历目录

	while((direntp = readdir(dirp)) != NULL)
	{
		if(strstr(direntp->d_name, path_mid))
		{
			strcpy(rom, direntp->d_name);
			printf("rom: %s\n",rom);
			read_g = 1;
		}
	}

	if(read_g == 0)
	{
		fprintf(stderr,"Func: %s, Line: %d, read dir error\n",__func__, __LINE__);
		closedir(dirp);
		return Ds18b20Driver_RET_FAIL;
	}

	closedir(dirp);
	
	strcat(path_head,rom);
	strcat(path_head,path_tail);

	fprintf(stderr,"Func: %s, Line: %d, path：%s\n",__func__, __LINE__, path_head); //test
	
	thiz->path = path_head;
	
	return Ds18b20Driver_RET_OK;
}

Ds18b20DriverRet ds18b20_read(Ds18b20Driver* thiz)
{
	size_t fd;
	char buf[100];
	char* temp;
	if(thiz == NULL)
	{
		fprintf(stderr,"Func: %s, Line: %d, invailded params\n",__func__, __LINE__);
		return Ds18b20Driver_RET_PARAMS;
	}
	
	if((fd = open(thiz->path, O_RDONLY)) < 0)
	{
		fprintf(stderr,"Func: %s, Line: %d, path: %s open failed\n",__func__, __LINE__, thiz->path);
		return Ds18b20Driver_RET_PARAMS;
	}
	
	if(read(fd, buf, sizeof(buf)) < 0)
	{
		fprintf(stderr,"Func: %s, Line: %d, read failed\n",__func__, __LINE__);
		return Ds18b20Driver_RET_FAIL;
	}
	
	temp = strchr(buf,'t');
	sscanf(temp,"t=%s",temp);
	thiz->temp = atof(temp)/1000;
	printf("temp : %3.3f °C\n", thiz->temp);
	
	sleep(1);
}