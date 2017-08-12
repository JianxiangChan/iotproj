#include "./hardware/gpio.h"
#include "./hardware/ds18b20.h"
#include <fcntl.h>
#include <stdio.h>

int main(int argc, char* argv[])
{
	GPIODriver* dr;
	Ds18b20Driver* dr_ds;
	
	dr = gpio_creat(LED_PIN, mode_output, 0);
	gpio_init(dr);
	

	dr_ds = ds18b20_creat();

	ds18b20_init(dr_ds);

	//ds18b20_destory(dr_ds);
	while(1)
	{
	 	led_open(dr);
		ds18b20_read(dr_ds);
		usleep(500 * 1000);
		led_close(dr);
		usleep(500 * 1000); 
	}
}