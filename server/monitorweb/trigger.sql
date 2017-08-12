# noinspection SqlNoDataSourceInspectionForFile

DROP TRIGGER
IF EXISTS t_after_insert_on_tab_all_sensor_data;

CREATE TRIGGER t_after_insert_on_tab_all_sensor_data AFTER INSERT ON tab_all_sensor_data FOR EACH ROW
BEGIN

SET @cnt = (SELECT count(*)	FROM tab_all_sensor_data_last WHERE mac = NEW.mac);
IF @cnt = 0 THEN
	INSERT INTO tab_all_sensor_data_last (
		mac,
		temperature,
		led,
		create_date
	)
VALUES
	(
		NEW.mac,
		NEW.temperature,
		NEW.led,
		NEW.create_date
	);
ELSEIF @cnt > 0 THEN
		UPDATE tab_all_sensor_data_last
		SET mac = NEW.mac,
		temperature = NEW.temperature,
		led = NEW.led,
		create_date = NEW.create_date
		WHERE
			mac = NEW.mac;
END IF;
END