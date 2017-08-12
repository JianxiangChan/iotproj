objects = main.o gpio.o ds18b20.o JSON_pkg.o
VPATH = ./hardware:./lib:./packets:./config

wifi_iot: $(objects)
	gcc -o wifi_iot $(objects)
main.o: gpio.h ds18b20.h main.c
gpio.o: gpio.h gpio.c
ds18b20.o: ds18b20.c ds18b20.h
JSON_pkg.o: JSON_pkg.h JSON_pkg.c ds18b20.h gpio.h  config.h

.PHONY: clean
clean:
	-rm wifi_iot $(objects)