CREATE TABLE backup_monitors(
  id BIGINT NOT NULL AUTO_INCREMENT,
  backup_status_id BIGINT NOT NULL,
  capacity TINYINT NOT NULL,
  battery TINYINT,
  connection BOOLEAN DEFAULT FALSE,
  device_id BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(id),
  CONSTRAINT FK_backup_monitors_device_id
  FOREIGN KEY (device_id) REFERENCES devices(id),
  CONSTRAINT FK_backup_monitors_backup_status_id
  FOREIGN KEY (backup_status_id) REFERENCES backup_statuses(id)
);