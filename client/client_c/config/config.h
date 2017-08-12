#ifndef __CONFIG_H
#define __CONFIG_H

//mqtt config
const char mqtt_server[] = "tcp://121.42.35.23:1883";
const char mqtt_usrname[] = "monitor";
const char mqtt_password[] = "monitor";
const char mqtt_subtopic[] = "node/b827ebd212b7/downlink";
const char mqtt_pubtopic[] = "node/b827ebd212b7/uplink";

//mac
char mac[] = "b827ebd212b7";

#endif