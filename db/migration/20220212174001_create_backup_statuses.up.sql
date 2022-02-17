CREATE TABLE backup_statuses(
  id BIGINT NOT NULL AUTO_INCREMENT,
  backup_schedule_id BIGINT NOT NULL,
  date_backup DATETIME NOT NULL,
  is_success BOOLEAN DEFAULT FALSE,
  PRIMARY KEY (id),
  CONSTRAINT FK_backup_schedule_id
  FOREIGN KEY (backup_schedule_id) REFERENCES backup_schedules(id)
);