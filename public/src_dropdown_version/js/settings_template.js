var settings_template = {
    "fw_version": {
        "type": "rangeredge_tracker",
        "major": "3",
        "minor": "2"
    },
    "hardware": {
        "type": "rangeredge_nrf52840",
        "version": {
            "major": "1",
            "minor": "4"
        }
    },
    "settings": {
        // "port": "port_settings",
        // "type": "rangeredge_tracker",
        "lr_gps_interval": {
            "id": "0x01",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "ublox_send_interval": {
            "id": "0x02",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 172800,
            "length": 4,
            "conversion": "uint32"
        },
        "status_send_interval": {
            "id": "0x03",
            "enabled": true,
            "default": 3600,
            "min": 1,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "satellite_send_interval": {
            "id": "0x04",
            "enabled": true,
            "default": 86400,
            "min": 0,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "gps_init_lon": {
            "id": "0x05",
            "enabled": true,
            "default": 156447700,
            "min": -1800000000,
            "max": 1800000000,
            "length": 4,
            "conversion": "int32"
        },
        "gps_init_lat": {
            "id": "0x06",
            "enabled": true,
            "default": 465556280,
            "min": -900000000,
            "max": 900000000,
            "length": 4,
            "conversion": "int32"
        },
        "init_time": {
            "id": "0x07",
            "enabled": true,
            "default": 1606314575,
            "min": 1606314575,
            "max": 4294967295,
            "length": 4,
            "conversion": "uint32"
        },
        "ble_adv": {
            "id": "0x08",
            "enabled": true,
            "default": true,
            "min": false,
            "max": true,
            "length": 1,
            "conversion": "bool"
        },
        "lr_sat_data": {
            "id": "0x09",
            "enabled": true,
            "default": true,
            "min": false,
            "max": true,
            "length": 1,
            "conversion": "bool"
        },
        "gps_sat_data": {
            "id": "0x0A",
            "enabled": true,
            "default": true,
            "min": false,
            "max": true,
            "length": 1,
            "conversion": "bool"
        },
        "data_log": {
            "id": "0x0B",
            "enabled": true,
            "default": true,
            "min": false,
            "max": true,
            "length": 1,
            "conversion": "bool"
        },
        "lr_send_flag": {
            "id": "0x0C",
            "enabled": true,
            "default": 4227859967,
            "min": 0,
            "max": 4294967295,
            "length": 4,
            "conversion": "uint32"
        },
        "flash_store_flag": {
            "id": "0x0D",
            "enabled": true,
            "default": 67110683,
            "min": 0,
            "max": 4294967295,
            "length": 4,
            "conversion": "uint32"
        },
        "lr_adr": {
            "id": "0x0E",
            "enabled": true,
            "default": 3,
            "min": 0,
            "max": 15,
            "length": 1,
            "conversion": "uint8"
        },
        "lr_region": {
            "id": "0x0F",
            "enabled": true,
            "default": 1,
            "min": 1,
            "max": 11,
            "length": 1,
            "conversion": "uint8"
        },
        "app_key": {
            "id": "0x10",
            "enabled": true,
            "default": "{0x8B, 0xCD, 0x49, 0x42, 0x11, 0x67, 0xDD, 0x03, 0xBA, 0xD3, 0xAE, 0xEA, 0x98, 0xEF, 0xE4, 0x09}",
            "min": "{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00}",
            "max": "{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00}",
            "length": 16,
            "conversion": "byte_array"
        },
        "device_eui": {
            "id": "0x11",
            "enabled": true,
            "default": "{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00}",
            "min": "{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00}",
            "max": "{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00}",
            "length": 8,
            "conversion": "byte_array"
        },
        "app_eui": {
            "id": "0x12",
            "enabled": true,
            "default": "{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00}",
            "min": "{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00}",
            "max": "{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00}",
            "length": 8,
            "conversion": "byte_array"
        },
        "horizontal_accuracy": {
            "id": "0x13",
            "enabled": true,
            "default": 50,
            "min": 1,
            "max": 1000,
            "length": 4,
            "conversion": "uint32"
        },
        "cold_fix_retry": {
            "id": "0x14",
            "enabled": true,
            "default": 200,
            "min": 1,
            "max": 255,
            "length": 1,
            "conversion": "uint8"
        },
        "hot_fix_retry": {
            "id": "0x15",
            "enabled": true,
            "default": 4,
            "min": 1,
            "max": 255,
            "length": 1,
            "conversion": "uint8"
        },
        "cold_fix_timeout": {
            "id": "0x16",
            "enabled": true,
            "default": 200,
            "min": 1,
            "max": 600,
            "length": 2,
            "conversion": "uint16"
        },
        "hot_fix_timeout": {
            "id": "0x17",
            "enabled": true,
            "default": 45,
            "min": 1,
            "max": 600,
            "length": 2,
            "conversion": "uint16"
        },
        "ble_advertisement_interval": {
            "id": "0x18",
            "enabled": true,
            "default": 500,
            "min": 100,
            "max": 10000,
            "length": 4,
            "conversion": "uint32"
        },
        "wifi_scan_interval": {
            "id": "0x19",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "wifi_scan_aggregated_interval": {
            "id": "0x1A",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "ble_scan_duration": {
            "id": "0x1B",
            "enabled": true,
            "default": 600,
            "min": 50,
            "max": 10000,
            "length": 4,
            "conversion": "uint32"
        },
        "ble_scan_interval": {
            "id": "0x1C",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "ble_scan_aggregated_interval": {
            "id": "0x1D",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "ble_scan_filter": {
            "id": "0x1E",
            "enabled": true,
            "default": 1,
            "min": 0,
            "max": 3,
            "length": 1,
            "conversion": "uint8"
        },
        "ble_auto_disconnect": {
            "id": "0x1F",
            "enabled": true,
            "default": 600,
            "min": 0,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "device_name": {
            "id": "0x21",
            "enabled": true,
            "default": "",
            "min": "",
            "max": "",
            "length": 8,
            "conversion": "string"
        },
        "lr_join_flag": {
            "id": "0x22",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 4294967295,
            "length": 4,
            "conversion": "uint32"
        },
        "lr_confirm_flag": {
            "id": "0x23",
            "enabled": true,
            "default": 8,
            "min": 0,
            "max": 4294967295,
            "length": 4,
            "conversion": "uint32"
        },
        "lr_max_confirm_fail": {
            "id": "0x24",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 1000,
            "length": 2,
            "conversion": "uint16"
        },
        "gps_backoff_factor": {
            "id": "0x25",
            "enabled": true,
            "default": 15,
            "min": 10,
            "max": 100,
            "length": 1,
            "conversion": "uint8"
        },
        "ublox_send_interval_2": {
            "id": "0x26",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "ublox_switch_interval": {
            "id": "0x27",
            "enabled": true,
            "default": 11,
            "min": 0,
            "max": 12,
            "length": 1,
            "conversion": "uint8"
        },
        "ublox_min_fix_time": {
            "id": "0x28",
            "enabled": true,
            "default": 5,
            "min": 0,
            "max": 30,
            "length": 1,
            "conversion": "uint8"
        },
        "ublox_multiple_intervals": {
            "id": "0x29",
            "enabled": true,
            "default": false,
            "min": false,
            "max": true,
            "length": 1,
            "conversion": "bool"
        },
        "device_pin": {
            "id": "0x2A",
            "enabled": true,
            "default": "{0x00,0x00,0x00,0x00}",
            "min": "{0x00,0x00,0x00,0x00}",
            "max": "{0x00,0x00,0x00,0x00}",
            "length": 4,
            "conversion": "byte_array"
        },
        "ublox_active_tracking": {
            "id": "0x2B",
            "enabled": true,
            "default": false,
            "min": false,
            "max": true,
            "length": 1,
            "conversion": "bool"
        },
        "led_enabled": {
            "id": "0x2C",
            "enabled": true,
            "default": true,
            "min": false,
            "max": true,
            "length": 1,
            "conversion": "bool"
        },
        "motion_ths": {
            "id": "0x2D",
            "enabled": true,
            "default": 6,
            "min": 0,
            "max": 63,
            "length": 1,
            "conversion": "uint8"
        },
        "enable_motion_trig_gps": {
            "id": "0x2E",
            "enabled": true,
            "default": false,
            "min": false,
            "max": true,
            "length": 1,
            "conversion": "bool"
        },
        "gps_triggered_interval": {
            "id": "0x2F",
            "enabled": true,
            "default": 600,
            "min": 0,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "gps_skipped_triggered_interval": {
            "id": "0x30",
            "enabled": true,
            "default": 5,
            "min": 0,
            "max": 255,
            "length": 1,
            "conversion": "uint8"
        },
        "lr_messaging_retry_interval": {
            "id": "0x31",
            "enabled": true,
            "default": 60,
            "min": 0,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "lr_messaging_retry_count": {
            "id": "0x32",
            "enabled": true,
            "default": 2,
            "min": 1,
            "max": 5,
            "length": 1,
            "conversion": "uint8"
        },
        "ublox_leave_on": {
            "id": "0x33",
            "enabled": true,
            "default": 15,
            "min": 0,
            "max": 60,
            "length": 1,
            "conversion": "uint8"
        },
        "memfault_send_interval": {
            "id": "0x34",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "rejoin_interval": {
            "id": "0x35",
            "enabled": true,
            "default": 43200,
            "min": 600,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "check_error_interval": {
            "id": "0x36",
            "enabled": true,
            "default": 86400,
            "min": 0,
            "max": 2678400,
            "length": 4,
            "conversion": "uint32"
        },
        "gnss_constellation_to_use": {
            "id": "0x37",
            "enabled": true,
            "default": 3,
            "min": 1,
            "max": 3,
            "length": 1,
            "conversion": "uint8"
        },
        "rf_scan_enabled": {
            "id": "0x38",
            "enabled": true,
            "default": false,
            "min": false,
            "max": true,
            "length": 1,
            "conversion": "bool"
        },
        "sat_send_flag": {
            "id": "0x39",
            "enabled": true,
            "default": 138,
            "min": 0,
            "max": 4294967295,
            "length": 4,
            "conversion": "uint32"
        },
        "satellite_enabled": {
            "id": "0x3A",
            "enabled": true,
            "default": false,
            "min": false,
            "max": true,
            "length": 1,
            "conversion": "bool"
        },
        "satellite_retry": {
            "id": "0x3B",
            "enabled": true,
            "default": 10,
            "min": 1,
            "max": 15,
            "length": 1,
            "conversion": "uint8"
        },
        "rf_scan_interval": {
            "id": "0x3C",
            "enabled": true,
            "default": 60,
            "min": 30,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "rf_scan_duration": {
            "id": "0x3D",
            "enabled": true,
            "default": 60,
            "min": 0,
            "max": 86400,
            "length": 4,
            "conversion": "uint32"
        },
        "rf_scan_config": {
            "id": "0x3E",
            "enabled": true,
            "default": "{0x01, 0x0e, 0x0a, 0x0a, 0x70, 0x03, 0x93, 0x03, 0x14, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}",
            "min": "{}",
            "max": "{}",
            "length": 46,
            "conversion": "byte_array"
        }
    },
    "commands": {
        // "port": "port_commands",
        "cmd_join": {
            "id": "0xA0",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_reset": {
            "id": "0xA1",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_send_all_val": {
            "id": "0xA2",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_send_single_val": {
            "id": "0xA3",
            "length": 1,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_send_status": {
            "id": "0xA4",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_send_position": {
            "id": "0xA5",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_reset_gps": {
            "id": "0xA6",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_send_all_settings": {
            "id": "0xA7",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_send_single_setting": {
            "id": "0xA8",
            "length": 1,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_reset_initial_position": {
            "id": "0xA9",
            "length": 1,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_reset_initial_time": {
            "id": "0xAA",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_clear_nvs": {
            "id": "0xAB",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_reset_to_def_settings": {
            "id": "0xAC",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_send_status_lr": {
            "id": "0xAD",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_send_lr_fix": {
            "id": "0xAE",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_set_location_and_time": {
            "id": "0xAF",
            "length": 12,
            "conversion": "uint32",
            "value": 0
        },
        "cmd_get_rf_scan": {
            "id": "0xB0",
            "length": 16,
            "conversion": "uint32",
            "value": 0
        },
        "cmd_get_wifi_scan": {
            "id": "0xB1",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_get_ble_scan": {
            "id": "0xB2",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_get_lr_satellite_data": {
            "id": "0xB4",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_get_ublox_satellite_data": {
            "id": "0xB5",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_almanac_update": {
            "id": "0xB6",
            "length": 250,
            "conversion": "byte_array",
            "value": 0
        },
        "cmd_get_mac": {
            "id": "0xB7",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_get_ublox_fix": {
            "id": "0xB8",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_reset_lr": {
            "id": "0xB9",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_flash_clear": {
            "id": "0xBA",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_flash_get_all": {
            "id": "0xBB",
            "length": 1,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_flash_get_from_head": {
            "id": "0xBC",
            "length": 12,
            "conversion": "uint32",
            "value": 0
        },
        "cmd_lr_max_payload": {
            "id": "0xBD",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_set_operation_mode_com_th": {
            "id": "0xBE",
            "length": 1,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_disable_flash_th": {
            "id": "0xBF",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_disable_bt_th": {
            "id": "0xC0",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_set_operation_mode_main_th": {
            "id": "0xC1",
            "length": 1,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_check_pin": {
            "id": "0xC2",
            "length": 16,
            "conversion": "byte_array",
            "value": 0
        },
        "cmd_set_hibernation_mode": {
            "id": "0xC3",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_send_lr_message": {
            "id": "0xC4",
            "length": 46,
            "conversion": "string",
            "value": 0
        },
        "cmd_read_all_lr_messages": {
            "id": "0xC5",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_send_sat_buffer": {
            "id": "0xC6",
            "length": 0,
            "conversion": "uint8",
            "value": 0
        },
        "cmd_pause_rf_scan": {
            "id": "0xC7",
            "length": 1,
            "conversion": "uint8",
            "value": 0
        }
    },
    "values": {
        "port": "port_values",
        "reset_reason": {
            "id": "0xD0",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 4294967295,
            "length": 4,
            "conversion": "uint32"
        },
        "gps_lon": {
            "id": "0xD1",
            "enabled": true,
            "default": 156447700,
            "min": -1800000000,
            "max": 1800000000,
            "length": 4,
            "conversion": "int32"
        },
        "gps_lat": {
            "id": "0xD2",
            "enabled": true,
            "default": 465556280,
            "min": -900000000,
            "max": 900000000,
            "length": 4,
            "conversion": "int32"
        },
        "gps_alt": {
            "id": "0xD3",
            "enabled": true,
            "default": 253000,
            "min": -400000,
            "max": 8000000,
            "length": 4,
            "conversion": "int32"
        },
        "lis2_acc_x": {
            "id": "0xD4",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "lis2_acc_y": {
            "id": "0xD5",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "lis2_acc_z": {
            "id": "0xD6",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "batt_mV": {
            "id": "0xD7",
            "enabled": true,
            "default": 2500,
            "min": 1500,
            "max": 4500,
            "length": 4,
            "conversion": "int32"
        },
        "ublox_time": {
            "id": "0xD8",
            "enabled": true,
            "default": 1621861202,
            "min": 1621861202,
            "max": 4294967295,
            "length": 4,
            "conversion": "uint32"
        },
        "lr_satellites": {
            "id": "0xD9",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 255,
            "length": 1,
            "conversion": "uint8"
        },
        "mcu_temp": {
            "id": "0xDA",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "charge_mV": {
            "id": "0xDB",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 32000,
            "length": 4,
            "conversion": "int32"
        },
        "lsm303_acc_x": {
            "id": "0xDC",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "lsm303_acc_y": {
            "id": "0xDD",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "lsm303_acc_z": {
            "id": "0xDE",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "lsm303_mag_x": {
            "id": "0xDF",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "lsm303_mag_y": {
            "id": "0xE0",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "lsm303_mag_z": {
            "id": "0xE1",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "lsm6_acc_x": {
            "id": "0xE2",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "lsm6_acc_y": {
            "id": "0xE3",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "lsm6_acc_z": {
            "id": "0xE4",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "lsm6_gyro_x": {
            "id": "0xE5",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "lsm6_gyro_y": {
            "id": "0xE6",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "lsm6_gyro_z": {
            "id": "0xE7",
            "enabled": true,
            "default": 0,
            "min": -100,
            "max": 100,
            "length": 4,
            "conversion": "float"
        },
        "flash_nr_msg": {
            "id": "0xE8",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 2147483647,
            "length": 4,
            "conversion": "uint32"
        },
        "last_position_time": {
            "id": "0xE9",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 4294967295,
            "length": 4,
            "conversion": "uint32"
        },
        "last_accel_int_time": {
            "id": "0xEA",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 4294967295,
            "length": 4,
            "conversion": "uint32"
        },
        "n_mes": {
            "id": "0xEB",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 5,
            "length": 1,
            "conversion": "uint8"
        },
        "almanac_age": {
            "id": "0xEC",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 65535,
            "length": 2,
            "conversion": "uint16"
        },
        "factory_device_name": {
            "id": "0xED",
            "enabled": true,
            "default": "",
            "min": "",
            "max": "",
            "length": 8,
            "conversion": "string"
        },
        "satellite_resend_try": {
            "id": "0xEF",
            "enabled": true,
            "default": 0,
            "min": 0,
            "max": 30,
            "length": 1,
            "conversion": "uint8"
        }
    },
    "messages": {
        "msg_lr_messaging": {
            "port": "port_lr_messaging",
            "id": "0x90",
            "length": 250,
            "conversion": "byte_array"
        },
        "msg_memfault_chunks": {
            "port": "port_memfault",
            "id": "0x91",
            "length": 250,
            "conversion": "byte_array"
        },
        "msg_lr_messaging_timestamp": {
            "port": "port_lr_messaging",
            "id": "0xF0",
            "length": 250,
            "conversion": "byte_array"
        },
        "msg_gnss": {
            "port": "port_lr_gps",
            "id": "0xF1",
            "length": 250,
            "conversion": "byte_array"
        },
        "msg_ublox_location": {
            "port": "port_ublox_gps",
            "id": "0xF2",
            "length": 30,
            "conversion": "byte_array"
        },
        "msg_cmd_confirm": {
            "port": "port_messages",
            "id": "0xF3",
            "length": 2,
            "conversion": "byte_array"
        },
        "msg_status": {
            "port": "port_status",
            "id": "0xF4",
            "length": 14,
            "conversion": "byte_array"
        },
        "msg_lr_satellites": {
            "port": "port_lr_sat_data",
            "id": "0xF5",
            "length": 250,
            "conversion": "byte_array"
        },
        "msg_ublox_satellites": {
            "port": "port_ublox_sat_data",
            "id": "0xF6",
            "length": 250,
            "conversion": "byte_array"
        },
        "msg_wifi_scan_aggregated": {
            "port": "port_wifi_scan_aggregated",
            "id": "0xF7",
            "length": 250,
            "conversion": "byte_array"
        },
        "msg_wifi_scan": {
            "port": "port_wifi_scan",
            "id": "0xF8",
            "length": 250,
            "conversion": "byte_array"
        },
        "msg_ble_scan_aggregated": {
            "port": "port_ble_scan_aggregated",
            "id": "0xF9",
            "length": 250,
            "conversion": "byte_array"
        },
        "msg_ble_scan": {
            "port": "port_ble_scan",
            "id": "0xFA",
            "length": 250,
            "conversion": "byte_array"
        },
        "msg_rf_scan": {
            "port": "port_rf_scan",
            "id": "0xFB",
            "length": 250,
            "conversion": "byte_array"
        },
        "msg_mac_id": {
            "port": "port_messages",
            "id": "0xFD",
            "length": 6,
            "conversion": "byte_array"
        },
        "msg_last_position": {
            "port": "port_messages",
            "id": "0xFE",
            "length": 16,
            "conversion": "byte_array"
        },
        "msg_read_flash": {
            "port": "port_flash_log",
            "id": "0xFF",
            "length": 250,
            "conversion": "byte_array"
        }
    },
    "ports": {
        "port_lr_gps": 1,
        "port_ublox_gps": 2,
        "port_settings": 3,
        "port_status": 4,
        "port_lr_sat_data": 5,
        "port_wifi_scan_aggregated": 6,
        "port_ble_scan_aggregated": 7,
        "port_rf_scan": 8,
        "port_ublox_sat_data": 9,
        "port_wifi_scan": 10,
        "port_ble_scan": 11,
        "port_memfault": 27,
        "port_lr_messaging": 28,
        "port_flash_log": 29,
        "port_values": 30,
        "port_messages": 31,
        "port_commands": 32
    }
}

export default settings_template;